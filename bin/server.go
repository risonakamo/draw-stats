package main

import (
	"path/filepath"
	"runtime"

	"github.com/gofiber/fiber/v2"
)

func main() {
	var here string=getHereDir()

    var app *fiber.App=fiber.New(fiber.Config{
        CaseSensitive:true,
        EnablePrintRoutes:false,
    })

    app.Static("/",filepath.Join(here,"../time-stats-web/build"))

    app.Listen(":4200")
}

// get directory of main function
func getHereDir() string {
    var selfFilepath string
    _,selfFilepath,_,_=runtime.Caller(0)

    return filepath.Dir(selfFilepath)
}