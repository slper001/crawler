package parser

import (
	"github.com/antchfx/htmlquery"
	"github.com/crawler/engine"
	"strings"
)
/*
获得某个大的行政区的子区域列表信息
比如南山区
就获得南山区下的桃园、大新、侨城东。。。等子区域的url列表信息
*/

func ParseAreaList(contents []byte) engine.ParseResult {
	doc, _ := htmlquery.Parse(strings.NewReader(string(contents)))
	result := engine.ParseResult{}

	zone := 1   //用1代表南山区，依次递进其他区域

	///南山区对应于html/body/div[3]/div/div[4]/div[2]/dl[2]/dd/div/div[1]/div/a
	//福田区对应于/html/body/div[3]/div/div[4]/div[2]/dl[2]/dd/div/div[2]/div/a
	//罗湖区对应于/html/body/div[3]/div/div[4]/div[2]/dl[2]/dd/div/div[3]/div/a
	// 依次递进 宝安区div[4]，龙岗区div[5]，龙华区div[6]
	for _,v:= range htmlquery.Find(doc,"/html/body/div[3]/div/div[4]/div[2]/dl[2]/dd/div/div[1]/div/a"){
		result.Items = append(result.Items,zone)
		result.Requests = append(result.Requests,engine.Request{
			Url:        htmlquery.SelectAttr(v,"href"),
			ParserFunc: ParseArea,
		})
	}
	return result
}
