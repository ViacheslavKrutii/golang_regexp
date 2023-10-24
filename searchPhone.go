package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func searchPhoneNum(src string) {
	file, err := os.Open(src)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	r, err := regexp.Compile(`^\(?[\d]{3}\)?[ -.]?[\d]{3}[ -.]?[\d]{4}$`)

	if err != nil {
		log.Fatal(err)
	}

	for scanner.Scan() {
		if r.MatchString(scanner.Text()) {
			fmt.Println(scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
