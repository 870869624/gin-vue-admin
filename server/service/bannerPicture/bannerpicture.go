
package bannerPicture

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/bannerPicture"
    bannerPictureReq "github.com/flipped-aurora/gin-vue-admin/server/model/bannerPicture/request"
)

type BannerPictureService struct {}
// CreateBannerPicture 创建横幅图记录
// Author [yourname](https://github.com/yourname)
func (BPService *BannerPictureService) CreateBannerPicture(BP *bannerPicture.BannerPicture) (err error) {
	err = global.GVA_DB.Create(BP).Error
	return err
}

// DeleteBannerPicture 删除横幅图记录
// Author [yourname](https://github.com/yourname)
func (BPService *BannerPictureService)DeleteBannerPicture(ID string) (err error) {
	err = global.GVA_DB.Delete(&bannerPicture.BannerPicture{},"id = ?",ID).Error
	return err
}

// DeleteBannerPictureByIds 批量删除横幅图记录
// Author [yourname](https://github.com/yourname)
func (BPService *BannerPictureService)DeleteBannerPictureByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]bannerPicture.BannerPicture{},"id in ?",IDs).Error
	return err
}

// UpdateBannerPicture 更新横幅图记录
// Author [yourname](https://github.com/yourname)
func (BPService *BannerPictureService)UpdateBannerPicture(BP bannerPicture.BannerPicture) (err error) {
	err = global.GVA_DB.Model(&bannerPicture.BannerPicture{}).Where("id = ?",BP.ID).Updates(&BP).Error
	return err
}

// GetBannerPicture 根据ID获取横幅图记录
// Author [yourname](https://github.com/yourname)
func (BPService *BannerPictureService)GetBannerPicture(ID string) (BP bannerPicture.BannerPicture, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&BP).Error
	return
}

// GetBannerPictureInfoList 分页获取横幅图记录
// Author [yourname](https://github.com/yourname)
func (BPService *BannerPictureService)GetBannerPictureInfoList(info bannerPictureReq.BannerPictureSearch) (list []bannerPicture.BannerPicture, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&bannerPicture.BannerPicture{})
    var BPs []bannerPicture.BannerPicture
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
         	orderMap["position"] = true
       if orderMap[info.Sort] {
          OrderStr = info.Sort
          if info.Order == "descending" {
             OrderStr = OrderStr + " desc"
          }
          db = db.Order(OrderStr)
       }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&BPs).Error
	return  BPs, total, err
}
func (BPService *BannerPictureService)GetBannerPicturePublic() {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}