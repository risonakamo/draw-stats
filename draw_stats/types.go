package draw_stats

import "time"

// all possible tag type strings
type TagType string
const (
    ITEM_TAG TagType="item"
    CATEGORY_TAG TagType="category"
    DATE_TAG TagType="date"
)

type TagValue string

// identifier tags for a time event.
// key: name of the tag
// val: the tag value
type TagsDict map[TagType]TagValue

// time events grouped by tag value. which tag the values come from is kept
// seperate
// key: a unique tag value
// val: list of events, all of which have the particular tag value
type TimeEventsByTagValue map[TagValue][]TimeEvent

// analysis for each tag value
// key: tag value
// val: the analysis
type TagValueAnalysisDict map[TagValue]TimeEventAnalysis

// a single event. a full log is just a list of time events
type TimeEvent struct {
    Tags TagsDict

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
    Tag TagType

    // analysis for each tag value
    // key: tag value
    // val: the analysis
    ValuesAnalysis TagValueAnalysisDict

    // the time events keyed by their values
    // key: the tag value
    // val: the time events that have the particular value
    KeyedEvents TimeEventsByTagValue
}