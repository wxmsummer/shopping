module shopping/user

go 1.13

replace shopping/user/proto/user => ./proto/user

require (
	github.com/jinzhu/gorm v1.9.12
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.7.0
	github.com/micro/go-plugins/registry/consul v0.0.0-20200119172437-4fe21aa238fd
	github.com/micro/go-plugins/registry/consul/v2 v2.5.0
	github.com/micro/go-plugins/registry/etcd v0.0.0-20200119172437-4fe21aa238fd
	golang.org/x/crypto v0.0.0-20200323165209-0ec3e9974c59
	google.golang.org/protobuf v1.23.0 // indirect
	shopping/user/proto/user v0.0.0-00010101000000-000000000000
)
