package main

import (
	"net/http"
	"testing"

	"github.com/dnaeon/go-vcr/recorder"
)

func TestHelloWorld(t *testing.T) {}

func TestFetchRulesWithVCR(t *testing.T) {
	r, err := recorder.New("fixtures/rules")
	if err != nil {
		t.Fatal(err)
	}
	defer r.Stop()
	client := http.Client{
		Transport: r,
	}

	rules, err := fetchRules(*RulesUrlArg, client)
	if err != nil {
		t.Fatalf("Failed to fetch url %s", err)
	}

	exp := 5
	if len(rules) != exp {
		t.Errorf("Expected %d reports!. Got: %d", exp, len(rules))
	}

	for n, rule := range rules {
		if (rule == Rule{}) {
			t.Errorf("Rule #%d is serialized empty: %v", n, rule)
		}
	}
}

func TestFetchPeopleWithVCR(t *testing.T) {
	r, err := recorder.New("fixtures/people")
	if err != nil {
		t.Fatal(err)
	}
	defer r.Stop()
	client := http.Client{
		Transport: r,
	}

	people, err := fetchPeople(*PeopleUrlArg, client)
	if err != nil {
		t.Fatalf("Failed to fetch url %s", err)
	}

	exp := 4
	if len(people) != exp {
		t.Errorf("Expected %d people!. Got: %d", exp, len(people))
	}

	for n, people := range people {
		if people.InjuryDateS == "" {
			t.Errorf("People #%d is serialized empty: %v", n, people)
		}
	}
}
