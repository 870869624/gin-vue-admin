// 自动生成模板Airdrop
package Airdrop

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/funct"
	unixTime "github.com/flipped-aurora/gin-vue-admin/server/utils/ut"
)

// 空投项目 结构体  Airdrop
type Airdrop struct {
	global.GVA_MODEL
	AirdropName    *string    `json:"airdropName" form:"airdropName" gorm:"column:airdrop_name;comment:空投名称;" binding:"required"`              //空投名称
	AirdropPicture string     `json:"airdropPicture" form:"airdropPicture" gorm:"column:airdrop_picture;comment:空投项目图片;" binding:"required"`   //空投项目图片
	AirdropValue   *float64   `json:"airdropValue" form:"airdropValue" gorm:"column:airdrop_value;comment:空投项目价值;" binding:"required"`         //空投项目价值
	AirdropEndtime *time.Time `json:"airdropEndtime" form:"airdropEndtime" gorm:"column:airdrop_endtime;comment:空投项目结束日期;" binding:"required"` //空投项目结束日期
	PublicChainId  *string    `json:"publicChainId" form:"publicChainId" gorm:"column:public_chain_id;comment:公链id;" binding:"required"`       //公链id
	AirdropIsPass  *bool      `json:"airdropIsPass" form:"airdropIsPass" gorm:"default:false;column:airdrop_is_pass;comment:;"`                //是否通过审核
	AirdropIsShow  *bool      `json:"airdropIsShow" form:"airdropIsShow" gorm:"default:true;column:airdrop_is_show;comment:;"`                 //是否展示
	AirdropUrl     *string    `json:"airdropUrl" form:"airdropUrl" gorm:"column:airdrop_url;comment:空有项目网址;" binding:"required"`               //空有项目网址
	Brief          *string    `json:"brief" form:"brief" gorm:"column:brief;comment:简介;"`                                                      //简介
	Detail         *string    `json:"detail" form:"detail" gorm:"column:detail;comment:详情描述;type:text;"`                                       //详情描述
	XLink          *string    `json:"xLink" form:"xLink" gorm:"column:x_link;comment:x链接;"`                                                    //x链接
	TgLink         *string    `json:"tgLink" form:"tgLink" gorm:"column:tg_link;comment:TG链接;"`                                                //TG链接
	DiscordLink    *string    `json:"discordLink" form:"discordLink" gorm:"column:discord_link;comment:discord链接;"`                            //discord链接
}

type MobileAirdropResp struct {
	ID             uint               `gorm:"primarykey" json:"ID"` // 主键ID
	CreatedAt      unixTime.UnixTime  // 创建时间
	UpdatedAt      unixTime.UnixTime  // 更新时间
	DeletedAt      unixTime.UnixTime  `gorm:"index" json:"-"`
	AirdropName    *string            `json:"airdropName" form:"airdropName" gorm:"column:airdrop_name;comment:空投名称;" binding:"required"`              //空投名称
	AirdropPicture string             `json:"airdropPicture" form:"airdropPicture" gorm:"column:airdrop_picture;comment:空投项目图片;" binding:"required"`   //空投项目图片
	AirdropValue   *float64           `json:"airdropValue" form:"airdropValue" gorm:"column:airdrop_value;comment:空投项目价值;" binding:"required"`         //空投项目价值
	AirdropEndtime *unixTime.UnixTime `json:"airdropEndtime" form:"airdropEndtime" gorm:"column:airdrop_endtime;comment:空投项目结束日期;" binding:"required"` //空投项目结束日期
	UserId         *int               `json:"userId" form:"userId" gorm:"column:user_id;comment:提交者用户id;"`                                             //提交者用户id
	PublicChainId  *string            `json:"publicChainId" form:"publicChainId" gorm:"column:public_chain_id;comment:公链id;" binding:"required"`       //公链id
	AirdropIsPass  *bool              `json:"airdropIsPass" form:"airdropIsPass" gorm:"default:false;column:airdrop_is_pass;comment:;"`                //是否通过审核
	AirdropIsShow  *bool              `json:"airdropIsShow" form:"airdropIsShow" gorm:"default:true;column:airdrop_is_show;comment:;"`                 //是否展示
	AirdropUrl     *string            `json:"url" form:"url" gorm:"column:airdrop_url;comment:空投项目网址;" binding:"required"`                             //空投项目网址
	Brief          *string            `json:"brief" form:"brief" gorm:"column:brief;comment:简介;"`                                                      //简介
	Detail         *string            `json:"detail" form:"detail" gorm:"column:detail;comment:详情描述;type:text;"`                                       //详情描述
	XLink          *string            `json:"xLink" form:"xLink" gorm:"column:x_link;comment:x链接;"`                                                    //x链接
	TgLink         *string            `json:"tgLink" form:"tgLink" gorm:"column:tg_link;comment:TG链接;"`                                                //TG链接
	DiscordLink    *string            `json:"discordLink" form:"discordLink" gorm:"column:discord_link;comment:discord链接;"`                            //discord链接
	Name           *string            `json:"name" form:"name" gorm:"column:name;comment:公链名称;size:256;" binding:"required"`                           //公链名称
	Picture        string             `json:"picture" form:"picture" gorm:"column:picture;comment:公链图片;size:256;" binding:"required"`                  //公链图片
	Logo           string             `json:"logo" form:"logo" gorm:"column:logo;comment:公链logo;"`                                                     //公链logo
	Introduction   *string            `json:"introduction" form:"introduction" gorm:"column:introduction;comment:256;" binding:"required"`             //公链简介
}

type AirdropResp1 struct {
	AirdropName    *string            `json:"airdropName" form:"airdropName" gorm:"column:airdrop_name;comment:空投名称;" binding:"required"`              //空投名称
	AirdropPicture string             `json:"airdropPicture" form:"airdropPicture" gorm:"column:airdrop_picture;comment:空投项目图片;" binding:"required"`   //空投项目图片
	AirdropValue   *float64           `json:"airdropValue" form:"airdropValue" gorm:"column:airdrop_value;comment:空投项目价值;" binding:"required"`         //空投项目价值
	AirdropEndtime *unixTime.UnixTime `json:"airdropEndtime" form:"airdropEndtime" gorm:"column:airdrop_endtime;comment:空投项目结束日期;" binding:"required"` //空投项目结束日期
	PublicChainId  *string            `json:"publicChainId" form:"publicChainId" gorm:"column:public_chain_id;comment:公链id;" binding:"required"`       //公链id
	AirdropIsPass  *bool              `json:"airdropIsPass" form:"airdropIsPass" gorm:"default:false;column:airdrop_is_pass;comment:;"`                //是否通过审核
	AirdropIsShow  *bool              `json:"airdropIsShow" form:"airdropIsShow" gorm:"default:true;column:airdrop_is_show;comment:;"`                 //是否展示
	AirdropUrl     *string            `json:"airdropUrl" form:"airdropUrl" gorm:"column:airdrop_url;comment:空有项目网址;" binding:"required"`               //空有项目网址
	Brief          *string            `json:"brief" form:"brief" gorm:"column:brief;comment:简介;"`                                                      //简介
	Detail         *string            `json:"detail" form:"detail" gorm:"column:detail;comment:详情描述;type:text;"`                                       //详情描述
	XLink          *string            `json:"xLink" form:"xLink" gorm:"column:x_link;comment:x链接;"`                                                    //x链接
	TgLink         *string            `json:"tgLink" form:"tgLink" gorm:"column:tg_link;comment:TG链接;"`                                                //TG链接
	DiscordLink    *string            `json:"discordLink" form:"discordLink" gorm:"column:discord_link;comment:discord链接;"`                            //discord链接
}
type MobileAirdropCre struct {
	AirdropName    *string  `json:"airdropName" form:"airdropName" gorm:"column:airdrop_name;comment:空投名称;" binding:"required"`            //空投名称
	AirdropPicture string   `json:"airdropPicture" form:"airdropPicture" gorm:"column:airdrop_picture;comment:空投项目图片;" binding:"required"` //空投项目图片
	AirdropValue   *float64 `json:"airdropValue" form:"airdropValue" gorm:"column:airdrop_value;comment:空投项目价值;" binding:"required"`       //空投项目价值
	AirdropEndtime string   `json:"airdropEndtime" form:"airdropEndtime" gorm:"column:airdrop_endtime;comment:空投项目结束日期;" binding:"required"`
	UserId         *int     `json:"userId" form:"userId" gorm:"column:user_id;comment:提交者用户id;"`                                       //提交者用户id
	PublicChainId  *string  `json:"publicChainId" form:"publicChainId" gorm:"column:public_chain_id;comment:公链id;" binding:"required"` //公链id
	AirdropIsPass  *bool    `json:"airdropIsPass" form:"airdropIsPass" gorm:"default:false;column:airdrop_is_pass;comment:;"`          //是否通过审核
	AirdropIsShow  *bool    `json:"airdropIsShow" form:"airdropIsShow" gorm:"default:true;column:airdrop_is_show;comment:;"`           //是否展示
	AirdropUrl     *string  `json:"airdropUrl" form:"airdropUrl" gorm:"column:airdrop_url;comment:空投项目网址;" binding:"required"`         //空投项目网址
}

// TableName 空投项目 Airdrop自定义表名 candy_airdrop
func (a Airdrop) TableName() string {
	return "candy_airdrop"
}

func (a *Airdrop) ToArgs(a1 *MobileAirdropCre) {
	timeUnix := funct.StrUnixToTime(a1.AirdropEndtime)
	a.AirdropName = a1.AirdropName
	a.AirdropPicture = a1.AirdropPicture
	a.AirdropValue = a1.AirdropValue
	a.AirdropEndtime = &timeUnix
	a.PublicChainId = a1.PublicChainId
	a.AirdropIsPass = a1.AirdropIsPass
	a.AirdropIsShow = a1.AirdropIsShow
	a.AirdropUrl = a1.AirdropUrl
	a.GVA_MODEL = global.GVA_MODEL{}
}
