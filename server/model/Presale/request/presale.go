
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type PresaleSearch struct{
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    PresaleName  *string `json:"presaleName" form:"presaleName" `
    StartPresaleStartTime  *time.Time  `json:"startPresaleStartTime" form:"startPresaleStartTime"`
    EndPresaleStartTime  *time.Time  `json:"endPresaleStartTime" form:"endPresaleStartTime"`
    UserId  *int `json:"userId" form:"userId" `
    PresaleIsPass  *bool `json:"presaleIsPass" form:"presaleIsPass" `
    PresaleIsShow  *bool `json:"presaleIsShow" form:"presaleIsShow" `
    PublicChainId  *string `json:"publicChainId" form:"publicChainId" `
    request.PageInfo
}
