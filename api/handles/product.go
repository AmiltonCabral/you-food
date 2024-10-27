package handlers

import (
	"fmt"
	"net/http"
)

func (h Handler) ProductHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprint(w, "Hello from Home")
	if err != nil {
		fmt.Println(err)
	}
}
