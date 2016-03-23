package service

import (
	"testing"
)

func TestLPOP4Redis(t *testing.T) {

	var _service SmsService
	_service.LPOP4Redis("banerwai:sms:activeuser")
}
