package dmall

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
	"ysf/music"
	"ysf/wxPusher"
)
type response struct {
	Code string
	Data struct{
		ActivityId int
		Banner string
		RuleDetailList []map[string]string
		VendorId int
		WareInfos []WareInfo
	}
	Msg string
	Time int
}

type WareInfo struct {
	BtnUpText string
	ExtraQualify bool
	Img string
	monthlyQualify bool
	Name string
	Price int
	SkuId int
	stock int
}

type OrderSubmitResponse struct {
	Code string
	Data string
	Msg	string
	Time string
}

func SeriesWareList()  {
	//{"code":"0000","data":{"activityId":157601,"banner":"https://img.dmallcdn.com/presale/202209/79e890a4-a760-4cf8-b3af-e1547405e0f5","ruleDetailList":[{"content":"茅台预购2022年规则有效期：2022年9月1日-2022年9月30日</br>\n<b>【预购资格】需同时满足以下条件：</b></br>\n- <b>预购年龄：</b>会员满18周岁才能参与茅台预购</br>\n- <b>抢购前完善基本信息及实名认证：</b>抢购前需在多点APP上补充姓名、电话及身份证号码等信息，进行实名认证后方可参与抢购；</br>\n- <b>其他：</b>客户身份证在多点APP登记成功后，如需修改，需携带本人身份证原件（需与多点抢购时登记身份证件一致）到商超门店告知修改身份证的原因（每个自然年只有一次修改身份证的机会，同时一个身份证号仅可绑定一个会员手机号，如修改时有其他会员手机号下已绑定此身份证信息，则确认修改并成功后原绑定手机号下的身份证信息将自动清除）</br>\n<b>参与条件：</b></br>\n1、资格按月发放，发放资格当月注册时长不低于3个月，如5月发放资格，须注册时间为1月或1月前；</br>\n2、发放资格前3个月有连续消费记录，每月不低于1次，到家订单以完成收货时间为准。</br>\n<b>【详情说明】</b></br>\n1、满足以上条件的同一用户次月将获得预购资格，购买间隔期为30天。</br>\n2、消费订单包含：O2O（到家订单）、智能购（自助购/自由购/智能购物车）、电子扫码单、社区拼团；</br>\n3、每天设置每家店53度飞天茅台酒预约库存量，当日库存抢购完为止。</br>\n4、重百电子会员通过多点APP茅台预售页面，每天早上10点开始预约抢购。</br>\n5、购买53度500ml飞天茅台酒，不可使用重百卡及优惠券支付，不参与任何门店优惠活动，不参与多点积分的成长福利等活动。</br>\n6、重百电子会员到店提货时，销售人员拨打订单手机号，验证身份（包括多点app登记身份证号、本人手机号）并出示本人身份证原件，核销提货凭证后方可提货。非本人或验证信息不符的不得提货，领取资格至订单当日起为会员保留7天，逾期将做退货处理货款原路返还。</br>\n7、需会员本人现场验货并由专职销售人员登记防伪编码，酒品离柜后，概不退换。 </br>","title":"茅台预购规则调整通知"},{"content":"8、重百/新世纪超市员工通过多点APP消费，不获得预约抢购茅台资格。</br>\n9、同一用户的认定标准：同一登录账户，同一手机号，同一身份证号、同一终端设备号，同一支付账户，同一收货地址，同一IP或其他合理显示同一用户的情况均视为同一用户。</br>\n10、如有以下不正当操作的用户将取消茅台预约抢购资格：</br>\n1）如用户以参与营销活动或使用营销工具所获得的优惠权益进行盈利或非法获利，或者多点有合理理由怀疑用户存在不当使用优惠工具或优惠权益的，多点将取消用户的参与资格，并有权撤销相关违规交易、收回优惠权益（含已使用及未使用的），必要时将追究用户的法律责任。</br>\n2）如用户曾经存在、出现或经多点合理怀疑有违法违规或违背诚实信用原则的行为，用户将无法获取/使用全部或部分优惠权益， 并且多点有权追究用户的法律责任。本条用户违法违规或违背诚实信用原则的行为，包括但不限于：</br>\na）通过任何不正当手段或以违反诚实信用原则的方式参与活动的，如通过与其他用户、商家串通或利用机器等方式进行作弊、刷取活动道具或积分等扰乱活动秩序的行为；</br>\nb）通过任何不正当手段或以违反诚实信用原则的方式达成交易的，如利用秒杀器等工具下单、套取优惠利差、虚假下单交易、提供虚假交易信息等扰乱活动秩序、违反交易规则的行为；</br>\nc）用户欠缺交易意愿，如利用退款流程的便利性以实现获取优惠权益的目的，交易后出现退款的；</br>\nd）其他违反诚实信用原则的行为；<br>\n飞天43%vol 500ml贵州茅台酒（带杯），53%vol500ml陈年贵州茅台酒（15），飞天53%vol 200ml贵州茅台酒于2022年9月8日——2022年9月10日期间每日11:00放量，共计放量2160瓶，仅限购1瓶\n","title":"茅台预购规则调整通知"}],"vendorId":69,"wareInfos":[{"btnUpText":"","extraQualify":true,"img":"https://img.dmallcdn.com/20190802/1e839889-6020-4952-b007-8325595a21a8","monthlyQualify":true,"name":"飞天53%vol 500ml贵州茅台酒（带杯）","price":149900,"skuId":1000815503,"stock":0},{"btnUpText":"","extraQualify":false,"img":"https://img.dmallcdn.com/prod/69/20220406/fbf0cbc84a7a4958877d33b76b1992d5","monthlyQualify":false,"name":"飞天43%vol 500ml贵州茅台酒（带杯）","price":109900,"skuId":1000815504,"stock":0},{"btnUpText":"","extraQualify":false,"img":"https://img.dmallcdn.com/prod/69/20220831/75819a9135494375b0ff67cef649c24c","monthlyQualify":false,"name":"53%vol 500ml陈年贵州茅台酒（15年）","price":599900,"skuId":1000820056,"stock":0},{"btnUpText":"","extraQualify":false,"img":"https://img.dmallcdn.com/prod/69/20220831/a59c7b4c111f415d97e0cc4b659d6cdc","monthlyQualify":false,"name":"飞天53%vol 200ml贵州茅台酒","price":69900,"skuId":1022274481,"stock":0}]},"msg":"成功","time":1662620949712}

	//local ,_ := time.LoadLocation("Local")
	//endTime,_ := time.ParseInLocation("2006-01-02 15:04:05","2022-03-25 19:00:00",local)

	//current := 0  //当前数量
	//isPush := false  //是否已经推送消息


	for {
		//if endTime.Before(time.Now())  {
		//	break
		//}
		c := colly.NewCollector(
			// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
			//colly.UserAgent("Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36 NetType/WIFI MicroMessenger/7.0.20.1781(0x6700143B) WindowsWechat(0x63040026)"),
			//colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
			colly.UserAgent("Mozilla/5.0 (iPhone; CPU iPhone OS 14_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148Dmall/5.4.0"),
		)
		cookie := "addrId=; appMode=online; appVersion=5.5.5; businessCode=32521; dmTenantId=1; env=app; first_session_time=1617162749400; originBusinessFormat=1-4-8; platform=IOS; platformStoreGroup=; platformStoreGroupKey=2becf099e525de522e71fb90703d9f5f@MTI2Mi02OTM3Mg; polaris=; risk=0; session_count=621; session_id=5054FEBC37B842268031C70A2CA30E67; storeGroupKey=d8b3b5195fa03bff259650a9bde6b86a@MS0xNDI1Ny02OQ; storeId=14257; store_id=14257; tdc=26.17.0.104-3110844-3086078.1662868861789; ticketName=56711ADA76CBFAEAE48F46461F5965EE434C289932949280BCDB0768CB849FCB58E217AD5BFF8D20D37589A2B91D9AD89E3043CABDED263C501340E279CD5C9DB023E12F47E759B93EF8D405EEF1E12F26B64F4AA143AA529CCAE0A78741E4C981C1DA882AF6D4B5EB1E4D280968E81E3FCCE0263FA1066E6FAD580037E0745E; token=dca42e4f-3f51-40a0-b93b-d6c0d06ef148; userId=378745911; venderId=69; vender_id=69; webViewType=wkwebview; console_mode=0; inited=true; storeGroup=; updateTime=1662608957041; track_id=C9F83C56A2F00002DA901E8010301CD6; web_session_count=30; _utm_id=441058762; tempid=4bfbd9cb421eddb5e193eaaa323b0ac6"
		//cookies := setCookieRaw(cookie)
		//c.SetCookies("https://presale.dmall.com",cookies)
		c.OnRequest(func(r *colly.Request) {
			r.Headers.Set("cookie",cookie)
			r.Headers.Set("Accept","application/json, text/plain, */*")
			r.Headers.Set("Accept-Encoding","gzip, deflate, br")
			r.Headers.Set("Accept-Language","zh-cn")
			r.Headers.Set("Content-Length","349")
			r.Headers.Set("Content-Type","application/x-www-form-urlencoded")
			r.Headers.Set("Referer","https://static.dmall.com/kayak-project/vueacts/dist/index.html?tdc=26.17.0.104-3110844-3086078.259200000&venderId=69&venderType=3")
		})

		c.OnError(func(_ *colly.Response, err error) {
			log.Println("Something went wrong:", err)
		})

		c.OnResponse(func(r *colly.Response) {
			// 判断code
			if r.StatusCode == 200 {
				body := new(response)
				json.Unmarshal(r.Body,body)
				if len(body.Data.WareInfos) > 1 {
					wxPusher.Push("dmall")
					music.Play()
				}


				log.Println(body.Data.WareInfos)
				fmt.Println(string(r.Body))
			}
		})

		url := fmt.Sprintf("https://presale.dmall.com/maotai/seriesWareList")
		data := map[string]string{
			"param":"{\"vendorId\":\"69\"}",
			"token":"dca42e4f-3f51-40a0-b93b-d6c0d06ef148",
			"ticketName":"56711ADA76CBFAEAE48F46461F5965EE434C289932949280BCDB0768CB849FCB58E217AD5BFF8D20D37589A2B91D9AD89E3043CABDED263C501340E279CD5C9DB023E12F47E759B93EF8D405EEF1E12F26B64F4AA143AA529CCAE0A78741E4C981C1DA882AF6D4B5EB1E4D280968E81E3FCCE0263FA1066E6FAD580037E0745E",
		}
		_ = c.Post(url,data)

		time.Sleep(60*time.Second)
	}

}

func TradeInfo() {
	//{"code":"0000","data":{"optionalPeriods":[{"date":"2022-09-08","dateDisplay":"今日 周四","shipmentType":2,"timeList":[{"capacityFactor":0.0,"endTime":"20:00","isToday":true,"ruleHit":{"dimensionType":10,"hyPromiseCapacity":true,"preProduceTime":0,"ruleId":12503,"shipmentType":2,"systemVersion":"HY_PROMISE","timeInfoType":-1},"startTime":"16:00","timeInfoType":-1,"timeSlotId":-1,"timeStatus":1}]},{"date":"2022-09-09","dateDisplay":"明日 周五","shipmentType":2,"timeList":[{"capacityFactor":0.0,"endTime":"11:00","isToday":false,"ruleHit":{"dimensionType":10,"hyPromiseCapacity":true,"preProduceTime":0,"ruleId":12503,"shipmentType":2,"systemVersion":"HY_PROMISE","timeInfoType":-1},"startTime":"09:00","timeInfoType":-1,"timeSlotId":-1,"timeStatus":1},{"capacityFactor":0.0,"endTime":"16:00","isToday":false,"ruleHit":{"dimensionType":10,"hyPromiseCapacity":true,"preProduceTime":0,"ruleId":12503,"shipmentType":2,"systemVersion":"HY_PROMISE","timeInfoType":-1},"startTime":"11:00","timeInfoType":-1,"timeSlotId":-1,"timeStatus":1},{"capacityFactor":0.0,"endTime":"20:00","isToday":false,"ruleHit":{"dimensionType":10,"hyPromiseCapacity":true,"preProduceTime":0,"ruleId":12503,"shipmentType":2,"systemVersion":"HY_PROMISE","timeInfoType":-1},"startTime":"16:00","timeInfoType":-1,"timeSlotId":-1,"timeStatus":1}]},{"date":"2022-09-10","dateDisplay":"09月10日 周六","shipmentType":2,"timeList":[{"capacityFactor":0.0,"endTime":"11:00","isToday":false,"ruleHit":{"dimensionType":10,"hyPromiseCapacity":true,"preProduceTime":0,"ruleId":12503,"shipmentType":2,"systemVersion":"HY_PROMISE","timeInfoType":-1},"startTime":"09:00","timeInfoType":-1,"timeSlotId":-1,"timeStatus":1},{"capacityFactor":0.0,"endTime":"16:00","isToday":false,"ruleHit":{"dimensionType":10,"hyPromiseCapacity":true,"preProduceTime":0,"ruleId":12503,"shipmentType":2,"systemVersion":"HY_PROMISE","timeInfoType":-1},"startTime":"11:00","timeInfoType":-1,"timeSlotId":-1,"timeStatus":1},{"capacityFactor":0.0,"endTime":"20:00","isToday":false,"ruleHit":{"dimensionType":10,"hyPromiseCapacity":true,"preProduceTime":0,"ruleId":12503,"shipmentType":2,"systemVersion":"HY_PROMISE","timeInfoType":-1},"startTime":"16:00","timeInfoType":-1,"timeSlotId":-1,"timeStatus":1}]},{"date":"2022-09-11","dateDisplay":"09月11日 周日","shipmentType":2,"timeList":[{"capacityFactor":0.0,"endTime":"11:00","isToday":false,"ruleHit":{"dimensionType":10,"hyPromiseCapacity":true,"preProduceTime":0,"ruleId":12503,"shipmentType":2,"systemVersion":"HY_PROMISE","timeInfoType":-1},"startTime":"09:00","timeInfoType":-1,"timeSlotId":-1,"timeStatus":1},{"capacityFactor":0.0,"endTime":"16:00","isToday":false,"ruleHit":{"dimensionType":10,"hyPromiseCapacity":true,"preProduceTime":0,"ruleId":12503,"shipmentType":2,"systemVersion":"HY_PROMISE","timeInfoType":-1},"startTime":"11:00","timeInfoType":-1,"timeSlotId":-1,"timeStatus":1},{"capacityFactor":0.0,"endTime":"20:00","isToday":false,"ruleHit":{"dimensionType":10,"hyPromiseCapacity":true,"preProduceTime":0,"ruleId":12503,"shipmentType":2,"systemVersion":"HY_PROMISE","timeInfoType":-1},"startTime":"16:00","timeInfoType":-1,"timeSlotId":-1,"timeStatus":1}]},{"date":"2022-09-12","dateDisplay":"09月12日 周一","shipmentType":2,"timeList":[{"capacityFactor":0.0,"endTime":"11:00","isToday":false,"ruleHit":{"dimensionType":10,"hyPromiseCapacity":true,"preProduceTime":0,"ruleId":12503,"shipmentType":2,"systemVersion":"HY_PROMISE","timeInfoType":-1},"startTime":"09:00","timeInfoType":-1,"timeSlotId":-1,"timeStatus":1},{"capacityFactor":0.0,"endTime":"16:00","isToday":false,"ruleHit":{"dimensionType":10,"hyPromiseCapacity":true,"preProduceTime":0,"ruleId":12503,"shipmentType":2,"systemVersion":"HY_PROMISE","timeInfoType":-1},"startTime":"11:00","timeInfoType":-1,"timeSlotId":-1,"timeStatus":1},{"capacityFactor":0.0,"endTime":"20:00","isToday":false,"ruleHit":{"dimensionType":10,"hyPromiseCapacity":true,"preProduceTime":0,"ruleId":12503,"shipmentType":2,"systemVersion":"HY_PROMISE","timeInfoType":-1},"startTime":"16:00","timeInfoType":-1,"timeSlotId":-1,"timeStatus":1}]},{"date":"2022-09-13","dateDisplay":"09月13日 周二","shipmentType":2,"timeList":[{"capacityFactor":0.0,"endTime":"11:00","isToday":false,"ruleHit":{"dimensionType":10,"hyPromiseCapacity":true,"preProduceTime":0,"ruleId":12503,"shipmentType":2,"systemVersion":"HY_PROMISE","timeInfoType":-1},"startTime":"09:00","timeInfoType":-1,"timeSlotId":-1,"timeStatus":1},{"capacityFactor":0.0,"endTime":"16:00","isToday":false,"ruleHit":{"dimensionType":10,"hyPromiseCapacity":true,"preProduceTime":0,"ruleId":12503,"shipmentType":2,"systemVersion":"HY_PROMISE","timeInfoType":-1},"startTime":"11:00","timeInfoType":-1,"timeSlotId":-1,"timeStatus":1},{"capacityFactor":0.0,"endTime":"20:00","isToday":false,"ruleHit":{"dimensionType":10,"hyPromiseCapacity":true,"preProduceTime":0,"ruleId":12503,"shipmentType":2,"systemVersion":"HY_PROMISE","timeInfoType":-1},"startTime":"16:00","timeInfoType":-1,"timeSlotId":-1,"timeStatus":1}]},{"date":"2022-09-14","dateDisplay":"09月14日 周三","shipmentType":2,"timeList":[{"capacityFactor":0.0,"endTime":"11:00","isToday":false,"ruleHit":{"dimensionType":10,"hyPromiseCapacity":true,"preProduceTime":0,"ruleId":12503,"shipmentType":2,"systemVersion":"HY_PROMISE","timeInfoType":-1},"startTime":"09:00","timeInfoType":-1,"timeSlotId":-1,"timeStatus":1},{"capacityFactor":0.0,"endTime":"16:00","isToday":false,"ruleHit":{"dimensionType":10,"hyPromiseCapacity":true,"preProduceTime":0,"ruleId":12503,"shipmentType":2,"systemVersion":"HY_PROMISE","timeInfoType":-1},"startTime":"11:00","timeInfoType":-1,"timeSlotId":-1,"timeStatus":1},{"capacityFactor":0.0,"endTime":"20:00","isToday":false,"ruleHit":{"dimensionType":10,"hyPromiseCapacity":true,"preProduceTime":0,"ruleId":12503,"shipmentType":2,"systemVersion":"HY_PROMISE","timeInfoType":-1},"startTime":"16:00","timeInfoType":-1,"timeSlotId":-1,"timeStatus":1}]},{"date":"2022-09-15","dateDisplay":"09月15日 周四","shipmentType":2,"timeList":[{"capacityFactor":0.0,"endTime":"11:00","isToday":false,"ruleHit":{"dimensionType":10,"hyPromiseCapacity":true,"preProduceTime":0,"ruleId":12503,"shipmentType":2,"systemVersion":"HY_PROMISE","timeInfoType":-1},"startTime":"09:00","timeInfoType":-1,"timeSlotId":-1,"timeStatus":1},{"capacityFactor":0.0,"endTime":"16:00","isToday":false,"ruleHit":{"dimensionType":10,"hyPromiseCapacity":true,"preProduceTime":0,"ruleId":12503,"shipmentType":2,"systemVersion":"HY_PROMISE","timeInfoType":-1},"startTime":"11:00","timeInfoType":-1,"timeSlotId":-1,"timeStatus":1},{"capacityFactor":0.0,"endTime":"20:00","isToday":false,"ruleHit":{"dimensionType":10,"hyPromiseCapacity":true,"preProduceTime":0,"ruleId":12503,"shipmentType":2,"systemVersion":"HY_PROMISE","timeInfoType":-1},"startTime":"16:00","timeInfoType":-1,"timeSlotId":-1,"timeStatus":1}]},{"date":"2022-09-16","dateDisplay":"09月16日 周五","shipmentType":2,"timeList":[{"capacityFactor":0.0,"endTime":"11:00","isToday":false,"ruleHit":{"dimensionType":10,"hyPromiseCapacity":true,"preProduceTime":0,"ruleId":12503,"shipmentType":2,"systemVersion":"HY_PROMISE","timeInfoType":-1},"startTime":"09:00","timeInfoType":-1,"timeSlotId":-1,"timeStatus":1},{"capacityFactor":0.0,"endTime":"16:00","isToday":false,"ruleHit":{"dimensionType":10,"hyPromiseCapacity":true,"preProduceTime":0,"ruleId":12503,"shipmentType":2,"systemVersion":"HY_PROMISE","timeInfoType":-1},"startTime":"11:00","timeInfoType":-1,"timeSlotId":-1,"timeStatus":1},{"capacityFactor":0.0,"endTime":"20:00","isToday":false,"ruleHit":{"dimensionType":10,"hyPromiseCapacity":true,"preProduceTime":0,"ruleId":12503,"shipmentType":2,"systemVersion":"HY_PROMISE","timeInfoType":-1},"startTime":"16:00","timeInfoType":-1,"timeSlotId":-1,"timeStatus":1}]},{"date":"2022-09-17","dateDisplay":"09月17日 周六","shipmentType":2,"timeList":[{"capacityFactor":0.0,"endTime":"11:00","isToday":false,"ruleHit":{"dimensionType":10,"hyPromiseCapacity":true,"preProduceTime":0,"ruleId":12503,"shipmentType":2,"systemVersion":"HY_PROMISE","timeInfoType":-1},"startTime":"09:00","timeInfoType":-1,"timeSlotId":-1,"timeStatus":1},{"capacityFactor":0.0,"endTime":"16:00","isToday":false,"ruleHit":{"dimensionType":10,"hyPromiseCapacity":true,"preProduceTime":0,"ruleId":12503,"shipmentType":2,"systemVersion":"HY_PROMISE","timeInfoType":-1},"startTime":"11:00","timeInfoType":-1,"timeSlotId":-1,"timeStatus":1},{"capacityFactor":0.0,"endTime":"20:00","isToday":false,"ruleHit":{"dimensionType":10,"hyPromiseCapacity":true,"preProduceTime":0,"ruleId":12503,"shipmentType":2,"systemVersion":"HY_PROMISE","timeInfoType":-1},"startTime":"16:00","timeInfoType":-1,"timeSlotId":-1,"timeStatus":1}]},{"date":"2022-09-18","dateDisplay":"09月18日 周日","shipmentType":2,"timeList":[{"capacityFactor":0.0,"endTime":"11:00","isToday":false,"ruleHit":{"dimensionType":10,"hyPromiseCapacity":true,"preProduceTime":0,"ruleId":12503,"shipmentType":2,"systemVersion":"HY_PROMISE","timeInfoType":-1},"startTime":"09:00","timeInfoType":-1,"timeSlotId":-1,"timeStatus":1},{"capacityFactor":0.0,"endTime":"16:00","isToday":false,"ruleHit":{"dimensionType":10,"hyPromiseCapacity":true,"preProduceTime":0,"ruleId":12503,"shipmentType":2,"systemVersion":"HY_PROMISE","timeInfoType":-1},"startTime":"11:00","timeInfoType":-1,"timeSlotId":-1,"timeStatus":1},{"capacityFactor":0.0,"endTime":"20:00","isToday":false,"ruleHit":{"dimensionType":10,"hyPromiseCapacity":true,"preProduceTime":0,"ruleId":12503,"shipmentType":2,"systemVersion":"HY_PROMISE","timeInfoType":-1},"startTime":"16:00","timeInfoType":-1,"timeSlotId":-1,"timeStatus":1}]},{"date":"2022-09-19","dateDisplay":"09月19日 周一","shipmentType":2,"timeList":[{"capacityFactor":0.0,"endTime":"11:00","isToday":false,"ruleHit":{"dimensionType":10,"hyPromiseCapacity":true,"preProduceTime":0,"ruleId":12503,"shipmentType":2,"systemVersion":"HY_PROMISE","timeInfoType":-1},"startTime":"09:00","timeInfoType":-1,"timeSlotId":-1,"timeStatus":1},{"capacityFactor":0.0,"endTime":"16:00","isToday":false,"ruleHit":{"dimensionType":10,"hyPromiseCapacity":true,"preProduceTime":0,"ruleId":12503,"shipmentType":2,"systemVersion":"HY_PROMISE","timeInfoType":-1},"startTime":"11:00","timeInfoType":-1,"timeSlotId":-1,"timeStatus":1},{"capacityFactor":0.0,"endTime":"20:00","isToday":false,"ruleHit":{"dimensionType":10,"hyPromiseCapacity":true,"preProduceTime":0,"ruleId":12503,"shipmentType":2,"systemVersion":"HY_PROMISE","timeInfoType":-1},"startTime":"16:00","timeInfoType":-1,"timeSlotId":-1,"timeStatus":1}]},{"date":"2022-09-20","dateDisplay":"09月20日 周二","shipmentType":2,"timeList":[{"capacityFactor":0.0,"endTime":"11:00","isToday":false,"ruleHit":{"dimensionType":10,"hyPromiseCapacity":true,"preProduceTime":0,"ruleId":12503,"shipmentType":2,"systemVersion":"HY_PROMISE","timeInfoType":-1},"startTime":"09:00","timeInfoType":-1,"timeSlotId":-1,"timeStatus":1},{"capacityFactor":0.0,"endTime":"16:00","isToday":false,"ruleHit":{"dimensionType":10,"hyPromiseCapacity":true,"preProduceTime":0,"ruleId":12503,"shipmentType":2,"systemVersion":"HY_PROMISE","timeInfoType":-1},"startTime":"11:00","timeInfoType":-1,"timeSlotId":-1,"timeStatus":1},{"capacityFactor":0.0,"endTime":"20:00","isToday":false,"ruleHit":{"dimensionType":10,"hyPromiseCapacity":true,"preProduceTime":0,"ruleId":12503,"shipmentType":2,"systemVersion":"HY_PROMISE","timeInfoType":-1},"startTime":"16:00","timeInfoType":-1,"timeSlotId":-1,"timeStatus":1}]},{"date":"2022-09-21","dateDisplay":"09月21日 周三","shipmentType":2,"timeList":[{"capacityFactor":0.0,"endTime":"11:00","isToday":false,"ruleHit":{"dimensionType":10,"hyPromiseCapacity":true,"preProduceTime":0,"ruleId":12503,"shipmentType":2,"systemVersion":"HY_PROMISE","timeInfoType":-1},"startTime":"09:00","timeInfoType":-1,"timeSlotId":-1,"timeStatus":1},{"capacityFactor":0.0,"endTime":"16:00","isToday":false,"ruleHit":{"dimensionType":10,"hyPromiseCapacity":true,"preProduceTime":0,"ruleId":12503,"shipmentType":2,"systemVersion":"HY_PROMISE","timeInfoType":-1},"startTime":"11:00","timeInfoType":-1,"timeSlotId":-1,"timeStatus":1},{"capacityFactor":0.0,"endTime":"20:00","isToday":false,"ruleHit":{"dimensionType":10,"hyPromiseCapacity":true,"preProduceTime":0,"ruleId":12503,"shipmentType":2,"systemVersion":"HY_PROMISE","timeInfoType":-1},"startTime":"16:00","timeInfoType":-1,"timeSlotId":-1,"timeStatus":1}]}],"storeInfos":[{"distance":987,"distanceDouble":0.0,"latitude":29.600502,"longitude":106.500234,"sapId":"2266","storeAddr":"重庆市渝北区龙山街道新南路429号华融现代广场负一楼","storeId":14248,"storeName":"新世纪华融现代广场店","supportOfflineService":1},{"distance":3334,"distanceDouble":0.0,"latitude":29.633069,"longitude":106.480852,"sapId":"2238","storeAddr":"重庆市两江新区金通大道506号附1号","storeId":14252,"storeName":"新世纪康庄美地店","supportOfflineService":1},{"distance":4242,"distanceDouble":0.0,"latitude":29.571286,"longitude":106.497672,"sapId":"2318","storeAddr":"重庆市江北区大石坝红盛路37号","storeId":76342,"storeName":"新世纪东原店","supportOfflineService":1},{"distance":4979,"distanceDouble":0.0,"latitude":29.573139,"longitude":106.532278,"sapId":"2312","storeAddr":"观音桥步行街7号","storeId":69312,"storeName":"新世纪世纪新都超市","supportOfflineService":1},{"distance":6924,"distanceDouble":0.0,"latitude":29.557957,"longitude":106.461266,"sapId":"2178","storeAddr":"重庆市沙坪坝区渝碚路50号","storeId":14273,"storeName":"新世纪凯瑞商都店","supportOfflineService":1},{"distance":7356,"distanceDouble":0.0,"latitude":29.640764,"longitude":106.56878,"sapId":"2221","storeAddr":"两江新区金渝大道68号附58号","storeId":14254,"storeName":"新世纪金渝大道店","supportOfflineService":1},{"distance":8066,"distanceDouble":0.0,"latitude":29.537312,"longitude":106.512409,"sapId":"2315","storeAddr":"重庆市渝中区时代天街1号、2号负1层","storeId":76312,"storeName":"新世纪时代天街店","supportOfflineService":1},{"distance":8415,"distanceDouble":0.0,"latitude":29.545328,"longitude":106.455286,"sapId":"2157","storeAddr":"沙坪坝天星桥正街63号","storeId":14272,"storeName":"新世纪天星桥店","supportOfflineService":1},{"distance":9162,"distanceDouble":0.0,"latitude":29.557471,"longitude":106.575532,"sapId":"2311","storeAddr":"重庆市渝中区青年路18号","storeId":76292,"storeName":"新世纪解都超市","supportOfflineService":1},{"distance":9235,"distanceDouble":0.0,"latitude":29.527487,"longitude":106.485179,"sapId":"2161","storeAddr":"重庆市九龙坡区石桥铺石杨路17号","storeId":14228,"storeName":"新世纪石桥铺店","supportOfflineService":1},{"distance":10939,"distanceDouble":0.0,"latitude":29.511473,"longitude":106.514115,"sapId":"2313","storeAddr":"九龙坡区杨家坪正街26号附2号","storeId":69322,"storeName":"新世纪瑞成商都超市","supportOfflineService":1},{"distance":11169,"distanceDouble":0.0,"latitude":29.525676,"longitude":106.565891,"sapId":"2227","storeAddr":"重庆市南岸区南坪惠工路192号（曾记面庄对面入口）","storeId":14263,"storeName":"新世纪江南商都店","supportOfflineService":1},{"distance":14772,"distanceDouble":0.0,"latitude":29.477636,"longitude":106.48122,"sapId":"2314","storeAddr":"重庆市大渡口区春晖路街道松青路1011号裙楼负3-1号","storeId":76302,"storeName":"新世纪大渡口商都超市","supportOfflineService":1},{"distance":17291,"distanceDouble":0.0,"latitude":29.71729,"longitude":106.630584,"sapId":"2102","storeAddr":"重庆市渝北区双凤桥街道义学路2号金易都会负一楼","storeId":14308,"storeName":"重百渝北中心店","supportOfflineService":1},{"distance":23558,"distanceDouble":0.0,"latitude":29.400334,"longitude":106.542177,"sapId":"2147","storeAddr":"重庆市巴南区龙海大道5号1栋负一楼","storeId":14169,"storeName":"新世纪龙洲湾店","supportOfflineService":1},{"distance":23611,"distanceDouble":0.0,"latitude":29.798061,"longitude":106.389954,"sapId":"2262","storeAddr":"北碚区冯时行路300号","storeId":14178,"storeName":"新世纪北碚万达店","supportOfflineService":1},{"distance":61543,"distanceDouble":0.0,"latitude":29.363382,"longitude":105.932203,"sapId":"2322","storeAddr":"重庆市永川区昌州大道东段789号","storeId":76382,"storeName":"重百永川商场超市","supportOfflineService":1},{"distance":65313,"distanceDouble":0.0,"latitude":29.036385,"longitude":106.650571,"sapId":"2317","storeAddr":"重庆市綦江区文龙街道九龙大道47号","storeId":76332,"storeName":"新世纪綦江商都超市","supportOfflineService":1},{"distance":224232,"distanceDouble":0.0,"latitude":30.809616,"longitude":108.376999,"sapId":"2327","storeAddr":"重庆市万州区太白路2号","storeId":76432,"storeName":"新世纪万州商都超市","supportOfflineService":1},{"distance":324773,"distanceDouble":0.0,"latitude":31.015847,"longitude":109.467223,"sapId":"2320","storeAddr":"重庆市奉节县鱼复街道海成路78号","storeId":76362,"storeName":"新世纪奉节商都","supportOfflineService":1}],"tradeBannerText":"","userName":"吴磊","userPhone":"18883362533","wareInfo":{"img":"https://img.dmallcdn.com/prod/69/20220406/fbf0cbc84a7a4958877d33b76b1992d5","name":"飞天43%vol 500ml贵州茅台酒（带杯）","price":109900,"skuId":1000815504,"stock":0}},"msg":"成功","time":1662609309631}

	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		//colly.UserAgent("Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36 NetType/WIFI MicroMessenger/7.0.20.1781(0x6700143B) WindowsWechat(0x63040026)"),
		//colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
		colly.UserAgent("Mozilla/5.0 (iPhone; CPU iPhone OS 14_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148Dmall/5.5.5"),
	)
	cookie := "addrId=; appMode=online; appVersion=5.5.5; businessCode=1; dmTenantId=1; env=app; first_session_time=1617162749400; originBusinessFormat=1-4-8; platform=IOS; platformStoreGroup=; platformStoreGroupKey=c4fbfb37a0d47566fce18e447d596541@MTI2Mi02OTM3Mg; polaris=OK518pFO3EqVPZO6fgGNCCUs2fuDCAdaEFHiSiPcaFJJiKMAmBGOe6JKXxSrKn4Q3JOSwkpQpGICIHR5z0fEwA==; risk=0; session_count=620; session_id=50EA3BD7EE7B4FA5946830C7ACA6C400; storeGroupKey=ad4ab229a5d6a8ff326ce7a0045df860@MS0xNDI0NS02OQ; storeId=14245; store_id=14245; tdc=26.17.0.104-3110844-3086078.1662868214625; ticketName=56711ADA76CBFAEAE48F46461F5965EE434C289932949280BCDB0768CB849FCB58E217AD5BFF8D20D37589A2B91D9AD89E3043CABDED263C501340E279CD5C9DB023E12F47E759B93EF8D405EEF1E12F26B64F4AA143AA529CCAE0A78741E4C981C1DA882AF6D4B5EB1E4D280968E81E3FCCE0263FA1066E6FAD580037E0745E; token=dca42e4f-3f51-40a0-b93b-d6c0d06ef148; userId=378745911; venderId=69; vender_id=69; webViewType=wkwebview; console_mode=0; inited=true; storeGroup=; updateTime=1662608957041; track_id=C9F83C56A2F00002DA901E8010301CD6; web_session_count=30; _utm_id=441058762; tempid=4bfbd9cb421eddb5e193eaaa323b0ac6"
	//cookies := setCookieRaw(cookie)
	//c.SetCookies("https://presale.dmall.com",cookies)
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("cookie",cookie)
		r.Headers.Set("Accept","application/json, text/plain, */*")
		r.Headers.Set("Accept-Encoding","gzip, deflate, br")
		r.Headers.Set("Accept-Language","zh-cn")
		r.Headers.Set("Content-Length","890")
		r.Headers.Set("Content-Type","application/x-www-form-urlencoded")
		r.Headers.Set("Referer","https://static.dmall.com/kayak-project/vueacts/dist/index.html?tdc=26.17.0.104-3110844-3086078.259200000&venderId=69&venderType=3")
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		// 判断code
		if r.StatusCode == 200 {
			fmt.Println(string(r.Body))
		}
	})

	url := fmt.Sprintf("https://presale.dmall.com/maotai/tradeInfo")

	param := map[string]interface{}{
		"vendorId":"69",
		"longitude":106.50184678819444,
		"latitude":29.609269476996527,
		"erpStoreId":"",
		"skuId":"1022274481",
		"specialMark":1,
	}
	paramStr,_ := json.Marshal(param)
	data := map[string]string{
		"p" : "XcDSvwll7prbbsCWaZ48wz8iN5ZqhBdDm8XwtPkLPqQ38w3tIkcIw5FAuyL1DcnldmbZPMokcDVcjKIurCK8/D+d0uupwbDer6RGLNvvVie7GrgfPsWHmAemDdT0j1JsuNs3vfBMD4XgSWiXV9szRoerEknMy9SNCB0Q8aHixyIf6qqtrVMXaK/ueErxBTQ0YCDnAUKFWSVRbr6GuBmfI/zhnBj78faK0KF7hGMGDQKNTtBYbPeBSAKKQ/VJ3FwnxsV9ttMAoVlydB5FngGM1UtGShCTbwr3jjlzWI7p4V886502SQNhSL/aBMBRbiQN25EkeFCSHB8iKsfnzafpGw==",
		"t" : strconv.FormatInt(time.Now().UnixNano(),10),
		"param":string(paramStr),
		"token":"dca42e4f-3f51-40a0-b93b-d6c0d06ef148",
		"ticketName":"56711ADA76CBFAEAE48F46461F5965EE434C289932949280BCDB0768CB849FCB58E217AD5BFF8D20D37589A2B91D9AD89E3043CABDED263C501340E279CD5C9DB023E12F47E759B93EF8D405EEF1E12F26B64F4AA143AA529CCAE0A78741E4C981C1DA882AF6D4B5EB1E4D280968E81E3FCCE0263FA1066E6FAD580037E0745E",
	}
	_ = c.Post(url,data)
}


func OrderSubmit()  {

	isSuccess := false

	local ,_ := time.LoadLocation("Local")
	startTime,_ := time.ParseInLocation("2006-01-02 15:04:05","2022-09-09 11:00:00",local)

	for {

		// 判断时间
		for {
			if startTime.Before(time.Now())  {
				break
			}
			fmt.Println("还没到时间")
			time.Sleep(50*time.Millisecond)
		}

		c := colly.NewCollector(
			// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
			//colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
			colly.UserAgent("Mozilla/5.0 (iPhone; CPU iPhone OS 14_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148Dmall/5.5.5"),
		)
		cookie := "addrId=; appMode=online; appVersion=5.5.5; businessCode=1; dmTenantId=1; env=app; first_session_time=1617162749400; originBusinessFormat=1-4-8; platform=IOS; platformStoreGroup=; platformStoreGroupKey=c4fbfb37a0d47566fce18e447d596541@MTI2Mi02OTM3Mg; polaris=OK518pFO3EqVPZO6fgGNCCUs2fuDCAdaEFHiSiPcaFJJiKMAmBGOe6JKXxSrKn4QpAuk5NSlqG1ez78AliHzRQ==; risk=0; session_count=627; session_id=77BF4CC71F0449A6AC850AC5007227B2; storeGroupKey=ad4ab229a5d6a8ff326ce7a0045df860@MS0xNDI0NS02OQ; storeId=14245; store_id=14245; tdc=26.17.0.104-3110844-3086078.1662955181609; ticketName=56711ADA76CBFAEAE48F46461F5965EE434C289932949280BCDB0768CB849FCB58E217AD5BFF8D20D37589A2B91D9AD89E3043CABDED263C501340E279CD5C9DB023E12F47E759B93EF8D405EEF1E12F26B64F4AA143AA529CCAE0A78741E4C981C1DA882AF6D4B5EB1E4D280968E81E3FCCE0263FA1066E6FAD580037E0745E; token=dca42e4f-3f51-40a0-b93b-d6c0d06ef148; userId=378745911; venderId=69; vender_id=69; webViewType=wkwebview; track_id=C9F88FEFF7F000024FD3153012A0E700; web_session_count=35; inited=false; storeGroup=; updateTime=1662681664793; console_mode=0; _utm_id=441058762; tempid=4bfbd9cb421eddb5e193eaaa323b0ac6"
		//cookies := setCookieRaw(cookie)
		//c.SetCookies("https://presale.dmall.com",cookies)
		c.OnRequest(func(r *colly.Request) {
			r.Headers.Set("cookie",cookie)
			r.Headers.Set("Accept","application/json, text/plain, */*")
			r.Headers.Set("Accept-Encoding","gzip, deflate, br")
			r.Headers.Set("Accept-Language","zh-cn")
			r.Headers.Set("Content-Length","1266")
			r.Headers.Set("Content-Type","application/x-www-form-urlencoded")
			r.Headers.Set("Referer","https://static.dmall.com/kayak-project/vueacts/dist/index.html?tdc=26.17.0.104-3110844-3086078.259200000&venderId=69&venderType=3")
		})

		c.OnError(func(_ *colly.Response, err error) {
			log.Println("Something went wrong:", err)
		})

		c.OnResponse(func(r *colly.Response) {
			// 判断code
			if r.StatusCode == 200 {
				body := new(OrderSubmitResponse)
				json.Unmarshal(r.Body,body)
				if body.Code == "0000" {
					isSuccess = true
				}
				fmt.Println(string(r.Body))
			}
		})

		url := fmt.Sprintf("https://presale.dmall.com/maotai/orderSubmit")
		//{"vendorId":"69","erpStoreId":76332,"name":"吴磊","phone":"18883362533","shipmentDate":"2022-09-09","shipmentStartTime":"16:00","shipmentEndTime":"20:00","skuId":1000815504,"skuCount":1,"price":109900,"deviceId":"","longitude":106.50188666449652,"latitude":29.609061686197915,"specialMark":1,"orderOrigin":7,"channel":1,"appName":"com.dmall.dmall","dmTenantId":"1"}
		param := map[string]interface{}{
			"vendorId":"69",
			"erpStoreId":14248,
			"name":"吴磊",
			"phone":"18883362533",
			"shipmentDate":"2022-09-09",	//time.Now().Format("2006-01-02 15:04:05")
			"shipmentStartTime":"16:00",
			"shipmentEndTime":"20:00",
			"skuId":1022274481,
			"skuCount":1,
			"price":69900,
			"deviceId":"",
			"longitude":106.50184678819444,
			"latitude":29.609269476996527,
			"specialMark":1,
			"orderOrigin":7,
			"channel":1,
			"appName":"com.dmall.dmall",
			"dmTenantId":"1",
		}
		paramStr,_ := json.Marshal(param)
		data := map[string]string{
			"p" : "buPDLVhC3pPW0NrWdLaR7c4V9UrMM8qTBzN0MRPV73flGQCI2zPJ81fm19YfR9NL+FqBHhqGMeZUbVUVCZDSZ5rgruB8LAMEI/lG3BUa/myahfrOKi0kadvOdfoKb8xv1TQS5WHNoG7NXLzwIeYsGzwX6BYsPukymJFC5dqYNzrG3Hvo7bRvbcOr3fZUk0VZP3JpA5SBiln9SeW8nMBOplhhhreG751dfImgF/W3gA6xPucWGhF15cJ1UdXEFCVxZDYMK+yK/qzuY8FRKbaw6t3RiDEuZ2oc/aXhQMzs7bG18w1STONGvUjaP/MVLcXWbavooUVOVSJXPTX5yE6y6g==",
			"t" : strconv.FormatInt(time.Now().UnixNano(),10),
			"param":string(paramStr),
			"token":"dca42e4f-3f51-40a0-b93b-d6c0d06ef148",
			"ticketName":"56711ADA76CBFAEAE48F46461F5965EE434C289932949280BCDB0768CB849FCB58E217AD5BFF8D20D37589A2B91D9AD89E3043CABDED263C501340E279CD5C9DB023E12F47E759B93EF8D405EEF1E12F26B64F4AA143AA529CCAE0A78741E4C981C1DA882AF6D4B5EB1E4D280968E81E3FCCE0263FA1066E6FAD580037E0745E",
		}
		_ = c.Post(url,data)

		if isSuccess {
			break
		}
		time.Sleep(500*time.Millisecond)
	}


}


//{"vendorId":"69","longitude":106.50184678819444,"latitude":29.609269476996527,"erpStoreId":"","skuId":"1000815504","specialMark":1}
//{"vendorId":"69","erpStoreId":14248,"name":"吴磊","phone":"18883362533","shipmentDate":"2022-09-08","shipmentStartTime":"16:00","shipmentEndTime":"20:00","skuId":1000815504,"skuCount":1,"price":109900,"deviceId":"","longitude":106.50181477864584,"latitude":29.60911187065972,"specialMark":1,"orderOrigin":7,"channel":1,"appName":"com.dmall.dmall","dmTenantId":"1"}

//h11imx6YVKrxfsYQEsCAqprHHD4uyUOpQJua7hJxVTS+nfVsNf825q1WQRLkEAyclgIp4Mbl725YwjGFCnKz1KYGnA+FMxPCm9SvfVbWIJZe3rvoNLHh5y5dhYhywPoBJtjOU2QgmuLPIlWE0da1ZULtHlB2y+G9j+IGZ19g2k/KitAuJ3q8Bj2ld7ElbkPoe5t4vnO+q8c8feGO3mXRG7qNcw+y6/U8izyc2+l/AQ6YvHZPQIfAV20LAs6nTUO+IT5x7Kf1DST5pCrFGMteTqePVOQCq+9OjjHebDqYwPb/sm8bZqoj+4/2T6zOSjtCfpmPMFK/QZ6hegcg2us1Qg==
//LXEEbqrYIydQ+YKBXBQrfoO6d6dhxn8BXGtLRGgYf/da6HrBlfl/pfJOsZXuFpp1C5AY/hA/cnkQo9wtUbGzTipg2RMLGqhc+uCoKHEYB9JuFvmk4tOftSOiBxrqjHXMU4Ar2W25PE87XLkxG9LRB4rFMd2wLJ6CozPIOmPNsEo87VM+bWkTOGPl5HnRnBjM0rVXe7lTVLokZpu7P+Qh6NL4WkNDBzI6nKOg7lh+VlSnkIJ1knE4yehaKAdiIJRvSA3dikUkFWJahnGZDzntggpayw84vNt71ROi31i63QMHj++JGdB6ZeaHlW1//lxDoYSofxDkcjWGdlxpI8xZOA==

//1662609691740
//1662609309473


//56711ADA76CBFAEAE48F46461F5965EE434C289932949280BCDB0768CB849FCB58E217AD5BFF8D20D37589A2B91D9AD89E3043CABDED263C501340E279CD5C9DB023E12F47E759B93EF8D405EEF1E12F26B64F4AA143AA529CCAE0A78741E4C981C1DA882AF6D4B5EB1E4D280968E81E3FCCE0263FA1066E6FAD580037E0745E
//56711ADA76CBFAEAE48F46461F5965EE434C289932949280BCDB0768CB849FCB58E217AD5BFF8D20D37589A2B91D9AD89E3043CABDED263C501340E279CD5C9DB023E12F47E759B93EF8D405EEF1E12F26B64F4AA143AA529CCAE0A78741E4C981C1DA882AF6D4B5EB1E4D280968E81E3FCCE0263FA1066E6FAD580037E0745E


//dca42e4f-3f51-40a0-b93b-d6c0d06ef148
//dca42e4f-3f51-40a0-b93b-d6c0d06ef148


func setCookieRaw(cookieRaw string) []*http.Cookie {
	var cookies []*http.Cookie

	cookieList := strings.Split(cookieRaw,"; ")

	for _, item := range cookieList {
		keyValue := strings.Split(item,"=")
		name := keyValue[0]
		valueList := keyValue[1:]
		cookieItem := http.Cookie{
			Name:       name,
			Value:      strings.Join(valueList,"="),
		}
		cookies  = append(cookies,&cookieItem)
	}

	return cookies
}

