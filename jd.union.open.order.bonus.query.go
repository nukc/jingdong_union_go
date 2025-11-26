package jingdong_union_go

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/shopspring/decimal"
)

type JdUnionOpenOrderBonusQueryTopLevel struct {
	JdUnionOpenOrderBonusQueryResponse JdUnionOpenOrderBonusQueryResponse `json:"jd_union_open_order_bonus_query_responce"`
}

type JdUnionOpenOrderBonusQueryResponse struct {
	QueryResult string `json:"queryResult"`
	Code        string `json:"code"`
}

type JdUnionOpenOrderBonusQueryResult struct {
	Code      int64            `json:"code"`
	Data      []BonusOrderResp `json:"data"`
	Message   string           `json:"message"`
	RequestID string           `json:"requestId"`
}

type BonusOrderResp struct {
	EstimateCosPrice decimal.Decimal `json:"estimateCosPrice"` // 预估计佣金额
	CommissionRate   decimal.Decimal `json:"commissionRate"`   // 佣金比例
	OrderId          int             `json:"orderId"`          // 订单ID
	EstimateBonusFee decimal.Decimal `json:"estimateBonusFee"` // 预估奖励金额
	BonusText        string          `json:"bonusText"`        // 奖励状态文案
	EstimateFee      decimal.Decimal `json:"estimateFee"`      // 预估佣金
	Pid              string          `json:"pid"`              // 推广位ID
	OrderState       int             `json:"orderState"`       // 订单状态
	SkuName          string          `json:"skuName"`          // 商品名称
	ActivityId       int             `json:"activityId"`       // 活动ID
	OrderTime        int64           `json:"orderTime"`        // 下单时间
	PayPrice         decimal.Decimal `json:"payPrice"`         // 实际支付金额
	Id               int             `json:"id"`               // 记录ID
	ActualFee        decimal.Decimal `json:"actualFee"`        // 实际佣金
	FinalRate        decimal.Decimal `json:"finalRate"`        // 最终比例
	SkuId            int             `json:"skuId"`            // 商品ID
	ChannelId        int             `json:"channelId"`        // 渠道ID
	Ext1             string          `json:"ext1"`             // 扩展字段1
	BonusInvalidCode string          `json:"bonusInvalidCode"` // 奖励无效码
	FinishTime       int64           `json:"finishTime"`       // 完成时间
	UnionAlias       string          `json:"unionAlias"`       // 联盟别名
	SubsidyRate      decimal.Decimal `json:"subsidyRate"`      // 补贴比例
	UnionId          int             `json:"unionId"`          // 联盟ID
	ActivityName     string          `json:"activityName"`     // 活动名称
	BonusInvalidText string          `json:"bonusInvalidText"` // 奖励无效文案
	ParentId         int             `json:"parentId"`         // 父订单ID
	SubUnionId       string          `json:"subUnionId"`       // 子联盟ID
	SortValue        string          `json:"sortValue"`        // 排序值
	ItemId           string          `json:"itemId"`           // 商品项ID
	OrderText        string          `json:"orderText"`        // 订单状态文案
	PositionId       int             `json:"positionId"`       // 推广位ID
	ActualCosPrice   decimal.Decimal `json:"actualCosPrice"`   // 实际计佣金额
	BonusState       int             `json:"bonusState"`       // 奖励状态
	ActualBonusFee   decimal.Decimal `json:"actualBonusFee"`   // 实际奖励金额
	SubSideRate      decimal.Decimal `json:"subSideRate"`      // 补贴比例
	Account          string          `json:"account"`          // 账号
}

func (app *App) JdUnionOpenOrderBonusQuery(params map[string]interface{}) (result *JdUnionOpenOrderBonusQueryResult, err error) {

	body, err := app.Request("jd.union.open.order.bonus.query", "1.0", map[string]interface{}{"orderReq": params})
	resp := &JdUnionOpenOrderBonusQueryTopLevel{}
	if err != nil {
		log.Println(string(body))
		return
	}
	log.Printf("%v", string(body))
	if err = json.Unmarshal(body, resp); err != nil {
		return
	}
	if resp.JdUnionOpenOrderBonusQueryResponse.QueryResult != "" {
		result = &JdUnionOpenOrderBonusQueryResult{}
		if err = json.Unmarshal([]byte(resp.JdUnionOpenOrderBonusQueryResponse.QueryResult), result); err != nil {
			return
		}
	} else {
		err = errors.New("result is null")
	}
	return
}
