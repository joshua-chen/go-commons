/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-18 16:03:38
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-28 10:52:19
 */
package cors

import (
	"sync"

	corsmiddleware "github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12/context"

)
type Cors struct {
}
var (
	instance *Cors
	lock     *sync.Mutex = &sync.Mutex{}
)

func Instance() *Cors {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			instance = &Cors{}
		}
	}
	return instance
}
func (a *Cors) New() context.Handler {

	return New()
}

func New() context.Handler {
	return corsmiddleware.New(corsmiddleware.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowedMethods:   []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "X-Token", "Authorization"},
		AllowCredentials: true,
	})
}
