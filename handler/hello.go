package handler

import (
	"fmt"
	"io"
	"net/http"

	"github.com/greeflas/uber_fx/service"
	"github.com/sirupsen/logrus"
)

type HelloHandler struct {
	log          *logrus.Logger
	helloService service.Hello
}

func NewHelloHandler(log *logrus.Logger, helloService service.Hello) *HelloHandler {
	return &HelloHandler{
		log:          log,
		helloService: helloService,
	}
}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.log.Errorf("Failed to read request: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)

		return
	}

	if _, err := fmt.Fprintf(w, "%s\n", h.helloService.CreateMessage(string(body))); err != nil {
		h.log.Errorf("Failed to write response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)

		return
	}
}

func (*HelloHandler) Pattern() string {
	return "/hello"
}
