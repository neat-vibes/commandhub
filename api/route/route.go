package route

import (
	"time"

	"github.com/neat-vibes/commandhub/entities"

	"github.com/gin-gonic/gin"
)

// SetupRoute sets up the routes for the application using the provided Gin Engine.
// It configures the protected routes that require authentication.
func SetupRoute(timeout time.Duration, gin *gin.Engine, p entities.ICommandProcessor) {
	// Create a group for protected routes
	protectedRouter := gin.Group("")
	{
		// Register routes for the Command feature
		NewCommandRoute(timeout, protectedRouter, p)
	}
}
