package futuapi

import (
	"context"
	"fmt"
	"testing"

	"github.com/gandalf1024/go-futu-api/pb/qotcommon"
)

var context_back = context.Background()

func TestConnect(t *testing.T) {
	api := NewFutuAPI()
	defer api.Close(context_back)

	api.SetRecvNotify(true)

	nCh, err := api.SysNotify(context_back)
	if err != nil {
		panic(err)
	}

	err = api.Connect(context_back, ":11111")
	if err != nil {
		panic(err)
	}

	security := &Security{qotcommon.QotMarket_QotMarket_CNSZ_Security, "000001"}
	securitys := []*Security{{qotcommon.QotMarket_QotMarket_CNSZ_Security, "000001"}}
	subTypeList := []qotcommon.SubType{qotcommon.SubType_SubType_Ticker}
	_ = security

	err = api.Subscribe(context_back, securitys, subTypeList, true, true, false, false)
	fmt.Println("完成订阅")

	sub, err := api.QuerySubscription(context_back, true)
	if err != nil {
		panic(err)
	}
	for _, v := range sub.ConnSubInfos {
		for _, info := range v.SubInfos {
			for _, sec := range info.Securities {
				fmt.Println(sec.Code)
			}
		}
	}
	fmt.Println("查询订阅完成1")

	tCh, err := api.UpdateTicker(context_back)
	if err != nil {
		panic(err)
	}

	select {
	case notify := <-nCh:
		fmt.Println(notify.Notification.Event.Desc)
	case ticker := <-tCh:
		fmt.Println(ticker.Ticker.Tickers[0])
	}

	sub, err = api.QuerySubscription(context_back, true)
	if err != nil {
		panic(err)
	}
	for _, v := range sub.ConnSubInfos {
		for _, info := range v.SubInfos {
			for _, sec := range info.Securities {
				fmt.Println(sec.Code)
			}
		}
	}
	fmt.Println("查询订阅完成2")

	secs, err := api.GetUserSecurity(context_back, "全部")
	if err != nil {
		t.Error(err)
	} else {
		for i, sec := range secs {
			fmt.Println(i, sec.Basic.Name)
		}
	}
	fmt.Println("------------------------------------------")
	secs, err = api.GetUserSecurity(context_back, "ETF")
	if err != nil {
		t.Error(err)
	} else {
		for i, sec := range secs {
			fmt.Println(i, sec.Basic.Name)
		}
	}
}
