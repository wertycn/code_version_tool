package service_test

import (
	"F10-CLI/service"
	"testing"
)

// 同一业务，重复代码，不允许多次发布

// 重复代码，不同业务，可以发布

//
func TestLoadLocalVersionMap(t *testing.T) {
	service.LoadLocalVersionMap()
}

func TestSaveVersionMap(t *testing.T) {
	service.SaveVersionMap()
}


