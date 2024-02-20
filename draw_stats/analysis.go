package draw_stats

import "time"

// analyse list of time events
func analyseTimeEvents(events []TimeEvent) TimeEventAnalysis {
	var totalTime time.Duration=time.Duration(0)

	for i := range events {
        totalTime=totalTime+events[i].Duration
	}

    return TimeEventAnalysis {
        TotalTime: totalTime,
        AverageTime: time.Duration(
            totalTime.Nanoseconds()/int64(len(events)),
        ),

        Events: events,
    }
}

func genTagBreakdown(events []TimeEvent,tag TagType) TagBreakdown {

}