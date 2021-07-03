package main

import (
	. "ckrg"
	_ "ckrg/plugins/kancolle"
	_ "ckrg/plugins/niconico"
	_ "ckrg/plugins/pcrjp"
	"fmt"
)

func main() {
	//ip, _ := GetSelfIP()
	//isp, _ := GetLocalISP(ip)
	for _, p := range Plugins {
		res, _ := p.Check()
		fmt.Printf("%s\t%s\n", p.Name, func(r bool) string {
			if r {
				return "No"
			}
			return "Yes"
		}(res.Restricted))
	}
}
