package controller

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/neat-vibes/commandhub/entities"
)

type CommandController struct {
	CommandUsecase entities.ICommandUsecase
}

func (controller *CommandController) Execute(g *gin.Context) {
	var req entities.Request
	var res entities.Response

	err := g.ShouldBindJSON(&req)
	if err != nil {
		res.Err = err.Error()
		g.JSON(400, res)
		return
	}

	result, exitCode, err := controller.CommandUsecase.ExecuteCommand(g, &req)
	if err != nil {
		res.Err = err.Error()
		g.JSON(500, res)
		return
	}

	// Parsing result as JSON structure
	var parsedResult json.RawMessage
	err = json.Unmarshal([]byte(result), &parsedResult)
	if err != nil {
		parsedResult = json.RawMessage(result)
	}

	res.Result = parsedResult
	res.ExitCode = exitCode
	if err != nil {
		res.Err = err.Error()
	}

	g.JSON(200, res)
}
