
// 自动生成模板Platform
package PlatformV1
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 平台 结构体  Platform
type Platform struct {
    global.GVA_MODEL
    PlatformName  *string `json:"platformName" form:"platformName" gorm:"column:platform_name;comment:平台名称;" binding:"required"`  //平台名称 
    PlatformPicture  string `json:"platformPicture" form:"platformPicture" gorm:"column:platform_picture;comment:平台图片;" binding:"required"`  //平台图片 
    PlatformIsShow  *bool `json:"platformIsShow" form:"platformIsShow" gorm:"default:true;column:platform_is_show;comment:平台是否展示;"`  //平台是否展示 
    PublicChainId  *string `json:"publicChainId" form:"publicChainId" gorm:"column:public_chain_id;comment:公链ID;" binding:"required"`  //公链ID 
    PlatformUrl  *string `json:"platformUrl" form:"platformUrl" gorm:"column:platform_url;comment:平台链接;" binding:"required"`  //平台链接 
}


// TableName 平台 Platform自定义表名 candy_platform
func (Platform) TableName() string {
    return "candy_platform"
}



