// 自动生成模板Users
package Users

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 用户列表 结构体  Users
type Users struct {
	global.GVA_MODEL
	Username         *string  `json:"username" form:"username" gorm:"column:username;comment:用户名;" binding:"required"`           //用户名
	Avatar           string   `json:"avatar" form:"avatar" gorm:"column:avatar;comment:头像;"`                                     //头像
	MetaMaskMoney    *float64 `json:"metaMaskMoney" form:"metaMaskMoney" gorm:"column:meta_mask_money;comment:小狐狸钱包余额;"`         //小狐狸钱包余额
	TokenPocket      *string  `json:"tokenPocket" form:"tokenPocket" gorm:"column:token_pocket;comment:TP钱包地址;"`                 //TP钱包地址
	MetaMask         *string  `json:"metaMask" form:"metaMask" gorm:"column:meta_mask;comment:小狐狸钱包地址;"`                         //小狐狸钱包地址
	TokenPocketMoney *float64 `json:"tokenPocketMoney" form:"tokenPocketMoney" gorm:"column:token_pocket_money;comment:TP钱包余额;"` //TP钱包余额
	IsVip            bool     `json:"isVip" form:"isVip" gorm:"column:is_vip;comment:是否是VIP;"`                                   //是否是VIP
}

// TableName 用户列表 Users自定义表名 candy_users

func (Users) TableName() string {
	return "candy_users"
}

func (u *Users) GetUsername() string {
	return *u.Username
}

func (u *Users) GetUserId() uint {
	return u.ID
}

func (u *Users) GetUserInfo() any {
	return *u
}

func (u *Users) GetUser() error {
	return global.GVA_DB.Table(u.TableName()).Where("id = ?", u.GetUserId()).First(&u).Error
}
