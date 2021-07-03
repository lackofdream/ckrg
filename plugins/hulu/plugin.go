package niconico

import (
	. "ckrg"
)

const (
	id     = "hulu"
	name   = "Hulu"
	url    = "https://id.hulu.jp"
	substr = "login"
)

func init() {
	RegisterPlugin(id, name, CheckGetSubstring(url, substr, false))
}
