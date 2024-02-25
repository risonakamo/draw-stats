package main

import (
	"path/filepath"
	"runtime"
	"time-stats/time_stats"

	"github.com/gofiber/fiber/v2"
)

func main() {
	var here string=getHereDir()

    var dataDir string=filepath.Join(here,"../data")

    var app *fiber.App=fiber.New(fiber.Config{
        CaseSensitive:true,
        EnablePrintRoutes:false,
    })


    // --- apis ---
    // return list of available data files
    app.Get("/data-names",func (c *fiber.Ctx) error {
        var datalist []time_stats.DataFileInfo=time_stats.GetDataList(dataDir)

        return c.JSON(datalist)
    })


    // --- static ---
    app.Static("/",filepath.Join(here,"../time-stats-web/build"))

    app.Listen(":4200")
}

// get directory of main function
func getHereDir() string {
    var selfFilepath string
    _,selfFilepath,_,_=runtime.Caller(0)

    return filepath.Dir(selfFilepath)
}