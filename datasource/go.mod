module github.com/joshua-chen/go-commons/datasource

go 1.14

replace github.com/joshua-chen/go-commons/utils/yaml => ../utils/yaml

replace github.com/joshua-chen/go-commons/utils/security => ../utils/security

replace github.com/joshua-chen/go-commons/utils/security/aes => ../utils/security/aes

replace github.com/joshua-chen/go-commons/config => ../config

replace github.com/joshua-chen/go-commons/exception => ../exception

require (
	github.com/joshua-chen/go-commons/config v0.0.0-00010101000000-000000000000
	github.com/go-sql-driver/mysql v1.5.0
	github.com/kataras/golog v0.0.15
	github.com/xormplus/core v0.0.0-20200308074340-f3bce19d5f31
	github.com/xormplus/xorm v0.0.0-20200529061552-7d0d26c6f81c
)
