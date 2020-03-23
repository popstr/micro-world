module github.com/postr/micro-world/api-service

go 1.14

//replace github.com/popstr/micro-world/common => ../common
replace github.com/popstr/micro-world/name-service => ../name-service
replace github.com/popstr/micro-world/email-service => ../email-service

require (
	github.com/emicklei/go-restful v2.12.0+incompatible
	github.com/golang/protobuf v1.3.5
	github.com/labstack/echo v3.3.10+incompatible
	github.com/labstack/gommon v0.3.0 // indirect
	github.com/mattn/go-colorable v0.1.6 // indirect
	github.com/micro/go-micro/v2 v2.2.0
	github.com/valyala/fasttemplate v1.1.0 // indirect
	golang.org/x/crypto v0.0.0-20200323165209-0ec3e9974c59 // indirect
	golang.org/x/net v0.0.0-20200320220750-118fecf932d8 // indirect
	golang.org/x/sys v0.0.0-20200321134203-328b4cd54aae // indirect
	"github.com/popstr/micro-world/name-service" v0.0.0
	"github.com/popstr/micro-world/email-service" v0.0.0
)
