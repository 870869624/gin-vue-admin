package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type UsersSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	Username       *string    `json:"username" form:"username" `
	request.PageInfo
}

type Login struct {
	Username         string  `json:"username" form:"username" gorm:"column:username;comment:用户名;"`                              //用户名
	MetaMaskMoney    float64 `json:"metaMaskMoney" form:"metaMaskMoney" gorm:"column:meta_mask_money;comment:小狐狸钱包余额;"`         //小狐狸钱包余额
	TokenPocket      string  `json:"tokenPocket" form:"tokenPocket" gorm:"column:token_pocket;comment:TP钱包地址;"`                 //TP钱包地址
	MetaMask         string  `json:"metaMask" form:"metaMask" gorm:"column:meta_mask;comment:小狐狸钱包地址;"`                         //小狐狸钱包地址
	TokenPocketMoney float64 `json:"tokenPocketMoney" form:"tokenPocketMoney" gorm:"column:token_pocket_money;comment:TP钱包余额;"` //TP钱包余额
	Captcha          string  `json:"captcha"`                                                                                   // 验证码
	CaptchaId        string  `json:"captchaId"`                                                                                 // 验证码ID
}
