package main_test

import (
	"os"

	. "github.com/htmldrum/ds_solution"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DsSolution", func() {
	Describe("main.go", func() {
		Context("When given no additional CLI flags", func() {
			Describe("It prints a report of payments to stdout", func() {
				It("Runs without error", func() {
					PrintReport(*RulesUrlArg, *PeopleUrlArg, *OutputToCsvArg)
				})
			})
		})
		Context("When given the -o argument with a valid path", func() {
			Describe("It prints a report of payments in CSV format to that path", func() {
				It("Creates the csv file", func() {
					csv_path := "results.csv"
					fi, _ := os.Stat(csv_path)
					Expect(fi).To(BeNil())

					PrintReport(*RulesUrlArg, *PeopleUrlArg, csv_path)

					new_fi, _ := os.Stat(csv_path)
					Expect(new_fi).NotTo(BeNil())

					remove_err := os.Remove(csv_path)
					Expect(remove_err).To(BeNil())
				})
			})
		})
		Context("When given the -rulesUrlArg argument with a valid people file", func() {
			Describe("It prints a report of payments to stdout", func() {

				It("Runs without error", func() {
					rules_url := "https://raw.githubusercontent.com/DoneSafe/code_test/master/rules.json?q=equal_string"
					PrintReport(rules_url, *PeopleUrlArg, *OutputToCsvArg)
				})
			})
		})
		Context("When given the -rulesUrlArg argument with a valid rules file", func() {
			Describe("It prints a report of payments to stdout", func() {
				It("Runs without error", func() {
					people_url := "https://raw.githubusercontent.com/DoneSafe/code_test/master/people.json?should_not_matter"
					PrintReport(*RulesUrlArg, people_url, *OutputToCsvArg)
				})
			})
		})
	})
})
