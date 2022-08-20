package main

import (
	context "context"
	flag "flag"
	fmt "fmt"
	gohttp "net/http"
	os "os"
	signal "os/signal"
	atomic "sync/atomic"
	time "time"

	http "interview-test-free-fair/pkg/http"
	mariadb "interview-test-free-fair/pkg/mariadb"
	sys "interview-test-free-fair/pkg/sys"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var healthy int32
var env string
var serverPort int

func main() {
	sys.LogInfo("[Server initializing]")

	flag.StringVar(&env, "env", "prod", "Environment")
	flag.Parse()

	sys.LoadProperties(env)
	serverPort = sys.Properties.ServerPort

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	mariadb.Connect()
	sys.InitMetrics()

	server := buildServer()

	startServer(server)
	<-quit
	stopServer(server)
}

func buildServer() *gohttp.Server {
	mux := gohttp.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	mux.HandleFunc("/ping", ping)

	http.BuildHandlers(mux)

	return &gohttp.Server{
		Addr:         fmt.Sprintf(":%d", serverPort),
		Handler:      http.Tracing()(mux),
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 2 * time.Second,
		IdleTimeout:  10 * time.Second,
	}
}

func startServer(server *gohttp.Server) {
	go func() {
		if err := server.ListenAndServe(); err != nil && err != gohttp.ErrServerClosed {
			sys.LogFatal("[Could not listen on port %d] err:%v", serverPort, err)
		}
	}()

	atomic.StoreInt32(&healthy, 1)
	sys.LogInfo("[Server is ready to handle requests at port %d]", serverPort)
}

func stopServer(server *gohttp.Server) {
	sys.LogInfo("[Server shutting down]")

	atomic.StoreInt32(&healthy, 0)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer func() {
		mariadb.Disconnect()
		cancel()
	}()

	if err := server.Shutdown(ctx); err != nil {
		sys.LogFatal("[Could not gracefully shutdown the server] err:%v", err)
	}

	sys.LogInfo("[Server shutdown complete]")
}

func ping(w gohttp.ResponseWriter, r *gohttp.Request) {
	if atomic.LoadInt32(&healthy) == 1 {
		sys.HTTPResponseWithJSON(w, gohttp.StatusOK, "pong")
		return
	}

	sys.HTTPResponseWithCode(w, gohttp.StatusServiceUnavailable)
}
