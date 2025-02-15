package api

import (
	. "go-vue/core"
	"go-vue/model"
	"go-vue/result"
	"go-vue/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary 新增菜单
// @Tags 菜单相关接口
// @Produce json
// @Description 新增菜单
// @Param data body model.AddSysMenuDto true "data"
// @Success 200 {object} result.Result
// @router /api/sysMenu/add [post]
func CreateSysMenu(c *gin.Context) {
	// 绑定参数
	var dto model.AddSysMenuDto
	_ = c.BindJSON(&dto)
	// 查询菜单不能重复
	sysMenuByName := GetSysMenuByName(dto.MenuName)
	if sysMenuByName.ID != 0 {
		result.Failed(c, int(result.Apicode.SysMenuIsExist), result.Apicode.GetMessage(result.Apicode.SysMenuIsExist))
		return
	}
	// 目录
	if dto.MenuType == 1 {
		sysMenu := model.SysMenu{
			ParentId:   0,
			MenuName:   dto.MenuName,
			Icon:       dto.Icon,
			MenuType:   dto.MenuType,
			Url:        dto.Url,
			MenuStatus: dto.MenuStatus,
			Sort:       dto.Sort,
			CreateTime: utils.HTime{Time: time.Now()},
		}
		Db.Create(&sysMenu)
	} else if dto.MenuType == 2 {
		sysMenu := model.SysMenu{
			ParentId:   dto.ParentId,
			MenuName:   dto.MenuName,
			Icon:       dto.Icon,
			MenuType:   dto.MenuType,
			Url:        dto.Url,
			Value:      dto.Value,
			MenuStatus: dto.MenuStatus,
			Sort:       dto.Sort,
			CreateTime: utils.HTime{Time: time.Now()},
		}
		Db.Create(&sysMenu)

	} else if dto.MenuType == 3 {
		sysMenu := model.SysMenu{
			ParentId:   dto.ParentId,
			MenuName:   dto.MenuName,
			MenuType:   dto.MenuType,
			Value:      dto.Value,
			MenuStatus: dto.MenuStatus,
			Sort:       dto.Sort,
			CreateTime: utils.HTime{Time: time.Now()},
		}
		Db.Create(&sysMenu)

	}
	result.Success(c, true)
}

// @Summary 查询菜单列表
// @Tags 菜单相关接口
// @Produce json
// @Description 查询菜单列表
// @Param menuName query string false "菜单名称"
// @Param menuStatus query string false "菜单状态"
// @Success 200 {object} result.Result
// @router /api/sysMenu/list [get]
func GetSysMenuList(c *gin.Context) {
	MenuName := c.Query("menuName")
	MenuStatus := c.Query("menuStatus")
	var sysMenu []model.SysMenu
	curDb := Db.Order("sort") // .Table("sys_menu")
	if MenuName != "" {
		curDb = curDb.Where("menu_name = ?", MenuName)
	}
	if MenuStatus != "" {
		curDb = curDb.Where("menu_status = ?", MenuStatus)
	}
	curDb.Find(&sysMenu)
	result.Success(c, sysMenu)
}

// @Summary 根据id查询菜单
// @Tags 菜单相关接口
// @Produce json
// @Description 根据id查询菜单
// @Param id query int true "菜单id"
// @Success 200 {object} result.Result
// @router /api/sysMenu/info [get]
func GetSysMenu(c *gin.Context) {
	Id, _ := strconv.Atoi(c.Query("id"))
	var sysMenu model.SysMenu
	Db.First(&sysMenu, Id)
	result.Success(c, sysMenu)
}

// @Summary 修改菜单
// @Tags 菜单相关接口
// @Produce json
// @Description 修改菜单
// @Param data body model.UpdateSysMenuDto true "data"
// @Success 200 {object} result.Result
// @router /api/sysMenu/update [put]
func UpdateSysMenu(c *gin.Context) {
	var dto model.UpdateSysMenuDto
	_ = c.BindJSON(&dto)
	var sysMenu model.SysMenu
	Db.First(&sysMenu, dto.ID)
	sysMenu.ParentId = dto.ParentId
	sysMenu.MenuName = dto.MenuName
	sysMenu.Icon = dto.Icon
	sysMenu.Value = dto.Value
	sysMenu.MenuType = dto.MenuType
	sysMenu.Url = dto.Url
	sysMenu.MenuStatus = dto.MenuStatus
	sysMenu.Sort = dto.Sort
	Db.Save(&sysMenu)
	result.Success(c, true)
}

// @Summary 根据id删除菜单
// @Tags 菜单相关接口
// @Produce json
// @Description 根据id删除菜单
// @Param data body model.SysMenuIdDto true "data"
// @Success 200 {object} result.Result
// @router /api/sysMenu/delete [delete]
func DeleteSysMenu(c *gin.Context) {
	var dto model.SysMenuIdDto
	_ = c.BindJSON(&dto)
	// 菜单已分配不能删除
	sysRoleMenu := GetSysRoleMenu(dto.ID)
	if sysRoleMenu.MenuId > 0 {
		result.Failed(c, int(result.Apicode.DelSysMenuFailed), result.Apicode.GetMessage(result.Apicode.DelSysMenuFailed))
		return
	}
	// 存在子菜单不能删除
	sysMenu := GetChildSysMenu(dto.ID)
	if sysMenu.ID > 0 {
		result.Failed(c, int(result.Apicode.NorDeleteMenu), result.Apicode.GetMessage(result.Apicode.NorDeleteMenu))
		return
	}
	Db.Delete(&model.SysMenu{}, dto.ID)
	result.Success(c, true)
}

// 根据菜单名称查询菜单
func GetSysMenuByName(menuName string) (sysMenu model.SysMenu) {
	Db.Where("menu_name = ?", menuName).First(&sysMenu)
	return sysMenu
}

// 查询菜单是否已经被分配
func GetSysRoleMenu(id uint) (sysRoleMenu model.SysRoleMenu) {
	Db.Where("menu_id = ?", id).First(&sysRoleMenu)
	return sysRoleMenu
}

// 查询是否存在子菜单
func GetChildSysMenu(id uint) (sysMenu model.SysMenu) {
	Db.Where("parent_id = ?", id).First(&sysMenu)
	return sysMenu
}
