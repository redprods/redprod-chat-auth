package app

import "github.com/redprods/redprod-chat-auth/pkg/service"

func Run() {
	service := service.NewService()
	service.Run()
}
