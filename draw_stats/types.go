package draw_stats

import "time"

// identifier tags for a time event.
// key: name of the tag
// val: the tag value
type TagsDict map[string]string

// typed version of tags map
type TagsDictTyped struct {
    Item string
    Category string
    Date time.Time
}

// a single event. a full log is just a list of time events
type TimeEvent struct {
    Tags TagsDictTyped

    Start time.Time
    End time.Time

    Duration time.Duration
}




// ---- collection structures ----
// analysis of a list of time events. corresponds with a list of time events
type TimeEventAnalysis struct {
    TotalTime time.Duration
    AverageTime time.Duration

    Events []TimeEvent
}

// analysis of a list of events. focusing on a single tag, the events are grouped by the tag's
// unique values. then, for each unique value, stats are calculated
type TagBreakdown struct {
    Tag string

    // analysis for each tag value
    // key: tag value
    // val: the analysis
    ValuesAnalysis map[string]TimeEventAnalysis
}