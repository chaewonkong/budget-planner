package design

// go에서는 import하지 않은 의존성은 go.mod에 관리되지 않기 때문에 작성

import (
	// codegen
	_ "goa.design/goa/v3/codegen/generator"
	// plugin/zaplogger
	_ "goa.design/plugins/v3/zerologger"
	// grpc
	_ "github.com/grpc-ecosystem/go-grpc-middleware"
)
