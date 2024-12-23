package Information

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Information"
	InformationReq "github.com/flipped-aurora/gin-vue-admin/server/model/Information/request"
)

type InformationService struct{}

// CreateInformation 创建资讯记录
// Author [yourname](https://github.com/yourname)
func (informationService *InformationService) CreateInformation(information *Information.Information) (err error) {
	err = global.GVA_DB.Create(information).Error
	return err
}

// DeleteInformation 删除资讯记录
// Author [yourname](https://github.com/yourname)
func (informationService *InformationService) DeleteInformation(ID string) (err error) {
	err = global.GVA_DB.Unscoped().Delete(&Information.Information{}, "id = ?", ID).Error
	return err
}

// DeleteInformationByIds 批量删除资讯记录
// Author [yourname](https://github.com/yourname)
func (informationService *InformationService) DeleteInformationByIds(IDs []string) (err error) {
	err = global.GVA_DB.Unscoped().Delete(&[]Information.Information{}, "id in ?", IDs).Error
	return err
}

// UpdateInformation 更新资讯记录
// Author [yourname](https://github.com/yourname)
func (informationService *InformationService) UpdateInformation(information Information.Information) (err error) {
	err = global.GVA_DB.Model(&Information.Information{}).Where("id = ?", information.ID).Updates(&information).Error
	return err
}

// GetInformation 根据ID获取资讯记录
// Author [yourname](https://github.com/yourname)
func (informationService *InformationService) GetInformation(ID string) (information Information.Information, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&information).Error
	return
}

// GetInformationInfoList 分页获取资讯记录
// Author [yourname](https://github.com/yourname)
func (informationService *InformationService) GetInformationInfoList(info InformationReq.InformationSearch) (list []Information.Information, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&Information.Information{})
	var informations []Information.Information
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Title != nil && *info.Title != "" {
		db = db.Where("title LIKE ?", "%"+*info.Title+"%")
	}
	if info.Author != nil && *info.Author != "" {
		db = db.Where("author = ?", *info.Author)
	}
	if info.Is_Show != nil {
		db = db.Where("is__show = ?", *info.Is_Show)
	}
	if info.PublicChainId != nil && *info.PublicChainId != "" {
		db = db.Where("public_chain_id = ?", *info.PublicChainId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&informations).Error
	return informations, total, err
}
func (informationService *InformationService) GetInformationPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

func (informationService *InformationService) GetInformationMobile(ID string) (information Information.InformationOutput, err error) {
	err = global.GVA_DB.Model(&Information.Information{}).Where("id = ?", ID).First(&information).Error
	return
}

func (informationService *InformationService) GetInformationInfoListMobile(info InformationReq.InformationSearch) (list []Information.InformationOutput, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&Information.Information{})
	var informations []Information.InformationOutput
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Title != nil && *info.Title != "" {
		db = db.Where("title LIKE ?", "%"+*info.Title+"%")
	}
	if info.Author != nil && *info.Author != "" {
		db = db.Where("author = ?", *info.Author)
	}
	if info.Is_Show != nil {
		db = db.Where("is__show = ?", *info.Is_Show)
	}
	if info.PublicChainId != nil && *info.PublicChainId != "" {
		db = db.Where("public_chain_id = ?", *info.PublicChainId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}
	err = db.Find(&informations).Error
	return informations, total, err
}
