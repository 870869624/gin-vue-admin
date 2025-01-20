package Vip

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Vip"
	VipReq "github.com/flipped-aurora/gin-vue-admin/server/model/Vip/request"
)

type VipRecordService struct{}

// CreateVipRecord 创建会员记录记录
// Author [yourname](https://github.com/yourname)
func (vipRecordService *VipRecordService) CreateVipRecord(vipRecord *Vip.VipRecord) (err error) {
	err = global.GVA_DB.Create(vipRecord).Error
	return err
}

// DeleteVipRecord 删除会员记录记录
// Author [yourname](https://github.com/yourname)
func (vipRecordService *VipRecordService) DeleteVipRecord(ID string) (err error) {
	err = global.GVA_DB.Delete(&Vip.VipRecord{}, "id = ?", ID).Error
	return err
}

// DeleteVipRecordByIds 批量删除会员记录记录
// Author [yourname](https://github.com/yourname)
func (vipRecordService *VipRecordService) DeleteVipRecordByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]Vip.VipRecord{}, "id in ?", IDs).Error
	return err
}

// UpdateVipRecord 更新会员记录记录
// Author [yourname](https://github.com/yourname)
func (vipRecordService *VipRecordService) UpdateVipRecord(vipRecord Vip.VipRecord) (err error) {
	err = global.GVA_DB.Model(&Vip.VipRecord{}).Where("id = ?", vipRecord.ID).Updates(&vipRecord).Error
	return err
}

// GetVipRecord 根据ID获取会员记录记录
// Author [yourname](https://github.com/yourname)
func (vipRecordService *VipRecordService) GetVipRecord(ID string) (vipRecord Vip.VipRecord, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&vipRecord).Error
	return
}

// GetVipRecordInfoList 分页获取会员记录记录
// Author [yourname](https://github.com/yourname)
func (vipRecordService *VipRecordService) GetVipRecordInfoList(info VipReq.VipRecordSearch) (list []Vip.VipRecord, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&Vip.VipRecord{})
	var vipRecords []Vip.VipRecord
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.IsEffective != nil {
		db = db.Where("is_effective = ?", *info.IsEffective)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&vipRecords).Error
	return vipRecords, total, err
}
func (vipRecordService *VipRecordService) GetVipRecordPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

func (vipRecordService *VipRecordService) GetVipRecordMobile(ID string) (vipRecord Vip.VipRecord, err error) {
	err = global.GVA_DB.Model(&Vip.VipRecord{}).Where("user_id = ?", ID).First(&vipRecord).Error
	return
}

func (vipRecordService *VipRecordService) CreateVipRecordMobile(vipRecord *Vip.VipRecordMobile) (err error) {
	err = global.GVA_DB.Model(&Vip.VipRecord{}).Create(vipRecord).Error
	return err
}
