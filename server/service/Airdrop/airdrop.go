package Airdrop

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Airdrop"
	AirdropReq "github.com/flipped-aurora/gin-vue-admin/server/model/Airdrop/request"
)

type AirdropService struct{}

// CreateAirdrop 创建空投项目记录
// Author [yourname](https://github.com/yourname)
func (airdropService *AirdropService) CreateAirdrop(airdrop *Airdrop.Airdrop) (err error) {
	err = global.GVA_DB.Create(airdrop).Error
	return err
}

// DeleteAirdrop 删除空投项目记录
// Author [yourname](https://github.com/yourname)
func (airdropService *AirdropService) DeleteAirdrop(ID string) (err error) {
	err = global.GVA_DB.Delete(&Airdrop.Airdrop{}, "id = ?", ID).Error
	return err
}

// DeleteAirdropByIds 批量删除空投项目记录
// Author [yourname](https://github.com/yourname)
func (airdropService *AirdropService) DeleteAirdropByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]Airdrop.Airdrop{}, "id in ?", IDs).Error
	return err
}

// UpdateAirdrop 更新空投项目记录
// Author [yourname](https://github.com/yourname)
func (airdropService *AirdropService) UpdateAirdrop(airdrop Airdrop.Airdrop) (err error) {
	err = global.GVA_DB.Model(&Airdrop.Airdrop{}).Where("id = ?", airdrop.ID).Updates(&airdrop).Error
	return err
}

// GetAirdrop 根据ID获取空投项目记录
// Author [yourname](https://github.com/yourname)
func (airdropService *AirdropService) GetAirdrop(ID string) (airdrop Airdrop.Airdrop, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&airdrop).Error
	return
}

// GetAirdropInfoList 分页获取空投项目记录
// Author [yourname](https://github.com/yourname)
func (airdropService *AirdropService) GetAirdropInfoList(info AirdropReq.AirdropSearch) (list []Airdrop.Airdrop, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&Airdrop.Airdrop{})
	var airdrops []Airdrop.Airdrop
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.AirdropName != nil && *info.AirdropName != "" {
		db = db.Where("airdrop_name LIKE ?", "%"+*info.AirdropName+"%")
	}
	if info.AirdropValue != nil {
		db = db.Where("airdrop_value <> ?", *info.AirdropValue)
	}
	if info.StartAirdropEndtime != nil && info.EndAirdropEndtime != nil {
		db = db.Where("airdrop_endtime BETWEEN ? AND ? ", info.StartAirdropEndtime, info.EndAirdropEndtime)
	}
	if info.PublicChainId != nil && *info.PublicChainId != "" {
		db = db.Where("public_chain_id = ?", *info.PublicChainId)
	}
	if info.AirdropIsPass != nil {
		db = db.Where("airdrop_is_pass = ?", *info.AirdropIsPass)
	}
	if info.AirdropIsShow != nil {
		db = db.Where("airdrop_is_show = ?", *info.AirdropIsShow)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&airdrops).Error
	return airdrops, total, err
}
func (airdropService *AirdropService) GetAirdropPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

func (airdropService *AirdropService) MobileGetAirdropInfoList(info AirdropReq.AirdropSearch) (list []Airdrop.MobileAirdropResp, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Table(Airdrop.Airdrop.TableName(Airdrop.Airdrop{}) + " a").
		Select("a.id,a.created_at,a.updated_at,a.airdrop_endtime,a.airdrop_is_pass,a.airdrop_is_show,a.airdrop_name,a.airdrop_picture,a.airdrop_url,a.airdrop_value,cpc.logo,cpc.name,cpc.picture,a.public_chain_id,a.brief,a.detail,a.x_link,a.tg_link,a.discord_link,a.distribution_date")
	var airdrops []Airdrop.MobileAirdropResp
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("a.created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.AirdropName != nil && *info.AirdropName != "" {
		db = db.Where("a.airdrop_name LIKE ?", "%"+*info.AirdropName+"%")
	}
	if info.AirdropValue != nil {
		db = db.Where("a.airdrop_value <> ?", *info.AirdropValue)
	}
	if info.StartAirdropEndtime != nil && info.EndAirdropEndtime != nil {
		db = db.Where("a.airdrop_endtime BETWEEN ? AND ? ", info.StartAirdropEndtime, info.EndAirdropEndtime)
	}
	if info.PublicChainId != nil && *info.PublicChainId != "" {
		db = db.Where("a.public_chain_id = ?", *info.PublicChainId)
	}
	if info.AirdropIsPass != nil {
		db = db.Where("a.airdrop_is_pass = ?", *info.AirdropIsPass)
	}
	if info.AirdropIsShow != nil {
		db = db.Where("a.airdrop_is_show = ?", *info.AirdropIsShow)
	}
	db.Joins("LEFT JOIN candies_public_chain cpc ON a.public_chain_id = cpc.id")
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&airdrops).Error
	return airdrops, total, err
}

func (airdropService *AirdropService) CreateAirdropMobile(airdrop1 *Airdrop.MobileAirdropCre) (err error) {
	airdrop := &Airdrop.Airdrop{}
	airdrop.ToArgs(airdrop1)
	err = global.GVA_DB.Where(airdrop.TableName()).Create(airdrop).Error
	return err
}

func (airdropService *AirdropService) GetAirdropMobile(ID string) (airdrop Airdrop.MobileAirdropResp, err error) {

	db := global.GVA_DB.Table(Airdrop.Airdrop.TableName(Airdrop.Airdrop{}) + " a").
		Select("a.id,a.created_at,a.updated_at,a.airdrop_endtime,a.airdrop_is_pass,a.airdrop_is_show,a.airdrop_name,a.airdrop_picture,a.airdrop_url,a.airdrop_value,cpc.logo,cpc.name,cpc.picture,a.public_chain_id,a.brief,a.detail,a.x_link,a.tg_link,a.discord_link,a.distribution_date")
	var airdrops Airdrop.MobileAirdropResp

	db.Joins("LEFT JOIN candies_public_chain cpc ON a.public_chain_id = cpc.id").Where("a.id = ?", ID)

	err = db.First(&airdrops).Error
	return airdrops, err
}
