
// 自动生成模板Information
package Information
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 资讯 结构体  Information
type Information struct {
    global.GVA_MODEL
    Title  *string `json:"title" form:"title" gorm:"column:title;comment:标题;"`  //标题 
    Author  *string `json:"author" form:"author" gorm:"column:author;comment:作者;"`  //作者 
    Content  *string `json:"content" form:"content" gorm:"column:content;comment:正文;type:text;"`  //正文 
    Is_Show  *bool `json:"is_Show" form:"is_Show" gorm:"default:true;column:is__show;comment:是否展示;"`  //是否展示 
    PublicChainId  *string `json:"publicChainId" form:"publicChainId" gorm:"column:public_chain_id;comment:公链id;"`  //公链id 
}


// TableName 资讯 Information自定义表名 candy_information
func (Information) TableName() string {
    return "candy_information"
}



