package api

import (
	. "go-vue/core"
	"go-vue/model"
	"go-vue/result"
	"go-vue/utils"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary 新增角色
// @Tags 角色相关接口
// @Produce json
// @Description 新增角色
// @Param data body model.CreateSysRoleDto true "data"
// @Success 200 {object} result.Result
// @router /api/sysRole/add [post]
func CreateSysRole(c *gin.Context) {
	var dto model.CreateSysRoleDto
	// 绑定参数
	_ = c.BindJSON(&dto)
	// 查询角色是否重复
	sysRoleByName := GetSysRoleByName(dto.RoleName)
	if sysRoleByName.ID > 0 {
		result.Failed(c, int(result.Apicode.SysRoleIsExist), result.Apicode.GetMessage(result.Apicode.SysRoleIsExist))
		return
	}
	sysRole := model.SysRole{
		RoleName:    dto.RoleName,
		RoleKey:     dto.RoleKey,
		Status:      dto.Status,
		Description: dto.Description,
		CreateTime:  utils.HTime{Time: time.Now()},
	}
	Db.Create(&sysRole)
}

// @Summary 查询角色列表
// @Tags 角色相关接口
// @Produce json
// @Description 查询角色列表
// @Param roleName query string false "角色名称"
// @Param roleStatus query string false "角色状态"
// @Success 200 {object} result.Result
// @router /api/sysRole/list [get]
func GetSysRoleList(c *gin.Context) {
	roleName := c.Query("roleName")
	roleStatus := c.Query("roleStatus")
	var sysRoleList []model.SysRole
	curDb := Db.Table("sys_role")
	if roleName != "" {
		curDb = curDb.Where("roleName = ?")
	}
	if roleStatus != "" {
		curDb = curDb.Where("roleStatus = ?")
	}
	curDb.Find(&sysRoleList)
	result.Success(c, sysRoleList)
}

func GetSysRoleByName(roleName string) (sysRole model.SysRole) {
	Db.Where("role_name == ?", roleName).First(&sysRole)
	return sysRole
}
