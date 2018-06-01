package models

import (
	"strconv"
	"net/http"
	"strings"
	"time"
	"io/ioutil"
	"github.com/tidwall/gjson"
)
type KuWo struct {
}

func (ctx *KuWo)Search(q,p string)(searchResult []SearchResult){
	var err error
	pagenum,err:=strconv.Atoi(p)
	checkError(err)
	bodyStr:="all="+q+"&ft=music&itemset=web_2013&pn="+strconv.Itoa(pagenum-1)+"&rn=10&rformat=json&encoding=utf8"
	req,err:=http.NewRequest("GET","http://search.kuwo.cn/r.s?"+bodyStr,nil)
	checkError(err)
	req.Header.Set("Referer","http://player.kuwo.cn/webmusic/play")
	client:=http.Client{Timeout:time.Second*10}
	resp,err:=client.Do(req)
	checkError(err)
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	bodyS:=string(body)
	bodyS=strings.Replace(bodyS,"'","\"",-1)
	result := gjson.Parse(bodyS).Get("abslist")
	result.ForEach(func(key, value gjson.Result) bool {
		res:=value.String()
		var netr SearchResult
		netr.Type="kuwo"
		netr.Name = gjson.Get(res, "SONGNAME").String()
		netr.Id = gjson.Get(res, "MUSICRID").String()
		netr.Id=strings.Replace(netr.Id,"MUSIC_","",-1)
		netr.Author = gjson.Get(res, "ARTIST").String()
		netr.Url = "no"
		searchResult=append(searchResult, netr)
		return true // keep iterating
	})
	return
}
func (ctx *KuWo)GetDesc(id string)(descResults DescResult){
	req,err:=http.NewRequest("GET","http://player.kuwo.cn/webmusic/st/getNewMuiseByRid?rid=MUSIC_"+id,nil)
	checkError(err)
	req.Header.Set("Referer","http://m.kuwo.cn/yinyue/"+id)
	client:=http.Client{Timeout:time.Second*10}
	resp,err:=client.Do(req)
	checkError(err)
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	bodyS:=string(body)
	ss:=strings.Split(bodyS,"mp3path")
	if ss[1]!=""&&ss[1]!="><"{
		descResults.Id=id
		descResults.Type="kuwo"
		mp3d1s:=strings.Split(bodyS,"mp3dl")
		descResults.Url="http://"+mp3d1s[1]+"/resource/"+ss[1]
		descResults.Url=strings.Replace(descResults.Url,">","",-1)
		descResults.Url=strings.Replace(descResults.Url,"</","",-1)
	}
	return
}