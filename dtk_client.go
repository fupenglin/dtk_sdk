package dtk_sdk

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type DtkClient struct {
}

func (ctx DtkClient) getSign(param string) string {
	fmt.Println(param)
	data := []byte(param)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return strings.ToUpper(md5str)
}

func (ctx DtkClient) Get(url string, params *DtkParams, key string) (string, error) {
	param := params.String(true) + "&sign=" + ctx.getSign(params.String(false)+"&key="+key)
	client := &http.Client{}
	url = url + "?" + param
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	return string(body), err
}
