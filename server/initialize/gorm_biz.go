package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Airdrop"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Information"
	"github.com/flipped-aurora/gin-vue-admin/server/model/InvitationRecord"
	"github.com/flipped-aurora/gin-vue-admin/server/model/PublicChain"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Users"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(PublicChain.PublicChain{}, Users.Users{}, InvitationRecord.InvitationRecord{}, Information.Information{}, Airdrop.Airdrop{})
	if err != nil {
		return err
	}
	return nil
}
