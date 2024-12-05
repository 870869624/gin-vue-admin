
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type VoteRecordSearch struct{
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    VoteId  *int `json:"voteId" form:"voteId" `
    UserId  *int `json:"userId" form:"userId" `
    request.PageInfo
}
