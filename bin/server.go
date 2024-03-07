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

    // return data from a requested data file. generates analysis on the data file.
    // todo: missing filtering
    app.Post("/get-data",func (c *fiber.Ctx) error {
        var body time_stats.GetDataRequest
        var e error=c.BodyParser(&body)

        if e!=nil {
            panic(e)
        }

        var fullFilePath string=filepath.Join(dataDir,body.Filename)

        var timeEvents []time_stats.TimeEvent
        timeEvents,e=time_stats.ParseSheetTsv(fullFilePath,true)

        if e!=nil {
            return e
        }

        time_stats.AddDateTags(timeEvents)

        var analysis time_stats.TimeEventAnalysis=time_stats.AnalyseTimeEvents(timeEvents)

        var tagAnalysis time_stats.TagBreakdownsDict=time_stats.TagBreakdownForAllTags(timeEvents)

        var response time_stats.GetDataResponse=time_stats.GetDataResponse {
            TopAnalysis: analysis,
            TagsAnalysis: tagAnalysis,
        }

        return c.JSON(response)
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