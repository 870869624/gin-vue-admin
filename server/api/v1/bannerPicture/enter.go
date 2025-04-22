package bannerPicture

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ BannerPictureApi }

var BPService = service.ServiceGroupApp.BannerPictureServiceGroup.BannerPictureService
