module shopping/web

go 1.13

replace github.com/wxmsummer/shopping/user/proto/user => ../user/proto/user

require (
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.7.0
	github.com/micro/go-plugins/registry/consul/v2 v2.5.0
	github.com/wxmsummer/shopping/user/proto/user v0.0.0-00010101000000-000000000000
)
