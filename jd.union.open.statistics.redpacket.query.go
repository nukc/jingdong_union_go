package jingdong_union_go

import (
	"encoding/json"
	"log"
)

type JdUnionOpenStatisticsRedpacketQueryResponse struct {
	JdUnionOpenStatisticsRedpacketQueryResponce JdUnionOpenStatisticsRedpacketQueryResponce `json:"jd_union_open_statistics_redpacket_query_responce"`
}

type JdUnionOpenStatisticsRedpacketQueryResponce struct {
	QueryResult JdUnionOpenStatisticsRedpacketQueryQueryResult `json:"queryResult"`
}

type JdUnionOpenStatisticsRedpacketQueryQueryResult struct {
	Code       string                                       `json:"code"`
	Message    string                                       `json:"message"`
	TotalCount string                                       `json:"totalCount"`
	Data       JdUnionOpenStatisticsRedpacketQueryQueryData `json:"data"`
}

type JdUnionOpenStatisticsRedpacketQueryQueryData struct {
	RedPacketEffectDataResp RedPacketEffectDataResp `json:"redPacketEffectDataResp"`
}

type RedPacketEffectDataResp struct {
	ShowNum       string `json:"showNum"`
	PromotionName string `json:"promotionName"`
	PositionId    string `json:"positionId"`
	UseNum        string `json:"useNum"`
	ActId         string `json:"actId"`
	GiveNum       string `json:"giveNum"`
	OrderNum      string `json:"orderNum"`
	OrderFee      string `json:"orderFee"`
	OrderPrice    string `json:"orderPrice"`
	TjDate        string `json:"tjDate"`
	ChannelId     string `json:"channelId"`
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

	result = &resp.JdUnionOpenStatisticsRedpacketQueryResponce.QueryResult

	return
}
