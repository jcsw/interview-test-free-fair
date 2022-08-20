package main

import (
	context "context"
	flag "flag"
	fmt "fmt"
	http "net/http"
	os "os"
	signal "os/signal"
	atomic "sync/atomic"
	time "time"

	mariadb "interview-test-free-fair/pkg/infra/mariadb"
	sys "interview-test-free-fair/pkg/sys"
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

	server := buildServer()

	mariadb.Connect()
	sys.InitMetrics()

	startServer(server)
	<-quit
	stopServer(server)
}

func buildServer() *http.Server {
	mux := http.NewServeMux()

	initRouters(mux)

	return &http.Server{
		Addr:         fmt.Sprintf(":%d", serverPort),
		Handler:      tracing()(mux),
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 2 * time.Second,
		IdleTimeout:  10 * time.Second,
	}
}

func startServer(server *http.Server) {
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			sys.LogFatal("[Could not listen on port %d] err:%v", serverPort, err)
		}
	}()

	atomic.StoreInt32(&healthy, 1)
	sys.LogInfo("[Server is ready to handle requests at port %d]", serverPort)
}

func stopServer(server *http.Server) {
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

func ping(w http.ResponseWriter, r *http.Request) {
	if atomic.LoadInt32(&healthy) == 1 {
		sys.HTTPResponseWithJSON(w, http.StatusOK, "pong")
		return
	}

	sys.HTTPResponseWithCode(w, http.StatusServiceUnavailable)
}
