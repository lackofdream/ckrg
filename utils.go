package ckrg

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func GetSelfIP() (string, error) {
	resp, err := http.Get("https://api.ip.sb/ip")
	if err != nil {
		return "", err
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	resp.Body.Close()

	return strings.TrimSpace(string(data)), nil
}

func GetLocalISP(ip string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.ip.sb/geoip/%s", ip))
	if err != nil {
		return "", err
	}

	var geoip map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&geoip)
	resp.Body.Close()

	return geoip["organization"].(string), nil
}
