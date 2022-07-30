package pcrjp

import (
	. "ckrg"
)

const (
	id   = "kancolle"
	name = "艦隊これくしょん"
	url  = "http://203.104.209.7/gadget_html5/js/kcs_const.js"
)

func init() {
	RegisterPlugin(id, name, CheckGetCode(name, url, map[int]bool{200: false, 403: true, 503: true}, ""))
}
