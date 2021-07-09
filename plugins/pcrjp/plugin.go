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
	RegisterPlugin(id, name, CheckGetCode(name, url, map[int]bool{404: false, 403: true}, UaDalvik))
}
