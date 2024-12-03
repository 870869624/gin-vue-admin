
package Users

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Users"
    UsersReq "github.com/flipped-aurora/gin-vue-admin/server/model/Users/request"
)

type UsersService struct {}
// CreateUsers 创建用户列表记录
// Author [yourname](https://github.com/yourname)
func (usersService *UsersService) CreateUsers(users *Users.Users) (err error) {
	err = global.GVA_DB.Create(users).Error
	return err
}

// DeleteUsers 删除用户列表记录
// Author [yourname](https://github.com/yourname)
func (usersService *UsersService)DeleteUsers(ID string) (err error) {
	err = global.GVA_DB.Delete(&Users.Users{},"id = ?",ID).Error
	return err
}

// DeleteUsersByIds 批量删除用户列表记录
// Author [yourname](https://github.com/yourname)
func (usersService *UsersService)DeleteUsersByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]Users.Users{},"id in ?",IDs).Error
	return err
}

// UpdateUsers 更新用户列表记录
// Author [yourname](https://github.com/yourname)
func (usersService *UsersService)UpdateUsers(users Users.Users) (err error) {
	err = global.GVA_DB.Model(&Users.Users{}).Where("id = ?",users.ID).Updates(&users).Error
	return err
}

// GetUsers 根据ID获取用户列表记录
// Author [yourname](https://github.com/yourname)
func (usersService *UsersService)GetUsers(ID string) (users Users.Users, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&users).Error
	return
}

// GetUsersInfoList 分页获取用户列表记录
// Author [yourname](https://github.com/yourname)
func (usersService *UsersService)GetUsersInfoList(info UsersReq.UsersSearch) (list []Users.Users, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&Users.Users{})
    var userss []Users.Users
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Username != nil && *info.Username != "" {
        db = db.Where("username = ?",*info.Username)
    }
    if info.Email != nil && *info.Email != "" {
        db = db.Where("email = ?",*info.Email)
    }
    if info.PhoneNumber != nil && *info.PhoneNumber != "" {
        db = db.Where("phone_number = ?",*info.PhoneNumber)
    }
    if info.Age != nil {
        db = db.Where("age <> ?",*info.Age)
    }
    if info.Gender != nil && *info.Gender != "" {
        db = db.Where("gender = ?",*info.Gender)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&userss).Error
	return  userss, total, err
}
func (usersService *UsersService)GetUsersPublic() {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}