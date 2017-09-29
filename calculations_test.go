package main

import (
	"testing"
	"time"
)

func TestWeeksSinceInjuryDate(t *testing.T) {
	today := time.Date(2017, time.September, 29, 0, 0, 0, 0, time.UTC)
	tt := []struct {
		People People
		output uint
	}{
		{
			People{
				InjuryDateS: "2017/09/29",
			},
			0,
		},
		{
			People{
				InjuryDateS: "2017/09/23",
			},
			0,
		},
		{
			People{
				InjuryDateS: "2017/09/22",
			},
			1,
		},
		{
			People{
				InjuryDateS: "2016/09/29",
			},
			52,
		},
		{
			People{
				InjuryDateS: "2016/09/30",
			},
			52,
		},
		{
			People{
				InjuryDateS: "2016/09/22",
			},
			53,
		},
	}

	for n, tc := range tt {
		if r := tc.People.WeeksSinceInjuryDate(today); r != tc.output {
			t.Errorf("Expected %d, got %d for case #%d", tc.output, r, n+1)
		}
	}
}

func TestruleMatchingInjuryDate(t *testing.T) {
	rules := []Rule{
		Rule{ApplicableWeeks: "1-26"},
		Rule{ApplicableWeeks: "26-52"},
		Rule{ApplicableWeeks: "53-79"},
		Rule{ApplicableWeeks: "80-104"},
		Rule{ApplicableWeeks: "104+"},
	}
	today := time.Date(2017, time.September, 29, 0, 0, 0, 0, time.UTC)

	tt := []struct {
		MatchingRule *Rule
		People       People
		Today        time.Time
	}{
		{
			nil,
			People{InjuryDateS: "2017/09/29"},
			today,
		},
		{
			&rules[0],
			People{InjuryDateS: "2017/09/20"},
			today,
		},
		{
			&rules[1],
			People{InjuryDateS: "2017/09/29"},
			today,
		},
		{
			&rules[2],
			People{InjuryDateS: "2016/09/22"},
			today,
		},
		{
			&rules[3],
			People{InjuryDateS: "2015/09/22"},
			today,
		},
		{
			&rules[4],
			People{InjuryDateS: "2014/09/22"},
			today,
		},
		{
			&rules[4],
			People{InjuryDateS: "1999/01/01"},
			today,
		},
	}

	for n, tc := range tt {
		r, err := tc.People.ruleMatchingInjuryDate(rules, tc.Today)
		if tc.MatchingRule == nil && err == nil {
			t.Errorf("Expected error, got %v for case #%d", err, n+1)
		} else if r != *tc.MatchingRule {
			t.Errorf("Expected %d, got %d for case #%d", tc.MatchingRule, r, n+1)
		}
	}
}

func TestNormalPay(t *testing.T) {
	tt := []struct {
		People
		Expected Currency
	}{
		{
			People{
				HourlyRate:  75.003,
				NormalHours: 35.0,
			},
			Currency(2625.105),
		},
		{
			People{
				HourlyRate:  50.0,
				NormalHours: 40.0,
			},
			Currency(2000.0),
		},
	}

	for n, tc := range tt {
		if c := tc.People.normalPay(); c != tc.Expected {
			t.Errorf("Expected %f, got %f for case #%d", tc.Expected, c, n+1)
		}
	}
}

func TestOvertimePay(t *testing.T) {
	tt := []struct {
		Rule
		People
		Expected Currency
	}{
		{
			Rule{
				OvertimeIncluded: true,
			},
			People{
				OvertimeRate:  150.0000,
				OvertimeHours: 7.3,
			},
			Currency(1095.0),
		},
		{
			Rule{
				OvertimeIncluded: false,
			},
			People{
				OvertimeRate:  150.0000,
				OvertimeHours: 7.3,
			},
			Currency(0.0),
		},
	}

	for n, tc := range tt {
		if c := tc.People.overtimePay(tc.Rule); c != tc.Expected {
			t.Errorf("Expected %f, got %f for case #%d", tc.Expected, c, n+1)
		}
	}
}

func TestPayoutPayable(t *testing.T) {
	tt := []struct {
		People
		Rule
		normal_pay   Currency
		overtime_pay Currency
		Expected     Currency
	}{
		{
			People{},
			Rule{
				PercentagePayableU: 80,
			},
			Currency(2625.105000),
			Currency(1095.000000),
			Currency(2976.0840000000003),
		},
		{
			People{},
			Rule{
				PercentagePayableU: 80,
			},
			Currency(1875.0),
			Currency(0),
			Currency(1500.0),
		},
	}

	for n, tc := range tt {
		if c := tc.People.payoutPayable(tc.normal_pay, tc.overtime_pay, tc.Rule); c != tc.Expected {
			t.Errorf("Expected %f, got %f for case #%d", tc.Expected, c, n+1)
		}
	}
}
