// functions for adding the date tag to items

package time_stats

import (
	"time"
)

// mutate all time events in a list of time events, adding a date tag to
// all of them
func AddDateTags(events []TimeEvent) {
    for i := range events {
        addDateTag(&events[i])
    }
}

// from a date, generate a day tag. day tag is the date the date is resolved to.
// resolution rules: the day is the same as the date's day, except if the time is
// before 8am, it is moved backwards a day
func genDateTag(date time.Time) TagValue {
    if date.Hour()<8 {
        date=date.Add(-24*time.Hour)
    }

    return TagValue(date.Format("01/02"))
}

// mutate a given timeevent to add to it a date tag
func addDateTag(event *TimeEvent) {
    event.Tags[DATE_TAG]=genDateTag(event.Start)
}