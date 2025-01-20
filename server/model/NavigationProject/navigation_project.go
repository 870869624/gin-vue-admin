
// 自动生成模板NavigationProject
package NavigationProject
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 导航栏项目 结构体  NavigationProject
type NavigationProject struct {
    global.GVA_MODEL
    Name  *string `json:"name" form:"name" gorm:"column:name;comment:项目名称;" binding:"required"`  //项目名称 
    Picture  string `json:"picture" form:"picture" gorm:"column:picture;comment:项目图片;" binding:"required"`  //项目图片 
    Url  *string `json:"url" form:"url" gorm:"column:url;comment:项目链接;" binding:"required"`  //项目链接 
    Bar  *string `json:"bar" form:"bar" gorm:"column:bar;comment:项目栏分类;" binding:"required"`  //项目栏分类 
    Brief  *string `json:"brief" form:"brief" gorm:"column:brief;comment:;" binding:"required"`  //描述 
}


// TableName 导航栏项目 NavigationProject自定义表名 candy_navigation_project
func (NavigationProject) TableName() string {
    return "candy_navigation_project"
}



