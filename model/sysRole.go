package model

import "go-vue/utils"

type SysRole struct {
	ID          uint        `gorm:"column:id;comment:'主键';primaryKey;NOT NULL" json:"id"`
	RoleName    string      `gorm:"column:role_name;varchar(64);comment:'角色名称';NOT NULL" json:"roleName"`
	RoleKey     string      `gorm:"column:role_key;varchar(64);comment:'权限字符串';NOT NULL" json:"roleKey"`
	Status      int         `grom:"column:status;default:1;comment:'账号启动状态: 1->启用,2->禁用';NOT NULL" json:"status"`
	Description string      `gorm:"column:description;varchar(500);comment:'描述'" json:"description"`
	CreateTime  utils.HTime `grom:"column:create_time;comment:'创建时间';NOT NULL" json:"createTime"`
}

func (SysRole) TableName() string {
	return "sys_role"
}

type CreateSysRoleDto struct {
	RoleName    string `json:"roleName"`
	RoleKey     string `json:"roleKey"`
	Status      int    `json:"status"`
	Description string `json:"description"`
}

type UpdateSysRoleDto struct {
	ID          uint   `json:"id"`
	RoleName    string `json:"roleName"`
	RoleKey     string `json:"roleKey"`
	Status      int    `json:"status"`
	Description string `json:"description"`
}

type SysRoleIdDto struct {
	ID uint `json:"id"`
}

type UpdateSysRoleStatusDto struct {
	ID     uint `json:"id"`
	Status int  `json:"status"`
}

type SysRoleVo struct {
	ID       uint   `json:"id"`
	RoleName string `json:"roleName"`
}
