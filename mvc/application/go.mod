module commons/mvc/application

go 1.14

require (
	commons/config v0.0.0-00010101000000-000000000000
	commons/middleware v0.0.0-00010101000000-000000000000
	commons/middleware/auth v0.0.0-00010101000000-000000000000
	commons/middleware/cors v0.0.0-00010101000000-000000000000
	commons/middleware/recover v0.0.0-00010101000000-000000000000
	commons/mvc/context/response v0.0.0-00010101000000-000000000000
	github.com/betacraft/yaag v1.0.0
	github.com/iris-contrib/swagger/v12 v12.0.1
	github.com/kataras/iris/v12 v12.1.8
	github.com/nats-io/nats-server/v2 v2.1.7 // indirect
)

replace commons/config => ../../config

replace commons/exception => ../../exception

replace commons/datasource => ../../datasource

replace commons/utils => ../../utils

replace commons/utils/yaml => ../../utils/yaml

replace commons/middleware/casbin => ../../middleware/casbin

replace commons/middleware/jwt => ../../middleware/jwt

replace commons/middleware/models => ../../middleware/models

replace commons/middleware => ../../middleware

replace commons/utils/security => ../../utils/security

replace commons/utils/security/aes => ../../utils/security/aes

replace commons/mvc/context/response => ../context/response

replace commons/mvc/context/response/msg => ../context/response/msg

replace commons/mvc/context/request => ../context/request

replace commons/mvc/models => ../models

replace commons/mvc/context => ../context

replace commons/middleware/recover => ../../middleware/recover

replace commons/middleware/cors => ../../middleware/cors

replace commons/middleware/auth => ../../middleware/auth
