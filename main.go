package main

import (
	"net/http"
	"os"

	"github.com/metagogs/gogs"
	"github.com/metagogs/gogs/acceptor"
	"github.com/metagogs/gogs/config"
	"github.com/szpnygo/gowebrtc-signal/internal/server"
	"github.com/szpnygo/gowebrtc-signal/internal/svc"
	"github.com/szpnygo/gowebrtc-signal/model"
)

func main() {

	config := config.NewConfig()
	authKey := os.Getenv("AUTH_KEY")

	app := gogs.NewApp(config)
	app.AddAcceptor(acceptor.NewWSAcceptor(&acceptor.AcceptorConfig{
		HttpPort: 8888,
		Name:     "base",
		Groups: []*acceptor.AcceptorGroupConfig{
			{
				GroupName: "base",
				MiddlewareFunc: []acceptor.MiddlewareFunc{
					func(next http.Handler) http.Handler {
						return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
							// check the url param auth is equal to authKey
							if len(authKey) > 0 && r.URL.Query().Get("auth") != authKey {
								w.WriteHeader(http.StatusUnauthorized)
								return
							}
							next.ServeHTTP(w, r)
						})
					},
				},
			},
		},
	}))

	ctx := svc.NewServiceContext(app)
	srv := server.NewServer(ctx)

	model.RegisterAllComponents(app, srv)

	defer app.Shutdown()
	app.Start()
}
