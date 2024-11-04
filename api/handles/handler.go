package handlers

import controller "github.com/amiltoncabral/youFood/controllers"

type Handler struct {
	c controller.Controller
}

func New(c controller.Controller) Handler {
	return Handler{c: c}
}
