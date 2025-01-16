package handlers

import (
	controller "github.com/amiltoncabral/youFood/controllers"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
)

type Handler struct {
	c          controller.Controller
	instrument Instrument
}

type Instrument struct {
	orderCounter metric.Int64Counter
}

func New(c controller.Controller) Handler {
	return Handler{c, *getInstrument()}
}

func getInstrument() *Instrument {
	meter := otel.Meter("")

	orderCounter, err := meter.Int64Counter(
		"order.counter",
		metric.WithDescription("The number of orders created"),
		metric.WithUnit("{order}"),
	)

	if err != nil {
		panic(err)
	}

	return &Instrument{orderCounter}
}
