
package InvitationRecord

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/InvitationRecord"
    InvitationRecordReq "github.com/flipped-aurora/gin-vue-admin/server/model/InvitationRecord/request"
)

type InvitationRecordService struct {}
// CreateInvitationRecord 创建邀请记录记录
// Author [yourname](https://github.com/yourname)
func (IRService *InvitationRecordService) CreateInvitationRecord(IR *InvitationRecord.InvitationRecord) (err error) {
	err = global.GVA_DB.Create(IR).Error
	return err
}

// DeleteInvitationRecord 删除邀请记录记录
// Author [yourname](https://github.com/yourname)
func (IRService *InvitationRecordService)DeleteInvitationRecord(ID string) (err error) {
	err = global.GVA_DB.Delete(&InvitationRecord.InvitationRecord{},"id = ?",ID).Error
	return err
}

// DeleteInvitationRecordByIds 批量删除邀请记录记录
// Author [yourname](https://github.com/yourname)
func (IRService *InvitationRecordService)DeleteInvitationRecordByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]InvitationRecord.InvitationRecord{},"id in ?",IDs).Error
	return err
}

// UpdateInvitationRecord 更新邀请记录记录
// Author [yourname](https://github.com/yourname)
func (IRService *InvitationRecordService)UpdateInvitationRecord(IR InvitationRecord.InvitationRecord) (err error) {
	err = global.GVA_DB.Model(&InvitationRecord.InvitationRecord{}).Where("id = ?",IR.ID).Updates(&IR).Error
	return err
}

// GetInvitationRecord 根据ID获取邀请记录记录
// Author [yourname](https://github.com/yourname)
func (IRService *InvitationRecordService)GetInvitationRecord(ID string) (IR InvitationRecord.InvitationRecord, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&IR).Error
	return
}

// GetInvitationRecordInfoList 分页获取邀请记录记录
// Author [yourname](https://github.com/yourname)
func (IRService *InvitationRecordService)GetInvitationRecordInfoList(info InvitationRecordReq.InvitationRecordSearch) (list []InvitationRecord.InvitationRecord, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&InvitationRecord.InvitationRecord{})
    var IRs []InvitationRecord.InvitationRecord
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.User_id != nil {
        db = db.Where("user_id = ?",*info.User_id)
    }
    if info.Username != nil && *info.Username != "" {
        db = db.Where("username = ?",*info.Username)
    }
    if info.NewUserId != nil {
        db = db.Where("new_user_id = ?",*info.NewUserId)
    }
    if info.NewUsername != nil && *info.NewUsername != "" {
        db = db.Where("new_username = ?",*info.NewUsername)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&IRs).Error
	return  IRs, total, err
}
func (IRService *InvitationRecordService)GetInvitationRecordPublic() {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}