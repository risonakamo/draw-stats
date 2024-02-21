package time_stats

import (
	"testing"
	"time-stats/time_stats"

	"github.com/davecgh/go-spew/spew"
)

// test parsing data and analysing
func Test_parseSheetTsv(t *testing.T) {
    var result []time_stats.TimeEvent=time_stats.ParseSheetTsv("data1.tsv")

	var analysis time_stats.TimeEventAnalysis=time_stats.AnalyseTimeEvents(result)

	var tagAnalysis time_stats.TagBreakdownsDict=time_stats.TagBreakdownForAllTags(result)

	spew.Dump(analysis)
	spew.Dump(tagAnalysis)
}