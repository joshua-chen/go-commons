package query

import (
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joshua-chen/go-commons/mvc/context/request"
	_ "github.com/kataras/golog"
	_ "github.com/xormplus/core"
	"github.com/xormplus/xorm"

)

type Query struct {
	engine       *xorm.Engine
	lastSQL      string
	CountSession *xorm.Session
	RowsSession  *xorm.Session
}

var (
	instance *Query
	lock     *sync.Mutex = &sync.Mutex{}
)

func Instance() *Query {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			instance = &Query{}
		}
	}
	return instance
}
func New(engine *xorm.Engine) *Query {

	return &Query{
		engine: engine,
	}
}

//
func (q *Query) LimitSQL(limit, start int, sql string, args ...interface{}) *Query {

	q.CountSession = q.engine.SQL(sql, args)
	q.RowsSession = q.engine.SQL(sql, args).Limit(limit, start)

	return q
}

//
func (q *Query) FindAndCount(rowsSlicePtr interface{}) (int64, error) {

	count, err := q.CountSession.Query().Count()
	err = q.RowsSession.Find(rowsSlicePtr)
	return int64(count), err
}

//
func (q *Query) PaginationSQL(page *request.Pagination, sql string, args ...interface{}) *Query {

	q.LimitSQL(page.Limit, page.Offset, sql, args)

	if page.SortName != "" {
		switch page.SortOrder {
		case "asc":
			q.CountSession.Asc(page.SortName)
			q.RowsSession.Asc(page.SortName)
		case "desc":
			q.CountSession.Desc(page.SortName)
			q.RowsSession.Desc(page.SortName)
		}
	}

	return q
}
