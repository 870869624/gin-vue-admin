package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Airdrop"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Information"
	"github.com/flipped-aurora/gin-vue-admin/server/model/InvitationRecord"
	"github.com/flipped-aurora/gin-vue-admin/server/model/PlatformV1"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Presale"
	"github.com/flipped-aurora/gin-vue-admin/server/model/PublicChain"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Users"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Vip"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Vote"
	"github.com/flipped-aurora/gin-vue-admin/server/model/VoteRecord"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(PublicChain.PublicChain{}, InvitationRecord.InvitationRecord{}, Airdrop.Airdrop{}, Vote.Vote{}, VoteRecord.VoteRecord{}, Presale.Presale{}, Information.Information{}, PlatformV1.Platform{}, Vip.VipRecord{}, Users.Users{})
	if err != nil {
		return err
	}
	return nil
}
