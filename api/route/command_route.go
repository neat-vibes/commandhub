package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/neat-vibes/commandhub/api/controller"
	"github.com/neat-vibes/commandhub/entities"
	"github.com/neat-vibes/commandhub/usecase"
)

// NewCommandRoute registers the routes related to the Command feature under the provided Gin RouterGroup.
// It creates an instance of the CommandController and associates it with the CommandUsecase.
// The "/command" route is registered as a POST route, and it maps to the Execute method of the CommandController.
// This route is responsible for executing commands received from the client.
func NewCommandRoute(timeout time.Duration, group *gin.RouterGroup, p entities.ICommandProcessor) {
	// Create an instance of the CommandController
	cc := controller.CommandController{
		CommandUsecase: usecase.NewCommandUsecase(time.Duration(5)*time.Second, p),
	}

	// Register the "/command" route as a POST route under the provided RouterGroup
	// and map it to the Execute method of the CommandController
	group.POST("/command", cc.Execute)
}
