/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-18 09:21:47
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-29 00:12:34
 */
package datasource

import (
	"fmt"
	"sync"

	"github.com/joshua-chen/go-commons/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/golog"
	"github.com/xormplus/core"
	"github.com/xormplus/xorm"

)

var (
	masterEngine *xorm.Engine
	slaveEngine  *xorm.Engine
	lock         sync.Mutex
)
//
func MasterEngineByFeature(featureName string) {
	var prefix string
	for _, v := range config.DBConfig.TablePrefixes {
		if v.FeatureName == featureName {
			prefix = v.PrefixName
			break;
		}
	}
	engine := MasterEngine()
	tbMapper := core.NewPrefixMapper(core.GonicMapper{}, prefix)
	engine.SetTableMapper(tbMapper)
}

// 主库，单例
func MasterEngine() *xorm.Engine {
	if masterEngine != nil {
		return masterEngine
	}

	lock.Lock()
	defer lock.Unlock()

	if masterEngine != nil {
		return masterEngine
	}

	master := config.DBConfig.Master
	engine, err := xorm.NewEngine(master.Dialect, GetConnURL(&master))
	if err != nil {
		golog.Fatalf("@@@ Instance Master DB error!! %s", err)
		return nil
	}
	configure(engine, &master)
	engine.SetMapper(core.GonicMapper{})
	config.RegisterSql(engine)
	masterEngine = engine
	return masterEngine
}
func SlaveEngineByFeature(featureName string) {
	var prefix string
	for _, v := range config.DBConfig.TablePrefixes {
		if v.FeatureName == featureName {
			prefix = v.PrefixName
			break;
		}
	}
	engine := SlaveEngine()
	tbMapper := core.NewPrefixMapper(core.GonicMapper{}, prefix)
	engine.SetTableMapper(tbMapper)
}

// 从库，单例
func SlaveEngine() *xorm.Engine {
	if slaveEngine != nil {
		return slaveEngine
	}

	lock.Lock()
	defer lock.Unlock()

	if slaveEngine != nil {
		return slaveEngine
	}

	slave := config.DBConfig.Slave
	engine, err := xorm.NewEngine(slave.Dialect, GetConnURL(&slave))
	if err != nil {
		golog.Fatalf("@@@ Instance Slave DB error!! %s", err)
		return nil
	}
	configure(engine, &slave)
	config.RegisterSql(engine)

	slaveEngine = engine
	return engine
}

//
func configure(engine *xorm.Engine, info *config.DBInfo) {
	engine.ShowSQL(info.ShowSQL)
	engine.SetTZLocation(config.SysTimeLocation)
	if info.MaxIdleConns > 0 {
		engine.SetMaxIdleConns(info.MaxIdleConns)
	}
	if info.MaxOpenConns > 0 {
		engine.SetMaxOpenConns(info.MaxOpenConns)
	}

	// 性能优化的时候才考虑，加上本机的SQL缓存
	//cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	//engine.SetDefaultCacher(cacher)
}

// 获取数据库连接的url
// true：master主库
func GetConnURL(info *config.DBInfo) (url string) {
	//db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	url = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s",
		info.User,
		info.Password,
		info.Host,
		info.Port,
		info.Database,
		info.Charset)
	//golog.Infof("@@@ DB conn==>> %s", url)
	return
}
