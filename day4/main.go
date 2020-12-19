package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
)

/*
--- Day 4: Passport Processing ---
You arrive at the airport only to realize that you grabbed your North Pole Credentials instead of your passport. While these documents are extremely similar, North Pole Credentials aren't issued by a country and therefore aren't actually valid documentation for travel in most of the world.

It seems like you're not the only one having problems, though; a very long line has formed for the automatic passport scanners, and the delay could upset your travel itinerary.

Due to some questionable network security, you realize you might be able to solve both of these problems at the same time.

The automatic passport scanners are slow because they're having trouble detecting which passports have all required fields. The expected fields are as follows:

byr (Birth Year)
iyr (Issue Year)
eyr (Expiration Year)
hgt (Height)
hcl (Hair Color)
ecl (Eye Color)
pid (Passport ID)
cid (Country ID)
Passport data is validated in batch files (your puzzle input). Each passport is represented as a sequence of key:value pairs separated by spaces or newlines. Passports are separated by blank lines.

Here is an example batch file containing four passports:

ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in
The first passport is valid - all eight fields are present. The second passport is invalid - it is missing hgt (the Height field).

The third passport is interesting; the only missing field is cid, so it looks like data from North Pole Credentials, not a passport at all! Surely, nobody would mind if you made the system temporarily ignore missing cid fields. Treat this "passport" as valid.

The fourth passport is missing two fields, cid and byr. Missing cid is fine, but missing any other field is not, so this passport is invalid.

According to the above rules, your improved system would report 2 valid passports.

Count the number of valid passports - those that have all required fields. Treat cid as optional. In your batch file, how many passports are valid?
*/

type Passport struct {
	byr *string // (Birth Year)
	iyr *string // (Issue Year)
	eyr *string // (Expiration Year)
	hgt *string // (Height)
	hcl *string // (Hair Color)
	ecl *string // (Eye Color)
	pid *string // (Passport ID)
	cid *string // (Country ID)
}

func (p *Passport) isValid() bool {
	return p.byr != nil &&
		p.iyr != nil &&
		p.eyr != nil &&
		p.hgt != nil &&
		p.hcl != nil &&
		p.ecl != nil &&
		p.pid != nil
}

// field is in format key:val
func (p *Passport) applyField(field []byte) error {
	fieldParts := bytes.Split(field, []byte(":"))

	if len(fieldParts) != 2 {
		return errors.New("Field should be in form of key:value")
	}

	switch string(fieldParts[0]) {
	case "byr":
		p.byr = stringPtr(fieldParts[1])
	case "iyr":
		p.iyr = stringPtr(fieldParts[1])
	case "eyr":
		p.eyr = stringPtr(fieldParts[1])
	case "hgt":
		p.hgt = stringPtr(fieldParts[1])
	case "hcl":
		p.hcl = stringPtr(fieldParts[1])
	case "ecl":
		p.ecl = stringPtr(fieldParts[1])
	case "pid":
		p.pid = stringPtr(fieldParts[1])
	case "cid":
		p.cid = stringPtr(fieldParts[1])
	default:
		return errors.New("Key is invalid format")
	}
	return nil
}

func stringPtr(b []byte) *string {
	s := string(b)
	return &s
}

func parsePassports() ([]*Passport, error) {
	var passports []*Passport
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	//passports split by newlines
	passportSections := bytes.Split(content, []byte("\n\n"))
	for _, passport := range passportSections {
		p := &Passport{}
		passportParts := bytes.Fields(passport)
		for _, part := range passportParts {
			if err := p.applyField(part); err != nil {
				return nil, err
			}
		}
		passports = append(passports, p)
	}

	return passports, nil
}

func part1() {
	passports, err := parsePassports()
	if err != nil {
		log.Fatal(err)
	}
	validPassports := 0
	for _, passport := range passports {
		if passport.isValid() {
			validPassports += 1
		}
	}
	fmt.Println(validPassports)
}

func main() {
	part1()
}
