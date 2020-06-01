module commons/middleware/casbin

go 1.14

require (
	commons/config v0.0.0-00010101000000-000000000000
	commons/datasource v0.0.0-00010101000000-000000000000 //indirect
	commons/middleware/jwt v0.0.0-00010101000000-000000000000
	commons/middleware/models v0.0.0-00010101000000-000000000000
	commons/mvc/context/response v0.0.0-00010101000000-000000000000
	commons/mvc/context/response/msg v0.0.0-00010101000000-000000000000
	commons/utils/security v0.0.0-00010101000000-000000000000 // indirect
	commons/utils/security/aes v0.0.0-00010101000000-000000000000
	github.com/casbin/casbin v1.9.1
	github.com/go-sql-driver/mysql v1.5.0
	github.com/kataras/golog v0.0.15
	github.com/kataras/iris/v12 v12.1.8
	github.com/lib/pq v1.5.2
	github.com/xormplus/xorm v0.0.0-20200529061552-7d0d26c6f81c
)

replace commons/config => ../../config

replace commons/datasource => ../../datasource

replace commons/mvc/context => ../../mvc/context

replace commons/mvc/context/request => ../../mvc/context/request

replace commons/mvc/context/response => ../../mvc/context/response

replace commons/mvc/context/response/msg => ../../mvc/context/response/msg

replace commons/mvc/models => ../../mvc/models

replace commons/utils => ../../utils

replace commons/utils/security/aes => ../../utils/security/aes

replace commons/utils/security => ../../utils/security

replace commons/middleware/jwt => ../jwt

replace commons/middleware/models => ../models

replace commons/middleware/auth => ../auth

replace commons/exception => ../../exception

replace commons/utils/yaml => ../../utils/yaml
