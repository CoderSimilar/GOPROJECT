package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	// "strings"
)

type DictRequest struct {
	From 		string `form:"from"`
	To   		string `form:"to"`
	Query    	string `form:"query"`
	Transtype 	string `form:"transtype"`
}

func main() {
	client := &http.Client{}
	request := DictRequest{From: "zh", To: "en", Query: "你好", Transtype: "translang"}
	buf, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}
	var data = bytes.NewReader(buf)
	// var data = strings.NewReader(`from=zh&to=en&query=%E4%BD%A0%E5%A5%BD&transtype=translang&simple_means_flag=3&sign=232427.485594&token=6e635729fdf8b605b0733fba4fb1920d&domain=common&ts=1690282456632`)
	req, err := http.NewRequest("POST", "https://fanyi.baidu.com/v2transapi?from=zh&to=en", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("Acs-Token", "1690282440257_1690282456652_h39zIFzY5KyboIssX6C8jMnus9vlXr0iq0ZPG8oxg++NnqFZniiZ4r5vj5q5dbzG+03fV4VIZ0qNjUZr0DPGPQZBT+2W2aWJlYdA+KapCoRzliyUWLWqTDGkQWLBf0SPCq0Ag0i8UKxGZGvSoE7RFNquNESTZUaebORnDq88c41q8rOiCPmUCWR6oEAOCmqBRgi4GmxL6qRmISC653/SJTux3x7i5dV8Deufz5AJULFB7vMJoefOHneB+pZqKb1LW12Vy2UXfS0RQw9F1J1NCgInFhPhXuc+2PAeTq+wQHZ6c6aNtgd63xk35Jh6KQ+H0pVJOqxzm4QfMtjFFC/Z1e0X1tPJClIuCUWgPwb5sqcb5J4jrh1gccAXOtBJZ35x7t2EPCwlhSAmiCB04mq5oPR31jNNWSmSYc5V8fTJ9sbXK+Ea3P20MvWFPzHhzje39isYLWD1SEo1eEKljUQNQCmcwUaMAUnpLEg1B/k2JzP/K/+8BWsE8FzbeB+CnIop/7w4H6C70yRk4iSO3qtK3g==")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Cookie", "BIDUPSID=C75AA344FC9A3889E70270A74358CE58; PSTM=1690163741; BAIDUID=D82697B36B1BF87FEF3B567729BB1BB5:FG=1; BAIDUID_BFESS=D82697B36B1BF87FEF3B567729BB1BB5:FG=1; ZFY=dxt5ckOQWIQgG3MqGFrFYI75NP1V9eXCYbSw:B7:B7i64:C; REALTIME_TRANS_SWITCH=1; FANYI_WORD_SWITCH=1; HISTORY_SWITCH=1; SOUND_SPD_SWITCH=1; SOUND_PREFER_SWITCH=1; Hm_lvt_64ecd82404c51e03dc91cb9e8c025574=1690282441; Hm_lpvt_64ecd82404c51e03dc91cb9e8c025574=1690282441; ab_sr=1.0.1_OWFmMjFmZTU5Nzc5ZjhlMWExMmI5NGI1ZGU1MGU4MzRlMTZhZDMzNmZmZmJlMTAzZTgyM2JiMTkzNjNkNGU5MTk1M2Y1YTVhODk2OTgzNTBhODlmYmUzZGI2YzcwMjRlOGJmZDRjNWU1NDljZTAxYTA4YTYwYmMzNmY3MTA4YzYwYjcxYmY1OTc3ZTY5MDJjMzE2OWJlNDczMzQyODgwMg==")
	req.Header.Set("Origin", "https://fanyi.baidu.com")
	req.Header.Set("Referer", "https://fanyi.baidu.com/")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36 Edg/115.0.1901.183")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("sec-ch-ua", `"Not/A)Brand";v="99", "Microsoft Edge";v="115", "Chromium";v="115"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
}