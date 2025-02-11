package result

// Codes定义状态
type Codes struct {
	Message          map[uint]string
	Success          uint
	Failed           uint
	SysMenuIsExist   uint
	DelSysMenuFailed uint
	NorDeleteMenu    uint
	SysRoleIsExist   uint
	DelSysRoleFaild  uint
}

//ApiCode状态码，并实例化
var Apicode = &Codes{
	Success:          200,
	Failed:           501,
	SysMenuIsExist:   600,
	DelSysMenuFailed: 601,
	NorDeleteMenu:    602,
	SysRoleIsExist:   603,
	DelSysRoleFaild:  604,
}

//状态信息初始化
func init() {
	Apicode.Message = map[uint]string{
		Apicode.Success:          "成功",
		Apicode.Failed:           "失败",
		Apicode.SysMenuIsExist:   "菜单名称已存在，请重新输入",
		Apicode.DelSysMenuFailed: "菜单已分配不能删除",
		Apicode.NorDeleteMenu:    "存在子菜单不能删除",
		Apicode.SysRoleIsExist:   "系统角色已存在，请重新输入",
		Apicode.DelSysRoleFaild:  "角色已分配不能删除",
	}
}

//GetMessage外部调用
func (c *Codes) GetMessage(code uint) string {
	message, ok := c.Message[code]
	if !ok {
		return ""
	}
	return message
}
