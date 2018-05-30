package xmrstak

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"bitbucket.org/minerstats/output"
)

var defaultStruct map[string]interface{}

func HitXMRStak(host_l string, port_l string, buf *[]byte) {
	fullhost := "http://" + host_l + ":" + port_l + "/api.json"
	resp, err := http.Get(fullhost)
	if err != nil {
		*buf = output.MakeJSONError("xmrstak", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		*buf = output.MakeJSONError("xmrstak", err)
		return
	}
	json.Unmarshal(body, &defaultStruct)
	hrtotal := defaultStruct["hashrate"].(map[string]interface{})["total"].([]interface{})[0].(float64)
	numMiners := len(defaultStruct["connection"].(map[string]interface{}))
	hrstring := strconv.FormatFloat(hrtotal, 'f', 2, 64) + " H/s"
	js, err := output.MakeJSON_full("xmrstak", hrtotal, hrstring, numMiners, 0)
	if err != nil {
		*buf = output.MakeJSONError("xmrstak", err)
		return
	}
	*buf = js

}
