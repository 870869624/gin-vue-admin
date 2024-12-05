
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type VoteSearch struct{
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    VoteName  *string `json:"voteName" form:"voteName" `
    VoteIsPass  *bool `json:"voteIsPass" form:"voteIsPass" `
    VoteIsShow  *bool `json:"voteIsShow" form:"voteIsShow" `
    PublicChainId  *string `json:"publicChainId" form:"publicChainId" `
    VoteNum  *int `json:"voteNum" form:"voteNum" `
    request.PageInfo
}
