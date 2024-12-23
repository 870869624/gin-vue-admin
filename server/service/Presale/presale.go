package Presale

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Presale"
	PresaleReq "github.com/flipped-aurora/gin-vue-admin/server/model/Presale/request"
)

type PresaleService struct{}

// CreatePresale 创建预售项目记录
// Author [yourname](https://github.com/yourname)
func (presaleService *PresaleService) CreatePresale(presale *Presale.Presale) (err error) {
	err = global.GVA_DB.Create(presale).Error
	return err
}

// DeletePresale 删除预售项目记录
// Author [yourname](https://github.com/yourname)
func (presaleService *PresaleService) DeletePresale(ID string) (err error) {
	err = global.GVA_DB.Unscoped().Delete(&Presale.Presale{}, "id = ?", ID).Error
	return err
}

// DeletePresaleByIds 批量删除预售项目记录
// Author [yourname](https://github.com/yourname)
func (presaleService *PresaleService) DeletePresaleByIds(IDs []string) (err error) {
	err = global.GVA_DB.Unscoped().Delete(&[]Presale.Presale{}, "id in ?", IDs).Error
	return err
}

// UpdatePresale 更新预售项目记录
// Author [yourname](https://github.com/yourname)
func (presaleService *PresaleService) UpdatePresale(presale Presale.Presale) (err error) {
	err = global.GVA_DB.Model(&Presale.Presale{}).Where("id = ?", presale.ID).Updates(&presale).Error
	return err
}

// GetPresale 根据ID获取预售项目记录
// Author [yourname](https://github.com/yourname)
func (presaleService *PresaleService) GetPresale(ID string) (presale Presale.Presale, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&presale).Error
	return
}

// GetPresaleInfoList 分页获取预售项目记录
// Author [yourname](https://github.com/yourname)
func (presaleService *PresaleService) GetPresaleInfoList(info PresaleReq.PresaleSearch) (list []Presale.Presale, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&Presale.Presale{})
	var presales []Presale.Presale
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.PresaleName != nil && *info.PresaleName != "" {
		db = db.Where("presale_name LIKE ?", "%"+*info.PresaleName+"%")
	}
	if info.StartPresaleStartTime != nil && info.EndPresaleStartTime != nil {
		db = db.Where("presale_start_time BETWEEN ? AND ? ", info.StartPresaleStartTime, info.EndPresaleStartTime)
	}
	if info.UserId != nil {
		db = db.Where("user_id = ?", *info.UserId)
	}
	if info.PresaleIsPass != nil {
		db = db.Where("presale_is_pass = ?", *info.PresaleIsPass)
	}
	if info.PresaleIsShow != nil {
		db = db.Where("presale_is_show = ?", *info.PresaleIsShow)
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

	err = db.Find(&presales).Error
	return presales, total, err
}
func (presaleService *PresaleService) GetPresalePublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

func (presaleService *PresaleService) MobileGetPresaleInfoList(info PresaleReq.PresaleSearch) (list []Presale.MobilePresale, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Table(Presale.Presale.TableName(Presale.Presale{}) + " p").
		Select("p.id,p.created_at,p.updated_at,p.presale_is_pass,p.presale_is_show,p.presale_name,p.presale_picture,p.presaleurl,p.public_chain_id,p.presale_start_time,p.user_id,cpc.logo,cpc.name,cpc.picture")
	var presales []Presale.MobilePresale
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("p.created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.PresaleName != nil && *info.PresaleName != "" {
		db = db.Where("p.presale_name LIKE ?", "%"+*info.PresaleName+"%")
	}
	if info.StartPresaleStartTime != nil && info.EndPresaleStartTime != nil {
		db = db.Where("p.presale_start_time BETWEEN ? AND ? ", info.StartPresaleStartTime, info.EndPresaleStartTime)
	}
	if info.UserId != nil {
		db = db.Where("p.user_id = ?", *info.UserId)
	}
	if info.PresaleIsPass != nil {
		db = db.Where("p.presale_is_pass = ?", *info.PresaleIsPass)
	}
	if info.PresaleIsShow != nil {
		db = db.Where("p.presale_is_show = ?", *info.PresaleIsShow)
	}
	if info.PublicChainId != nil && *info.PublicChainId != "" {
		db = db.Where("p.public_chain_id = ?", *info.PublicChainId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}
	db.Joins("left join candies_public_chain cpc on p.public_chain_id = cpc.id")
	err = db.Find(&presales).Error
	return presales, total, err
}

func (presaleService *PresaleService) CreatePresaleMobile(presale1 *Presale.MobilePresaleCre) (err error) {
	presale := &Presale.Presale{}
	presale.ToArgs(presale1)
	err = global.GVA_DB.Where(presale.TableName()).Create(presale).Error
	return err
}
