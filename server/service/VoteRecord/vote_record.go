package VoteRecord

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/VoteRecord"
	VoteRecordReq "github.com/flipped-aurora/gin-vue-admin/server/model/VoteRecord/request"
)

type VoteRecordService struct{}

// CreateVoteRecord 创建投票记录记录
// Author [yourname](https://github.com/yourname)
func (voteRecordService *VoteRecordService) CreateVoteRecord(voteRecord *VoteRecord.VoteRecord) (err error) {
	err = global.GVA_DB.Create(voteRecord).Error
	return err
}

// DeleteVoteRecord 删除投票记录记录
// Author [yourname](https://github.com/yourname)
func (voteRecordService *VoteRecordService) DeleteVoteRecord(ID string) (err error) {
	err = global.GVA_DB.Unscoped().Delete(&VoteRecord.VoteRecord{}, "id = ?", ID).Error
	return err
}

// DeleteVoteRecordByIds 批量删除投票记录记录
// Author [yourname](https://github.com/yourname)
func (voteRecordService *VoteRecordService) DeleteVoteRecordByIds(IDs []string) (err error) {
	err = global.GVA_DB.Unscoped().Delete(&[]VoteRecord.VoteRecord{}, "id in ?", IDs).Error
	return err
}

// UpdateVoteRecord 更新投票记录记录
// Author [yourname](https://github.com/yourname)
func (voteRecordService *VoteRecordService) UpdateVoteRecord(voteRecord VoteRecord.VoteRecord) (err error) {
	err = global.GVA_DB.Model(&VoteRecord.VoteRecord{}).Where("id = ?", voteRecord.ID).Updates(&voteRecord).Error
	return err
}

// GetVoteRecord 根据ID获取投票记录记录
// Author [yourname](https://github.com/yourname)
func (voteRecordService *VoteRecordService) GetVoteRecord(ID string) (voteRecord VoteRecord.VoteRecord, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&voteRecord).Error
	return
}

// GetVoteRecordInfoList 分页获取投票记录记录
// Author [yourname](https://github.com/yourname)
func (voteRecordService *VoteRecordService) GetVoteRecordInfoList(info VoteRecordReq.VoteRecordSearch) (list []VoteRecord.VoteRecord, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&VoteRecord.VoteRecord{})
	var voteRecords []VoteRecord.VoteRecord
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.VoteId != nil {
		db = db.Where("vote_id = ?", *info.VoteId)
	}
	if info.UserId != nil {
		db = db.Where("user_id = ?", *info.UserId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&voteRecords).Error
	return voteRecords, total, err
}
func (voteRecordService *VoteRecordService) GetVoteRecordPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

func (voteRecordService *VoteRecordService) MobileGetVoteRecord(ID string) (voteRecord VoteRecord.VoteRecord, err error) {
	err = global.GVA_DB.Where("vote_id = ?", ID).First(&voteRecord).Error
	return
}

func (voteRecordService *VoteRecordService) MobileCountMyVoteRecord(ID string) (int64, error) {
	var voteRecord int64
	if err := global.GVA_DB.Debug().Model(&VoteRecord.VoteRecord{}).Where("created_at >= CURDATE() AND created_at < CURDATE() + INTERVAL 1 DAY and user_id = ?", ID).Count(&voteRecord).Error; err != nil {
		return 0, err
	}
	return voteRecord, nil
}

type Resp struct {
	Ranking      int    `json:"ranking" gorm:"column:ranking"`
	TotalVotes   int    `json:"total_votes" gorm:"column:total_votes"`
	VoteId       string `json:"vote_id" gorm:"column:vote_id"`
	MyVoteRecord int64  `json:"my_vote_record" gorm:"column:my_vote_record"`
}

func (voteRecordService *VoteRecordService) MobileCountAllVoteRecord(ID string) (*Resp, error) {
	var res Resp
	sql := "SELECT vote_id, total_votes, ranking " +
		"FROM(" +
		"SELECT vote_id, total_votes," +
		"@rank := @rank + 1 AS ranking " +
		"FROM(" +
		"SELECT vote_id, COUNT(user_id) AS total_votes " +
		"FROM candy_vote_record " +
		"GROUP BY vote_id " +
		"ORDER BY total_votes DESC " +
		")AS SubQuery, " +
		"(SELECT @rank := 0) AS Init " +
		")AS OuterQuery " +
		"WHERE vote_id =  " + ID

	if err := global.GVA_DB.Debug().Raw(sql).Scan(&res).Error; err != nil {
		return nil, err
	}
	return &res, nil
}
