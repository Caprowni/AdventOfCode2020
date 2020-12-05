package main

import (
	"fmt"
	"strings"
)

func main() {

	acceptedFields := []string{"ecl", "pid", "eyr", "hcl", "byr", "iyr", "cid", "hgt"}

	passport := []string{
		"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm",
		"iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884 hcl:#cfa07d byr:1929",
		"hcl:#ae17e1 iyr:2013 eyr:2024 ecl:brn pid:760753108 byr:1931 hgt:179cm",
		"hcl:#cfa07d eyr:2025 pid:166559648 iyr:2011 ecl:brn hgt:59in"}
	valid := 0
	for i := 0; i < len(passport); i++ {
		var missingFields []string
		for _, field := range acceptedFields {
			if !strings.Contains(passport[i], field) {
				missingFields = append(missingFields, field)
			}
		}
		if len(missingFields) > 0 {
			if len(missingFields) == 1 && missingFields[0] == "cid" {
				valid++
			}
		} else {
			valid++
		}
	}
	fmt.Println(valid)

}
