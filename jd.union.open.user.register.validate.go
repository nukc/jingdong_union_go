package jingdong_union_go

import (
	"encoding/json"
	"errors"
	"log"
)

type JdUnionOpenUserRegisterValidateTopLevel struct {
	JdUnionOpenUserRegisterValidateResponse JdUnionOpenUserRegisterValidateResponse `json:"jd_union_open_user_register_validate_responce"`
}

type JdUnionOpenUserRegisterValidateResponse struct {
	ValidateResult string `json:"validateResult"`
	Code           string `json:"code"`
}

type JdUnionOpenUserRegisterValidateResult struct {
	Code    int64                               `json:"code"`
	Message string                              `json:"message"`
	Data    JdUnionOpenUserRegisterValidateData `json:"data"`
}

type JdUnionOpenUserRegisterValidateData struct {
	UserResp UserStateResp `json:"userResp"`
}

type UserStateResp struct {
	JdUser int64 `json:"jdUser"` // 1：实时不满足要求，2：实时满足要求
}

func (app *App) JdUnionOpenUserRegisterValidate(params map[string]interface{}) (result *JdUnionOpenUserRegisterValidateResult, err error) {

	body, err := app.Request("jd.union.open.user.register.validate", "1.0", map[string]interface{}{"userStateReq": params})
	resp := &JdUnionOpenUserRegisterValidateTopLevel{}
	if err != nil {
		log.Println(string(body))
		return
	}
	log.Printf("%v", string(body))
	if err = json.Unmarshal(body, resp); err != nil {
		return
	}
	if resp.JdUnionOpenUserRegisterValidateResponse.ValidateResult != "" {
		result = &JdUnionOpenUserRegisterValidateResult{}
		if err = json.Unmarshal([]byte(resp.JdUnionOpenUserRegisterValidateResponse.ValidateResult), result); err != nil {
			return
		}
	} else {
		err = errors.New("result is null")
	}
	return
}
