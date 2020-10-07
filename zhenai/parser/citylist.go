package parser

import (
	"goSetUp/crawler/engine"
	"regexp"
)

const cityListRe = `<a href=\"(http://www.zhenai.com/zhenghun/[0-9a-z]+)\"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte, _ string) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		url := string(m[1])
		//result.Items = append(result.Items, "City: "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        url,
			ParserFunc: ParseCity,
		})
	}
	return result
}
