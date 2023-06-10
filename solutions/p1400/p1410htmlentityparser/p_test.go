package p1410htmlentityparser

import "strings"

func entityParser(text string) string {
	r := strings.NewReplacer(
		"&quot;", "\"",
		"&apos;", "'",
		"&gt;", ">",
		"&lt;", "<",
		"&frasl;", "/",
	)
	r2 := strings.NewReplacer("&amp;", "&")
	text = r.Replace(text)
	text = r2.Replace(text)
	return text
}
