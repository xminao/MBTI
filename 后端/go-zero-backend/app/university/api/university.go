package main

import (
	"flag"
	"fmt"
	"net/http"

	"backend/app/university/api/internal/config"
	"backend/app/university/api/internal/handler"
	"backend/app/university/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/university-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf, rest.WithCustomCors(nil, notAllowedFn, "http://localhost:8081"))
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

// 解决跨域问题
func notAllowedFn(w http.ResponseWriter) {
	w.Header().Add("Access-Control-Allow-Headers", "x-token")
}
