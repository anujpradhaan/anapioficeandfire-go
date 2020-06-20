package resttemplate

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func MakeGetCall(resourceUrl string, response interface{}, queryParams map[string]string) {
	httpUrl, error := url.Parse(resourceUrl)
	params := url.Values{}
	for key, value := range queryParams {
		params.Add(key, value)
	}
	httpUrl.RawQuery = params.Encode()
	fmt.Println(httpUrl.String())
	resp, err := http.Get(httpUrl.String())
	if error != nil {
		panic(error)
	}
	//fmt.Println(resp.Header)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	//log.Println(string(body))
	json.Unmarshal(body, response)
}
