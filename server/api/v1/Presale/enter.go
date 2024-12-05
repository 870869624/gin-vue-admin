package Presale

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ PresaleApi }

var presaleService = service.ServiceGroupApp.PresaleServiceGroup.PresaleService
