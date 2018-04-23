package models

import (
	"net/http"
	"strconv"
	"time"
	"io/ioutil"
	"strings"
	"github.com/tidwall/gjson"
)

func Query(page,query string)(respSearch []SearchDetail){
	var err error
	pagenum,err:=strconv.Atoi(page)
	checkError(err)
	bodyStr:="s="+query+"&type=1&offset="+strconv.Itoa( pagenum* 10 - 10)+"&limit=10"
	req,err:=http.NewRequest("POST","http://music.163.com/api/cloudsearch/pc",strings.NewReader(bodyStr))
	checkError(err)
	req.Header.Set("Referer","http://music.163.com/")
	req.Header.Set("Content-Type","application/x-www-form-urlencoded")
	client:=http.Client{Timeout:time.Second*5}
	resp,err:=client.Do(req)
	checkError(err)
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	result := gjson.Get(string(body), "result.songs")
	result.ForEach(func(key, value gjson.Result) bool {
		res:=value.String()
		var netr SearchDetail
		netr.Type="netease"
		netr.Name = gjson.Get(res, "name").String()
		netr.Id = int(gjson.Get(res, "id").Int())
		netr.Author = gjson.Get(res, "ar.0.name").String()
		netr.Url = "http://music.163.com/song/media/outer/url?id="+strconv.Itoa(netr.Id)+".mp3"
		respSearch=append(respSearch, netr)
		return true // keep iterating
	})
	return
}
func checkError(err error)  {
	if err!=nil{
		panic(err)
	}
}