package PlatformV1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/PlatformV1"
	PlatformV1Req "github.com/flipped-aurora/gin-vue-admin/server/model/PlatformV1/request"
)

type PlatformService struct{}

// CreatePlatform 创建平台记录
// Author [yourname](https://github.com/yourname)
func (platformService *PlatformService) CreatePlatform(platform *PlatformV1.Platform) (err error) {
	err = global.GVA_DB.Create(platform).Error
	return err
}

// DeletePlatform 删除平台记录
// Author [yourname](https://github.com/yourname)
func (platformService *PlatformService) DeletePlatform(ID string) (err error) {
	err = global.GVA_DB.Delete(&PlatformV1.Platform{}, "id = ?", ID).Error
	return err
}

// DeletePlatformByIds 批量删除平台记录
// Author [yourname](https://github.com/yourname)
func (platformService *PlatformService) DeletePlatformByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]PlatformV1.Platform{}, "id in ?", IDs).Error
	return err
}

// UpdatePlatform 更新平台记录
// Author [yourname](https://github.com/yourname)
func (platformService *PlatformService) UpdatePlatform(platform PlatformV1.Platform) (err error) {
	err = global.GVA_DB.Model(&PlatformV1.Platform{}).Where("id = ?", platform.ID).Updates(&platform).Error
	return err
}

// GetPlatform 根据ID获取平台记录
// Author [yourname](https://github.com/yourname)
func (platformService *PlatformService) GetPlatform(ID string) (platform PlatformV1.Platform, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&platform).Error
	return
}

// GetPlatformInfoList 分页获取平台记录
// Author [yourname](https://github.com/yourname)
func (platformService *PlatformService) GetPlatformInfoList(info PlatformV1Req.PlatformSearch) (list []PlatformV1.Platform, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&PlatformV1.Platform{})
	var platforms []PlatformV1.Platform
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.PlatformName != nil && *info.PlatformName != "" {
		db = db.Where("platform_name LIKE ?", "%"+*info.PlatformName+"%")
	}
	if info.PlatformIsShow != nil {
		db = db.Where("platform_is_show = ?", *info.PlatformIsShow)
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

	err = db.Find(&platforms).Error
	return platforms, total, err
}
func (platformService *PlatformService) GetPlatformPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

func (platformService *PlatformService) GetPlatformInfoListMobile(info PlatformV1Req.PlatformSearch) (list []PlatformV1.Platform, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&PlatformV1.Platform{})
	var platforms []PlatformV1.Platform
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.PlatformName != nil && *info.PlatformName != "" {
		db = db.Where("platform_name LIKE ?", "%"+*info.PlatformName+"%")
	}
	if info.PlatformIsShow != nil {
		db = db.Where("platform_is_show = ?", *info.PlatformIsShow)
	}
	if info.PublicChainId != nil && *info.PublicChainId != "" {
		db = db.Where("public_chain_id = ?", *info.PublicChainId)
	}
	db.Where("`deleted_at` IS NULL")
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&platforms).Error
	return platforms, total, err
}
