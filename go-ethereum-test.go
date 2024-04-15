package main

import (
	"fmt"
	controller "go-ethereum-test/Controller"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	fmt.Println("Hello, world.")
	client, err := ethclient.Dial("wss://wss.calibration.node.glif.io/apigw/lotus/rpc/v1")
	if err != nil {
		fmt.Println("error : ")
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", client)
	controllerAddr := common.HexToAddress("0xD2bD8A9c17298835416d3d9163c4599d7AD83E2e")
	controllerInstance, err := controller.NewController(controllerAddr, client)
	if err != nil {
		fmt.Println("error : ")
		fmt.Println(err)
	}

	owner, err := controllerInstance.Owner(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(owner)

}
