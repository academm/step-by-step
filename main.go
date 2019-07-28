package main

import (
        "net/http"
	"github.com/Sirupsen/logrus"
	"github.com/takama/router"
        "os"
	common_handlers "github.com/academm/handlers"
	"github.com/academm/step-by-step/version"
)

var log = logrus.New()

// Run server: go build && step-by-step
// Try requests: curl http://127.0.0.1:8000/test
func main() {
        port := os.Getenv("SERVICE_PORT")
	if len(port) == 0 {
		log.Fatal("Required parameter service port is not set")
	}
	r := router.New()
	r.Logger = logger
	r.GET("/", home)

	// Readiness and liveness probes for Kubernetes
	r.GET("/info", func(c *router.Control) {
		common_handlers.Info(c, version.RELEASE, version.REPO, version.COMMIT)
	})
	r.GET("/healthz", func(c *router.Control) {
		c.Code(http.StatusOK).Body(http.StatusText(http.StatusOK))
	})

        r.Listen("0.0.0.0:" + port)
}
