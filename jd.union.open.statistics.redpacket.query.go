package jingdong_union_go

import (
	"encoding/json"
	"errors"
	"log"
)

type JdUnionOpenStatisticsRedpacketQueryResponse struct {
	JdUnionOpenStatisticsRedpacketQueryResponce JdUnionOpenStatisticsRedpacketQueryResponce `json:"jd_union_open_statistics_redpacket_query_responce"`
}

type JdUnionOpenStatisticsRedpacketQueryResponce struct {
	Code        string `json:"code"`
	QueryResult string `json:"queryResult"`
}

type JdUnionOpenStatisticsRedpacketQueryQueryResult struct {
	Code       int                       `json:"code"`
	Message    string                    `json:"message"`
	RequestID  string                    `json:"requestId"`
	TotalCount int                       `json:"totalCount"`
	Data       []RedPacketEffectDataResp `json:"data"`
}

type RedPacketEffectDataResp struct {
	ActId         int64   `json:"actId"`
	GiveNum       int     `json:"giveNum"`
	OrderFee      float64 `json:"orderFee"`
	OrderNum      int     `json:"orderNum"`
	OrderPrice    float64 `json:"orderPrice"`
	PositionId    int64   `json:"positionId"`
	PromotionName string  `json:"promotionName"`
	ShowNum       int     `json:"showNum"`
	TjDate        string  `json:"tjDate"`
	UseNum        int     `json:"useNum"`
	ChannelId     int64   `json:"channelId"`
}

// JdUnionOpenStatisticsRedpacketQuery 京享红包效果数据
func (app *App) JdUnionOpenStatisticsRedpacketQuery(params map[string]interface{}) (result *JdUnionOpenStatisticsRedpacketQueryQueryResult, err error) {

	body, err := app.Request("jd.union.open.statistics.redpacket.query", "1.0", map[string]interface{}{"effectDataReq": params})

	resp := &JdUnionOpenStatisticsRedpacketQueryResponse{}
	if err != nil {
		log.Println(string(body))
		return
	}

	if err = json.Unmarshal(body, resp); err != nil {
		return
	}

	if resp.JdUnionOpenStatisticsRedpacketQueryResponce.QueryResult != "" {
		result = &JdUnionOpenStatisticsRedpacketQueryQueryResult{}
		if err = json.Unmarshal([]byte(resp.JdUnionOpenStatisticsRedpacketQueryResponce.QueryResult), result); err != nil {
			return
		}

	} else {
		err = errors.New("result is null")
	}

	return
}
