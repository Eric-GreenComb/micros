package service

import ()

// RenderService is the abstract representation of this service.
type RenderService interface {
	Ping() string
	// Parameters:
	//  - Tplname
	//  - KeyMmap
	RenderTpl(tplname string, key_mmap map[string]string) string
}
