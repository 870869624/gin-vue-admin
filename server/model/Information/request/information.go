package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type InformationSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	Title          *string    `json:"title" form:"title" `
	Author         *string    `json:"author" form:"author" `
	Is_Show        *bool      `json:"is_Show" form:"is_Show" `
	PublicChainId  *string    `json:"publicChainId" form:"publicChainId" gorm:"column:public_chain_id;comment:公链id;"` //公链id
	request.PageInfo
}
