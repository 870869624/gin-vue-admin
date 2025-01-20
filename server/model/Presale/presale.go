// 自动生成模板Presale
package Presale

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/funct"
	unixTime "github.com/flipped-aurora/gin-vue-admin/server/utils/ut"
)

// 预售项目 结构体  Presale
type Presale struct {
	global.GVA_MODEL
	PresaleName      *string    `json:"presaleName" form:"presaleName" gorm:"column:presale_name;comment:预售项目名称;" binding:"required"`                          //预售项目名称
	PresalePicture   string     `json:"presalePicture" form:"presalePicture" gorm:"column:presale_picture;comment:预售项目图片;" binding:"required"`                 //预售项目图片
	PresaleStartTime *time.Time `json:"presaleStartTime" form:"presaleStartTime" gorm:"column:presale_start_time;comment:预售开始时间;" binding:"required"`          //预售开始时间
	PresaleIsPass    *bool      `json:"presaleIsPass" form:"presaleIsPass" gorm:"default:false;column:presale_is_pass;comment:预售项目是否审核通过;" binding:"required"` //预售项目是否审核通过
	PresaleIsShow    *bool      `json:"presaleIsShow" form:"presaleIsShow" gorm:"default:false;column:presale_is_show;comment:预售项目是否展示;" binding:"required"`   //预售项目是否展示
	PublicChainId    *string    `json:"publicChainId" form:"publicChainId" gorm:"column:public_chain_id;comment:公链ID;" binding:"required"`                     //公链ID
	Presaleurl       *string    `json:"presaleurl" form:"presaleurl" gorm:"column:presaleurl;comment:预售项目链接;" binding:"required"`                              //预售项目链接
	Brief            *string    `json:"brief" form:"brief" gorm:"column:brief;comment:简介;"`                                                                    //简介
	Detail           *string    `json:"detail" form:"detail" gorm:"column:detail;comment:详情描述;type:text;"`                                                     //详情描述
	XLink            *string    `json:"xLink" form:"xLink" gorm:"column:x_link;comment:x链接;"`                                                                  //x链接
	TgLink           *string    `json:"tgLink" form:"tgLink" gorm:"column:tg_link;comment:tg链接;"`                                                              //tg链接
	DiscordLink      *string    `json:"discordLink" form:"discordLink" gorm:"column:discord_link;comment:discord链接;"`                                          //discord链接
}

type MobilePresale struct {
	ID               uint               `gorm:"primarykey" json:"ID"` // 主键ID
	CreatedAt        unixTime.UnixTime  // 创建时间
	UpdatedAt        unixTime.UnixTime  // 更新时间
	DeletedAt        unixTime.UnixTime  `gorm:"index" json:"-"`
	PresaleName      *string            `json:"presaleName" form:"presaleName" gorm:"column:presale_name;comment:预售项目名称;" binding:"required"`                 //预售项目名称
	PresalePicture   string             `json:"presalePicture" form:"presalePicture" gorm:"column:presale_picture;comment:预售项目图片;" binding:"required"`        //预售项目图片
	PresaleStartTime *unixTime.UnixTime `json:"presaleStartTime" form:"presaleStartTime" gorm:"column:presale_start_time;comment:预售开始时间;" binding:"required"` //预售开始时间
	UserId           *int               `json:"userId" form:"userId" gorm:"column:user_id;comment:上传者用户ID;"`                                                  //上传者用户ID
	PresaleIsPass    *bool              `json:"presaleIsPass" form:"presaleIsPass" gorm:"default:false;column:presale_is_pass;comment:预售项目是否审核通过;" `          //预售项目是否审核通过
	PresaleIsShow    *bool              `json:"presaleIsShow" form:"presaleIsShow" gorm:"default:false;column:presale_is_show;comment:预售项目是否展示;" `            //预售项目是否展示
	PublicChainId    *string            `json:"publicChainId" form:"publicChainId" gorm:"column:public_chain_id;comment:公链ID;" binding:"required"`            //公链ID
	Presaleurl       *string            `json:"url" form:"url" gorm:"column:presaleurl;comment:预售项目链接;" binding:"required"`                                   //预售项目链接
	Brief            *string            `json:"brief" form:"brief" gorm:"column:brief;comment:简介;"`                                                           //简介
	Detail           *string            `json:"detail" form:"detail" gorm:"column:detail;comment:详情描述;type:text;"`                                            //详情描述
	XLink            *string            `json:"xLink" form:"xLink" gorm:"column:x_link;comment:x链接;"`                                                         //x链接
	TgLink           *string            `json:"tgLink" form:"tgLink" gorm:"column:tg_link;comment:tg链接;"`                                                     //tg链接
	DiscordLink      *string            `json:"discordLink" form:"discordLink" gorm:"column:discord_link;comment:discord链接;"`                                 //discord链接
	Name             *string            `json:"name" form:"name" gorm:"column:name;comment:公链名称;size:256;" binding:"required"`                                //公链名称
	Picture          string             `json:"picture" form:"picture" gorm:"column:picture;comment:公链图片;size:256;" binding:"required"`                       //公链图片
	Logo             string             `json:"logo" form:"logo" gorm:"column:logo;comment:公链logo;"`                                                          //公链logo
	Introduction     *string            `json:"introduction" form:"introduction" gorm:"column:introduction;comment:256;" binding:"required"`                  //公链简介
}
type PresaleResp struct {
	PresaleName      *string            `json:"presaleName" form:"presaleName" gorm:"column:presale_name;comment:预售项目名称;" binding:"required"`                          //预售项目名称
	PresalePicture   string             `json:"presalePicture" form:"presalePicture" gorm:"column:presale_picture;comment:预售项目图片;" binding:"required"`                 //预售项目图片
	PresaleStartTime *unixTime.UnixTime `json:"presaleStartTime" form:"presaleStartTime" gorm:"column:presale_start_time;comment:预售开始时间;" binding:"required"`          //预售开始时间
	PresaleIsPass    *bool              `json:"presaleIsPass" form:"presaleIsPass" gorm:"default:false;column:presale_is_pass;comment:预售项目是否审核通过;" binding:"required"` //预售项目是否审核通过
	PresaleIsShow    *bool              `json:"presaleIsShow" form:"presaleIsShow" gorm:"default:false;column:presale_is_show;comment:预售项目是否展示;" binding:"required"`   //预售项目是否展示
	PublicChainId    *string            `json:"publicChainId" form:"publicChainId" gorm:"column:public_chain_id;comment:公链ID;" binding:"required"`                     //公链ID
	Presaleurl       *string            `json:"presaleurl" form:"presaleurl" gorm:"column:presaleurl;comment:预售项目链接;" binding:"required"`                              //预售项目链接
	Brief            *string            `json:"brief" form:"brief" gorm:"column:brief;comment:简介;"`                                                                    //简介
	Detail           *string            `json:"detail" form:"detail" gorm:"column:detail;comment:详情描述;type:text;"`                                                     //详情描述
	XLink            *string            `json:"xLink" form:"xLink" gorm:"column:x_link;comment:x链接;"`                                                                  //x链接
	TgLink           *string            `json:"tgLink" form:"tgLink" gorm:"column:tg_link;comment:tg链接;"`                                                              //tg链接
	DiscordLink      *string            `json:"discordLink" form:"discordLink" gorm:"column:discord_link;comment:discord链接;"`                                          //discord链接
}
type MobilePresaleCre struct {
	PresaleName      *string `json:"presaleName" form:"presaleName" gorm:"column:presale_name;comment:预售项目名称;" binding:"required"`          //预售项目名称
	PresalePicture   string  `json:"presalePicture" form:"presalePicture" gorm:"column:presale_picture;comment:预售项目图片;" binding:"required"` //预售项目图片
	PresaleStartTime string  `json:"presaleStartTime" form:"presaleStartTime" gorm:"column:presale_start_time;comment:预售开始时间;" binding:"required"`
	UserId           *int    `json:"userId" form:"userId" gorm:"column:user_id;comment:上传者用户ID;"`                                         //上传者用户ID
	PresaleIsPass    *bool   `json:"presaleIsPass" form:"presaleIsPass" gorm:"default:false;column:presale_is_pass;comment:预售项目是否审核通过;" ` //预售项目是否审核通过
	PresaleIsShow    *bool   `json:"presaleIsShow" form:"presaleIsShow" gorm:"default:false;column:presale_is_show;comment:预售项目是否展示;" `   //预售项目是否展示
	PublicChainId    *string `json:"publicChainId" form:"publicChainId" gorm:"column:public_chain_id;comment:公链ID;" binding:"required"`   //公链ID
	Presaleurl       *string `json:"presaleurl" form:"presaleurl" gorm:"column:presaleurl;comment:预售项目链接;" binding:"required"`            //预售项目链接
}

// TableName 预售项目 Presale自定义表名 candy_presale
func (Presale) TableName() string {
	return "candy_presale"
}

func (m *Presale) ToArgs(m1 *MobilePresaleCre) {
	timeUnix := funct.StrUnixToTime(m1.PresaleStartTime)
	m.PresaleName = m1.PresaleName
	m.PresalePicture = m1.PresalePicture
	m.PresaleStartTime = &timeUnix
	m.PresaleIsPass = m1.PresaleIsPass
	m.PresaleIsShow = m1.PresaleIsShow
	m.PublicChainId = m1.PublicChainId
	m.Presaleurl = m1.Presaleurl
	m.GVA_MODEL = global.GVA_MODEL{}
}
