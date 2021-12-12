package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"strings"

	noansi "github.com/ELPanaJose/api-deno-compiler/src/routes/others"
	"github.com/labstack/echo"
)

func ExecuteCommandUnixNoAnsi(c echo.Context, command string) {

	cmd := exec.Command("sh", "-c", command+`&`+` sleep 2;kill $! 2>&1`)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	peo := cmd.Run()
	if peo != nil {
		fmt.Println(peo)
	}
	out2 := strings.ReplaceAll(stdout.String()+stderr.String(), "sh: 1: kill: No such process", "")
	output := noansi.NoAnsi(out2)
	c.Response().Header().Set("Content-Type", "application/json")
	c.Response().WriteHeader(http.StatusCreated)
	json.NewEncoder(c.Response()).Encode(output)
}

func ExecuteCommandUnixColor(c echo.Context, command string) {

	cmd := exec.Command("sh", "-c", command+`&`+` sleep 2;kill $! 2>&1`)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	peo := cmd.Run()
	if peo != nil {
		fmt.Println(peo)
	}
	out2 := strings.ReplaceAll(stdout.String()+stderr.String(), "sh: 1: kill: No such process", "")
	c.Response().Header().Set("Content-Type", "application/json")
	c.Response().WriteHeader(http.StatusCreated)
	json.NewEncoder(c.Response()).Encode(out2)
}
