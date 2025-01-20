
package NavigationProject

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/NavigationProject"
    NavigationProjectReq "github.com/flipped-aurora/gin-vue-admin/server/model/NavigationProject/request"
)

type NavigationProjectService struct {}
// CreateNavigationProject 创建导航栏项目记录
// Author [yourname](https://github.com/yourname)
func (NPService *NavigationProjectService) CreateNavigationProject(NP *NavigationProject.NavigationProject) (err error) {
	err = global.GVA_DB.Create(NP).Error
	return err
}

// DeleteNavigationProject 删除导航栏项目记录
// Author [yourname](https://github.com/yourname)
func (NPService *NavigationProjectService)DeleteNavigationProject(ID string) (err error) {
	err = global.GVA_DB.Delete(&NavigationProject.NavigationProject{},"id = ?",ID).Error
	return err
}

// DeleteNavigationProjectByIds 批量删除导航栏项目记录
// Author [yourname](https://github.com/yourname)
func (NPService *NavigationProjectService)DeleteNavigationProjectByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]NavigationProject.NavigationProject{},"id in ?",IDs).Error
	return err
}

// UpdateNavigationProject 更新导航栏项目记录
// Author [yourname](https://github.com/yourname)
func (NPService *NavigationProjectService)UpdateNavigationProject(NP NavigationProject.NavigationProject) (err error) {
	err = global.GVA_DB.Model(&NavigationProject.NavigationProject{}).Where("id = ?",NP.ID).Updates(&NP).Error
	return err
}

// GetNavigationProject 根据ID获取导航栏项目记录
// Author [yourname](https://github.com/yourname)
func (NPService *NavigationProjectService)GetNavigationProject(ID string) (NP NavigationProject.NavigationProject, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&NP).Error
	return
}

// GetNavigationProjectInfoList 分页获取导航栏项目记录
// Author [yourname](https://github.com/yourname)
func (NPService *NavigationProjectService)GetNavigationProjectInfoList(info NavigationProjectReq.NavigationProjectSearch) (list []NavigationProject.NavigationProject, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&NavigationProject.NavigationProject{})
    var NPs []NavigationProject.NavigationProject
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Name != nil && *info.Name != "" {
        db = db.Where("name = ?",*info.Name)
    }
    if info.Bar != nil && *info.Bar != "" {
        db = db.Where("bar = ?",*info.Bar)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&NPs).Error
	return  NPs, total, err
}
func (NPService *NavigationProjectService)GetNavigationProjectPublic() {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}