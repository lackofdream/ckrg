package pcrjp

import (
	. "ckrg"
)

const (
	id   = "kancolle"
	name = "艦隊これくしょん"
	url  = "http://203.104.209.7/kcscontents/"
)

func init() {
	RegisterPlugin(id, name, CheckGetCode(url, map[int]bool{200: false, 403: true}, ""))
}
