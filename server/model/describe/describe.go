
// 自动生成模板Describe
package describe
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 文档说明 结构体  Describe
type Describe struct {
    global.GVA_MODEL
    Name  *string `json:"name" form:"name" gorm:"column:name;comment:标题名称;" binding:"required"`  //标题名称 
    Content  *string `json:"content" form:"content" gorm:"column:content;comment:详细内容;type:text;" binding:"required"`  //详细内容 
    Status  *bool `json:"status" form:"status" gorm:"column:status;comment:是否展示;"`  //状态 
}


// TableName 文档说明 Describe自定义表名 candy_describe
func (Describe) TableName() string {
    return "candy_describe"
}



