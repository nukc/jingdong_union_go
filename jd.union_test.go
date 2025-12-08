package jingdong_union_go

import (
	"fmt"
	"log"
	"testing"
)

var app = &App{
	ID:     "",
	Key:    "",
	Secret: "",
}

// 获取商品类目
func TestOpenCategoryGoodsGet(t *testing.T) {
	res, err := app.JdUnionOpenCategoryGoodsGet(map[string]interface{}{
		"parentId": 0,
		"grade":    0,
	})
	log.Println(res, err)
}

// 获取活动列表
func TestOpenActivityQuery(t *testing.T) {
	res, err := app.JdUnionOpenActivityQuery(map[string]interface{}{
		"pageIndex":  1,
		"pageSize":   50,
		"poolId":     6, //活动物料ID，1：营销日历热门会场；2：营销日历热门榜单；6：PC站长端官方活动
		"activeDate": "20210421",
	})
	log.Println(len(res.Data), res.TotalCount, err)
}

func TestOpenGoodsJingfenQuery(t *testing.T) {
	res, err := app.JdUnionOpenGoodsJingfenQuery(map[string]interface{}{
		"eliteId":   1,
		"sortName":  "price",
		"sort":      "asc",
		"pageIndex": 1,
		"pageSize":  10,
	})
	log.Println(res, err)
}

func TestOpenGoodsQuery(t *testing.T) {

	//单品查询
	res, err := app.JdUnionOpenGoodsQuery(map[string]interface{}{
		"skuIds":   []int{30881878056},
		"isCoupon": "1",
	})
	log.Println(res, err)

	//列表查询
	res, err = app.JdUnionOpenGoodsQuery(map[string]interface{}{
		"sort":                 "asc",   // asc desc
		"sortName":             "price", //price：单价, commissionShare：佣金比例, commission：佣金， inOrderCount30Days：30天引单量， inOrderComm30Days：30天支出佣金
		"isCoupon":             1,
		"commissionShareStart": 20, //佣金比例开始区间
		"pageIndex":            1,
		"pageSize":             10,
		"cid1":                 1315,
	})
	log.Println(res, err)
}

// 获取单品信息
func TestOpenGoodsPromotionGoodsInfoQuery(t *testing.T) {
	res, err := app.JdUnionOpenGoodsPromotiongoodsinfoQuery(map[string]interface{}{
		"skuIds": "30881878056",
	})
	log.Println(res, err)
}

// 获取通用推广链接
func TestOpenPromotionCommonGet(t *testing.T) {
	res, err := app.JdUnionOpenPromotionCommonGet(map[string]interface{}{
		"subUnionId": "test_subunionid",
		"ext1":       "test_ext",
		"siteId":     app.ID,
		"materialId": "https://wqitem.jd.com/item/view?sku=43415523405",
		"positionId": 1000,
		"couponUrl":  "http://coupon.jd.com/ilink/couponSendFront/send_index.action?key=02d2b6ff587c42fda6d4cac7ff1c2d6a&roleId=20498843&to=mall.jd.com/index-821028.html",
	})
	log.Println(res, err)
}

// 获取商品订单
func TestOpenOrderQuery(t *testing.T) {

	//单品查询
	res, err := app.JdUnionOpenOrderQuery(map[string]interface{}{
		"type":     "1", //1 下单时间  2 完成时间 3 更新时间
		"time":     "201906141811",
		"pageNo":   1,
		"pagesize": 500,
	})

	log.Println(res, err)
}

// 通过subUnionid获取推广链接
// https://wqitem.jd.com/item/view?sku=
func TestOpenPromotionBySubUnionIdGet(t *testing.T) {
	skuId := 43415523405
	res, err := app.JdUnionOpenPromotionBysubunionidGet(map[string]interface{}{
		"subUnionId": "{\"u\":\"1020\"}",
		"positionId": 1000,
		"chainType":  3,
		"materialId": fmt.Sprintf("https://wqitem.jd.com/item/view?sku=%d", skuId),
		"couponUrl":  "http://coupon.jd.com/ilink/couponSendFront/send_index.action?key=02d2b6ff587c42fda6d4cac7ff1c2d6a&roleId=20498843&to=mall.jd.com/index-821028.html",
	})
	log.Println(res, err)
}

// 查询渠道下订单行信息
// https://jos.jd.com/apilistnewdetail?apiGroupId461&apiId=20843
func TestOpenOrangeOrderQuery(t *testing.T) {

	pageIndex := 1
	for {
		res, err := app.JdOpenOrangeOrderQuery("",
			map[string]interface{}{
				"type":      1, //1 下单时间
				"startTime": "2024-10-30 00:00:00",
				"endTime":   "2024-10-31 00:00:00",
				"pageIndex": pageIndex,
				"pageSize":  500,
			})
		if err != nil {
			log.Println(res, err)
			break
		}
		if len(res.Result) == 0 {
			break
		}
		for _, item := range res.Result {
			if item.ReBuyDayW30 != "0" {
				log.Println(item.OrderID)
			} else if item.ReBuyMonth1 != "0" {
				log.Println(item.OrderID)
			}
		}
		pageIndex++
	}

}

func TestOpenStatisticsRedpacketQuery(t *testing.T) {
	result, err := app.JdUnionOpenStatisticsRedpacketQuery(map[string]interface{}{
		"startDate": "2025-06-18",
		"endDate":   "2025-06-18",
		"pageIndex": 1,
		"pageSize":  100,
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}

func TestOpenPositionCreate(t *testing.T) {
	result, err := app.JdUnionOpenPositionCreate(map[string]interface{}{
		"unionId":       "",
		"key":           "",
		"unionType":     4,
		"type":          6,
		"spaceNameList": []string{"test"},
		"siteId":        0,
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}

func TestJdUnionOpenShPromotionGet(t *testing.T) {
	result, err := app.JdUnionOpenShPromotionGet(map[string]interface{}{
		"account":    "",
		"materialId": "",
		"taskId":     "",
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}

// 查询奖励订单
func TestJdUnionOpenOrderBonusQuery(t *testing.T) {
	result, err := app.JdUnionOpenOrderBonusQuery(map[string]interface{}{
		"optType":   2,             // 1.下单时间拉取、2.更新时间拉取
		"startTime": 1763976720000, // 订单开始时间，时间戳（毫秒）
		"endTime":   1763976820000, // 订单结束时间，时间戳（毫秒），与startTime间隔不超过10秒
		"pageNo":    1,             // 页码，默认值为1
		"pageSize":  100,           // 每页数量，上限100
		"sortValue": "",            // 第一页传空，后续页传入上一页最后一条记录的sortValue
		// "activityId": 1,          // 奖励活动ID（可选）
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
	for _, item := range result.Data {
		t.Logf("订单ID: %d, 商品名称: %s, 活动名称: %s",
			item.OrderId,
			item.SkuName,
			item.ActivityName)
	}
}

// 用户注册验证
func TestJdUnionOpenUserRegisterValidate(t *testing.T) {
	result, err := app.JdUnionOpenUserRegisterValidate(map[string]interface{}{
		"userId":     "512E10A8817CA3F964523933210D0CC7", // 用户ID
		"userIdType": 131072,                             // 用户ID类型，支持多种类型组合
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
	t.Logf("用户状态: %d", result.Data.UserResp.JdUser)
}
