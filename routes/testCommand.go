package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strings"

	noansi "github.com/ELPanaJose/api-deno-compiler/src/routes/others"
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
	if input == "" {
		json.NewEncoder(c.Response()).Encode("Error, Empty Request.")
	} else {

		fmt.Println(input)
		var stdout, stderr bytes.Buffer
		// sleep 1 second and kill the process
		cmd := exec.Command("sh", "-c", input+`&`+` sleep 0.5;kill $! 2>&1`)
		// show the output
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr
		peo := cmd.Run()
		if peo != nil {
			fmt.Println(err)
		}
		// capture the stderr and stdout
		executedOut := stdout.String() + stderr.String()
		out2 := strings.ReplaceAll(executedOut, "sh: 1: kill: No such process", "")
		output := noansi.NoAnsi(out2)

		// Print the Output
		fmt.Println(executedOut)
		c.Response().Header().Set("Content-Type", "application/json")
		c.Response().WriteHeader(http.StatusCreated)
		// send the response with the headers
		json.NewEncoder(c.Response()).Encode(output)

	}
	return nil
}
