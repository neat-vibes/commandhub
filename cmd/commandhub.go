package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/neat-vibes/commandhub/api/route"
	"github.com/neat-vibes/commandhub/processor"
)

func main() {
	// Parse flags
	sAddress := flag.String("a", "127.0.0.1", "server address")
	sPort := flag.String("p", "8080", "server port")
	cmdTimeout := flag.Int("t", 5, "command timeout (seconds)")
	flag.Parse()

	// Create command processor
	p := processor.NewCommandProcessor()

	// Setup gin route
	r := gin.New()
	route.SetupRoute(time.Duration(*cmdTimeout)*time.Second, r, p)

	// Run server
	r.Run(fmt.Sprintf("%s:%s", *sAddress, *sPort))
}
