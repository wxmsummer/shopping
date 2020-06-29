module shopping/comment

go 1.13

replace shopping/comment/proto/comment => ./proto/comment

require (
	github.com/jinzhu/gorm v1.9.12
	github.com/micro/go-micro/v2 v2.7.0
	github.com/micro/go-plugins/registry/consul/v2 v2.5.0
	github.com/wxmsummer/shopping/comment/proto/comment v0.0.0-20200626085206-ff0f0785ac72 // indirect
	shopping/comment/proto/comment v0.0.0-00010101000000-000000000000
)
