package main

import (
	"math"
	"testing"
	"time"
)

func TestInuryDateParsing(t *testing.T) {
	tt := []struct {
		person People
		output time.Time
	}{
		{
			People{
				"Ebony Boycott",
				75.0030,
				150.0000,
				35.0,
				7.3,
				"2016/05/31",
			},
			time.Date(2016, time.May, 31, 0, 0, 0, 0, time.UTC),
		},
		{
			People{
				"Jason Lanning",
				40.0055,
				90.9876,
				40.0,
				12.4,
				"2013/01/01",
			},
			time.Date(2013, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tc := range tt {
		idDate := tc.person.InjuryDate()
		if idDate != tc.output {
			t.Errorf("Expected %v, got %v", tc.output, idDate)
		}
	}
}

func TestRuleGetBounds(t *testing.T) {
	tt := []struct {
		Rule
		bounds [2]uint
	}{
		{
			Rule{ApplicableWeeks: "1-26"},
			[2]uint{1, 26},
		},
		{

			Rule{ApplicableWeeks: "26-52"},
			[2]uint{26, 52},
		},
		{
			Rule{ApplicableWeeks: "53-79"},
			[2]uint{53, 79},
		},
		{
			Rule{ApplicableWeeks: "80-104"},
			[2]uint{80, 104},
		},
		{
			Rule{ApplicableWeeks: "104+"},
			[2]uint{104, math.MaxUint32},
		},
	}

	for _, tc := range tt {
		min, max := tc.getBounds()
		if min != tc.bounds[0] {
			t.Errorf("Expected min %v, got %v", tc.bounds[0], min)
		}
		if max != tc.bounds[1] {
			t.Errorf("Expected max %v, got %v", tc.bounds[1], max)
		}
	}
}
