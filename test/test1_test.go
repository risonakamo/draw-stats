package time_stats

import (
	"encoding/json"
	"fmt"
	"testing"
	"time-stats/time_stats"

	"github.com/davecgh/go-spew/spew"
)

// general test 1
func Test_test1(t *testing.T) {
    var result []time_stats.TimeEvent=time_stats.ParseSheetTsv("data1.tsv")

	var analysis time_stats.TimeEventAnalysis=time_stats.AnalyseTimeEvents(result)

	var tagAnalysis time_stats.TagBreakdownsDict=time_stats.TagBreakdownForAllTags(result)

	spew.Dump(analysis)
	spew.Dump(tagAnalysis)
}

func Test_test2(t *testing.T) {
    var result []time_stats.TimeEvent=time_stats.ParseSheetTsv("data1.tsv")

	fmt.Println("total",len(result))

	var filteredByItem1 []time_stats.TimeEvent=time_stats.FilterEvents(result,time_stats.FilterDict{
		time_stats.ITEM_TAG:"1",
	})

	if len(filteredByItem1)>=len(result) {
		t.Error("after filtering, was the same length. should be impossible")
	}

	fmt.Println("after filter",len(filteredByItem1))

	var item1analysis time_stats.TimeEventAnalysis=time_stats.AnalyseTimeEvents(filteredByItem1)
	var item1TagAnalysis time_stats.TagBreakdownsDict=time_stats.TagBreakdownForAllTags(filteredByItem1)

	fmt.Println()
	spew.Dump(item1analysis)
	fmt.Println()
	spew.Dump(item1TagAnalysis)

	// pretty.Println(item1analysis)
	// pretty.Println(item1TagAnalysis)

	data,_:=json.MarshalIndent(item1TagAnalysis,""," ")
	fmt.Println(string(data))
}