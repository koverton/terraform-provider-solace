package main

import (
	"github.com/koverton/semp_client"
	"log"
)

func logSempMeta(op string, meta semp_client.SempMeta) {
	log.Printf("ENTITY:%s; Response Code:%d, ErrorCode:%d, Status:%s, Description:%s\n" ,
		op,
		meta.ResponseCode,
		meta.Error_.Code,
		meta.Error_.Status,
		meta.Error_.Description)
}
