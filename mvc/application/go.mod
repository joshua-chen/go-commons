module github.com/joshua-chen/go-commons/mvc/application

go 1.14

require (
	github.com/joshua-chen/go-commons/config v0.0.0-00010101000000-000000000000
	github.com/joshua-chen/go-commons/middleware v0.0.0-00010101000000-000000000000
	github.com/joshua-chen/go-commons/middleware/auth v0.0.0-00010101000000-000000000000
	github.com/joshua-chen/go-commons/middleware/cors v0.0.0-00010101000000-000000000000
	github.com/joshua-chen/go-commons/middleware/recover v0.0.0-00010101000000-000000000000
	github.com/joshua-chen/go-commons/mvc/context/response v0.0.0-00010101000000-000000000000
	github.com/betacraft/yaag v1.0.0
	github.com/iris-contrib/swagger/v12 v12.0.1
	github.com/kataras/iris/v12 v12.1.8
	github.com/nats-io/nats-server/v2 v2.1.7 // indirect
)

replace github.com/joshua-chen/go-commons/config => ../../config

replace github.com/joshua-chen/go-commons/exception => ../../exception

replace github.com/joshua-chen/go-commons/datasource => ../../datasource

replace github.com/joshua-chen/go-commons/utils => ../../utils

replace github.com/joshua-chen/go-commons/utils/yaml => ../../utils/yaml

replace github.com/joshua-chen/go-commons/middleware/casbin => ../../middleware/casbin

replace github.com/joshua-chen/go-commons/middleware/jwt => ../../middleware/jwt

replace github.com/joshua-chen/go-commons/middleware/models => ../../middleware/models

replace github.com/joshua-chen/go-commons/middleware => ../../middleware

replace github.com/joshua-chen/go-commons/utils/security => ../../utils/security

replace github.com/joshua-chen/go-commons/utils/security/aes => ../../utils/security/aes

replace github.com/joshua-chen/go-commons/mvc/context/response => ../context/response

replace github.com/joshua-chen/go-commons/mvc/context/response/msg => ../context/response/msg

replace github.com/joshua-chen/go-commons/mvc/context/request => ../context/request

replace github.com/joshua-chen/go-commons/mvc/models => ../models

replace github.com/joshua-chen/go-commons/mvc/context => ../context

replace github.com/joshua-chen/go-commons/middleware/recover => ../../middleware/recover

replace github.com/joshua-chen/go-commons/middleware/cors => ../../middleware/cors

replace github.com/joshua-chen/go-commons/middleware/auth => ../../middleware/auth
