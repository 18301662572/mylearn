package grpc_web

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net/http"
	"strings"
)

//GRPC和Web服务共存
//https://www.jishuchi.com/read/GO/695

//GRPC构建在HTTP/2协议之上，因此我们可以将GRPC服务和普通的Web服务架设在同一个端口之上。
// 因为目前Go语言版本的GRPC实现还不够完善，只有启用了TLS协议之后才能将GRPC和Web服务运行在同一个端口。

//启用http服务
func stathttp() {
	var port = "localhost:1234"
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, "hello")
	})
	//启用TLS，需要先配置TLS 服务器证书
	http.ListenAndServeTLS(port, "server.crt", "server.key",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			mux.ServeHTTP(w, r)
			return
		}),
	)
}

//单独启用带证书的GRPC服务
func statrpcTSL() {
	creds, err := credentials.NewServerTLSFromFile("server.crt", "server.key")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer(grpc.Creds(creds))
	//...
	fmt.Println(grpcServer)
}

//因为GRPC服务已经实现了ServeHTTP方法，
//可以直接作为Web路由处理对象。如果将GRPC和Web服务放在一起，会导致GRPC和Web路径的冲突，在处理时我们需要区分两类服务。
//main()同时支持Web和GRPC协议的路由处理函数
func main() {
	//...
	var port = "localhost:1234"
	mux := http.NewServeMux()      //http服务
	grpcServer := grpc.NewServer() //grpc服务
	http.ListenAndServeTLS(port, "server.crt", "server.key",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.ProtoMajor != 2 {
				mux.ServeHTTP(w, r)
				return
			}
			if strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
				grpcServer.ServeHTTP(w, r) // GRPC Server
				return
			}
			mux.ServeHTTP(w, r)
			return
		}),
	)
}

//首先GRPC是建立在HTTP/2版本之上，如果HTTP不是HTTP/2协议则必然无法提供GRPC支持。
// 同时，每个GRPC调用请求的Content-Type类型会被标注为”application/grpc”类型。
//这样我们就可以在GRPC端口上同时提供Web服务了。
