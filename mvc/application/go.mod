module go-commons/mvc/application

go 1.14

require (
	go-commons/config v0.0.0-00010101000000-000000000000
	go-commons/middleware v0.0.0-00010101000000-000000000000
	go-commons/middleware/auth v0.0.0-00010101000000-000000000000
	go-commons/middleware/cors v0.0.0-00010101000000-000000000000
	go-commons/middleware/recover v0.0.0-00010101000000-000000000000
	go-commons/mvc/context/response v0.0.0-00010101000000-000000000000
	github.com/betacraft/yaag v1.0.0
	github.com/iris-contrib/swagger/v12 v12.0.1
	github.com/kataras/iris/v12 v12.1.8
	github.com/nats-io/nats-server/v2 v2.1.7 // indirect
)

replace go-commons/config => ../../config

replace go-commons/exception => ../../exception

replace go-commons/datasource => ../../datasource

replace go-commons/utils => ../../utils

replace go-commons/utils/yaml => ../../utils/yaml

replace go-commons/middleware/casbin => ../../middleware/casbin

replace go-commons/middleware/jwt => ../../middleware/jwt

replace go-commons/middleware/models => ../../middleware/models

replace go-commons/middleware => ../../middleware

replace go-commons/utils/security => ../../utils/security

replace go-commons/utils/security/aes => ../../utils/security/aes

replace go-commons/mvc/context/response => ../context/response

replace go-commons/mvc/context/response/msg => ../context/response/msg

replace go-commons/mvc/context/request => ../context/request

replace go-commons/mvc/models => ../models

replace go-commons/mvc/context => ../context

replace go-commons/middleware/recover => ../../middleware/recover

replace go-commons/middleware/cors => ../../middleware/cors

replace go-commons/middleware/auth => ../../middleware/auth
