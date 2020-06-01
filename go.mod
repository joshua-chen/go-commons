module commons

go 1.14

replace commons/config => ./config

replace commons/mvc/context => ./mvc/context

replace commons/mvc/context/request => ./mvc/context/request

replace commons/mvc/context/response => ./mvc/context/response

replace commons/mvc/context/response/msg => ./mvc/context/response/msg

replace commons/mvc/models => ./mvc/models

replace commons/utils => ./utils

replace commons/middleware/jwt => ./middleware/jwt

replace commons/middleware/casbin => ./middleware/casbin

replace commons/middleware/auth => ./middleware/auth

replace commons/exception => ./exception

replace commons/utils/yaml => ./utils/yaml
