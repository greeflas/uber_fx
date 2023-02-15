package handler

import (
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

type EchoHandler struct {
	log *logrus.Logger
}

func NewEchoHandler(log *logrus.Logger) *EchoHandler {
	return &EchoHandler{log: log}
}

func (h *EchoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, err := io.Copy(w, r.Body); err != nil {
		h.log.Errorf("Failed to handle request: %v", err)
	}
}

func (*EchoHandler) Pattern() string {
	return "/echo"
}
