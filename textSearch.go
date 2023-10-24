package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func textSearch(src string) {
	file, err := os.Open(src)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	capitalizedWords := `[Є-ЯҐ][а-їґ]+`
	r, err := regexp.Compile(capitalizedWords)

	if err != nil {
		log.Fatal(err)
	}

	for scanner.Scan() {
		find := r.FindAll([]byte(scanner.Text()), -1)
		for _, v := range find {
			fmt.Println(string(v))
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
