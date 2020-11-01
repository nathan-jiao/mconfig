package main

import (
	log "github.com/mhchlib/logger"
	"github.com/mhchlib/mconfig-api/api/v1/sdk"
	"github.com/mhchlib/mconfig/service"
	"github.com/micro/go-micro/v2"
)

func init() {
}

func main() {
	defer service.InitMconfig()()
	mService := micro.NewService()
	mService.Init()
	err := sdk.RegisterMConfigHandler(mService.Server(), &service.MConfig{})
	if err != nil {
		log.Fatal(err)
	}
	err = mService.Run()
	if err != nil {
		log.Fatal(err)
	}
}
