package jingdong_union_go

import (
	"encoding/json"
	"errors"
	"log"
)

type JdOpenOrangeOrderQueryResponseTopLevel struct {
	JdOpenOrangeOrderQueryResponse JdOpenOrangeOrderQueryResponse `json:"jingdong_open_orange_order_query_responce"`
}

type JdOpenOrangeOrderQueryResponse struct {
	Result    JdOpenOrangeOrderQueryResult `json:"queryResult"`
	Code      string                       `json:"code"`
	RequestID string                       `json:"requestId"`
}

type JdOpenOrangeOrderQueryResult struct {
	Code    int64          `json:"code"`
	Result  []*OrangeOrder `json:"result"`
	Message string         `json:"message"`
}

type OrangeOrder struct {
	ReBuyDayW30  string `json:"reBuyDayW30"`
	ReBuyMonth1  string `json:"reBuyMonth1"`
	EUID         string `json:"euId"`
	SpName       string `json:"spName"`
	OrderID      string `json:"orderId"`
	ReBuyDay4    string `json:"reBuyDay4"`
	SkuName      string `json:"skuName"`
	ReBuyDay3    string `json:"reBuyDay3"`
	ReBuyDay2    string `json:"reBuyDay2"`
	OrderTime    string `json:"orderTime"`
	ReBuyDay1    string `json:"reBuyDay1"`
	ReBuyDay7    string `json:"reBuyDay7"`
	ReBuyDay6    string `json:"reBuyDay6"`
	ReBuyDay5    string `json:"reBuyDay5"`
	ReBuyDayW15  string `json:"reBuyDayW15"`
	StoreName    string `json:"storeName"`
	RefStatus    string `json:"refStatus"`
	SkuID        string `json:"skuId"`
	EmployeeName string `json:"employeeName"`
	GyxRealtime  string `json:"gyxRealtime"`
	ReBuyDayW3   string `json:"reBuyDayW3"`
	ReBuyDayW2   string `json:"reBuyDayW2"`
	ReBuyDayW5   string `json:"reBuyDayW5"`
	ReBuyDayW4   string `json:"reBuyDayW4"`
	ReBuyDayW7   string `json:"reBuyDayW7"`
	ReBuyDayW6   string `json:"reBuyDayW6"`
	EmployeeID   string `json:"employeeId"`
	StoreID      string `json:"storeId"`
	SpID         string `json:"spId"`
	ParentID     string `json:"parentId"`
	SubUnionID   string `json:"subUnionId"`
	EmployeePin  string `json:"employeePin"`
}

func (app *App) JdOpenOrangeOrderQuery(accessToken string, params map[string]interface{}) (result *JdOpenOrangeOrderQueryResult, err error) {
	body, err := app.Request("jingdong.open.orange.order.query", "2.0", map[string]interface{}{"request": params, "access_token": accessToken})

	resp := &JdOpenOrangeOrderQueryResponseTopLevel{}
	if err != nil {
		log.Println(string(body))
		return
	}

	if err = json.Unmarshal(body, resp); err != nil {
		return
	}
	if resp.JdOpenOrangeOrderQueryResponse.Code == "0" {
		return &resp.JdOpenOrangeOrderQueryResponse.Result, nil
	} else {
		err = errors.New("result is null")
	}
	return
}
