package tuya

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil" // 1.15 limitation
	"log"
	"net/http"
)

type TuyaResponse struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Success bool   `json:"success"`
	T       int64  `json:"t"`
	Tid     string `json:"tid"`
}

func PerformRequest(method string, path string, body []byte) (TuyaResponse, error) {
	client := &http.Client{}
	req, _ := http.NewRequest(method, Host+path, bytes.NewReader(body))

	buildHeader(req, body)
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return TuyaResponse{}, err
	}
	defer resp.Body.Close()
	bs, _ := ioutil.ReadAll(resp.Body)
	response := string(bs)
	log.Println("resp:", response)

	var respStu TuyaResponse
	err = json.Unmarshal(bs, &respStu)
	if err != nil {
		fmt.Println("Failed to unmarshal response:", err)
		return TuyaResponse{}, err
	}

	if !respStu.Success {
		return respStu, fmt.Errorf(response)
	}
	return respStu, nil
}
