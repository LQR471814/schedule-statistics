package main

import (
	"calutil/internal/calendar"
	"testing"
	"time"
)

func TestLayerEvents(t *testing.T) {
	type testCase struct {
		input  []calendar.Event
		expect []calendar.Event
	}

	datetime := func(hour, minute int) time.Time {
		return time.Date(2000, time.January, 1, hour, minute, 0, 0, time.Local)
	}

	table := []testCase{
		{
			input: []calendar.Event{
				calendar.Event{Uid: "1", Name: "A", Tags: nil, Start: datetime(9, 30), End: datetime(10, 0)},
				calendar.Event{Uid: "2", Name: "B", Tags: nil, Start: datetime(9, 30), End: datetime(9, 45)},
				calendar.Event{Uid: "3", Name: "C", Tags: nil, Start: datetime(9, 50), End: datetime(10, 0)},
			},
			expect: []calendar.Event{
				calendar.Event{Uid: "1", Name: "B", Tags: nil, Start: datetime(9, 30), End: datetime(9, 45)},
				calendar.Event{Uid: "2", Name: "A", Tags: nil, Start: datetime(9, 45), End: datetime(9, 50)},
				calendar.Event{Uid: "3", Name: "C", Tags: nil, Start: datetime(9, 50), End: datetime(10, 0)},
			},
		},
		{
			input: []calendar.Event{
				calendar.Event{Uid: "1", Name: "A", Tags: nil, Start: datetime(9, 0), End: datetime(9, 15)},
				calendar.Event{Uid: "2", Name: "B", Tags: nil, Start: datetime(9, 15), End: datetime(9, 30)},
				calendar.Event{Uid: "3", Name: "C", Tags: nil, Start: datetime(9, 30), End: datetime(9, 45)},
			},
			expect: []calendar.Event{
				calendar.Event{Uid: "1", Name: "A", Tags: nil, Start: datetime(9, 0), End: datetime(9, 15)},
				calendar.Event{Uid: "2", Name: "B", Tags: nil, Start: datetime(9, 15), End: datetime(9, 30)},
				calendar.Event{Uid: "3", Name: "C", Tags: nil, Start: datetime(9, 30), End: datetime(9, 45)},
			},
		},
		{
			input: []calendar.Event{
				calendar.Event{Uid: "1", Name: "A", Tags: nil, Start: datetime(9, 0), End: datetime(9, 15)},
				calendar.Event{Uid: "2", Name: "B", Tags: nil, Start: datetime(9, 15), End: datetime(9, 45)},
				calendar.Event{Uid: "3", Name: "C", Tags: nil, Start: datetime(9, 30), End: datetime(9, 40)},
			},
			expect: []calendar.Event{
				calendar.Event{Uid: "1", Name: "A", Tags: nil, Start: datetime(9, 0), End: datetime(9, 15)},
				calendar.Event{Uid: "2", Name: "B", Tags: nil, Start: datetime(9, 15), End: datetime(9, 30)},
				calendar.Event{Uid: "3", Name: "C", Tags: nil, Start: datetime(9, 30), End: datetime(9, 40)},
				calendar.Event{Uid: "2", Name: "B", Tags: nil, Start: datetime(9, 40), End: datetime(9, 45)},
			},
		},
	}

	for _, test := range table {
		result := make([]calendar.Event, len(test.input))
		copy(result, test.input)
		DeoverlapEvents(&result)
		if len(test.expect) != len(result) {
			t.Fatalf(
				"event lists are not equal\n\nInput: %s\n\nExpected: %s\n\nResult: %s",
				prettyPrint(test.input),
				prettyPrint(test.expect),
				prettyPrint(result),
			)
		}
		for i := range test.expect {
			if !equalEvents(test.expect[i], result[i]) {
				t.Log(prettyPrint(result))
				t.Fatalf(
					"events are not equal\n\nInput: %s\n\nExpected: %s\n\nResult: %s",
					prettyPrint(test.input),
					prettyPrint(test.expect[i]),
					prettyPrint(result[i]),
				)
			}
		}
	}
}

func equalEvents(a, b calendar.Event) bool {
	return a.Name == b.Name &&
		a.Start.Equal(b.Start) &&
		a.End.Equal(b.End)
}
