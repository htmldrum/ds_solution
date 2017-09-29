package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

var (
	payoutHeadings = []string{"Name", "Normal Pay", "Overtime Pay", "Payout Payable", "Applicable Weeks"}
)

const (
	HOURS_IN_DAY = 24.0
	DAYS_IN_WEEK = 7.0
)

type Currency float64

func (c Currency) String() string {
	return fmt.Sprintf("%.2f", c)
}

// Reflects the calculated payout based on the inputs
// Serializes inputs for caclulation along with the value returned
type Payout struct {
	People
	Rule
	NormalPay     Currency
	OvertimePay   Currency
	PayoutPayable Currency
}

func (p Payout) Row() []string {
	var r []string

	r = append(r, p.People.Name)
	r = append(r, p.NormalPay.String())
	r = append(r, p.OvertimePay.String())
	r = append(r, p.PayoutPayable.String())
	r = append(r, p.Rule.ApplicableWeeks)
	return r
}

// This is the set of rules used to calculate payments for people. Each rule has:
// - applicableWeeks - the number of weeks from the date of injury this level of payment is used at
// - percentagePayable - what percentage of an employee's qualifiying wage is due back to the employee
// - overtimeIncluded - whether overtime hours are part of the qualifying wage
type Rule struct {
	ApplicableWeeks    string `json:"applicableWeeks"`
	PercentagePayableU uint   `json:"percentagePayable"`
	OvertimeIncluded   bool   `json:"overtimeIncluded"`
}

func (r Rule) PercentagePayable() float64 {
	return float64(r.PercentagePayableU) / 100.0
}

func (r Rule) getBounds() (min uint, max uint) {
	var (
		min_i int
		min_s string
		max_i int
		max_s string
		err   error
	)

	if string(r.ApplicableWeeks[len(r.ApplicableWeeks)-1]) == "+" {
		min_s = r.ApplicableWeeks[0 : len(r.ApplicableWeeks)-1]
		if min_i, err = strconv.Atoi(min_s); err != nil {
			panic(fmt.Sprintf("Invalid input to atoi: %s. %v", min_s, err))
		}
		min = uint(min_i)
		max = math.MaxUint32
		return
	}

	components := strings.Split(r.ApplicableWeeks, "-")
	min_s = components[0]
	if min_i, err = strconv.Atoi(min_s); err != nil {
		panic(fmt.Sprintf("Invalid input to atoi: %s. %v", min_s, err))
	}
	min = uint(min_i)
	max_s = components[1]
	if max_i, err = strconv.Atoi(max_s); err != nil {
		panic(fmt.Sprintf("Invalid input to atoi: %s. %v", max_s, err))
	}
	max = uint(max_i)

	return
}

// This is the list of people who need workers compensation payments. Each person has:
// name - the employee's name
//     hourlyRate - how much they are paid per hour
//     overtimeRate - how much they are paid per overtime hour
//     normalHours - the number of hours per week they usually work
//     overtimeHours - the number of hours they get paid overtime for
//     injuryDate - when they were injured
type People struct {
	Name          string  `json:"name"`
	HourlyRate    float64 `json:"hourlyRate"`
	OvertimeRate  float64 `json:"overtimeRate"`
	NormalHours   float64 `json:"normalHours"`
	OvertimeHours float64 `json:"overtimeHours"`
	InjuryDateS   string  `json:"injuryDate"`
}

type PeopleResponse struct {
	People []People `json:"people"`
}

type RulesResponse struct {
	Rules []Rule `json:"rules"`
}

func (p People) InjuryDate() time.Time {
	var (
		year  int
		month time.Month
		day   int
	)

	components := strings.Split(p.InjuryDateS, "/")
	year, _ = strconv.Atoi(components[0])
	if mInt, err := strconv.Atoi(components[1]); err == nil {
		month = time.Month(mInt)
	}
	day, _ = strconv.Atoi(components[2])

	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}
