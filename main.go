package main

import (
	"flag"
)

func main() {
	txtFilename := flag.String("txt", "1689007675141_numbers.txt", "a txt file in the format of 'phone numbers list'")
	flag.Parse()

	searchPhoneNum(*txtFilename)
	rename := "1689007676028_text.txt"
	txtFilename = &rename
	textSearch(*txtFilename)
}
