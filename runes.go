package golib

import "unicode"

// [Unicode字符类 ](https://blog.csdn.net/weixin_36082485/article/details/53154065)
// IsMark determines whether the rune is a marker
func IsMark(r rune) bool {
	return unicode.Is(unicode.Mn, r) || unicode.Is(unicode.Me, r) || unicode.Is(unicode.Mc, r)
}
