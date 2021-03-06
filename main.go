package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/syariatifaris/arkeus/core/net"
	"github.com/syariatifaris/arkeus/core/framework/handler"
    "github.com/syariatifaris/arkeus/core/log/arklog"

	"github.com/syariatifaris/kumparan/app/core"
)

func main() {
	var (
		router   net.Router
		handlers []handler.THandler
	)

	di := core.NewDependencies()
	di.GetAddedDependencies(&handlers)
	di.GetAddedDependency(&router)

	for _, h := range handlers {
		arklog.INFO.Println("Registering", h.Name())
		h.RegisterHandlers(router)
	}

	var err error
	errChan := make(chan error)
	go func() {
		err = net.Serve(
			&http.Server{Addr: "0.0.0.0:9093", Handler: router.GetHandler()},
		)
		arklog.ERROR.Println(err)
	}()

	term := make(chan os.Signal, 1)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM)
	select {

	case <-term:
		arklog.WARN.Println("Signal terminate detected")
	case err := <-errChan:
		arklog.ERROR.Println("Server error: ", err.Error())
	}
}
