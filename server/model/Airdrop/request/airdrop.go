
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type AirdropSearch struct{
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    AirdropName  *string `json:"airdropName" form:"airdropName" `
    AirdropValue  *float64 `json:"airdropValue" form:"airdropValue" `
    StartAirdropEndtime  *time.Time  `json:"startAirdropEndtime" form:"startAirdropEndtime"`
    EndAirdropEndtime  *time.Time  `json:"endAirdropEndtime" form:"endAirdropEndtime"`
    PublicChainId  *string `json:"publicChainId" form:"publicChainId" `
    AirdropIsPass  *bool `json:"airdropIsPass" form:"airdropIsPass" `
    AirdropIsShow  *bool `json:"airdropIsShow" form:"airdropIsShow" `
    request.PageInfo
}
