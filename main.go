package main

import (
	"demo/gosvc_test/service"
	_ "github.com/icattlecoder/godaemon"
	"github.com/judwhite/go-svc/svc"
	"log"
)

func main() {
	var clons float64
	clons = 1.0
	if clons >= 1.0 {
		log.Println(clons)
	}
	if err := svc.Run(&service.Service{}); err != nil {
		log.Println(err)
	}
}
