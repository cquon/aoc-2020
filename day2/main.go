package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
--- Day 2: Password Philosophy ---
Your flight departs in a few days from the coastal airport; the easiest way down to the coast from here is via toboggan.

The shopkeeper at the North Pole Toboggan Rental Shop is having a bad day. "Something's wrong with our computers; we can't log in!" You ask if you can take a look.

Their password database seems to be a little corrupted: some of the passwords wouldn't have been allowed by the Official Toboggan Corporate Policy that was in effect when they were chosen.

To try to debug the problem, they have created a list (your puzzle input) of passwords (according to the corrupted database) and the corporate policy when that password was set.

For example, suppose you have the following list:

1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc
Each line gives the password policy and then the password. The password policy indicates the lowest and highest number of times a given letter must appear for the password to be valid. For example, 1-3 a means that the password must contain a at least 1 time and at most 3 times.

In the above example, 2 passwords are valid. The middle password, cdefg, is not; it contains no instances of b, but needs at least 1. The first and third passwords are valid: they contain one a or nine c, both within the limits of their respective policies.

How many passwords are valid according to their policies?
*/

type passwordPolicy struct {
	char     byte
	minCount int
	maxCount int
}

func (p *passwordPolicy) isValidPassword(password string) bool {
	charCount := 0
	for i := 0; i < len(password); i++ {
		if password[i] == p.char {
			charCount++
		}
	}
	if charCount >= p.minCount && charCount <= p.maxCount {
		return true
	}
	return false
}

func main() {
	validPasswords := 0
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputLine := scanner.Text()
		inputParts := strings.Split(inputLine, ":")
		if len(inputParts) != 2 {
			panic("Input malformed")
		}

		policySection := inputParts[0]
		policyParts := strings.Split(policySection, " ")
		if len(inputParts) != 2 {
			panic("Input malformed")
		}
		policyBounds := strings.Split(policyParts[0], "-")
		if len(inputParts) != 2 {
			panic("Input malformed")
		}
		lowBound, err := strconv.Atoi(policyBounds[0])
		if err != nil {
			panic("Input malformed")
		}
		upperBound, err := strconv.Atoi(policyBounds[1])
		if err != nil {
			panic("Input malformed")
		}

		policy := &passwordPolicy{
			char:     policyParts[1][0],
			minCount: lowBound,
			maxCount: upperBound,
		}

		// password after the ":" has a space before it starts
		password := inputParts[1][1:]

		if policy.isValidPassword(password) {
			validPasswords++
		}
	}
	fmt.Println(validPasswords)
}
