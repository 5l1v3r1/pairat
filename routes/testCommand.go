package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"

	"github.com/labstack/echo"
)

type command struct {
	Command string
}

func UploadCommand(c echo.Context) error {
	// get the request
	var inputCommand command
	reqBody, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		fmt.Fprintf(c.Response(), "Error")
	}

	json.Unmarshal([]byte(reqBody), &inputCommand)
	input := inputCommand.Command
	fmt.Println(input)
	var stdout, stderr bytes.Buffer
	cmd := exec.Command("sh", "-c", input)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	peo := cmd.Run()
	if peo != nil {
		fmt.Println(err)
	}
	// capture the stderr and stdout
	executedOut := stdout.String() + stderr.String()
	c.Response().Header().Set("Content-Type", "application/json")
	c.Response().WriteHeader(http.StatusCreated)
	// send the response with the headers
	json.NewEncoder(c.Response()).Encode(executedOut)
	return nil

}
