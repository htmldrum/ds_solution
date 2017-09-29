package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/htmldrum/ds_solution/writers"
)

var (
	RulesUrlArg    = flag.String("rules", "https://raw.githubusercontent.com/DoneSafe/code_test/master/rules.json", "URL of rules.json")
	PeopleUrlArg   = flag.String("people", "https://raw.githubusercontent.com/DoneSafe/code_test/master/people.json", "URL of people.json")
	OutputToCsvArg = flag.String("o", "", "Specify path to output report as CSV")
)

func main() {
	PrintReport(*RulesUrlArg, *PeopleUrlArg, *OutputToCsvArg)
	return
}

func PrintReport(rulesUrl string, peopleUrl string, outputToCsv string) {
	var (
		payouts []writers.RowPrintable
		writer  writers.ReportWriter
		today   = getToday()
		c       = http.Client{}
	)

	rules, err := fetchRules(rulesUrl, c)
	if err != nil {
		panic(err.Error())
	}

	people, err := fetchPeople(peopleUrl, c)
	if err != nil {
		panic(err.Error())
	}

	for _, person := range people {
		if payout, err := person.CalculatePayout(rules, today); err == nil {
			payouts = append(payouts, payout)
		} else {
			panic(err.Error())
		}
	}

	preamble := "Report"
	postamble := fmt.Sprintf("The report was done at %v\n", time.Now())

	if outputToCsv == "" {
		writer = writers.NewStdOut(preamble, postamble, payoutHeadings)
	} else {
		if err, writer = writers.NewCSV(preamble, postamble, payoutHeadings, outputToCsv); err != nil {
			panic(err.Error())
		}
	}
	if err := writer.Write(payouts); err != nil {
		panic(err.Error())
	}
}

func fetchBytesFromURL(url string, client http.Client) ([]byte, error) {
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func fetchRules(url string, client http.Client) ([]Rule, error) {
	var rr RulesResponse

	b, err := fetchBytesFromURL(url, client)
	if err != nil {
		return nil, err
	}

	if err = json.NewDecoder(bytes.NewReader(b)).Decode(&rr); err != nil {
		return nil, err
	}

	return rr.Rules, nil
}

func fetchPeople(url string, client http.Client) ([]People, error) {
	var pr PeopleResponse

	b, err := fetchBytesFromURL(url, client)
	if err != nil {
		return nil, err
	}

	if err = json.NewDecoder(bytes.NewReader(b)).Decode(&pr); err != nil {
		return nil, err
	}

	return pr.People, nil
}

func getToday() time.Time {
	now := time.Now().UTC()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
}
