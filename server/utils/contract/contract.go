package contract

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/rpc"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
)

func GetServe(from string) (*common.Address, *utils.CdsAbi, *ethclient.Client, error) {
	rpcClient, _ := rpc.Dial("https://bsc-dataseed2.defibit.io/")
	client := ethclient.NewClient(rpcClient)
	contractAddr := common.HexToAddress("0x7045e3f0456DaAD3176e1B51cbd94e86B44CA99D")

	instance, err := utils.NewCdsAbi(contractAddr, client)
	if err != nil {
		return nil, nil, nil, err
	}

	fromAddress := common.HexToAddress(from)
	return &fromAddress, instance, client, nil
}

func GetAddress(key string) (string, error) {
	privateKeyHex := key

	// 将十六进制私钥字符串转换为私钥对象
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatalf("无法将十六进制字符串转换为私钥对象: %v", err)
	}

	// 从私钥获取公钥
	publicKey := privateKey.Public()
	publicKeyECDSA := publicKey.(*ecdsa.PublicKey)

	// 从公钥获取以太坊地址
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Printf("以太坊地址: %s\n", address.Hex())

	return address.Hex(), nil
}
