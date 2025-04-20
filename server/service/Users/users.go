package Users

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/InvitationRecord"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Users"
	UsersReq "github.com/flipped-aurora/gin-vue-admin/server/model/Users/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Vip"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UsersService struct{}

// CreateUsers 创建用户列表记录
// Author [yourname](https://github.com/yourname)
func (usersService *UsersService) CreateUsers(users *Users.Users) (err error) {
	err = global.GVA_DB.Create(users).Error
	return err
}

// DeleteUsers 删除用户列表记录
// Author [yourname](https://github.com/yourname)
func (usersService *UsersService) DeleteUsers(ID string) (err error) {
	err = global.GVA_DB.Delete(&Users.Users{}, "id = ?", ID).Error
	return err
}

// DeleteUsersByIds 批量删除用户列表记录
// Author [yourname](https://github.com/yourname)
func (usersService *UsersService) DeleteUsersByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]Users.Users{}, "id in ?", IDs).Error
	return err
}

// UpdateUsers 更新用户列表记录
// Author [yourname](https://github.com/yourname)
func (usersService *UsersService) UpdateUsers(users Users.Users) (err error) {
	err = global.GVA_DB.Model(&Users.Users{}).Where("id = ?", users.ID).Updates(&users).Error
	return err
}

// GetUsers 根据ID获取用户列表记录
// Author [yourname](https://github.com/yourname)
func (usersService *UsersService) GetUsers(ID string) (users Users.Users, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&users).Error
	return
}

// GetUsersInfoList 分页获取用户列表记录
// Author [yourname](https://github.com/yourname)
func (usersService *UsersService) GetUsersInfoList(info UsersReq.UsersSearch) (list []Users.Users, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&Users.Users{})
	var userss []Users.Users
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Username != nil && *info.Username != "" {
		db = db.Where("username = ?", *info.Username)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&userss).Error
	return userss, total, err
}
func (usersService *UsersService) GetUsersPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

func (usersService *UsersService) Login(u *Users.Users, l *UsersReq.Login) (userInter *Users.Users, err error) {
	if nil == global.GVA_DB {
		return nil, fmt.Errorf("db not init")
	}
	var user Users.Users
	if u.MetaMask != nil {
		err = global.GVA_DB.Debug().Where("meta_mask = ?", u.MetaMask).First(&user).Error
		if err == nil {
			err = global.GVA_DB.Debug().Table(user.TableName()).Where("id = ?", user.ID).Update("meta_mask_money", u.MetaMaskMoney).Error
			if err != nil {
				return nil, err
			}

			userId := int(user.ID)
			vipRecord := Vip.VipRecord{
				UserId: &userId,
			}
			if err := vipRecord.GetUserId(); err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					user.IsVip = false
				}
			} else {
				user.IsVip = true
			}

		} else if errors.Is(err, gorm.ErrRecordNotFound) {
			err = global.GVA_DB.Debug().Where("username = ?", u.Username).First(&user).Error
			if err == nil {
				err = global.GVA_DB.Debug().Table(user.TableName()).Where("id = ?", user.ID).Update("meta_mask_money", u.MetaMaskMoney).Error
				if err != nil {
					return nil, err
				}
				err = global.GVA_DB.Debug().Table(user.TableName()).Where("id = ?", user.ID).Update("meta_mask", u.MetaMask).Error
				if err != nil {
					return nil, err
				}

				userId := int(user.ID)
				vipRecord := Vip.VipRecord{
					UserId: &userId,
				}
				if err := vipRecord.GetUserId(); err != nil {
					if errors.Is(err, gorm.ErrRecordNotFound) {
						user.IsVip = false
						goto next
					}
					return nil, err
				} else {
					user.IsVip = true
				}
			next:
				return &user, nil
			}

			id := generate11DigitID()
			uuidString := strconv.Itoa(id)
			user := &Users.Users{Username: &uuidString, Avatar: "", MetaMask: u.MetaMask, TokenPocket: u.TokenPocket, MetaMaskMoney: u.MetaMaskMoney, TokenPocketMoney: u.TokenPocketMoney}
			user, err = usersService.Register(user)
			if err != nil {
				global.GVA_LOG.Error("注册失败!", zap.Error(err))
				return
			}
			return user, nil
		}
		return &user, err
	} else if u.TokenPocket != nil {
		err = global.GVA_DB.Debug().Where("token_pocket = ?", u.TokenPocket).First(&user).Error
		if err == nil {
			err = global.GVA_DB.Debug().Table(user.TableName()).Where("id = ?", user.ID).Update("token_pocket_money", u.TokenPocketMoney).Error
			if err != nil {
				return nil, err
			}

			userId := int(user.ID)
			vipRecord := Vip.VipRecord{
				UserId: &userId,
			}
			if err := vipRecord.GetUserId(); err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					user.IsVip = false
				}
			} else {
				user.IsVip = true
			}
		} else if errors.Is(err, gorm.ErrRecordNotFound) {
			err = global.GVA_DB.Debug().Where("username = ?", u.Username).First(&user).Error
			if err == nil {
				err = global.GVA_DB.Debug().Table(user.TableName()).Where("id = ?", user.ID).Update("token_pocket_money", u.TokenPocketMoney).Error
				if err != nil {
					return nil, err
				}
				err = global.GVA_DB.Debug().Table(user.TableName()).Where("id = ?", user.ID).Update("token_pocket", u.TokenPocket).Error
				if err != nil {
					return nil, err
				}

				userId := int(user.ID)
				vipRecord := Vip.VipRecord{
					UserId: &userId,
				}
				if err := vipRecord.GetUserId(); err != nil {
					if errors.Is(err, gorm.ErrRecordNotFound) {
						user.IsVip = false
						goto next1
					}
					return nil, err
				} else {
					user.IsVip = true
				}

			next1:
				return &user, nil
			}

			id := generate11DigitID()
			uuidString := strconv.Itoa(id)
			user := &Users.Users{Username: &uuidString, Avatar: "", MetaMask: u.MetaMask, TokenPocket: u.TokenPocket, MetaMaskMoney: u.MetaMaskMoney, TokenPocketMoney: u.TokenPocketMoney}
			user, err = usersService.Register(user)
			if err != nil {
				global.GVA_LOG.Error("注册失败!", zap.Error(err))
				return
			}

			IR := &InvitationRecord.InvitationRecord{
				InviteCode: &l.InviteCode,
			}
			if err := IR.GetByCode(); err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return nil, errors.New("邀请码错误")
				}
				return nil, err
			} else {
				global.GVA_DB.Table(IR.TableName()).Where("invite_code = ?", l.InviteCode).Update("new_user_id", user.ID)
			}

			return user, nil
		}
		return &user, err
	}
	return nil, errors.New("用户不存在")
}

func (usersService *UsersService) Register(u *Users.Users) (userInter *Users.Users, err error) {
	var user Users.Users
	if u.MetaMask != nil {
		if !errors.Is(global.GVA_DB.Where("username = ? and meta_mask = ?", u.Username, u.MetaMask).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
			return userInter, errors.New("用户名已注册")
		}
	} else if u.TokenPocket != nil {
		if !errors.Is(global.GVA_DB.Where("username = ? and token_pocket = ?", u.Username, u.TokenPocket).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
			return userInter, errors.New("用户名已注册")
		}
	}

	// 否则 附加uuid 密码hash加密 注册
	err = global.GVA_DB.Create(&u).Error
	return u, err
}

func (usersService *UsersService) ChangeUsername(u *Users.Users, newPassword string) (userInter *Users.Users, err error) {
	var user Users.Users
	if err = global.GVA_DB.Where("id = ?", u.ID).First(&user).Error; err != nil {
		return nil, err
	}

	*user.Username = newPassword
	err = global.GVA_DB.Save(&user).Error
	return &user, err
}

func generate11DigitID() int {
	// 获取当前时间的Unix毫秒时间戳
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	// 将时间戳转换为11位数字ID
	id := int(timestamp % 1e11)
	return id
}
