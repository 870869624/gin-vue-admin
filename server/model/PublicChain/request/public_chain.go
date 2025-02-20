
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type PublicChainSearch struct{
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    Name  *string `json:"name" form:"name" `
    Is_Show  *bool `json:"is_Show" form:"is_Show" `
    request.PageInfo
}
