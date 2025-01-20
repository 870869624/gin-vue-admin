
package NavigationBar

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/NavigationBar"
    NavigationBarReq "github.com/flipped-aurora/gin-vue-admin/server/model/NavigationBar/request"
)

type NavigationBarService struct {}
// CreateNavigationBar 创建导航栏分类记录
// Author [yourname](https://github.com/yourname)
func (NAService *NavigationBarService) CreateNavigationBar(NA *NavigationBar.NavigationBar) (err error) {
	err = global.GVA_DB.Create(NA).Error
	return err
}

// DeleteNavigationBar 删除导航栏分类记录
// Author [yourname](https://github.com/yourname)
func (NAService *NavigationBarService)DeleteNavigationBar(ID string) (err error) {
	err = global.GVA_DB.Delete(&NavigationBar.NavigationBar{},"id = ?",ID).Error
	return err
}

// DeleteNavigationBarByIds 批量删除导航栏分类记录
// Author [yourname](https://github.com/yourname)
func (NAService *NavigationBarService)DeleteNavigationBarByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]NavigationBar.NavigationBar{},"id in ?",IDs).Error
	return err
}

// UpdateNavigationBar 更新导航栏分类记录
// Author [yourname](https://github.com/yourname)
func (NAService *NavigationBarService)UpdateNavigationBar(NA NavigationBar.NavigationBar) (err error) {
	err = global.GVA_DB.Model(&NavigationBar.NavigationBar{}).Where("id = ?",NA.ID).Updates(&NA).Error
	return err
}

// GetNavigationBar 根据ID获取导航栏分类记录
// Author [yourname](https://github.com/yourname)
func (NAService *NavigationBarService)GetNavigationBar(ID string) (NA NavigationBar.NavigationBar, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&NA).Error
	return
}

// GetNavigationBarInfoList 分页获取导航栏分类记录
// Author [yourname](https://github.com/yourname)
func (NAService *NavigationBarService)GetNavigationBarInfoList(info NavigationBarReq.NavigationBarSearch) (list []NavigationBar.NavigationBar, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&NavigationBar.NavigationBar{})
    var NAs []NavigationBar.NavigationBar
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&NAs).Error
	return  NAs, total, err
}
func (NAService *NavigationBarService)GetNavigationBarPublic() {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}