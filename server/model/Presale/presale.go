// 自动生成模板Presale
package Presale

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 预售项目 结构体  Presale
type Presale struct {
	global.GVA_MODEL
	PresaleName      *string    `json:"presaleName" form:"presaleName" gorm:"column:presale_name;comment:预售项目名称;" binding:"required"`                 //预售项目名称
	PresalePicture   string     `json:"presalePicture" form:"presalePicture" gorm:"column:presale_picture;comment:预售项目图片;" binding:"required"`        //预售项目图片
	PresaleStartTime *time.Time `json:"presaleStartTime" form:"presaleStartTime" gorm:"column:presale_start_time;comment:预售开始时间;" binding:"required"` //预售开始时间
	UserId           *int       `json:"userId" form:"userId" gorm:"column:user_id;comment:上传者用户ID;"`                                                  //上传者用户ID
	PresaleIsPass    *bool      `json:"presaleIsPass" form:"presaleIsPass" gorm:"default:false;column:presale_is_pass;comment:预售项目是否审核通过;"`           //预售项目是否审核通过
	PresaleIsShow    *bool      `json:"presaleIsShow" form:"presaleIsShow" gorm:"default:false;column:presale_is_show;comment:预售项目是否展示;"`             //预售项目是否展示
	PublicChainId    *string    `json:"publicChainId" form:"publicChainId" gorm:"column:public_chain_id;comment:公链ID;" binding:"required"`            //公链ID
	Presaleurl       *string    `json:"presaleurl" form:"presaleurl" gorm:"column:presaleurl;comment:预售项目链接;" binding:"required"`                     //预售项目链接
}

type MobilePresale struct {
	global.GVA_MODEL
	PresaleName      *string    `json:"presaleName" form:"presaleName" gorm:"column:presale_name;comment:预售项目名称;" binding:"required"`                 //预售项目名称
	PresalePicture   string     `json:"presalePicture" form:"presalePicture" gorm:"column:presale_picture;comment:预售项目图片;" binding:"required"`        //预售项目图片
	PresaleStartTime *time.Time `json:"presaleStartTime" form:"presaleStartTime" gorm:"column:presale_start_time;comment:预售开始时间;" binding:"required"` //预售开始时间
	UserId           *int       `json:"userId" form:"userId" gorm:"column:user_id;comment:上传者用户ID;"`                                                  //上传者用户ID
	PresaleIsPass    *bool      `json:"presaleIsPass" form:"presaleIsPass" gorm:"default:false;column:presale_is_pass;comment:预售项目是否审核通过;" `          //预售项目是否审核通过
	PresaleIsShow    *bool      `json:"presaleIsShow" form:"presaleIsShow" gorm:"default:false;column:presale_is_show;comment:预售项目是否展示;" `            //预售项目是否展示
	PublicChainId    *string    `json:"publicChainId" form:"publicChainId" gorm:"column:public_chain_id;comment:公链ID;" binding:"required"`            //公链ID
	Presaleurl       *string    `json:"presaleurl" form:"presaleurl" gorm:"column:presaleurl;comment:预售项目链接;" binding:"required"`                     //预售项目链接
	Name             *string    `json:"name" form:"name" gorm:"column:name;comment:公链名称;size:256;" binding:"required"`                                //公链名称
	Picture          string     `json:"picture" form:"picture" gorm:"column:picture;comment:公链图片;size:256;" binding:"required"`                       //公链图片
	Logo             string     `json:"logo" form:"logo" gorm:"column:logo;comment:公链logo;"`                                                          //公链logo
	Introduction     *string    `json:"introduction" form:"introduction" gorm:"column:introduction;comment:256;" binding:"required"`                  //公链简介
	Url              *string    `json:"url" form:"url" gorm:"column:url;comment:公链网址;size:256;" binding:"required"`                                   //公链网址
}

// TableName 预售项目 Presale自定义表名 candy_presale
func (Presale) TableName() string {
	return "candy_presale"
}
