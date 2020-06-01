module commons/middleware/models

go 1.14

replace commons/datasource => ../../datasource

replace commons/config => ../../config

replace commons/exception => ../../exception

replace commons/utils/yaml => ../../utils/yaml

replace commons/mvc/context/request => ../../mvc/context/request

require (
	commons/datasource v0.0.0-00010101000000-000000000000 //indirect
	commons/mvc/context/request v0.0.0-00010101000000-000000000000 //indirect
	github.com/ajg/form v1.5.1 // indirect
	github.com/fasthttp-contrib/websocket v0.0.0-20160511215533-1f3b11f56072 // indirect
	github.com/google/go-querystring v1.0.0 // indirect
	github.com/imkira/go-interpol v1.1.0 // indirect
	github.com/k0kubun/colorstring v0.0.0-20150214042306-9440f1994b88 // indirect
	github.com/mattn/go-colorable v0.1.6 // indirect
	github.com/moul/http2curl v1.0.0 // indirect
	github.com/sergi/go-diff v1.1.0 // indirect
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/valyala/fasthttp v1.13.1 // indirect
	github.com/xeipuuv/gojsonschema v1.2.0 // indirect
	github.com/yalp/jsonpath v0.0.0-20180802001716-5cc68e5049a0 // indirect
	github.com/yudai/gojsondiff v1.0.0 // indirect
	github.com/yudai/golcs v0.0.0-20170316035057-ecda9a501e82 // indirect
	github.com/yudai/pp v2.0.1+incompatible // indirect
)
