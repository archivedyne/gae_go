package main

import (
	"fmt"
	"regexp"
)

func main() {
	testReg("sitemap.2.xml")
	testReg("sitemap..xml")
	testReg("sitemap.1.xml.gz")

	// Make xml.gzip file from struct
}

func testReg(str string) {
	reg := regexp.MustCompile(`^sitemap(?:\.[\d]+\.{1}|\.{1}){1}xml(?:\.gz)?$`)
	if reg.MatchString(str) {
		fmt.Printf("Matched: %s \n", str)
	} else {
		fmt.Printf("UNMatched: %s \n", str)
	}
}
