package Airdrop

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ AirdropApi }

var airdropService = service.ServiceGroupApp.AirdropServiceGroup.AirdropService
