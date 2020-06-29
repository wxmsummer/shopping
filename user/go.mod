module shopping/user

go 1.13

replace shopping/user/proto/user => ./proto/user

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/jinzhu/gorm v1.9.12
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.9.0
	github.com/micro/go-plugins/registry/consul v0.0.0-20200119172437-4fe21aa238fd
	github.com/micro/go-plugins/registry/consul/v2 v2.5.0
	golang.org/x/crypto v0.0.0-20200510223506-06a226fb4e37
	shopping/user/proto/user v0.0.0-00010101000000-000000000000
)
