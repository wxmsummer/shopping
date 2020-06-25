module shopping/comment

go 1.13

replace shopping/comment/proto/comment => ./proto/comment

require (
	github.com/jinzhu/gorm v1.9.12
	github.com/micro/go-micro/v2 v2.7.0
	github.com/micro/go-plugins/registry/consul/v2 v2.5.0
)
