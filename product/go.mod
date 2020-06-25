module shopping/product

go 1.13

replace shopping/product/proto/product => ./proto/product

require (
	github.com/jinzhu/gorm v1.9.12
	github.com/micro/go-micro/v2 v2.7.0
	github.com/micro/go-plugins/registry/consul/v2 v2.5.0
	shopping/product/proto/product v0.0.0-00010101000000-000000000000
)
