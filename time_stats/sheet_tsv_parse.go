// functions for parsing Sheet Tsv time data format

package time_stats

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"time"
)

// parse sheet tsv format
// [0] -> item tag. if not present, uses the last previous item tag
//   - if first one is missing, raises error
// [1] -> category tag. if not present, use the last previous category tag
// [2] -> start time
// [3] -> end time
// [4] -> duration, but only for confirmation between diffing start/end
//
// any other lines are ignored.
// the first line must have item and category. after that, lines that omit
// item or category will use the previous row's item or category
func ParseSheetTsv(filepath string) []TimeEvent {
    var file *os.File
    var e error
    file,e=os.Open(filepath)

    if e!=nil {
        panic(e)
    }

    var reader *csv.Reader=csv.NewReader(file)
    reader.Comma='\t'

    var lastItem string=""
    var lastCategory string=""

    var timeEvents []TimeEvent=[]TimeEvent{}

    for {
        var record []string
        record,e=reader.Read()

        if e==io.EOF {
            break
        }

        if len(record)<4 {
            fmt.Println("not enough items in record, need 4 items")
            fmt.Println("the record:")
            fmt.Println(record)
            continue
        }



        // --- grab row items into vars ---
        var item string=record[0]
        var category string=record[1]
        var startStr string=record[2]
        var endStr string=record[3]



        // --- time calculations ---
        var startTime time.Time
        var endTime time.Time

        var calculatedDuration time.Duration

        startTime,e=parseSheetTsvTime(startStr)

        if e!=nil {
            fmt.Println("failed to parse start time of row. skipping")
            fmt.Println("the row:")
            fmt.Println(record)
            continue
        }

        endTime,e=parseSheetTsvTime(endStr)

        if e!=nil {
            fmt.Println("failed to parse end time of row. skipping")
            fmt.Println("the row:")
            fmt.Println(record)
            continue
        }

        calculatedDuration=endTime.Sub(startTime)



        // --- tag calculations ---
        // ensure we are able to replace the item and category with the previous
        // values if we need to. if we can't then that's an issue with the data.
        if len(lastItem)==0 && len(item)==0 {
            panic("first item is missing item tag")
        }

        if len(lastCategory)==0 && len(category)==0 {
            panic("first item is missing category tag")
        }

        // fill in the item and category with the previous item or category, if either
        // of them does not exist in this current row
        if len(item)==0 {
            item=lastItem
        }

        if len(category)==0 {
            category=lastCategory
        }

        // fill in the last item/category with the current
        lastItem=item
        lastCategory=category



        // --- final event creation ---
        timeEvents=append(timeEvents,TimeEvent {
            Tags: TagsDict{
                ITEM_TAG:TagValue(item),
                CATEGORY_TAG:TagValue(category),
            },

            Start: startTime,
            End: endTime,

            Duration: calculatedDuration,
        })
    }

    return timeEvents
}

// parse special time format from Sheet Tsv time data
func parseSheetTsvTime(timestr string) (time.Time,error) {
    var parsedTime time.Time
    var e error
    parsedTime,e=time.Parse("01/02 15:04",timestr)

    return parsedTime,e
}