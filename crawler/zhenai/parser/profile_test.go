package parser

import (
	"fmt"
	"io/ioutil"
	"learnGo/crawler/model"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParseProfile(contents)
	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 element; but was %v", result.Items)
	}
	fmt.Println(result.Items[0])
	profile := result.Items[0].(model.Profile)
	expected := model.Profile{
		Name:     "简单就好",
		Age:      28,
		Height:   175,
		Weight:   70,
		Income:   "5-8千",
		Gender:   "男",
		Xinzuo:   "射手座",
		Marriage: "未婚",
	}
	if profile != expected {
		t.Errorf("expected %v, but was %v", expected, profile)
	}
}
