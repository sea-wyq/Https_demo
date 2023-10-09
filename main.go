package main

import (
	"log"
	"net/http"

	"github.com/emicklei/go-restful"
)

func main() {
	wsContainer := restful.NewContainer()
	wsContainer.Router(restful.CurlyRouter{})

	ws := new(restful.WebService)
	ws.
		Path("/api").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON) // 可单独设置每一个方法

	//最终地址 https://127.0.0.1:7443/api/hello
	ws.Route(ws.GET("/hello").To(hello))
	wsContainer.Add(ws)
	log.Printf("start listening on localhost:7443")
	server := &http.Server{Addr: ":7443", Handler: wsContainer}

	// log.Fatal(server.ListenAndServe())
	log.Fatal(server.ListenAndServeTLS("./cert/server.crt", "./cert/server.key"))

}

// rest请求处理方法
func hello(request *restful.Request, response *restful.Response) {
	response.Write([]byte("Hello, TLS!"))
}
