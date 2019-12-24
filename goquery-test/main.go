/**
 * Created by sunyh-vm on 2019/12/20
 * Description:
 */

package main

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func main() {
	ExampleScrapy()
}

func ExampleScrapy() {
	res, err := http.Get("https://www.tianyancha.com/getBrand/list.html?type=searchList&pn=2&name=大数据&originalId=b36db1368506")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("body > div > div.brand-search-table.pt10 > table > tbody > tr").Each(func(i int, selection *goquery.Selection) {
		nameSelector := selection.Find("td > table > tbody > tr > td:nth-child(2) > a")
		productName := nameSelector.Text()
		log.Println(productName)
		productLink, existLink := nameSelector.Attr("href")
		if existLink {
			log.Println(productLink)
		} else {
			log.Println("product link parser error")
		}

		establishDate := selection.Find("td:nth-child(3)").Text()
		log.Println(establishDate)

		introduction := selection.Find("td:nth-child(5) > div > div").Text()
		log.Println(introduction)
	})
}
