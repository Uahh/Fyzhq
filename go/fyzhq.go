package fyzhq

import (
	"fmt"
	"os"

	"github.com/mozillazg/go-pinyin"
	"github.com/tidwall/gjson"
)

var worddict gjson.Result

func init() {
	data, err := os.ReadFile("../data/language.json")
	if err != nil {
		fmt.Println(err)
	}
	worddict = gjson.ParseBytes(data)
}

func Transform(text string, lan string) (result string) {
	p := pinyin.Pinyin(text, pinyin.NewArgs())
	if len(p) == 0 {
		return text
	}
	switch lan {
	case "jp":
		for _, v := range p {
			result += worddict.Get("Japanese." + v[0]).String()
		}
	case "en":
		for _, v := range p {
			result += worddict.Get("English." + v[0]).String()
		}
	case "fr":
		for _, v := range p {
			result += worddict.Get("French." + v[0]).String()
		}
	case "gm":
		for _, v := range p {
			result += worddict.Get("German." + v[0]).String()
		}
	case "ru":
		for _, v := range p {
			result += worddict.Get("Russian." + v[0]).String()
		}
	case "kr":
		for _, v := range p {
			result += worddict.Get("Korean." + v[0]).String()
		}
	case "th":
		for _, v := range p {
			result += worddict.Get("Thai." + v[0]).String()
		}
	default:
		result = text
	}
	return
}
