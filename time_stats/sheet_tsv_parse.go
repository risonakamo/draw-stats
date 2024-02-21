// functions for parsing Sheet Tsv time data format

package time_stats

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/davecgh/go-spew/spew"
)

// parse sheet tsv format
// [0] -> item tag. if not present, uses the last previous item tag
//   - if first one is missing, raises error
// [1] -> category tag. if not present, use the last previous category tag
// [2] -> start time
// [3] -> end time
// [4] -> duration, but only for confirmation between diffing start/end
// any other lines are ignored.
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

    for {
        var record []string
        record,e=reader.Read()

        if e==io.EOF {
            break
        }

        if len(record)<5 {
            fmt.Println("not enough lines in record, need 5 lines")
            fmt.Println("the record:")
            spew.Dump(record)
        }

        var item string=record[0]
        var category string=record[1]
        var startStr string=record[2]
        var endStr string=record[3]
        var durationStr string=record[4]

        var startTime time.Time
        var endTime time.Time

        startTime,e=parseSheetTsvTime(startStr)

        if e!=nil {
            panic(e)
        }

        endTime,e=parseSheetTsvTime(endStr)

        if e!=nil {
            panic(e)
        }



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


    }

    return []TimeEvent{}
}

// parse special time format from Sheet Tsv time data
func parseSheetTsvTime(timestr string) (time.Time,error) {
    var parsedTime time.Time
    var e error
    parsedTime,e=time.Parse("01/02 15:04",timestr)

    return parsedTime,e
}