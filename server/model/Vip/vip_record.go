// 自动生成模板VipRecord
package Vip

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 会员记录 结构体  VipRecord
type VipRecord struct {
	global.GVA_MODEL
	UserId      *int    `json:"userId" form:"userId" gorm:"column:user_id;comment:充值者ID;" binding:"required"`                             //充值者ID
	IsEffective *bool   `json:"isEffective" form:"isEffective" gorm:"default:false;column:is_effective;comment:是否有效;" binding:"required"` //是否有效
	BlockHash   *string `json:"blockHash" form:"blockHash" gorm:"column:block_hash;comment:交易hash;" binding:"required"`                   //交易hash
	BlockNumber *string `json:"blockNumber" form:"blockNumber" gorm:"column:block_number;comment:交易网络数;" binding:"required"`              //交易网络数
}

type VipRecordMobile struct {
	global.GVA_MODEL
	UserId      *int    `json:"userId" form:"userId" gorm:"column:user_id;comment:充值者ID;" binding:"-"`                             //充值者ID
	IsEffective *bool   `json:"isEffective" form:"isEffective" gorm:"default:false;column:is_effective;comment:是否有效;" binding:"-"` //是否有效
	BlockHash   *string `json:"blockHash" form:"blockHash" gorm:"column:block_hash;comment:交易hash;" binding:"required"`            //交易hash
	BlockNumber *string `json:"blockNumber" form:"blockNumber" gorm:"column:block_number;comment:交易网络数;" binding:"required"`       //交易网络数
}

// TableName 会员记录 VipRecord自定义表名 candy_vip_record
func (VipRecord) TableName() string {
	return "candy_vip_record"
}

func (v *VipRecord) GetUserId() error {
	return global.GVA_DB.Table(v.TableName()).Where("user_id = ?", v.UserId).First(&v).Error
}
