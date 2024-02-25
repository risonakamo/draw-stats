package time_stats

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

// collection of tag breakdowns. associated with a list of events
type TagBreakdownsDict map[TagType]TagBreakdown

// a dict that can be used as a filter. includes the tags you would like to filter
// on, and the value of that tag. when applied to a list of events, only the events
// that have the proper tag-tagvalue pairs will be kept.
//
// the filter dict can have multiple filters, which are all AND'd - all items must have
// all of the tag-value pairs to be included in the result
// key: tag
// val: tag value (making a tag-value pair with the key)
type FilterDict map[TagType]TagValue

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
    TotalTime time.Duration `json:"totalTime"`
    AverageTime time.Duration `json:"averageTime"`

    EarliestEventDate time.Time `json:"earliestEventDate"`
}

// analysis of a list of events. focusing on a single tag, the events are grouped by the tag's
// unique values. then, for each unique value, stats are calculated
type TagBreakdown struct {
    Tag TagType `json:"tag"`

    // analysis for each tag value
    // key: tag value
    // val: the analysis
    ValuesAnalysis TagValueAnalysisDict `json:"valuesAnalysis"`

    // average time per unique value in the values analysis
    AverageTime time.Duration `json:"averageTime"`

    // total time of all the items in this breakdown
    TotalTime time.Duration `json:"totalTime"`
}