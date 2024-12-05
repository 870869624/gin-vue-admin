
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type PlatformSearch struct{
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    PlatformName  *string `json:"platformName" form:"platformName" `
    PlatformIsShow  *bool `json:"platformIsShow" form:"platformIsShow" `
    PublicChainId  *string `json:"publicChainId" form:"publicChainId" `
    request.PageInfo
}
