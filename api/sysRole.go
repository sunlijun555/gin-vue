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
	result.Success(c, sysRole)
}

// @Summary 分页查询角色列表
// @Tags 角色相关接口
// @Produce json
// @Description 分页查询角色列表
// @Param pageNum query int false "分页数量"
// @Param pageSize query int false "每页大小"
// @Param roleName query string false "角色名称"
// @Param roleStatus query string false "角色状态"
// @Param beginTime query string false "开始时间"
// @Param endTime query string false "结束时间"
// @Success 200 {object} result.Result
// @router /api/sysRole/list [get]
func GetSysRoleList(c *gin.Context) {
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	roleName := c.Query("roleName")
	roleStatus := c.Query("roleStatus")
	BeginTime := c.Query("beginTime")
	EndTime := c.Query("endTime")
	if pageSize < 1 {
		pageSize = 10
	}
	if pageNum < 1 {
		pageNum = 1
	}
	var sysRoleList []model.SysRole
	var count int64
	curDb := Db.Table("sys_role")
	if roleName != "" {
		curDb = curDb.Where("roleName = ?")
	}
	if BeginTime != "" && EndTime != "" {
		curDb = curDb.Where("create_time BETWEEN ? AND ?", BeginTime, EndTime)
	}
	if roleStatus != "" {
		curDb = curDb.Where("roleStatus = ?")
	}
	curDb.Count(&count).Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("create_time DESC").Find(&sysRoleList)
	result.Success(c, map[string]interface{}{"total": count, "pageSize": pageSize, "pageNum": pageNum, "list": sysRoleList})
}

// @Summary 根据id查角色
// @Tags 角色相关接口
// @Produce json
// @Description 根据id查询角色
// @Param id query int true "角色id"
// @Success 200 {object} result.Result
// @router /api/sysRole/info [get]
func GetSysRole(c *gin.Context) {
	Id, _ := strconv.Atoi(c.Query("id"))
	sysRole := model.SysRole{}
	Db.First(&sysRole, Id)
	result.Success(c, sysRole)
}

// @Summary 修改角色
// @Tags 角色相关接口
// @Produce json
// @Description 修改角色
// @Param data body model.UpdateSysRoleDto true "data"
// @Success 200 {object} result.Result
// @router /api/sysRole/update [put]
func UpdateSysRole(c *gin.Context) {
	var dto model.UpdateSysRoleDto
	var sysRole model.SysRole
	// 绑定参数，将上下文中的body数据填充到dto
	_ = c.BindJSON(&dto)
	Db.First(&sysRole, dto.ID)
	sysRole.RoleName = dto.RoleName
	sysRole.RoleKey = dto.RoleKey
	sysRole.Status = dto.Status
	if dto.Description != "" {
		sysRole.Description = dto.Description

	}
	Db.Save(&sysRole)
	result.Success(c, true)
}

// @Summary 根据id删除角色
// @Tags 角色相关接口
// @Produce json
// @Description 根据id删除角色
// @Param data body model.SysRoleIdDto true "data"
// @Success 200 {object} result.Result
// @router /api/sysRole/delete [delete]
func DeleteSysRole(c *gin.Context) {
	var dto model.SysRoleIdDto
	// 绑定参数
	_ = c.BindJSON(&dto)
	// 已分配给用户的角色不能删除
	sysAdminRole := GetSysAdminRoleById(int(dto.ID))
	if sysAdminRole.RoleId > 0 {
		result.Failed(c, int(result.Apicode.DelSysRoleFaild), result.Apicode.GetMessage(result.Apicode.DelSysRoleFaild))
		return
	}
	var sysRole model.SysRole
	Db.Table("sys_role").Delete(&sysRole, dto.ID)
	// 删除角色关联的菜单
	Db.Table("sys_role_menu").Where("role_id = ?", dto.ID).Delete(&model.SysRoleMenu{})
	result.Success(c, true)
}

// @Summary 修改角色状态
// @Tags 角色相关接口
// @Produce json
// @Description 修改角色状态
// @Param data body model.UpdateSysRoleStatusDto true "data"
// @Success 200 {object} result.Result
// @router /api/sysRole/updateStatus [put]
func UpdateSysRoleStatus(c *gin.Context) {
	var dto model.UpdateSysRoleStatusDto
	_ = c.BindJSON(&dto)
	var sysRole model.SysRole
	Db.First(&sysRole, dto.ID)
	sysRole.Status = dto.Status
	Db.Save(&sysRole)
	result.Success(c, true)
}

// @Summary 查询角色下拉列表
// @Tags 角色相关接口
// @Produce json
// @Description 查询角色下拉列表
// @Success 200 {object} result.Result
// @router /api/sysRole/vo/list [get]
func GetSysRoleVoList(c *gin.Context) {
	var sysRoleVo []model.SysRoleVo
	Db.Table("sys_role").Select("id, role_name").Scan(&sysRoleVo)
	result.Success(c, sysRoleVo)
}

func GetSysRoleByName(roleName string) (sysRole model.SysRole) {
	Db.Where("role_name == ?", roleName).First(&sysRole)
	return sysRole
}

func GetSysAdminRoleById(sysRoleId int) (sysAdminRole model.SysAdminRole) {
	Db.Where("role_id = ?", sysRoleId).First(&sysAdminRole)
	return sysAdminRole
}
