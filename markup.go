package wiki
/*
 *  Filename:    markup.go
 *  Package:     wiki
 *  Author:      Bryan Matsuo <bmatsuo@soe.ucsc.edu>
 *  Created:     Mon Jul  4 05:49:19 PDT 2011
 *  Description: Markup wiki content so it can be markdown'ed.
 */
import (
    //"fmt"
    //"strings"
    "github.com/russross/blackfriday"
)

func markdown(title, content string) string {
	// set up options
	var extensions = 0
	extensions |= blackfriday.EXTENSION_NO_INTRA_EMPHASIS
	extensions |= blackfriday.EXTENSION_TABLES
	extensions |= blackfriday.EXTENSION_FENCED_CODE
	extensions |= blackfriday.EXTENSION_AUTOLINK
	extensions |= blackfriday.EXTENSION_STRIKETHROUGH
	extensions |= blackfriday.EXTENSION_SPACE_HEADERS

    var htmlflags = 0
    htmlflags |= blackfriday.HTML_USE_SMARTYPANTS
    htmlflags |= blackfriday.HTML_COMPLETE_PAGE
    htmlflags |= blackfriday.HTML_TOC

    var renderer = blackfriday.HtmlRenderer(htmlflags, title, "")
    var buff = make([]byte, len(content))
    copy(buff, content)
    return string(blackfriday.Markdown(buff, renderer, extensions))
}
