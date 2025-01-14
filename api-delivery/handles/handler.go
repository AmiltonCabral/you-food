package handlers

import (
	controller "youfood-delivery/controllers"
)

type Handler struct {
	c controller.Controller
}

func New(c controller.Controller) Handler {
	return Handler{c}
}
