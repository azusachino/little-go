package zhenai

import (
	"learn_project/crawler/engine"
	"regexp"
)

const cityListRe = `<a target="_blank" href="(http://www.zhenai.com/zhenhun/[a-zA-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParseFunc: engine.NilParser,
		})
	}
	return result
}
