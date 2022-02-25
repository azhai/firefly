package db

// 查询符合条件的所有行
func GetAllRoles() (objs []*Role, err error) {
	err = engine.Find(&objs)
	return
}

// 查询某个用户的所有角色名
func GetUserRoles(uid string) (roles []string) {
	Table(new(UserRole)).Where("user_uid = ?", uid).Cols("role_name").Find(&roles)
	return
}

// 查询属于某个角色的所有用户
func (m UserRole) GetRoleUsers(roleName string) (users []*User) {
	var uids []string
	Table(new(UserRole)).Where("role_name = ?", roleName).Cols("user_uid").Find(&uids)
	engine.Where("uid IN (?)", uids).Find(&users)
	return
}
