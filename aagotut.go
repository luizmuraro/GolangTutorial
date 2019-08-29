package main

// <sitemapindex xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
// 	<sitemap>
// 		<loc>
// 			https://www.washingtonpost.com/news-sitemaps/politics.xml
// 		</loc>
// 	</sitemap>
// 	<sitemap>
// 		<loc>
// 			https://www.washingtonpost.com/news-sitemaps/opinions.xml
// 		</loc>
// 	</sitemap>



import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/xml"
	"strings"
)

type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Titles []string `xml:"url>news>title"`
	Keywords []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

type NewsMap struct {
	Keywords string
	Location string
}

func main() {
	var s SitemapIndex
	var n News
	newsmap := make(map[string]NewsMap)
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &s)

	// fmt.Println(s.Locations)
	
	for _, Location := range s.Locations {
		Location = strings.TrimSpace(Location)
		resp, err := http.Get(Location)
		if err != nil {
			fmt.Println(err)
		}
		bytes, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes, &n)
		for idx, _ := range n.Titles {
			newsmap[n.Titles[idx]] = NewsMap{n.Keywords[idx], n.Locations[idx]}
		}
	}
	for idx, data := range newsmap {
		fmt.Println("\n\n\n", idx)
		fmt.Println("\n", data.Keywords)
		fmt.Println("\n", data.Location)
	}
}