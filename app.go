package jingdong_union_go

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

type JDUnion struct {
	ID string
}

type App struct {
	ID     string
	Name   string
	Key    string
	Secret string
}

type JdUnionErrResp struct {
	ErrorResponse ErrorResponse `json:"errorResponse"`
}

type ErrorResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

const RouterURL = "https://api.jd.com/routerjson"
const RequestMethod = "POST"

func (app *App) Request(method string, v string, paramJSON map[string]interface{}) ([]byte, error) {
	// common params
	params := map[string]interface{}{}
	params["method"] = method
	params["app_key"] = app.Key
	params["format"] = "json"
	params["v"] = v
	params["sign_method"] = "md5"
	params["timestamp"] = time.Now().Format("2006-01-02 15:04:05")
	if paramJSON["access_token"] != nil {
		params["access_token"] = paramJSON["access_token"]
		delete(paramJSON, "access_token")
	}

	// api params
	paramJSONStr, _ := json.Marshal(paramJSON)
	params["360buy_param_json"] = string(paramJSONStr)
	params["sign"] = GetSign(app.Secret, params)
	log.Printf("Request: %s, %v", RouterURL, params)
	resp, err := http.PostForm(RouterURL, app.Values(params))
	log.Printf("Responce:%v %v", resp, err)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	//log.Printf("Responce Body:%v ", string(body))
	if err != nil {
		return nil, err
	}

	jdErr := JdUnionErrResp{}
	if err := json.Unmarshal(body, &jdErr); err != nil {
		return nil, err
	}

	if jdErr.ErrorResponse.Code != "" {
		return nil, errors.New(jdErr.ErrorResponse.Msg)
	}
	return body, nil
}

func (app *App) Values(params map[string]interface{}) url.Values {
	vals := url.Values{}
	for key, val := range params {
		vals.Add(key, GetString(val))
	}
	return vals
}
