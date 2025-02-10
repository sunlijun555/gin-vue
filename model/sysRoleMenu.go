package model

type SysRoleMenu struct {
	RoleId uint `gorm:"column:role_id;comment:'角色id';NOT NULL" json:"roleId"`
	MenuId uint `gorm:"column:menu_id;comment:'菜单id';NOT NULL" json:"menuId"`
}

func (SysRoleMenu) TableName() string {
	return "sys_role_menu"
}
