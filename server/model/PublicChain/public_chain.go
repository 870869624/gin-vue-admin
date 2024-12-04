// 自动生成模板PublicChain
package PublicChain

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 公链 结构体  PublicChain
type PublicChain struct {
	global.GVA_MODEL
	Name         *string `json:"name" form:"name" gorm:"column:name;comment:公链名称;size:256;" binding:"required"`               //公链名称
	Picture      string  `json:"picture" form:"picture" gorm:"column:picture;comment:公链图片;size:256;" binding:"required"`      //公链图片
	Logo         string  `json:"logo" form:"logo" gorm:"column:logo;comment:公链logo;"`                                         //公链logo
	Introduction *string `json:"introduction" form:"introduction" gorm:"column:introduction;comment:256;" binding:"required"` //公链简介
	Url          *string `json:"url" form:"url" gorm:"column:url;comment:公链网址;size:256;" binding:"required"`                  //公链网址
	Is_Show      *bool   `json:"is_Show" form:"is_Show" gorm:"default:true;column:is__show;comment:;"`                        //是否展示
}

// TableName 公链 PublicChain自定义表名 candies_public_chain
func (PublicChain) TableName() string {
	return "candies_public_chain"
}
