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

--- Part Two ---
While it appears you validated the passwords correctly, they don't seem to be what the Official Toboggan Corporate Authentication System is expecting.

The shopkeeper suddenly realizes that he just accidentally explained the password policy rules from his old job at the sled rental place down the street! The Official Toboggan Corporate Policy actually works a little differently.

Each policy actually describes two positions in the password, where 1 means the first character, 2 means the second character, and so on. (Be careful; Toboggan Corporate Policies have no concept of "index zero"!) Exactly one of these positions must contain the given letter. Other occurrences of the letter are irrelevant for the purposes of policy enforcement.

Given the same example list from above:

1-3 a: abcde is valid: position 1 contains a and position 3 does not.
1-3 b: cdefg is invalid: neither position 1 nor position 3 contains b.
2-9 c: ccccccccc is invalid: both position 2 and position 9 contain c.
How many passwords are valid according to the new interpretation of the policies?
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

type passwordPolicy2 struct {
	char        byte
	firstIndex  int
	secondIndex int
}

func (p *passwordPolicy2) isValidPassword2(password string) bool {
	firstIndexMatch := password[p.firstIndex-1] == p.char
	secondIndexMatch := password[p.secondIndex-1] == p.char
	return (firstIndexMatch || secondIndexMatch) && !(firstIndexMatch && secondIndexMatch)
}

func part1() {
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

func part2() {
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
		firstIndex, err := strconv.Atoi(policyBounds[0])
		if err != nil {
			panic("Input malformed")
		}
		secondIndex, err := strconv.Atoi(policyBounds[1])
		if err != nil {
			panic("Input malformed")
		}

		policy := &passwordPolicy2{
			char:        policyParts[1][0],
			firstIndex:  firstIndex,
			secondIndex: secondIndex,
		}

		// password after the ":" has a space before it starts
		password := inputParts[1][1:]

		if policy.isValidPassword2(password) {
			validPasswords++
		}
	}
	fmt.Println(validPasswords)
}

func main() {
	part1()
	part2()
}
