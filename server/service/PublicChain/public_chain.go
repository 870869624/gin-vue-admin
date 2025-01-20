package PublicChain

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/PublicChain"
	PublicChainReq "github.com/flipped-aurora/gin-vue-admin/server/model/PublicChain/request"
)

type PublicChainService struct{}

// CreatePublicChain 创建公链记录
// Author [yourname](https://github.com/yourname)
func (PCService *PublicChainService) CreatePublicChain(PC *PublicChain.PublicChain) (err error) {
	err = global.GVA_DB.Create(PC).Error
	return err
}

// DeletePublicChain 删除公链记录
// Author [yourname](https://github.com/yourname)
func (PCService *PublicChainService) DeletePublicChain(ID string) (err error) {
	err = global.GVA_DB.Delete(&PublicChain.PublicChain{}, "id = ?", ID).Error
	return err
}

// DeletePublicChainByIds 批量删除公链记录
// Author [yourname](https://github.com/yourname)
func (PCService *PublicChainService) DeletePublicChainByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]PublicChain.PublicChain{}, "id in ?", IDs).Error
	return err
}

// UpdatePublicChain 更新公链记录
// Author [yourname](https://github.com/yourname)
func (PCService *PublicChainService) UpdatePublicChain(PC PublicChain.PublicChain) (err error) {
	err = global.GVA_DB.Model(&PublicChain.PublicChain{}).Where("id = ?", PC.ID).Updates(&PC).Error
	return err
}

// GetPublicChain 根据ID获取公链记录
// Author [yourname](https://github.com/yourname)
func (PCService *PublicChainService) GetPublicChain(ID string) (PC PublicChain.PublicChain, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&PC).Error
	return
}

// GetPublicChainInfoList 分页获取公链记录
// Author [yourname](https://github.com/yourname)
func (PCService *PublicChainService) GetPublicChainInfoList(info PublicChainReq.PublicChainSearch) (list []PublicChain.PublicChain, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&PublicChain.PublicChain{})
	var PCs []PublicChain.PublicChain
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Name != nil && *info.Name != "" {
		db = db.Where("name = ?", *info.Name)
	}
	if info.Is_Show != nil {
		db = db.Where("is__show = ?", *info.Is_Show)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&PCs).Error
	return PCs, total, err
}
func (PCService *PublicChainService) GetPublicChainPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

func (PCService *PublicChainService) GetPublicChainInfoListMobile(info PublicChainReq.PublicChainSearch) (list []PublicChain.PublicChain, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&PublicChain.PublicChain{})
	var PCs []PublicChain.PublicChain
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Name != nil && *info.Name != "" {
		db = db.Where("name = ?", *info.Name)
	}
	if info.Is_Show != nil {
		db = db.Where("is__show = ?", *info.Is_Show)
	}
	db.Where("`deleted_at` IS NULL and is__show = ?", true)

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&PCs).Error
	return PCs, total, err
}
