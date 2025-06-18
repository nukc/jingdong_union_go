package jingdong_union_go

import (
	"encoding/json"
	"errors"
	"log"
)

type JdUnionOpenPositionCreateResponse struct {
	JdUnionOpenPositionCreateResponce JdUnionOpenPositionCreateResponce `json:"jd_union_open_position_create_responce"`
}

type JdUnionOpenPositionCreateResponce struct {
	Code         string `json:"code"`
	CreateResult string `json:"createResult"`
}

type JdUnionOpenPositionCreateResult struct {
	Code    string                        `json:"code"`
	Message string                        `json:"message"`
	Data    JdUnionOpenPositionCreateData `json:"data"`
}

type JdUnionOpenPositionCreateData struct {
	UnionId    string     `json:"unionId"`
	SiteId     string     `json:"siteId"`
	Pid        PidObject  `json:"pid"`
	Type       string     `json:"type"`
	ResultList ResultList `json:"resultList"`
}

type PidObject struct {
	Pid string `json:"pid"`
}

type ResultList struct {
	Result string `json:"result"`
}

// JdUnionOpenPositionCreate 创建推广位
func (app *App) JdUnionOpenPositionCreate(params map[string]interface{}) (result *JdUnionOpenPositionCreateResult, err error) {

	body, err := app.Request("jd.union.open.statistics.redpacket.query", "1.0", map[string]interface{}{"effectDataReq": params})

	resp := &JdUnionOpenPositionCreateResponse{}
	if err != nil {
		log.Println(string(body))
		return
	}

	if err = json.Unmarshal(body, resp); err != nil {
		return
	}

	if resp.JdUnionOpenPositionCreateResponce.CreateResult != "" {
		result = &JdUnionOpenPositionCreateResult{}
		if err = json.Unmarshal([]byte(resp.JdUnionOpenPositionCreateResponce.CreateResult), result); err != nil {
			return
		}

	} else {
		err = errors.New("result is null")
	}

	return
}
