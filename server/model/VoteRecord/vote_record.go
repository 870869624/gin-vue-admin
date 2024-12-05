
// 自动生成模板VoteRecord
package VoteRecord
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 投票记录 结构体  VoteRecord
type VoteRecord struct {
    global.GVA_MODEL
    VoteId  *int `json:"voteId" form:"voteId" gorm:"column:vote_id;comment:投票项目ID;" binding:"required"`  //投票项目ID 
    UserId  *int `json:"userId" form:"userId" gorm:"column:user_id;comment:投票者ID;" binding:"required"`  //投票者ID 
}


// TableName 投票记录 VoteRecord自定义表名 candy_vote_record
func (VoteRecord) TableName() string {
    return "candy_vote_record"
}



