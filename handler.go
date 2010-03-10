package wiki

import (
	"fmt"
	"io/ioutil"
	"regexp"
	web "github.com/hoisie/web.go"
	"template"
)

var urlPrefix string

var linkRe, titleRe *regexp.Regexp

func init() {
	linkRe = regexp.MustCompile("\\[[a-zA-Z0-9]+\\]")
	titleRe = regexp.MustCompile("[^a-zA-Z0-9]+")
}

func viewHandler(ctx *web.Context, title string) {
	page, err := loadPage(title)
	if err != nil {
		redirect(ctx, "edit", title)
		return
	}
	renderTmpl(ctx, "view", page.title, makeLinks(page.body))
}

func editHandler(ctx *web.Context, title string) {
	page, err := loadPage(title)
	if err != nil {
		page = makePage(title, "")
	}
	renderTmpl(ctx, "edit", page.title, page.body)
}

func saveHandler(ctx *web.Context, title string) {
	body, ok := ctx.Request.Params["body"]
	if !ok {
		ctx.Abort(500, "No body supplied.")
		return
	}
	page := makePage(title, body[0])
	page.save()
	redirect(ctx, "view", title)
}

func renderTmpl(ctx *web.Context, tmpl, title, body string) {
	d := map[string]string{
		"prefix": urlPrefix,
		"title":  title,
		"body":   body,
	}
	filename := "tmpl/"+tmpl+".html"
	html, err := ioutil.ReadFile(filename)
	if err != nil {
		ctx.Abort(500, "Unable to Read template file: "+filename)
		return
	}
	t, err := template.Parse(string(html), nil)
	if err != nil {
		ctx.Abort(500, "Unable to Parse template file: "+filename)
		return
	}
	err = t.Execute(d, ctx)
	if err != nil {
		ctx.Abort(500, "Unable to Execute template")
	}
}

func redirect(ctx *web.Context, handler, title string) {
	ctx.Redirect(302, urlPrefix+handler+"/"+safeTitle(title))
}

func makeLinks(body string) string {
	return linkRe.ReplaceAllStringFunc(body, func(match string) string {
		inner := match[1 : len(match)-1]
		return fmt.Sprintf("<a href=\"/view/%s\">%s</a>", inner, inner)
	})
}

// prefix should be something like "/" or "/wiki/"
func RegisterHandlers(prefix string) {
	urlPrefix = prefix
	web.Get(urlPrefix, func(ctx *web.Context) { redirect(ctx, "view", "FrontPage") })
	web.Get(urlPrefix+"view/(.+)", viewHandler)
	web.Get(urlPrefix+"edit/(.+)", editHandler)
	web.Post(urlPrefix+"save/(.+)", saveHandler)
}
