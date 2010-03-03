
package wiki

import (
	"mustache"
	"web"
)

var urlPrefix string

func viewHandler(ctx *web.Context, title string) {
	page, err := loadPage(title)
	if err != nil {
		redirect(ctx, "edit", title)
		return
	}
	renderTmpl(ctx, "view", page)
}

func editHandler(ctx *web.Context, title string) {
	page, err := loadPage(title)
	if err != nil {
		page = makePage(title, "")
	}
	renderTmpl(ctx, "edit", page)
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

func renderTmpl(ctx *web.Context, tmpl string, page *page) {
	d := map[string]string{
		"prefix": urlPrefix,
		"title": page.title,
		"body": page.body,
	}
	content, _ := mustache.RenderFile("tmpl/"+tmpl+".mustache", d)
	ctx.WriteString(content)
}

func redirect(ctx *web.Context, handler, title string) {
	ctx.Redirect(302, urlPrefix + handler + "/" + safeTitle(title))
}

// prefix should be something like "/" or "/wiki/"
func RegisterHandlers(prefix string) {
	urlPrefix = prefix
	web.Get(urlPrefix, func(ctx *web.Context) {
		redirect(ctx, "view", "FrontPage")
	})
	web.Get(urlPrefix + "view/(.*)", viewHandler)
	web.Get(urlPrefix + "edit/(.*)", editHandler)
	web.Post(urlPrefix + "save/(.*)", saveHandler)
}

