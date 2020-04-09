package zhenai

import (
	"errors"
	"learn_project/crawler/engine"
	"learn_project/crawler/model"
	"regexp"
)

var ageRe = regexp.MustCompile(`<td><span class="label">年龄: </span>([\d]+)岁</td>`)

func ParseProfile(contents []byte) engine.ParseResult{
	profile := model.Profile{}
	age, err := extractString(contents, ageRe)
	if err != nil {
		profile.Age = age
	}
	return engine.ParseResult{
		Items: []interface{}{profile},
	}
}

func extractString(contents []byte, re *regexp.Regexp) (string, error){
	match := re.FindSubmatch(contents)

	if len(match) >=2 {
		return string(match[1]), nil
	} else {
		return "", errors.New("wrong match")
	}
}