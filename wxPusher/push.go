package wxPusher

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var appId string = "AT_nOCJxifRddZZhC2sLJGeEYLtYDNIUlgU"

var uid = "UID_1vn6xPjOlFyZPKIQxUkCzveIqpyN"

type Payload struct {
	AppToken    string   `json:"appToken"`
	Content     string   `json:"content"`
	Summary     string   `json:"summary"`
	ContentType int      `json:"contentType"`
	TopicIds    []string `json:"topicIds"`
	Uids        []string `json:"uids"`
	Url         string   `json:"url"`
}

func Push(summary string) {
	uid := []string{
		"UID_1vn6xPjOlFyZPKIQxUkCzveIqpyN", //myself
		//"UID_1l7ITMCWRYiAl11U6C6unL55MmUA",
		//"UID_IIEICEbDUYLKfK9rLRz8PnZDMiZy",  //向俊
	}
	postUrl := "http://wxpusher.zjiecode.com/api/send/message"
	form := new(Payload)
	form.AppToken = "AT_KKgNkezIWUwRa8Uh1Co8BBeaOWy93VCo"
	form.Content = "notice"
	form.Summary = summary
	form.ContentType = 1
	form.Uids = uid

	data, err := json.Marshal(form)
	fmt.Println(string(data))
	body := bytes.NewBufferString(string(data))
	rsp, err := http.Post(postUrl, "application/json", body)
	if err != nil {
		panic(err)
	}
	defer rsp.Body.Close()
	body_byte, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body_byte))
}
