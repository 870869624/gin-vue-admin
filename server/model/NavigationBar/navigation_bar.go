
// 自动生成模板NavigationBar
package NavigationBar
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 导航栏分类 结构体  NavigationBar
type NavigationBar struct {
    global.GVA_MODEL
    Name  *string `json:"name" form:"name" gorm:"column:name;comment:导航栏分类名称;" binding:"required"`  //导航栏分类名称 
    Picture  string `json:"picture" form:"picture" gorm:"column:picture;comment:分类图片;" binding:"required"`  //分类图片 
}


// TableName 导航栏分类 NavigationBar自定义表名 candy_navigation_bar
func (NavigationBar) TableName() string {
    return "candy_navigation_bar"
}



