package parser

import (
	"goSetUp/crawler/engine"
	"goSetUp/crawler/model"
	"regexp"
	"strconv"
)

//Name       string
//	Age        int
//	Height     int
//	Weight     int
//	Income     string
//	Marriage   string
//	Origin     string
//var nameRe = regexp.MustCompile(`<h1 data-v-11592496="" class="nickName"[^>]*>([^<]+)</h1>`)
var ageRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([\d]+)岁</div>`)
var heightRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([\d]+)cm</div>`)
var weightRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([\d]+)kg</div>`)
var incomeRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>月收入:([^<]+)万</div>`)
var marriageRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([^<]+婚)</div>`)
var originRe = regexp.MustCompile(`<div class="m-btn pink"[^>]*>籍贯:([^<]+)</div>`)
var idRe = regexp.MustCompile(`https://album.zhenai.com/u/([\d]+)`)

func ParseProfile(contents []byte, url string, name string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name
	//profile.Name = extractString(contents, nameRe)
	if age, err := strconv.Atoi(extractString(contents, ageRe)); err == nil {
		profile.Age = age
	}
	if height, err := strconv.Atoi(extractString(contents, heightRe)); err == nil {
		profile.Height = height
	}
	if weight, err := strconv.Atoi(extractString(contents, weightRe)); err == nil {
		profile.Weight = weight
	}
	profile.Income = extractString(contents, incomeRe)
	marriage := extractString(contents, marriageRe)
	if marriage == "" {
		profile.Marriage = "离异"
	} else {
		profile.Marriage = marriage
	}

	profile.Origin = extractString(contents, originRe)
	result := engine.ParseResult{
		Items: []engine.Item{
			{
				Url:     url,
				Type:    "zhenai",
				Id:      extractString([]byte(url), idRe),
				Payload: profile,
			},
		},
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
