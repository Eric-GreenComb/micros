package service

import ()

// RenderService is the abstract representation of this service.
type RenderService interface {
	Ping() string
	RenderHello(tmpl, name string) string
}
