package service

import ()

// RenderService is the abstract representation of this service.
type RenderService interface {
	RenderHello(tmpl, name string) string
}
