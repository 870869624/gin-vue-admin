package bannerPicture

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ BannerPictureRouter }

var BPApi = api.ApiGroupApp.BannerPictureApiGroup.BannerPictureApi
