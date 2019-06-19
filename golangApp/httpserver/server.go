package httpserver

import (
	"context"
	"log"
	"net/http"
	"sync"
	"time"

	"../common"
	"github.com/gin-gonic/gin"
)

type RESTServer struct {
	httpServer      *http.Server
	router          *gin.Engine
	cpuStateCurrent common.CpuState
	cpuHistory      [10]common.CpuState
	mutexHistory    sync.Mutex
	mutexCurrent    sync.Mutex
}

func (r *RESTServer) UpdateCPUState(state <-chan common.CpuState) {
	go func() {
		iterator := 0
		for val := range state {
			if iterator > 9 {
				iterator = 0
			}
			r.mutexHistory.Lock()
			r.cpuHistory[iterator] = val
			r.mutexHistory.Unlock()
			r.mutexCurrent.Lock()
			r.cpuStateCurrent = val
			r.mutexCurrent.Unlock()
			iterator++
		}
	}()
}
func (r *RESTServer) currentState(c *gin.Context) {
	r.mutexCurrent.Lock()
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
	c.JSON(http.StatusOK, r.cpuStateCurrent)
	r.mutexCurrent.Unlock()
}
func (r *RESTServer) historyState(c *gin.Context) {
	r.mutexHistory.Lock()
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
	c.JSON(http.StatusOK, r.cpuHistory)
	r.mutexHistory.Unlock()
}

func (r *RESTServer) InitWeb(addr string) {
	r.router = gin.Default()

	v1 := r.router.Group("/api/v1/")
	{
		v1.GET("/cpu/state/current", r.currentState)
		v1.GET("/cpu/state/history", r.historyState)
	}

	r.httpServer = &http.Server{
		Addr:    addr,
		Handler: r.router,
	}

}

func (r *RESTServer) ListenAndServer() {
	go func() {
		// service connections

		if err := r.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
}

func (r *RESTServer) Shutdown() {
	log.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := r.httpServer.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
}
