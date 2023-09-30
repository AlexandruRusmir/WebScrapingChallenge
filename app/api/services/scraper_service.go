package services

import (
	"api/util"
	"errors"

	"github.com/PuerkitoBio/goquery"
)

func GetDefaultTags() []string {
	return []string{
		"a", "abbr", "address", "article", "aside", "b", "blockquote", "body", "button",
		"caption", "cite", "code", "div", "dl", "dt", "dd", "em", "figcaption", "figure",
		"footer", "form", "h1", "h2", "h3", "h4", "h5", "h6", "header", "i", "label", "li",
		"main", "nav", "ol", "p", "pre", "q", "s", "section", "select", "small", "span",
		"strong", "table", "tbody", "td", "tfoot", "th", "thead", "time", "tr", "u", "ul", "img"}
}

func handleImgTag(item *goquery.Selection, data map[string]interface{}) {
	imgAttributes := make(map[string]string)
	src, srcExists := item.Attr("src")
	if srcExists {
		imgAttributes["src"] = src
	}

	alt, altExists := item.Attr("alt")
	if altExists {
		imgAttributes["alt"] = alt
	}

	if srcExists || altExists {
		imgData := data["img"].([]map[string]string)
		imgData = append(imgData, imgAttributes)
		data["img"] = imgData
	}
}

func handleATag(item *goquery.Selection, data map[string]interface{}) {
	href, exists := item.Attr("href")
	if exists {
		linkData := data["a"].([]map[string]string)
		linkData = append(linkData, map[string]string{
			"href": href,
			"text": item.Text(),
		})
		data["a"] = linkData
	}
}

func handleDefaultTag(tag string, item *goquery.Selection, data map[string]interface{}) error {
	defaultData, ok := data[tag].([]string)
	if !ok {
		return errors.New("value in data map is not of type []string")
	}
	defaultData = append(defaultData, item.Text())
	data[tag] = defaultData

	return nil
}

func GetWordCountFromDocument(url string) (int, error) {
	doc, err := util.FetchDocument(url)
	if err != nil {
		return 0, err
	}

	bodyContent := doc.Find("body").Text()

	tokens := util.Tokenize(bodyContent)

	return len(tokens), nil
}

func ScrapeContentWithWordCount(url string, tags []string) (map[string]interface{}, error) {
	contentData, err := ScrapeContent(url, tags)
	if err != nil {
		return nil, err
	}

	wordCount, err := GetWordCountFromDocument(url)
	if err != nil {
		return nil, err
	}

	result := map[string]interface{}{
		"wordCount": wordCount,
		"content":   contentData,
	}

	return result, nil
}

func ScrapeContent(url string, tags []string) (map[string]interface{}, error) {
	if len(tags) == 0 {
		tags = GetDefaultTags()
	}

	doc, err := util.FetchDocument(url)
	if err != nil {
		return nil, err
	}

	data := make(map[string]interface{})

	for _, tag := range tags {
		if _, exists := data[tag]; !exists {
			switch tag {
			case "img", "a":
				data[tag] = []map[string]string{}
			default:
				data[tag] = []string{}
			}
		}

		var err error
		doc.Find(tag).Each(func(index int, item *goquery.Selection) {
			switch tag {
			case "img":
				handleImgTag(item, data)
			case "a":
				handleATag(item, data)
			default:
				err = handleDefaultTag(tag, item, data)
			}
			if err != nil {
				return
			}
		})

		if err != nil {
			return nil, err
		}
	}

	return data, nil

}
