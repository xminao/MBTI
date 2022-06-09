package main

import (
	"backend/app/user/api/internal/config"
	"backend/app/user/api/internal/handler"
	"backend/app/user/api/internal/svc"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"net/http"
)

var configFile = flag.String("f", "etc/user-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf, rest.WithCustomCors(nil, notAllowedFn, "http://43.138.137.222:6789"))
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

// 解决跨域问题
func notAllowedFn(w http.ResponseWriter) {
	w.Header().Add("Access-Control-Allow-Headers", "x-token")
}
