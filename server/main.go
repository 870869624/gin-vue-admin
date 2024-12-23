package main

import (
	"github.com/flipped-aurora/gin-vue-admin/server/core"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize"
	_ "go.uber.org/automaxprocs"
	"go.uber.org/zap"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title                       Gin-Vue-Admin Swagger API接口文档
// @version                     v2.7.7
// @description                 使用gin+vue进行极速开发的全栈开发基础平台
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        x-token
// @BasePath                    /
func main() {
	global.GVA_VP = core.Viper() // 初始化Viper
	initialize.OtherInit()
	global.GVA_LOG = core.Zap() // 初始化zap日志库
	zap.ReplaceGlobals(global.GVA_LOG)
	global.GVA_DB = initialize.Gorm() // gorm连接数据库
	initialize.Timer()
	initialize.DBList()
	if global.GVA_DB != nil {
		initialize.RegisterTables() // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}
	core.RunWindowsServer()
	// v := "https://bsc-dataseed2.defibit.io/"
	// rpcClient, err := rpc.Dial(v)
	// if err != nil {
	// 	fmt.Println(err, "地址", v)
	// 	return
	// }
	// client := ethclient.NewClient(rpcClient)
	// contractAddr := common.HexToAddress("0x7045e3f0456DaAD3176e1B51cbd94e86B44CA99D")

	// instance, err := utils.NewCdsAbi(contractAddr, client)
	// if err != nil {
	// 	fmt.Println(err, "地址", v)
	// 	return
	// }

	// fromAddress := common.HexToAddress("0x2eEdBAE4373B976c47eB41Eb0c1A9f505405Ee3c")

	// balance, err := instance.BalanceOf(nil, fromAddress)
	// if err != nil {
	// 	fmt.Println(err, "获取金额错误")
	// 	return
	// }
	// fmt.Println(balance.String(), "余额")

	// toAddress := common.HexToAddress("0x04bCBCe3cE594915d6D57BB6F892b0a58726c904")
	// opts := &bind.TransactOpts{
	// 	From: fromAddress,
	// }
	// instance.Approve(opts, fromAddress, big.NewInt(1000000000000000000))
	// balance1, err := instance.TransferFrom(opts, fromAddress, toAddress, big.NewInt(1))
	// if err != nil {
	// 	fmt.Println(err, "支付错误")
	// 	return
	// }
	// fmt.Println(balance1.Hash())

}
