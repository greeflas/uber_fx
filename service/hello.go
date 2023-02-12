package service

import "fmt"

type Hello interface {
	CreateMessage(name string) string
}

type HelloService struct{}

func NewHelloService() *HelloService {
	return &HelloService{}
}

func (s *HelloService) CreateMessage(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

type HelloJSONDecorator struct {
	inner Hello
}

func NewHelloJSONDecorator(inner Hello) *HelloJSONDecorator {
	return &HelloJSONDecorator{inner: inner}
}

func (d *HelloJSONDecorator) CreateMessage(name string) string {
	hello := d.inner.CreateMessage(name)

	return fmt.Sprintf(`{"message": "%s"}`, hello)
}
