package parser

import (
	"github.com/antchfx/htmlquery"
	"github.com/crawler/engine"
	"github.com/crawler/model"
	"strconv"
	"strings"
)

func ParseRentDetail(contents []byte, url string) engine.ParseResult {
	doc, _ := htmlquery.Parse(strings.NewReader(string(contents)))
	attribute := model.Attribute{}

	name := htmlquery.Find(doc,"/html/body/div[3]/div[1]/div[2]/div[2]/div[1]/h1")
	attribute.Name= name[0].FirstChild.Data

	price := htmlquery.Find(doc,"/html/body/div[3]/div[1]/div[2]/div[2]/div[3]/div[2]/div[2]/div")
	if len(price)>0{
		p := price[0].FirstChild.Data
		p = strings.ReplaceAll(p," ","")
		p = strings.ReplaceAll(p,"\n","")
		Price, _ := strconv.Atoi(p)
		attribute.Price = Price
	}else {
		attribute.Price = 0
	}
	area := htmlquery.Find(doc,"/html/body/div[3]/div[1]/div[2]/div[2]/div[4]/div[1]/div[1]/label")
	attribute.Area = area[0].FirstChild.Data

	number := htmlquery.Find(doc,"/html/body/div[3]/div[1]/div[2]/div[2]/div[4]/div[1]/div[2]/label")
	 attribute.Number= number[0].FirstChild.Data

	structure := htmlquery.Find(doc, "/html/body/div[3]/div[1]/div[2]/div[2]/div[4]/div[1]/div[3]/label")
	Structure := structure[0].FirstChild.Data
	Structure = strings.ReplaceAll(Structure," ","")
	Structure = strings.ReplaceAll(Structure,"\n","")
	attribute.Structure = Structure

	pay := htmlquery.Find(doc, "/html/body/div[3]/div[1]/div[2]/div[2]/div[4]/div[1]/div[4]/label/a")
	attribute.Pay= pay[0].FirstChild.Data

	orientation := htmlquery.Find(doc,"/html/body/div[3]/div[1]/div[2]/div[2]/div[4]/div[2]/div[1]/label")
	attribute.Orientation=orientation[0].FirstChild.Data

	floor := htmlquery.Find(doc,"/html/body/div[3]/div[1]/div[2]/div[2]/div[4]/div[2]/div[2]/label")
	attribute.Floor= floor[0].FirstChild.Data

	region := htmlquery.Find(doc,"/html/body/div[3]/div[1]/div[2]/div[2]/div[4]/div[2]/div[3]/label/div")
	attribute.Region= htmlquery.SelectAttr(region[0],"title")

	metro := htmlquery.Find(doc,"/html/body/div[3]/div[1]/div[2]/div[2]/div[4]/div[2]/div[4]/label")
	attribute.Metro= htmlquery.SelectAttr(metro[0],"title")

	attribute.Url = url

	result := engine.ParseResult{
		Requests: nil,
		Items:    []interface{}{attribute},
	}
	return result
}
