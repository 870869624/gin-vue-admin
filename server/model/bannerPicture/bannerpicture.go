
// 自动生成模板BannerPicture
package bannerPicture
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 横幅图 结构体  BannerPicture
type BannerPicture struct {
    global.GVA_MODEL
    Title  *string `json:"title" form:"title" gorm:"column:title;comment:;" binding:"required"`  //标题 
    Image  string `json:"image" form:"image" gorm:"column:image;comment:;" binding:"required"`  //图片 
    URL  *string `json:"url" form:"url" gorm:"column:url;comment:;"`  //链接 
    Position  *int `json:"position" form:"position" gorm:"column:position;comment:;" binding:"required"`  //位置 
}


// TableName 横幅图 BannerPicture自定义表名 BannerPictures
func (BannerPicture) TableName() string {
    return "BannerPictures"
}



