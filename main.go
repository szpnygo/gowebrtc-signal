package main

import (
	"github.com/metagogs/gogs"
	"github.com/metagogs/gogs/acceptor"
	"github.com/metagogs/gogs/config"
	"github.com/szpnygo/gowebrtc-signal/internal/server"
	"github.com/szpnygo/gowebrtc-signal/internal/svc"
	"github.com/szpnygo/gowebrtc-signal/model"
)

func main() {

	config := config.NewConfig()

	app := gogs.NewApp(config)
	app.AddAcceptor(acceptor.NewWSAcceptor(&acceptor.AcceptorConfig{
		HttpPort: 8888,
		Name:     "base",
		Groups: []*acceptor.AcceptorGroupConfig{
			{
				GroupName: "base",
			},
		},
	}))

	ctx := svc.NewServiceContext(app)
	srv := server.NewServer(ctx)

	model.RegisterAllComponents(app, srv)

	defer app.Shutdown()
	app.Start()
}
