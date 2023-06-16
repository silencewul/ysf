package elm

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

//eyJhY2Nlc3NQb2ludCI6IjEwMDAxN181MDAwMDBfNCIsImNhcnJpZXIiOiLnlLXkv6EiLCJjbGllbnRJc3AiOiIxMDAwMTciLCJjb2RlIjoxMDAwLCJjdiI6MCwiZG5zIjpbeyJhaXNsZXMiOlt7ImN0byI6MTAwMDAsImhlYXJ0YmVhdCI6NDUwMDAsInBvcnQiOjgwLCJwcm90b2NvbCI6Imh0dHAyIiwicHVibGlja2V5IjoiY2RuIiwicmV0cnkiOjEsInJ0byI6MTAwMDAsInJ0dCI6IjBydHQifSx7ImN0byI6MTAwMDAsImhlYXJ0YmVhdCI6NDUwMDAsInBvcnQiOjQ0MywicHJvdG9jb2wiOiJodHRwMiIsInB1YmxpY2tleSI6ImNkbiIsInJldHJ5IjoxLCJydG8iOjEwMDAwLCJydHQiOiIwcnR0In0seyJjdG8iOjEwMDAwLCJoZWFydGJlYXQiOjAsInBvcnQiOjQ0MywicHJvdG9jb2wiOiJodHRwcyIsInJldHJ5IjoxLCJydG8iOjEwMDAwfSx7ImN0byI6MTAwMDAsImhlYXJ0YmVhdCI6MCwicG9ydCI6ODAsInByb3RvY29sIjoiaHR0cCIsInJldHJ5IjoxLCJydG8iOjEwMDAwfV0sImhvc3QiOiJndy5hbGljZG4uY29tIiwiaXBzIjpbIjExMy4xNDEuMTkwLjEyMiIsIjExNi4yMTEuMTgzLjIzNSJdLCJpc0hvdCI6MSwic2FmZUFpc2xlcyI6Imh0dHBzIiwic3RyYXRlZ2llcyI6W10sInR0bCI6MzAwLCJ1bSI6MSwidmVyc2lvbiI6IjU0MSJ9LHsiYWlzbGVzIjpbXSwiaG9zdCI6ImgtYWRhc2h4LnV0LmVsZS5tZSIsImlwcyI6W10sInN0cmF0ZWdpZXMiOltdLCJ0dGwiOjMwMCwidW0iOjN9XSwiaXAiOiIxNC4xMDcuNzkuOTIiLCJ1aWQiOiIyMjA0OTQ1ODg0MjQ0IiwidW5pdCI6ImNlbnRlciJ9

var encodeStr = "H4sIAAAAAAAAC8tIzcnJVyjPL8pJAQCFEUoNCwAAAA=="


func Stat()  {
	//base64.StdEncoding.EncodeToString([]byte(encodeStr))
	dstr,_ := base64.StdEncoding.DecodeString(encodeStr)
	
	//
	//var b bytes.Buffer
	//gz := gzip.NewWriter(&b)
	//if _, err := gz.Write([]byte("hello world")); err != nil {
	//	panic(err)
	//}
	//if err := gz.Flush(); err != nil {
	//	panic(err)
	//}
	//if err := gz.Close(); err != nil {
	//	panic(err)
	//}
	//str := base64.StdEncoding.EncodeToString(b.Bytes())
	//fmt.Println(str)
	//data, _ := base64.StdEncoding.DecodeString(str)
	//fmt.Println(string(data))
	rdata := bytes.NewReader(dstr)
	r, _ := gzip.NewReader(rdata)
	s, _ := ioutil.ReadAll(r)
	fmt.Println(string(s))

	//for {
	//	//if endTime.Before(time.Now())  {
	//	//	break
	//	//}
	//	c := colly.NewCollector(
	//		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
	//		//colly.UserAgent("Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36 NetType/WIFI MicroMessenger/7.0.20.1781(0x6700143B) WindowsWechat(0x63040026)"),
	//		//colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
	//		colly.UserAgent("MTOPSDK%2F1.9.3.48%20%28iOS%3B14.6%3BApple%3BiPhone12%2C1%29"),
	//	)
	//	cookie := "_m_h5_tk=2d8e28cadbcabe2638d0582cd92b94fe_1661246911239; _m_h5_tk_enc=b81abca0e244bcd6a084bf65a6e48b44; _tb_token_=eee653376bea3; cookie2=12eb8f9293d30eba048e2a84378ec399; csg=5a714103; munb=2204945884244; sgcookie=W100OJVNZci%2B6pddmqbCT%2FZ455aogFAubzluegEn3%2F9bfBrm3YGgPomTeLBPJn%2BCUmloydX0Xan%2FkJkmSaafrs9%2Fjqt3xZkb%2ByzRmu7vNkcwJCc%3D; t=11f38e46fb4395a3ef6057af8739671a; unb=2204945884244; isg=BJ6eJbSOcc19cKfZY9WcPM4g5TDgX2LZVpaSbUgnCuHcaz5FsO-y6cQZZfHnyFrx; UTUSER=968844706; l=eB_v1BV7g1xumw3tBOfahurza77OSCdYYuPzaNbMiOCPO0565xkNW6l8d-8BC31RhssMR3W8tjC2BeYBm3xonxvt-H3Lzzkmn; tfstk=cIZfBVbE1O62Wqj3fq6P7k2x0lodZQKSUMc3h_mvKzeAUXNfi-1EtCcTCe3xqY1..; DNT=0; SID=QQAAAAA5v2WiDQAyAAAyOTNkMzBlYmEwNIEAdBCDNMhNZB-QfEBm6K5q; USERID=968844706; x5check_ele=X4M0jNBbec%2Buyiqj6MWskFCO4b9mtI5g7aqpHJAtlg0%3D; t_eleuc4=id4=0%40BA%2FvuHCrrRkTbsWA6zTxKjPbhbG9LzEtlCEF6w%3D%3D; track_id=1635989856|c6b6f96c1e14e03a803a91726e97cc7a4c5f91a8ab25a9ff36|82b0e699c456e93b63098ae6307ddfd6; vip_edition=1.0.2; tzyy=a5ac7b42f0e295ae177c3caa38135ab0; ubt_ssid=59fbe99f-c0e6-45a0-92e1-eccb6f99ea1c_2021-11-04; cna=Qwn7GeqtSxECAXdVHv6YuSSH"
	//	//cookies := setCookieRaw(cookie)
	//	//c.SetCookies("https://presale.dmall.com",cookies)
	//	c.OnRequest(func(r *colly.Request) {
	//		r.Headers.Set("cookie",cookie)
	//		r.Headers.Set("Accept","*/*")
	//		r.Headers.Set("Accept-Encoding","gzip, deflate, br")
	//		r.Headers.Set("Accept-Language","zh-cn")
	//		r.Headers.Set("Content-Length","158007")
	//		r.Headers.Set("Content-Type","application/x-www-form-urlencoded; charset=utf-8")
	//
	//		r.Headers.Set("a-orange-dq","appKey%3D24894833%26appVersion%3D10.10.3%26clientAppIndexVersion%3D1120220822230700999")
	//		r.Headers.Set("a-orange-env","prod")
	//		r.Headers.Set("A-SLIDER-Q","appKey%3D24894833%26ver%3D1660802689775")
	//		r.Headers.Set("f-refer","mtop")
	//		r.Headers.Set("x-account-site","eleme")
	//		r.Headers.Set("x-alsc-pageid","a2f7o.14330926__YXQbInytsr0DAFzCF3Sw7ilm__OB9S2pq_")
	//		r.Headers.Set("x-app-conf-v","0")
	//		r.Headers.Set("x-appkey","24894833")
	//		r.Headers.Set("x-app-ver","10.10.3")
	//		r.Headers.Set("x-bx-version","6.5.51")
	//		r.Headers.Set("x-cmd-v","0%7C0")
	//		r.Headers.Set("x-decode-ua","true")
	//		r.Headers.Set("x-deviceInfo","bmV0X3R5cGU6V0lGSSBsYXRpdHVkZToyOS42MDkyODggbG9uZ2l0dWRlOjEwNi41MDE4MTA%3D")
	//		r.Headers.Set("x-ele-ua","RenderWay%2FminiProgram%20MiniAppId%2F2021001155698149%20MiniAppVersion%2F10.11.0%20DeviceId%2FYXQbInytsr0DAFzCF3Sw7ilm%20AppName%2Feleme%20Apple%2FiPhone12%2C1%20iOS%2F14.6%20Eleme%2F10.10.3%20MiniHostVersion%2F10.7.0%20channel%2Fmini_app%20subChannel%2Fminiapp.eleme%20subSubChannel%2Fios.default.scheme_bf34178265cc4f77929036b6ac3212bc%20BusinessComeFrom%2Fistore")
	//		r.Headers.Set("x-features","11")
	//		r.Headers.Set("x-ltraffic-src","%257B%2522afid%2522%253A%2522afc_launch%255Eme.ele.ios.eleme%255E1012_Initiactive%255E16F22C58-35D9-4D17-B4E3-91780FCE9A5F_1661236308990.252930%2522%252C%2522pvid%2522%253A%2522a2f7o.14330915__YXQbInytsr0DAFzCF3Sw7ilm__OB9RuaT_%2522%257D")
	//		r.Headers.Set("x-miniapp-env","%7B%22nbsource%22%3A%22online%22%2C%22nbsn%22%3A%22ONLINE%22%7D")
	//		r.Headers.Set("x-miniapp-id-taobao","2021001155698149")
	//		r.Headers.Set("x-miniapp-version","10.20220712.175646")
	//		r.Headers.Set("x-mini-wua","HHnB_dZEYfy8v9ePmmxH03pGrxOUJp8XNDKJsgm9VL%2F7TV3NMBA4Acvxl%2F4QxvYkHnNJc1v7AwDxtBtx7DtzpinovJBzWZdHfj9vZk6vAoPgNFRQgMSJ6019T%2Bnwo0y1rU33D88nYjayYLuo0nIYbdpUXFdg4Fh5B20QRO%2FdiphNkK7k%3D")
	//		r.Headers.Set("x-nq","WiFi")
	//		r.Headers.Set("x-page-name","TRVAppPageViewController")
	//		r.Headers.Set("x-pv","6.3")
	//		r.Headers.Set("x-sgext","JAVPX2yg5GgJSPvnU11ZaF96b358e2h7aGxteWhsfH5peW59bHlsf2dsb39teW9%2Fb39vf29%2Fb39vf3x%2FfH98f29sb39vf3x%2FfH58fnx%2BfH58fXx%2BfH8%3D")
	//		r.Headers.Set("x-sid","12eb8f9293d30eba048e2a84378ec399")
	//		r.Headers.Set("x-sign","izMRVI002xAAKVSZC8SZdPlu1KZEKUSZVrjj30AuUvA25gAqjB3nOTMep8F4G1wxMMD%2Fkro3%2BWZfI2Dd98gQ1Q0lBjlUiVSZVIlUmV")
	//		r.Headers.Set("x-t","1661236971")
	//		r.Headers.Set("x-ttid","2021001155698149%40eleme_iphone_10.10.3")
	//		r.Headers.Set("x-uid","2204945884244")
	//		r.Headers.Set("x-umt","J%2FVLPYhLOgZ4fzWCyCCKhkWasL7q6t%2Fw")
	//		r.Headers.Set("x-utdid","YXQbInytsr0DAFzCF3Sw7ilm")
	//		r.Headers.Set("Connection","keep-alive")
	//		r.Headers.Set("Host","alsc-buy2.ele.me")
	//	})
	//
	//	c.OnError(func(_ *colly.Response, err error) {
	//		log.Println("Something went wrong:", err)
	//	})
	//
	//	c.OnResponse(func(r *colly.Response) {
	//		// 判断code
	//		if r.StatusCode == 200 {
	//			body := new(response)
	//			json.Unmarshal(r.Body,body)
	//			log.Println(body.Data.WareInfos)
	//
	//
	//			//fmt.Println(string(r.Body))
	//		}
	//	})
	//
	//	wua := "HuWH_omDx%2FBAqnQ%2Fq4BScFwtFOmjf0IjDHHNSPc%2BqFwRuuSt17%2FyEJuvS7tCNXwD30BKsO%2FI6TDrgOYK82ASmFi9zmdyak55VRMdrlream9aTmhjEoVEzzbjuq7xTP4bzv9FPguMWaN5BtlIu8eDGP43yVlxOIQC%2F4%2FDU611%2FsfchG4I7vHHEU8C%2F2nMtJ8LQEGOasoYe7L9VaMFKpgMhveBIBwAtEakQkLRulfzACnqyqr6qyr%2B5vVuaB8F6oN11NaarwjMogMM3z0JtwNOfzGXfaKOWd2ls4GCL7x67EKAbPnY%3D"
	//	url := fmt.Sprintf("https://alsc-buy2.ele.me/gw/mtop.trade.elemeorder.create/1.0?type=originaljson&rnd=176BEF4F689C468A063ACBC2AF5ADB83&wua=%s",wua)
	//
	//	base64.StdEncoding.EncodeToString([]byte(wua))
	//
	//	data := map[string]string{
	//		"param":"{\"vendorId\":\"69\"}",
	//		"token":"5d8608fe-8b70-45d7-8762-229525c563ee",
	//		"ticketName":"57EA2312C6780CE6D2DDD460D3711A2B381A3AB41080A6C3A8ED4AD82AA633C8AEED76C96D9B79414CE75041843E8321BBADC00201C2578034E1D66F132EB17BC0682EFC71118DCC6BEBCF217AC7990D72D48B098F5BE5A288388FDC82AACE5AB22A45AFC55D419C9FE78A366728193BA02D605A80CD8F7861AB97432B5E7E94",
	//	}
	//	_ = c.Post(url,data)
	//
	//	time.Sleep(10*time.Second)
	//}
}