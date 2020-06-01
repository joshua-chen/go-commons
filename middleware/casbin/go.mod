module github.com/joshua-chen/go-commons/middleware/casbin

go 1.14

require (
	github.com/joshua-chen/go-commons/config v0.0.0-00010101000000-000000000000
	github.com/joshua-chen/go-commons/datasource v0.0.0-00010101000000-000000000000
	github.com/joshua-chen/go-commons/middleware/jwt v0.0.0-00010101000000-000000000000
	github.com/joshua-chen/go-commons/middleware/models v0.0.0-00010101000000-000000000000
	github.com/joshua-chen/go-commons/mvc/context/response v0.0.0-00010101000000-000000000000
	github.com/joshua-chen/go-commons/mvc/context/response/msg v0.0.0-00010101000000-000000000000
	github.com/joshua-chen/go-commons/utils/security v0.0.0-00010101000000-000000000000 // indirect
	github.com/joshua-chen/go-commons/utils/security/aes v0.0.0-00010101000000-000000000000
	github.com/casbin/casbin v1.9.1
	github.com/go-sql-driver/mysql v1.5.0
	github.com/kataras/golog v0.0.15
	github.com/kataras/iris/v12 v12.1.8
	github.com/lib/pq v1.5.2
	github.com/xormplus/xorm v0.0.0-20200529061552-7d0d26c6f81c
)

replace github.com/joshua-chen/go-commons/config => ../../config

replace github.com/joshua-chen/go-commons/datasource => ../../datasource
replace github.com/joshua-chen/go-commons/mvc/context => ../../mvc/context

replace github.com/joshua-chen/go-commons/mvc/context/request => ../../mvc/context/request

replace github.com/joshua-chen/go-commons/mvc/context/response => ../../mvc/context/response

replace github.com/joshua-chen/go-commons/mvc/context/response/msg => ../../mvc/context/response/msg

replace github.com/joshua-chen/go-commons/mvc/models => ../../mvc/models

replace github.com/joshua-chen/go-commons/utils => ../../utils

replace github.com/joshua-chen/go-commons/utils/security/aes => ../../utils/security/aes

replace github.com/joshua-chen/go-commons/utils/security => ../../utils/security

replace github.com/joshua-chen/go-commons/middleware/jwt => ../jwt

replace github.com/joshua-chen/go-commons/middleware/models => ../models

replace github.com/joshua-chen/go-commons/middleware/auth => ../auth

replace github.com/joshua-chen/go-commons/exception => ../../exception

replace github.com/joshua-chen/go-commons/utils/yaml => ../../utils/yaml
