module go-commons/middleware/jwt

go 1.14

require (
	go-commons/config v0.0.0-00010101000000-000000000000
	go-commons/datasource v0.0.0-00010101000000-000000000000 // indirect
	go-commons/middleware/models v0.0.0-00010101000000-000000000000
	go-commons/mvc/context/request v0.0.0-00010101000000-000000000000 // indirect
	go-commons/mvc/context/response v0.0.0-00010101000000-000000000000
	go-commons/mvc/context/response/msg v0.0.0-00010101000000-000000000000
	github.com/ajg/form v1.5.1 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fasthttp-contrib/websocket v0.0.0-20160511215533-1f3b11f56072 // indirect
	github.com/google/go-querystring v1.0.0 // indirect
	github.com/imkira/go-interpol v1.1.0 // indirect
	github.com/iris-contrib/middleware/cors v0.0.0-20191219204441-78279b78a367
	github.com/iris-contrib/middleware/jwt v0.0.0-20191219204441-78279b78a367
	github.com/jmespath/go-jmespath v0.3.0 // indirect
	github.com/k0kubun/colorstring v0.0.0-20150214042306-9440f1994b88 // indirect
	github.com/kataras/golog v0.0.15
	github.com/kataras/iris/v12 v12.1.8
	github.com/mattn/go-colorable v0.1.6 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/moul/http2curl v1.0.0 // indirect
	github.com/onsi/ginkgo v1.12.2 // indirect
	github.com/sergi/go-diff v1.1.0 // indirect
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/spf13/cast v1.3.0
	github.com/valyala/fasthttp v1.13.1 // indirect
	github.com/xeipuuv/gojsonschema v1.2.0 // indirect
	github.com/yalp/jsonpath v0.0.0-20180802001716-5cc68e5049a0 // indirect
	github.com/yudai/gojsondiff v1.0.0 // indirect
	github.com/yudai/golcs v0.0.0-20170316035057-ecda9a501e82 // indirect
	github.com/yudai/pp v2.0.1+incompatible // indirect
)

replace go-commons/config => ../../config

replace go-commons/datasource => ../../datasource

replace go-commons/mvc/context/request => ../../mvc/context/request

replace go-commons/utils/yaml => ../../utils/yaml

replace go-commons/exception => ../../exception

replace go-commons/middleware/models => ../models

replace go-commons/mvc/context/response => ../../mvc/context/response

replace go-commons/mvc/context/response/msg => ../../mvc/context/response/msg
