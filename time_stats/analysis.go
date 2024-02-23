// analysis functions

package time_stats

import (
	"fmt"
	"time"

	"github.com/davecgh/go-spew/spew"
	mapset "github.com/deckarep/golang-set/v2"
)

// analyse list of time events
func AnalyseTimeEvents(events []TimeEvent) TimeEventAnalysis {
	var totalTime time.Duration=time.Duration(0)

	for i := range events {
        totalTime=totalTime+events[i].Duration
	}

    return TimeEventAnalysis {
        TotalTime: totalTime,
        AverageTime: time.Duration(
            totalTime.Nanoseconds()/int64(len(events)),
        ),
    }
}

// for a list of events, generate tag breakdowns for all of the seen tags
func TagBreakdownForAllTags(events []TimeEvent) TagBreakdownsDict {
    var allTags []TagType=findAllTags(events)

    var breakdowns TagBreakdownsDict=make(TagBreakdownsDict)

    for i := range allTags {
        breakdowns[allTags[i]]=genTagBreakdown(events,allTags[i])
    }

    return breakdowns
}

// filter list of time events based on given filter dict. check filter dict
// docs for how to choose filters
func FilterEvents(events []TimeEvent,filter FilterDict) []TimeEvent {
    var filteredEvents []TimeEvent=events

    var filterTag TagType
    var filterTagValue TagValue
    for filterTag,filterTagValue = range filter {
        filteredEvents=filterByTag(events,filterTag,filterTagValue)
    }

    return filteredEvents
}

// generate a tag breakdown from a list of events for a certain target tag.
// all events must have the tag. events without the tag will print out a warning and
// be excluded from the breakdown
func genTagBreakdown(events []TimeEvent,targetTag TagType) TagBreakdown {
    var keyedEvents TimeEventsByTagValue=groupEventsByTagValue(events,targetTag)

    var analysisDict TagValueAnalysisDict=make(TagValueAnalysisDict)

    var tagValue TagValue
    for tagValue = range keyedEvents {
        analysisDict[tagValue]=AnalyseTimeEvents(keyedEvents[tagValue])
    }

    var totalTime time.Duration=AnalyseTimeEvents(events).TotalTime
    var avgTimePerTagValue=time.Duration(totalTime.Nanoseconds()/int64(len(analysisDict)))

    return TagBreakdown {
        Tag: targetTag,

        ValuesAnalysis: analysisDict,
        // KeyedEvents: keyedEvents,

        AverageTime: avgTimePerTagValue,
        TotalTime: totalTime,
    }
}

// group events in a list of events by their value for a certain target tag
func groupEventsByTagValue(events []TimeEvent,targetTag TagType) TimeEventsByTagValue {
    // events keyed by tag value
    var keyedEvents TimeEventsByTagValue=make(TimeEventsByTagValue)

    for i := range events {
        var tagValue TagValue
        var exists bool

        tagValue,exists=events[i].Tags[targetTag]

        if !exists {
            fmt.Println("event that was missing target tag")
            fmt.Println("missing tag: ",targetTag)
            fmt.Println("event:")
            spew.Dump(events[i])
            continue
        }

        _,exists=keyedEvents[tagValue]

        if !exists {
            keyedEvents[tagValue]=[]TimeEvent{}
        }

        keyedEvents[tagValue]=append(keyedEvents[tagValue],events[i])
    }

    return keyedEvents
}

// filter list of events to only those with a tag with the specified value
func filterByTag(events []TimeEvent,tag TagType,tagValue TagValue) []TimeEvent {
    var results []TimeEvent

    for i := range events {
        var itemTagValue TagValue
        var exists bool
        itemTagValue,exists=events[i].Tags[tag]
        if exists && itemTagValue==tagValue {
            results=append(results,events[i])
        }
    }

    return results
}

// of a list of events, get all unique tag types
func findAllTags(events []TimeEvent) []TagType {
    var seenTypes mapset.Set[TagType]=mapset.NewSet[TagType]()

    // for all events
    for i := range events {
        // for all tags in event
        var eventTag TagType
        for eventTag = range events[i].Tags {
            seenTypes.Add(eventTag)
        }
    }

    return seenTypes.ToSlice()
}