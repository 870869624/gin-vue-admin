// 自动生成模板Users
package Users

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type Login interface {
	GetUsername() string
	GetUserId() uint
	GetUserInfo() any
}

// 用户列表 结构体  Users
type Users struct {
	global.GVA_MODEL
	Username    *string `json:"username" form:"username" gorm:"column:username;comment:用户名;" binding:"required"`            //用户名
	Password    *string `json:"password" form:"password" gorm:"column:password;comment:;" binding:"required"`               //密码
	Email       *string `json:"email" form:"email" gorm:"column:email;comment:邮件;" binding:"required"`                      //邮件
	PhoneNumber *string `json:"phoneNumber" form:"phoneNumber" gorm:"column:phone_number;comment:电话号码;" binding:"required"` //电话号码
	Avatar      string  `json:"avatar" form:"avatar" gorm:"column:avatar;comment:头像;"`                                      //头像
	Brief       *string `json:"brief" form:"brief" gorm:"column:brief;comment:个人简介;"`                                       //个人简介
	Age         *int    `json:"age" form:"age" gorm:"column:age;comment:;"`                                                 //年龄
	Gender      *string `json:"gender" form:"gender" gorm:"column:gender;comment:性别;" binding:"required"`                   //性别
	Address     *string `json:"address" form:"address" gorm:"column:address;comment:住址;"`                                   //住址
	MetaMask    *string `json:"metaMask" form:"metaMask" gorm:"column:meta_mask;comment:小狐狸钱包地址;"`                          //小狐狸钱包地址
	TokenPocket *string `json:"tokenPocket" form:"tokenPocket" gorm:"column:token_pocket;comment:TP钱包地址;"`                  //TP钱包地址
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
