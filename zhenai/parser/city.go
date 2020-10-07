package parser

import (
	"goSetUp/crawler/engine"
	"regexp"
)

var cityRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)

func ParseCity(contents []byte, _ string) engine.ParseResult {

	matches := cityRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		name := string(m[2])
		//result.Items = append(result.Items, "User: "+name)
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ProfileParser(name),
		})
	}
	return result
}

func ProfileParser(name string) engine.ParserFunc {
	return func(c []byte, url string) engine.ParseResult {
		return ParseProfile(c, url, name)
	}
}
