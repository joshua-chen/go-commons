module github.com/joshua-chen/go-commons/config

go 1.14

require (
	github.com/joshua-chen/go-commons/exception v0.0.0-00010101000000-000000000000
	github.com/joshua-chen/go-commons/utils/yaml v0.0.0-00010101000000-000000000000
	github.com/Chronokeeper/anyxml v0.0.0-20160530174208-54457d8e98c6 // indirect
	github.com/CloudyKit/fastprinter v0.0.0-20200109182630-33d98a066a53 // indirect
	github.com/CloudyKit/jet v2.1.2+incompatible
	github.com/agrison/go-tablib v0.0.0-20160310143025-4930582c22ee // indirect
	github.com/agrison/mxj v0.0.0-20160310142625-1269f8afb3b4 // indirect
	github.com/bndr/gotabulate v1.1.2 // indirect
	github.com/clbanning/mxj v1.8.4 // indirect
	github.com/fatih/structs v1.1.0 // indirect
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/go-sql-driver/mysql v1.5.0
	github.com/kataras/golog v0.0.15
	github.com/mattn/go-sqlite3 v2.0.3+incompatible // indirect
	github.com/syndtr/goleveldb v1.0.0 // indirect
	github.com/tealeg/xlsx v1.0.5 // indirect
	github.com/xormplus/builder v0.0.0-20200331055651-240ff40009be // indirect
	github.com/xormplus/core v0.0.0-20200308074340-f3bce19d5f31
	github.com/xormplus/xorm v0.0.0-20200514184607-0f37421d8714
	gopkg.in/flosch/pongo2.v3 v3.0.0-20141028000813-5e81b817a0c4 // indirect
	gopkg.in/yaml.v2 v2.3.0
)

replace github.com/joshua-chen/go-commons/utils/yaml => ../utils/yaml

replace github.com/joshua-chen/go-commons/utils/security => ../utils/security

replace github.com/joshua-chen/go-commons/utils/security/aes => ../utils/security/aes

replace github.com/joshua-chen/go-commons/exception => ../exception
