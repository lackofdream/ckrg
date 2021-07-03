package pcrjp

import (
	. "ckrg"
)

const (
	id   = "pcrjp"
	name = "プリンセスコネクト！Re:Dive"
	url  = "https://api-priconne-redive.cygames.jp/"
)

func init() {
	RegisterPlugin(id, name, CheckGetCode(url, map[int]bool{404: false, 403: true}, UA_Dalvik))
}
