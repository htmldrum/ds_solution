package main

import (
	"fmt"
	"math"
	"time"
)

func (p People) CalculatePayout(rules []Rule, today time.Time) (Payout, error) {
	r, err := p.ruleMatchingInjuryDate(rules, today)
	if err != nil {
		return Payout{}, nil
	}

	normal_pay := p.normalPay()
	overtime_pay := p.overtimePay(r)
	payout_payable := p.payoutPayable(normal_pay, overtime_pay, r)

	s := Payout{
		People:        p,
		Rule:          r,
		NormalPay:     normal_pay,
		OvertimePay:   overtime_pay,
		PayoutPayable: payout_payable,
	}

	return s, nil
}

func (p People) ruleMatchingInjuryDate(rules []Rule, today time.Time) (Rule, error) {
	weeksSinceInjuryDate := p.WeeksSinceInjuryDate(today)

	for _, rule := range rules {
		min, max := rule.getBounds()
		if weeksSinceInjuryDate < max && weeksSinceInjuryDate >= min {
			return rule, nil
		}
	}

	return Rule{}, fmt.Errorf("No rules match")
}

// Returns integer number of weeks since injury date
// Calculation is rounded down to the nearest integer
// Given the granularity of the dates, the arbitrary 'now' Time is constructed
// without regard to hours, seconds, etc
func (p People) WeeksSinceInjuryDate(today time.Time) uint {
	d := today.Sub(p.InjuryDate())

	hours := d.Hours()
	days := hours / HOURS_IN_DAY
	weeks := days / DAYS_IN_WEEK
	floored := math.Floor(weeks)

	return uint(floored)
}

func (p People) normalPay() (c Currency) {
	c = Currency(p.HourlyRate * p.NormalHours)
	return
}

func (p People) overtimePay(r Rule) (c Currency) {
	if r.OvertimeIncluded {
		c = Currency(p.OvertimeRate * p.OvertimeHours)
	}
	return
}

func (p People) payoutPayable(normal_pay, overtime_pay Currency, r Rule) Currency {
	return Currency((float64(normal_pay) + float64(overtime_pay)) * r.PercentagePayable())
}
