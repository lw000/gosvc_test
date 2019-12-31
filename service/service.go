package service

import (
	"demo/gosvc_test/service/service1"
	"demo/gosvc_test/service/service2"
	"github.com/judwhite/go-svc/svc"
	"log"
)

type Service struct {
	serv1 service1.TimerService
	serv2 service2.TimerService
}

func (u *Service) Init(env svc.Environment) (err error) {
	if env.IsWindowsService() {

	} else {

	}

	return nil
}

func (u *Service) Start() error {
	u.serv1.Start()
	u.serv2.Start()
	return nil
}

func (u *Service) Stop() error {
	var err error
	err = u.serv1.Stop()
	if err != nil {
		log.Println(err)
	}

	err = u.serv2.Stop()
	if err != nil {
		log.Println(err)
	}

	log.Println("SERVICE·EXIT")

	return nil
}
