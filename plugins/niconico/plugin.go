package niconico

import (
	. "ckrg"
)

const (
	id     = "niconico"
	name   = "ニコニコ"
	url    = "https://www.nicovideo.jp/watch/so23017073"
	substr = "同じ地域"
)

func init() {
	RegisterPlugin(id, name, CheckGetSubstring(url, substr, true))
}
