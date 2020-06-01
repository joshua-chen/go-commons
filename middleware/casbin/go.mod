module go-commons/middleware/casbin

go 1.14

require (
	go-commons/config v0.0.0-00010101000000-000000000000
	go-commons/datasource v0.0.0-00010101000000-000000000000
	go-commons/middleware/jwt v0.0.0-00010101000000-000000000000
	go-commons/middleware/models v0.0.0-00010101000000-000000000000
	go-commons/mvc/context/response v0.0.0-00010101000000-000000000000
	go-commons/mvc/context/response/msg v0.0.0-00010101000000-000000000000
	go-commons/utils/security v0.0.0-00010101000000-000000000000 // indirect
	go-commons/utils/security/aes v0.0.0-00010101000000-000000000000
	github.com/casbin/casbin v1.9.1
	github.com/go-sql-driver/mysql v1.5.0
	github.com/kataras/golog v0.0.15
	github.com/kataras/iris/v12 v12.1.8
	github.com/lib/pq v1.5.2
	github.com/xormplus/xorm v0.0.0-20200529061552-7d0d26c6f81c
)

replace go-commons/config => ../../config

replace go-commons/datasource => ../../datasource
replace go-commons/mvc/context => ../../mvc/context

replace go-commons/mvc/context/request => ../../mvc/context/request

replace go-commons/mvc/context/response => ../../mvc/context/response

replace go-commons/mvc/context/response/msg => ../../mvc/context/response/msg

replace go-commons/mvc/models => ../../mvc/models

replace go-commons/utils => ../../utils

replace go-commons/utils/security/aes => ../../utils/security/aes

replace go-commons/utils/security => ../../utils/security

replace go-commons/middleware/jwt => ../jwt

replace go-commons/middleware/models => ../models

replace go-commons/middleware/auth => ../auth

replace go-commons/exception => ../../exception

replace go-commons/utils/yaml => ../../utils/yaml
