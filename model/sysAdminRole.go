package model

type SysAdminRole struct {
	RoleId  uint `gorm:"column:role_id;comment:'角色id';NOT NULL" json:"roleId"`
	AdminId uint `gorm:"column:admin_id;comment:'用户id';NOT NULL" json:"adminId"`
}

func (SysAdminRole) TableName() string {
	return "sys_admin_role"
}
