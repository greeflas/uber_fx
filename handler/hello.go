package handler

import (
	"fmt"
	"github.com/greeflas/uber_fx/service"
	"go.uber.org/zap"
	"io"
	"net/http"
)

type HelloHandler struct {
	log          *zap.Logger
	helloService service.Hello
}

func NewHelloHandler(log *zap.Logger, helloService service.Hello) *HelloHandler {
	return &HelloHandler{
		log:          log,
		helloService: helloService,
	}
}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.log.Error("Failed to read request", zap.Error(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)

		return
	}

	if _, err := fmt.Fprintf(w, "%s\n", h.helloService.CreateMessage(string(body))); err != nil {
		h.log.Error("Failed to write response", zap.Error(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)

		return
	}
}

func (*HelloHandler) Pattern() string {
	return "/hello"
}
