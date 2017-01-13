package main

import (
	"archive/zip"
	"bytes"
	"compress/gzip"
	"encoding/xml"
	"fmt"
	"os"
	"regexp"
)

func main() {
	TestName("sitemap.2.xml")
	TestName("sitemap..xml")
	TestName("sitemap.1.xml.gz")

	sm := NewSitemap([]string{"gawgwa", "gwweag"})
	b := makeXML(sm)
	gb := makeGzip(b)
	save(gb, "sitemap.xml.gz")

	makeZip()
}

const sitemapXmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"

type URLSet struct {
	XMLName xml.Name `xml:"urlset"`
	Xmlns   string   `xml:"xmlns,attr"`
	URLs    []*URL   `xml:"url"`
}

type URL struct {
	Loc       string  `xml:"loc"`
	ChageFreq string  `xml:"changefreq"`
	Priority  float32 `xml:"priority"`
}

func makeXML(sm *URLSet) *bytes.Buffer {
	var buf bytes.Buffer
	header := []byte(xml.Header)
	body, _ := xml.MarshalIndent(sm, "", " ")
	buf.Write(header)
	buf.Write(body)
	return &buf
}
func TestName(str string) {
	reg := regexp.MustCompile(`^sitemap(?:\.[\d]+\.{1}|\.{1}){1}xml(?:\.gz)?$`)
	if reg.MatchString(str) {
		fmt.Printf("Matched: %s \n", str)
	} else {
		fmt.Printf("UNMatched: %s \n", str)
	}
}
func NewSitemap(urlList []string) *URLSet {
	urls := make([]*URL, 0, len(urlList))
	for _, v := range urlList {
		var url URL
		url.Loc = v
		urls = append(urls, &url)
	}
	return &URLSet{
		Xmlns: sitemapXmlns,
		URLs:  urls,
	}
}

func save(b *bytes.Buffer, filename string) {
	fw, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	fw.Write(b.Bytes())
	fw.Close()
}

func makeGzip(src *bytes.Buffer) *bytes.Buffer {
	var buf bytes.Buffer
	gw := gzip.NewWriter(&buf)

	if _, err := gw.Write(src.Bytes()); err != nil {
		panic(err)
	}
	if err := gw.Flush(); err != nil {
		panic(err)
	}
	if err := gw.Close(); err != nil {
		panic(err)
	}
	return &buf
}
func makeZip() {

	zipFile, err := os.Create("sample.zip")
	if err != nil {
		panic(err)
	}
	defer zipFile.Close()
	archive := zip.NewWriter(zipFile)
	defer archive.Close()

	f, err := archive.Create("test.txt")
	if err != nil {
		panic(err)
	}
	_, err = f.Write([]byte("text \n text"))
	if err != nil {
		panic(err)
	}
	return
}
