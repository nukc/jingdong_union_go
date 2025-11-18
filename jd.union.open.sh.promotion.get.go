package jingdong_union_go

import (
	"encoding/json"
	"errors"
	"log"
)

type JdUnionOpenShPromotionGetTopLevel struct {
	JdUnionOpenShPromotionGetResponse JdUnionOpenShPromotionGetResponse `json:"jd_union_open_sh_promotion_get_responce"`
}

type JdUnionOpenShPromotionGetResponse struct {
	Result string `json:"getResult"`
	Code   string `json:"code"`
}

type JdUnionOpenShPromotionGetResult struct {
	Code      int64                      `json:"code"`
	Data      JdUnionOpenShPromotionData `json:"data"`
	Message   string                     `json:"message"`
	RequestID string                     `json:"requestId"`
}

type JdUnionOpenShPromotionData struct {
	ClickURL             string `json:"clickUrl"`
	ImpressionMonitorUrl string `json:"impressionMonitorUrl"`
	ClickMonitorUrl      string `json:"clickMonitorUrl"`
	AppUrl               string `json:"appUrl"`
}

func (app *App) JdUnionOpenShPromotionGet(params map[string]interface{}) (result *JdUnionOpenShPromotionGetResult, err error) {

	body, err := app.Request("jd.union.open.sh.promotion.get", "1.0", map[string]interface{}{"getCodeByTaskIdReq": params})
	resp := &JdUnionOpenShPromotionGetTopLevel{}
	if err != nil {
		log.Println(string(body))
		return
	}
	log.Printf("%v", string(body))
	if err = json.Unmarshal(body, resp); err != nil {
		return
	}
	if resp.JdUnionOpenShPromotionGetResponse.Result != "" {
		result = &JdUnionOpenShPromotionGetResult{}
		if err = json.Unmarshal([]byte(resp.JdUnionOpenShPromotionGetResponse.Result), result); err != nil {
			return
		}
	} else {
		err = errors.New("result is null")
	}
	return
}
