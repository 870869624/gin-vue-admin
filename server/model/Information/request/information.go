
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type InformationSearch struct{
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    Title  *string `json:"title" form:"title" `
    Author  *string `json:"author" form:"author" `
    Is_Show  *bool `json:"is_Show" form:"is_Show" `
    PublicChainId  *string `json:"publicChainId" form:"publicChainId" `
    request.PageInfo
}
