package wiki

import (
	"regexp"
	"io/ioutil"
	"os"
)

var reTitleIllegal = regexp.MustCompile("[^a-zA-Z0-9()]+")

type page struct {
	title string
	body  string
}

func safeTitle(title string) string {
	return reTitleIllegal.ReplaceAllString(title, "")
}

func filename(title string) string { return "data/" + title + ".txt" }

func makePage(title string, body string) *page {
	return &page{title: safeTitle(title), body: body}
}

func loadPage(title string) (*page, os.Error) {
	title = safeTitle(title)
	body, err := ioutil.ReadFile(filename(title))
	if err != nil {
		return nil, err
	}
	return &page{title: title, body: string(body)}, nil
}

func (p *page) save() os.Error {
	p.title = safeTitle(p.title)
	return ioutil.WriteFile(filename(p.title), []byte(p.body), 0600)
}
