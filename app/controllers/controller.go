package controllers

import "todo-golang/app/services"

type Controller struct {
	Service *services.Service
}

func NewController(service *services.Service) *Controller {
	return &Controller{
		Service: service,
	}
}
