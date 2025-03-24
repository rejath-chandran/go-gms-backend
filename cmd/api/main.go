package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)



const
(
	// Version of the application
	Version = "1.0.0"
)

type Config struct {
	// Port to listen on
	Port int
}

type App struct {
	// Config for the application
	Config Config
	Logger *slog.Logger
}

func main() {	

	var conf Config
	flag.IntVar(&conf.Port, "port", 8080, "Port to listen on")
	flag.Parse()
	logger:=slog.New(slog.NewTextHandler(os.Stdout,nil))


	app := App{
		Config: conf,
		Logger: logger,
	}

router:=httprouter.New()



router.HandlerFunc(http.MethodGet, "/v1/health", func(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
})


	server:=http.Server{
		Addr: fmt.Sprintf(":%d", app.Config.Port),
		Handler: router,
	}
	app.Logger.Info("Server started on port","PORT", app.Config.Port)
	err:=server.ListenAndServe()
	app.Logger.Error("Server stopped", "error", err.Error())
	os.Exit(1)
	
}