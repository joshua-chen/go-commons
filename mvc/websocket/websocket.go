/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-29 10:53:17
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-29 11:24:54
 */
package websocket

import (
	_ "commons/utils/security/aes"
	_ "errors"
	"log"
	_ "time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
	_ "github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/websocket"

)

type clientPage struct {
	Title string
	Host  string
}

func RegisterWs(app *iris.Application) {

	ws := websocket.New(websocket.DefaultGorillaUpgrader, websocket.Events{
		websocket.OnNativeMessage: func(nsConn *websocket.NSConn, msg websocket.Message) error {
			log.Printf("Server got: %s from [%s]", msg.Body, nsConn.Conn.ID())

			nsConn.Conn.Server().Broadcast(nsConn, msg)
			return nil
		},
	})

	ws.OnConnect = func(c *websocket.Conn) error {
		log.Printf("[%s] Connected to server!", c.ID())
		return nil
	}

	ws.OnDisconnect = func(c *websocket.Conn) {
		log.Printf("[%s] Disconnected from server", c.ID())
	}

	app.HandleDir("/js", "./static/js") // serve our custom javascript code.

	// register the server on an endpoint.
	// see the inline javascript code i the websockets.html, this endpoint is used to connect to the server.
	//app.Get("/my_endpoint", websocket.Handler(ws))

	//deps := hero.New()
	//service := services.NewUserService()
	//deps.Register(service)
	app.Get("/websocket", websocket.Handler(ws))

	//app.Get("/", func(ctx iris.Context) {
	//	ctx.View("client.html", clientPage{"Client Page", "localhost:8080"})
	//})

}

func WebSocket() {

}
