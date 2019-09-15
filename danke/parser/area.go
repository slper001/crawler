package parser

import (
	"github.com/antchfx/htmlquery"
	"github.com/crawler/engine"
	"strings"
)

func ParseArea(contents []byte)  engine.ParseResult{
	doc, _ := htmlquery.Parse(strings.NewReader(string(contents)))
	result := engine.ParseResult{}

	z := htmlquery.Find(doc,"/html/body/div[3]/div/div[5]/a[2]/em")
	zone := z[0].FirstChild.Data

	for _,v := range htmlquery.Find(doc,"/html/body/div[3]/div/div[7]/div[2]/div[1]/div/div[2]/a"){
		result.Items = append(result.Items,zone)
		result.Requests = append(result.Requests,engine.Request{
			Url:        htmlquery.SelectAttr(v,"href"),
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseRentDetail(c,htmlquery.SelectAttr(v,"href"))
			},
		})
	}
	return result
}
