package main

import (
	"fmt"

	"github.com/fupenglin/dtk_sdk"
)

const (
	appSecret string = "5ab6af96dc843971a5593066221051a8"
	appKey    string = "5dce6f7fcae08"
)

func main() {

	sdk := dtk_sdk.NewOpenPlatform(appKey, appSecret)
	result, err := sdk.GetSuperCategory()
	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}
}
