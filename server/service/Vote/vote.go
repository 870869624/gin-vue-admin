package Vote

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Vote"
	VoteReq "github.com/flipped-aurora/gin-vue-admin/server/model/Vote/request"
)

type VoteService struct{}

// CreateVote 创建投票记录
// Author [yourname](https://github.com/yourname)
func (voteService *VoteService) CreateVote(vote *Vote.Vote) (err error) {
	err = global.GVA_DB.Create(vote).Error
	return err
}

// DeleteVote 删除投票记录
// Author [yourname](https://github.com/yourname)
func (voteService *VoteService) DeleteVote(ID string) (err error) {
	err = global.GVA_DB.Delete(&Vote.Vote{}, "id = ?", ID).Error
	return err
}

// DeleteVoteByIds 批量删除投票记录
// Author [yourname](https://github.com/yourname)
func (voteService *VoteService) DeleteVoteByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]Vote.Vote{}, "id in ?", IDs).Error
	return err
}

// UpdateVote 更新投票记录
// Author [yourname](https://github.com/yourname)
func (voteService *VoteService) UpdateVote(vote Vote.Vote) (err error) {
	err = global.GVA_DB.Model(&Vote.Vote{}).Where("id = ?", vote.ID).Updates(&vote).Error
	return err
}

func (voteService *VoteService) SaveVote(vote *Vote.Vote) (err error) {
	err = global.GVA_DB.Save(vote).Error
	return err
}

// GetVote 根据ID获取投票记录
// Author [yourname](https://github.com/yourname)
func (voteService *VoteService) GetVote(ID string) (vote Vote.Vote, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&vote).Error
	return
}

// GetVoteInfoList 分页获取投票记录
// Author [yourname](https://github.com/yourname)
func (voteService *VoteService) GetVoteInfoList(info VoteReq.VoteSearch) (list []Vote.Vote, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&Vote.Vote{})
	var votes []Vote.Vote
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.VoteName != nil && *info.VoteName != "" {
		db = db.Where("vote_name LIKE ?", "%"+*info.VoteName+"%")
	}
	if info.VoteIsPass != nil {
		db = db.Where("vote_is_pass = ?", *info.VoteIsPass)
	}
	if info.VoteIsShow != nil {
		db = db.Where("vote_is_show = ?", *info.VoteIsShow)
	}
	if info.PublicChainId != nil && *info.PublicChainId != "" {
		db = db.Where("public_chain_id = ?", *info.PublicChainId)
	}
	if info.VoteNum != nil {
		db = db.Where("vote_num <> ?", *info.VoteNum)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&votes).Error
	return votes, total, err
}
func (voteService *VoteService) GetVotePublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

func (voteService *VoteService) MobileGetVoteInfoList(info VoteReq.VoteSearch) (list []Vote.MobileVote, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Table(Vote.Vote.TableName(Vote.Vote{}) + " v").
		Select("v.id,v.created_at,v.updated_at,v.vote_is_pass,v.vote_is_show,v.vote_name,v.vote_num,v.vote_picture,v.vote_url,cpc.logo,cpc.name,cpc.picture,v.public_chain_id,v.brief,v.detail,v.x_link,v.tg_link,v.discord_link")
	var votes []Vote.MobileVote
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("v.created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.VoteName != nil && *info.VoteName != "" {
		db = db.Where("v.vote_name LIKE ?", "%"+*info.VoteName+"%")
	}
	if info.VoteIsPass != nil {
		db = db.Where("v.vote_is_pass = ?", *info.VoteIsPass)
	}
	if info.VoteIsShow != nil {
		db = db.Where("v.vote_is_show = ?", *info.VoteIsShow)
	}
	if info.PublicChainId != nil && *info.PublicChainId != "" {
		db = db.Where("v.public_chain_id = ?", *info.PublicChainId)
	}
	if info.VoteNum != nil {
		db = db.Where("v.vote_num <> ?", *info.VoteNum)
	}
	db.Joins("left join candies_public_chain cpc ON v.public_chain_id = cpc.id")
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset).Order("v.vote_num desc")
	}

	err = db.Find(&votes).Error
	return votes, total, err
}

func (voteService *VoteService) GetVoteMobile(ID string) (vote Vote.MobileVote, err error) {
	db := global.GVA_DB.Table(Vote.Vote.TableName(Vote.Vote{}) + " v").
		Select("v.id,v.created_at,v.updated_at,v.vote_is_pass,v.vote_is_show,v.vote_name,v.vote_num,v.vote_picture,v.vote_url,cpc.logo,cpc.name,cpc.picture,v.public_chain_id,v.brief,v.detail,v.x_link,v.tg_link,v.discord_link")
	var votes Vote.MobileVote
	// 如果有条件搜索 下方会自动创建搜索语句

	db.Joins("left join candies_public_chain cpc ON v.public_chain_id = cpc.id").Where("v.id = ?", ID)

	err = db.First(&votes).Error
	return votes, err
}
