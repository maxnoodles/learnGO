package parser

import (
	"learnGo/crawler/engine"
	"learnGo/crawler/model"
	"regexp"
	"strconv"
)

var ageRe = regexp.MustCompile(`<div.*?>(\d+)岁</div>`)
var marriageRe = regexp.MustCompile(
	`<div.*?>(未婚|离异|丧偶|)</div>`)
var heightRe = regexp.MustCompile(`<div.*?>(\d+)cm</div>`)
var weightRe = regexp.MustCompile(`<div.*?>(\d+)kg</div>`)
var incomeRe = regexp.MustCompile(`<div.*?>月收入:(.*?)</div>`)
var genderRe = regexp.MustCompile(`<a.*?>.*?([男女]).*?</a>`)
var xinzuoRe = regexp.MustCompile(`<div.*?>(.*?)座</div>`)

//var educationRe = regexp.MustCompile(`<div.*?>(\d+)cm</div>`)
//var occupationRe = regexp.MustCompile(`<div.*?>(\d+)cm</div>`)
//var hokouRe = regexp.MustCompile(`<div.*?>(\d+)cm</div>`)
//var houseRe = regexp.MustCompile(`<div.*?>(\d+)cm</div>`)
//var carRe = regexp.MustCompile(`<div.*?>(\d+)cm</div>`)

// 获取用户的详细资料
func ParseProfile(contents []byte, name string) engine.ParseResult {
	profile := model.Profile{}
	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err == nil {
		profile.Age = age
	}

	profile.Marriage = extractString(contents, marriageRe)

	height, err := strconv.Atoi(extractString(contents, heightRe))
	if err == nil {
		profile.Height = height
	}

	weight, err := strconv.Atoi(extractString(contents, weightRe))
	if err == nil {
		profile.Weight = weight
	}

	profile.Income = extractString(contents, incomeRe)
	profile.Gender = extractString(contents, genderRe)
	profile.Xinzuo = extractString(contents, xinzuoRe)
	profile.Marriage = extractString(contents, marriageRe)
	//profile.Education = extractString(contents, educationRe)
	//profile.Occupation = extractString(contents, occupationRe)
	//profile.Hokou = extractString(contents, hokouRe)
	//profile.House = extractString(contents, houseRe)
	//profile.Car = extractString(contents, carRe)

	result := engine.ParseResult{
		Items: []interface{}{profile},
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
