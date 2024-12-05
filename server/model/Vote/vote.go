// 自动生成模板Vote
package Vote

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 投票 结构体  Vote
type Vote struct {
	global.GVA_MODEL
	VoteName      *string `json:"voteName" form:"voteName" gorm:"column:vote_name;comment:投票项目名称;" binding:"required"`               //投票项目名称
	VotePicture   string  `json:"votePicture" form:"votePicture" gorm:"column:vote_picture;comment:投票项目图片;" binding:"required"`      //投票项目图片
	UserId        *int    `json:"userId" form:"userId" gorm:"column:user_id;comment:提交者id;"`                                         //提交者id
	VoteIsPass    *bool   `json:"voteIsPass" form:"voteIsPass" gorm:"default:false;column:vote_is_pass;comment:投票是否通过审核;"`           //投票是否通过审核
	VoteIsShow    *bool   `json:"voteIsShow" form:"voteIsShow" gorm:"default:false;column:vote_is_show;comment:投票项目是否展示;"`           //投票项目是否展示
	PublicChainId *string `json:"publicChainId" form:"publicChainId" gorm:"column:public_chain_id;comment:公链id;" binding:"required"` //公链id
	VoteNum       *int    `json:"voteNum" form:"voteNum" gorm:"column:vote_num;comment:投票总数;"`                                       //投票总数
	VoteUrl       *string `json:"voteUrl" form:"voteUrl" gorm:"column:vote_url;comment:投票项目链接;"`                                     //投票项目链接
}

type MobileVote struct {
	global.GVA_MODEL
	VoteName      *string `json:"voteName" form:"voteName" gorm:"column:vote_name;comment:投票项目名称;" binding:"required"`               //投票项目名称
	VotePicture   string  `json:"votePicture" form:"votePicture" gorm:"column:vote_picture;comment:投票项目图片;" binding:"required"`      //投票项目图片
	UserId        *int    `json:"userId" form:"userId" gorm:"column:user_id;comment:提交者id;"`                                         //提交者id
	VoteIsPass    *bool   `json:"voteIsPass" form:"voteIsPass" gorm:"default:false;column:vote_is_pass;comment:投票是否通过审核;"`           //投票是否通过审核
	VoteIsShow    *bool   `json:"voteIsShow" form:"voteIsShow" gorm:"default:false;column:vote_is_show;comment:投票项目是否展示;"`           //投票项目是否展示
	PublicChainId *string `json:"publicChainId" form:"publicChainId" gorm:"column:public_chain_id;comment:公链id;" binding:"required"` //公链id
	VoteNum       *int    `json:"voteNum" form:"voteNum" gorm:"column:vote_num;comment:投票总数;"`                                       //投票总数
	VoteUrl       *string `json:"voteUrl" form:"voteUrl" gorm:"column:vote_url;comment:投票项目链接;"`                                     //投票项目链接
	Name          *string `json:"name" form:"name" gorm:"column:name;comment:公链名称;size:256;" binding:"required"`                     //公链名称
	Picture       string  `json:"picture" form:"picture" gorm:"column:picture;comment:公链图片;size:256;" binding:"required"`            //公链图片
	Logo          string  `json:"logo" form:"logo" gorm:"column:logo;comment:公链logo;"`                                               //公链logo
	Introduction  *string `json:"introduction" form:"introduction" gorm:"column:introduction;comment:256;" binding:"required"`       //公链简介
	Url           *string `json:"url" form:"url" gorm:"column:url;comment:公链网址;size:256;" binding:"required"`                        //公链网址
}

// TableName 投票 Vote自定义表名 candy_vote
func (Vote) TableName() string {
	return "candy_vote"
}
