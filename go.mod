module github.com/joshua-chen/go-commons

go 1.14

require (
	github.com/Chronokeeper/anyxml v0.0.0-20160530174208-54457d8e98c6 // indirect
	github.com/CloudyKit/jet v2.1.3-0.20180809161101-62edd43e4f88+incompatible
	github.com/agrison/go-tablib v0.0.0-20160310143025-4930582c22ee // indirect
	github.com/agrison/mxj v0.0.0-20160310142625-1269f8afb3b4 // indirect
	github.com/ajg/form v1.5.1 // indirect
	github.com/betacraft/yaag v1.0.0
	github.com/bndr/gotabulate v1.1.2 // indirect
	github.com/casbin/casbin v1.9.1
	github.com/clbanning/mxj v1.8.4 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fasthttp-contrib/websocket v0.0.0-20160511215533-1f3b11f56072 // indirect
	github.com/go-playground/locales v0.13.0
	github.com/go-playground/universal-translator v0.17.0
	github.com/go-playground/validator/v10 v10.4.1
	github.com/go-redis/redis v6.15.9+incompatible // indirect
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/snappy v0.0.2 // indirect
	github.com/google/uuid v1.1.2
	github.com/iris-contrib/middleware/cors v0.0.0-20191219204441-78279b78a367
	github.com/iris-contrib/swagger/v12 v12.0.1
	github.com/jmespath/go-jmespath v0.3.0
	github.com/k0kubun/colorstring v0.0.0-20150214042306-9440f1994b88 // indirect
	github.com/kataras/golog v0.0.15
	github.com/kataras/iris/v12 v12.1.8
	github.com/lib/pq v1.6.0
	github.com/mattn/go-colorable v0.1.6 // indirect
	github.com/mattn/go-sqlite3 v2.0.3+incompatible // indirect
	github.com/nats-io/nats-server/v2 v2.1.7 // indirect
	github.com/onsi/ginkgo v1.12.2 // indirect
	github.com/sergi/go-diff v1.1.0 // indirect
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/spf13/cast v1.3.0
	github.com/syndtr/goleveldb v1.0.0 // indirect
	github.com/tealeg/xlsx v1.0.5 // indirect
	github.com/valyala/fasthttp v1.13.1 // indirect
	github.com/xeipuuv/gojsonschema v1.2.0 // indirect
	github.com/xormplus/builder v0.0.0-20200331055651-240ff40009be // indirect
	github.com/xormplus/core v0.0.0-20200308074340-f3bce19d5f31
	github.com/xormplus/xorm v0.0.0-20201020065950-782d727a761a
	github.com/yudai/pp v2.0.1+incompatible // indirect
	golang.org/x/sys v0.0.0-20201207223542-d4d67f95c62d // indirect
	gopkg.in/flosch/pongo2.v3 v3.0.0-20141028000813-5e81b817a0c4 // indirect
	gopkg.in/yaml.v2 v2.4.0
)

replace github.com/joshua-chen/go-commons/config => ./config

replace github.com/joshua-chen/go-commons/datasource => ./datasource

replace github.com/joshua-chen/go-commons/exception => ./exception

replace github.com/joshua-chen/go-commons/mvc/context => ./mvc/context

replace github.com/joshua-chen/go-commons/mvc/context/request => ./mvc/context/request

replace github.com/joshua-chen/go-commons/mvc/context/response => ./mvc/context/response

replace github.com/joshua-chen/go-commons/mvc/context/response/msg => ./mvc/context/response/msg

replace github.com/joshua-chen/go-commons/mvc/models => ./mvc/models

replace github.com/joshua-chen/go-commons/mvc/route => ./mvc/route

replace github.com/joshua-chen/go-commons/utils => ./utils

replace github.com/joshua-chen/go-commons/middleware => ./middleware

replace github.com/joshua-chen/go-commons/middleware/jwt => ./middleware/jwt

replace github.com/joshua-chen/go-commons/middleware/casbin => ./middleware/casbin

replace github.com/joshua-chen/go-commons/middleware/auth => ./middleware/auth

replace github.com/joshua-chen/go-commons/middleware/recover => ./middleware/recover

replace github.com/joshua-chen/go-commons/middleware/perm => ./middleware/perm

replace github.com/joshua-chen/go-commons/utils/yaml => ./utils/yaml

replace github.com/joshua-chen/go-commons/utils/security => ./utils/security

replace github.com/joshua-chen/go-commons/utils/security/aes => ./utils/security/aes

replace github.com/joshua-chen/go-commons/utils/json => ./utils/json
