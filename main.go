package main

import (
	proto "MaisrForAdvancedSystems/go-biller/proto"
	srv "MaisrForAdvancedSystems/go-biller/service"
	"context"
	"log"
)

func main() {
	req := proto.BillRequest{}
	service:=srv.BillingService{}
	service.Charge(context.Background(),nil)
	log.Println(req)
}
