package utils

import (
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

// MarkdownToHTML 将markdown 转换为 html
func MarkdownToHTML(md string) string {
	myHTMLFlags := 0 |
		blackfriday.UseXHTML |
		blackfriday.Smartypants |
		blackfriday.SmartypantsFractions |
		blackfriday.SmartypantsDashes |
		blackfriday.SmartypantsLatexDashes

	myExtensions := 0 |
		blackfriday.NoIntraEmphasis |
		blackfriday.Tables |
		blackfriday.FencedCode |
		blackfriday.Autolink |
		blackfriday.Strikethrough |
		blackfriday.SpaceHeadings |
		blackfriday.HeadingIDs |
		blackfriday.BackslashLineBreak |
		blackfriday.DefinitionLists |
		blackfriday.HardLineBreak

	htmlRendererParameters := blackfriday.HTMLRendererParameters{
		Flags: myHTMLFlags,
	}

	renderer := blackfriday.NewHTMLRenderer(htmlRendererParameters)

	options := blackfriday.WithRenderer(renderer)
	options = blackfriday.WithExtensions(myExtensions)

	bytes := blackfriday.Run([]byte(md), options)
	theHTML := string(bytes)
	return bluemonday.UGCPolicy().Sanitize(theHTML)
}

// AvoidXSS 避免XSS
func AvoidXSS(theHTML string) string {
	return bluemonday.UGCPolicy().Sanitize(theHTML)
}
