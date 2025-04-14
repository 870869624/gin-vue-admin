
package describe

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/describe"
    describeReq "github.com/flipped-aurora/gin-vue-admin/server/model/describe/request"
)

type DescribeService struct {}
// CreateDescribe 创建文档说明记录
// Author [yourname](https://github.com/yourname)
func (DService *DescribeService) CreateDescribe(D *describe.Describe) (err error) {
	err = global.GVA_DB.Create(D).Error
	return err
}

// DeleteDescribe 删除文档说明记录
// Author [yourname](https://github.com/yourname)
func (DService *DescribeService)DeleteDescribe(ID string) (err error) {
	err = global.GVA_DB.Delete(&describe.Describe{},"id = ?",ID).Error
	return err
}

// DeleteDescribeByIds 批量删除文档说明记录
// Author [yourname](https://github.com/yourname)
func (DService *DescribeService)DeleteDescribeByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]describe.Describe{},"id in ?",IDs).Error
	return err
}

// UpdateDescribe 更新文档说明记录
// Author [yourname](https://github.com/yourname)
func (DService *DescribeService)UpdateDescribe(D describe.Describe) (err error) {
	err = global.GVA_DB.Model(&describe.Describe{}).Where("id = ?",D.ID).Updates(&D).Error
	return err
}

// GetDescribe 根据ID获取文档说明记录
// Author [yourname](https://github.com/yourname)
func (DService *DescribeService)GetDescribe(ID string) (D describe.Describe, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&D).Error
	return
}

// GetDescribeInfoList 分页获取文档说明记录
// Author [yourname](https://github.com/yourname)
func (DService *DescribeService)GetDescribeInfoList(info describeReq.DescribeSearch) (list []describe.Describe, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&describe.Describe{})
    var Ds []describe.Describe
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
         	orderMap["name"] = true
         	orderMap["status"] = true
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

	err = db.Find(&Ds).Error
	return  Ds, total, err
}
func (DService *DescribeService)GetDescribePublic() {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}