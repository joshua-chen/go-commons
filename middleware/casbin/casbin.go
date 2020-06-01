package casbin

import (
	"commons/config"
	"commons/datasource"
	"commons/middleware/jwt"
	"commons/mvc/context/response"
	"commons/mvc/context/response/msg"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/casbin/casbin"
	//"github.com/casbin/xorm-adapter"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris/v12/context"

)

var (
	adapter *Adapter // Your driver and data source.
	enforcer   *casbin.Enforcer

	adapterLook sync.Mutex
	enforcerLook   sync.Mutex

	rbacModel string
)

type Casbin struct {
}

var (
	instance *Casbin
	lock     *sync.Mutex = &sync.Mutex{}
)

func Instance() *Casbin {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			instance = &Casbin{}
		}
	}
	return instance
}
func (a *Casbin) Filter(ctx context.Context) bool {

	return Filter(ctx)
}

// Casbin is the casbins services which contains the casbins enforcer.
//type Casbin struct {
//	Enforcer *casbins.Enforcer
//}

// New returns the casbins service which receives a casbins enforcer.
//
// Adapt with its `Wrapper` for the entire application
// or its `ServeHTTP` for specific routes or parties.
//func New() *Casbin {
//	return &Casbin{Enforcer: enforcer}
//}

func SetRbacModel(rootID string) {
	rbacModel = fmt.Sprintf(`
[request_definition]
r = sub, obj, act, suf

[policy_definition]
p = sub, obj, act, suf

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub) && keyMatch(r.obj, p.obj) && regexMatch(r.suf, p.suf) && regexMatch(r.act, p.act) || r.sub == "%s"
`, rootID)
}

// 
func GetEnforcer() *casbin.Enforcer {
	if enforcer != nil {
		enforcer.LoadPolicy()
		return enforcer
	}
	enforcerLook.Lock()
	defer enforcerLook.Unlock()
	if enforcer != nil {
		enforcer.LoadPolicy()
		return enforcer
	}

	m := casbin.NewModel(rbacModel)
	//m.AddDef("r", "r", "sub, obj, act, suf")
	//m.AddDef("p", "p", "sub, obj, act, suf")
	//m.AddDef("g", "g", "_, _")
	//m.AddDef("e", "e", "some(where (p.eft == allow))")
	//m.AddDef("m", "m", `g(r.sub, p.sub) && keyMatch(r.obj, p.obj) && regexMatch(r.suf, p.suf) && regexMatch(r.act, p.act) || r.sub == "1"`)

	// Or you can use an existing DB "abc" like this:
	// The adapter will use the table named "casbin_rule".
	// If it doesn't exist, the adapter will create it automatically.
	// a := xormadapter.NewAdapter("mysql", "mysql_username:mysql_password@tcp(127.0.0.1:3306)/abc", true)
	// TODO use go-bindata fill
	//enforcer = casbin.NewEnforcer("conf/rbac_model.conf", singleAdapter())
	enforcer = casbin.NewEnforcer(m, singleAdapter())
	enforcer.EnableLog(true)
	return enforcer
}

func singleAdapter() *Adapter {
	if adapter != nil {
		return adapter
	}
	adapterLook.Lock()
	defer adapterLook.Unlock()
	if adapter != nil {
		return adapter
	}

	master := config.DBConfig.Master
	url := datasource.GetConnURL(&master)
	// Initialize a Gorm adapter and use it in a Casbin enforcer:
	// The adapter will use the MySQL database named "casbins".
	// If it doesn't exist, the adapter will create it automatically.
	// a := xormadapter.NewAdapter("mysql", "root:root@tcp(127.0.0.1:3306)/?charset=utf8&parseTime=True&loc=Local") // Your driver and data source.
	adapter = NewAdapter(master.Dialect, url, true) // Your driver and data source.
	return adapter
}

// casbin????
// ServeHTTP is the iris compatible casbins handler which should be passed to specific routes or parties.
// Usage:
// [...]
// app.Get("/dataset1/resource1", casbinMiddleware.ServeHTTP, myHandler)
// [...]
func Filter(ctx context.Context) bool {
	user, ok := jwt.ParseToken(ctx)
	if !ok {
		return false
	}

	uid := strconv.Itoa(int(user.Id))
	yes := GetEnforcer().Enforce(uid, ctx.Path(), ctx.Method(), ".*")
	if !yes {
		response.Unauthorized(ctx, msg.PermissionsLess, nil)
		ctx.StopExecution()
		return false
	}

	return true
	//ctx.Next()
}

// Wrapper is the router wrapper, prefer this method if you want to use casbins to your entire iris application.
// Usage:
// [...]
// app.WrapRouter(casbinMiddleware.Wrapper())
// app.Get("/dataset1/resource1", myHandler)
// [...]
func Wrapper() func(w http.ResponseWriter, r *http.Request, router http.HandlerFunc) {
	return func(w http.ResponseWriter, r *http.Request, router http.HandlerFunc) {
		//if !c.Check(r) {
		//	w.WriteHeader(http.StatusForbidden)
		//	w.Write([]byte("403 Forbidden"))
		//	return
		//}
		router(w, r)
	}
}
