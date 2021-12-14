package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"runtime"

	tools "github.com/ELPanaJose/pairat/src/tools"
	"github.com/labstack/echo"
)

type command struct {
	Command string
}

func UploadCommand(c echo.Context) error {
	switch runtime.GOOS {
	case "linux", "darwin":
		var inputCommand command
		reqBody, err := ioutil.ReadAll(c.Request().Body)
		if err != nil {
			fmt.Fprintf(c.Response(), "Error")
		}
		json.Unmarshal([]byte(reqBody), &inputCommand)
		if inputCommand.Command == "" {
			json.NewEncoder(c.Response()).Encode("Error, Empty Command.")
		} else {
			tools.ExecuteCommandUnixNoAnsi(c, inputCommand.Command)
		}

	case "windows":
		var inputCommand command
		reqBody, err := ioutil.ReadAll(c.Request().Body)
		if err != nil {
			fmt.Fprintf(c.Response(), "Error")
		}
		json.Unmarshal([]byte(reqBody), &inputCommand)
		if inputCommand.Command == "" {
			json.NewEncoder(c.Response()).Encode("Error, Empty Command.")
		} else {
			tools.ExecuteCommandWindowsNoAnsi(c, inputCommand.Command)
		}

	default:
		fmt.Println("no os detected")
	}
	return nil
}
