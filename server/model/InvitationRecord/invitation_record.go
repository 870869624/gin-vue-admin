// 自动生成模板InvitationRecord
package InvitationRecord

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 邀请记录 结构体  InvitationRecord
type InvitationRecord struct {
	global.GVA_MODEL
	User_id     *int    `json:"user_id" form:"user_id" gorm:"column:user_id;comment:用户id;" binding:"required"`                 //用户id
	Username    *string `json:"username" form:"username" gorm:"column:username;comment:邀请者用户名;" binding:"required"`            //邀请者用户名
	NewUserId   *int    `json:"newUserId" form:"newUserId" gorm:"column:new_user_id;comment:被邀请者用户id;" binding:"required"`     //被邀请者用户id
	NewUsername *string `json:"newUsername" form:"newUsername" gorm:"column:new_username;comment:被邀请者用户名;" binding:"required"` //被邀请者用户名
	InviteCode  *string `json:"inviteCode" form:"inviteCode" gorm:"column:invite_code;comment:邀请码;"`                           //邀请码
}

// TableName 邀请记录 InvitationRecord自定义表名 candy_invitation_record
func (InvitationRecord) TableName() string {
	return "candy_invitation_record"
}

func (IR *InvitationRecord) GetByCode() (err error) {
	err = global.GVA_DB.Where("invite_code = ?", *IR.InviteCode).First(&IR).Error
	return err
}
