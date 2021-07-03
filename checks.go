package ckrg

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const UA_Dalvik = "Dalvik/2.1.0 (Linux; U; Android 9; ALP-AL00 Build/HUAWEIALP-AL00)"

func CheckGetCode(url string, codeToRestricted map[int]bool, ua string) func() (Result, error) {
	return func() (Result, error) {
		client := &http.Client{}
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return Result{}, err
		}
		if ua != "" {
			req.Header.Set("User-Agent", ua)
		}
		resp, err := client.Do(req)
		if err != nil {
			return Result{}, err
		}
		for k, v := range codeToRestricted {
			if resp.StatusCode == k {
				return Result{Restricted: v}, nil
			}
		}
		return Result{}, errors.New(fmt.Sprintf("unknown status: %d", resp.StatusCode))
	}
}

func CheckGetSubstring(url string, substr string, restrictedIfContains bool) func() (Result, error) {
	return func() (Result, error) {
		resp, err := http.Get(url)
		if err != nil {
			return Result{}, err
		}
		text, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			return Result{}, err
		}
		if strings.Contains(string(text), substr) {
			return Result{
				Restricted: restrictedIfContains,
			}, nil
		}
		return Result{
			Restricted: !restrictedIfContains,
		}, nil
	}
}
