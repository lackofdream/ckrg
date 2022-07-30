package main

import (
	. "ckrg"
	_ "ckrg/plugins/erogamescape"
	_ "ckrg/plugins/kancolle"
	_ "ckrg/plugins/niconico"
	_ "ckrg/plugins/pcrjp"
	"fmt"
	"net/http"
	"os"
	"sort"
	"sync"
	"time"
)

func showIPAndISP() {
	ip, _ := GetSelfIP()
	isp, _ := GetLocalISP(ip)
	fmt.Printf("IP: %s\nISP: %s\n", ip, isp)
}

func check(plugin Plugin, resCh chan<- Result) {
	fmt.Printf("checking %s\n", plugin.Name)
	res, err := plugin.Check()
	fmt.Printf("%s checked\n", plugin.Name)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s\terror: %s\n", plugin.Name, err.Error())
	}
	resCh <- res
}

func receiveAndSortResult(resCh <-chan Result, wg *sync.WaitGroup) []string {
	var ret []string
	for res := range resCh {
		if res.Ok {
			ret = append(ret,
				fmt.Sprintf("%s\t%s\n", res.Name, func() string {
					if res.Restricted {
						return "No"
					}
					return "Yes"
				}()))
		}
		wg.Done()
	}
	sort.Strings(ret)
	return ret
}

func main() {
	http.DefaultClient.Timeout = 5 * time.Second
	showIPAndISP()
	wg := sync.WaitGroup{}
	resCh := make(chan Result, 8)
	for _, p := range Plugins {
		wg.Add(1)
		go check(p, resCh)
	}
	go func() {
		wg.Wait()
		close(resCh)
	}()
	res := receiveAndSortResult(resCh, &wg)
	for _, s := range res {
		fmt.Print(s)
	}
}
