package erogamescape

import (
	"ckrg"
	"time"
)

const (
	id   = "erogamescape"
	name = "ErogameScape"
	url  = "https://erogamescape.dyndns.org/~ap2/ero/toukei_kaiseki/"
)

func init() {
	ckrg.RegisterPlugin(id, name, ckrg.CheckGetTimeout(name, url, 2*time.Second))
}
