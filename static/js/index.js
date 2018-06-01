function addPages(){
    var f = document.getElementById("pages");
    var childs = f.childNodes;
    if(childs.length<1){
        var str = "<div class=\"mui-numbox page-child\"  data-numbox-min='1'>" +
            "  <button class=\"mui-btn mui-numbox-btn-minus\" type=\"button\">-</button>" +
            "  <input class=\"mui-numbox-input\" id='pageNum' type=\"number\" />" +
            "  <button class=\"mui-btn mui-numbox-btn-plus\" type=\"button\">+</button>" +
            "  </div>" +
            "  <button class=\"mui-btn page-child\" type=\"button\" id='pageGo'>Go</button>";
        f.innerHTML=str;
        mui("#pages").numbox();
        document.getElementById('pageGo').addEventListener("click",function (ev) {
            console.log("click");
            page = mui("#pages").numbox().getValue();
            goSearch(document.getElementById("search_input").value);
        });
    }
}
function addHtml(data){
    var str='';
    for(var i=0; i<data.length ;i++){
        var r = "<li class=\"mui-table-view-cell\">"+
            "<a href=\"javascript:void(0);\" title=\""+data[i].Url+"\" onclick='clickList(this"+","+data[i].Id+")'>"+data[i].Name+"</a>"+
            "<span class=\"mui-badge mui-badge-success\">"+data[i].Author+"</span>"+
            "</li>";
        str += r; //拼接str
    }
    document.getElementById('search_list').innerHTML=str;
    if(page===1){
        addPages();
        mui("#pages").numbox().setValue(1);
    }
}
function addMusicPlayer(url) {
    console.log("加载audioplayer成功");
    var d = document.createElement('div');
    d.setAttribute("id", "mypop");
    d.innerHTML= "<div class=\"mui-popup mui-popup-in\" style=\"display: block;width: 90%;\">\n" +
        "<div class=\"mui-popup-inner\">" +
        "<audio preload=\"auto\" controls>\n" +
        "<source src=\""+url+"\">\n" +
        "</audio></div>\n" +
        "<div class=\"mui-popup-buttons\">\n" +
        "<span id='my-popup-button-download'class=\"mui-popup-button\">下载</span>\n" +
        "<span id='my-popup-button-back' class=\"mui-popup-button\">返回</span>\n" +
        "</div>\n" +
        "</div>" +
        "<div class=\"mui-popup-backdrop mui-active\" style=\"display: block;\"></div>";
    document.body.appendChild(d);
    $( function() {
        $('audio').audioPlayer();
        $("#my-popup-button-download").click(function (ev) {
            window.open(url);
            //$('#mypop').remove();
        });
        $("#my-popup-button-back").click(function (ev) {
            $('#mypop').remove();
        });
    });
}
function chooseGo(url) {
    mui.confirm('', '选择您想要的操作', ['播放', '取消', '下载'], function(e) {
        switch (e.index){
            case 0:
                //mui.toast(that.getAttribute("title"));
                addMusicPlayer(url);
                break;
            case 1:
                break;
            default:
                window.open(url);
                break;
        }
    })
}
function clickList(that,id) {
    if(that.getAttribute("title")==="no"){
       getDesc(id);
    }else{
        chooseGo(that.getAttribute("title"));
    }
}
function goSearch(e){
    type = $("input[name='music_type']:checked").val();
    console.log(type);
    mui("#progressbar").progressbar().show();
    mui.ajax('/sou/search',{
        data:{q:e, p:page, type:type},
        dataType:'json',//服务器返回json格式数据
        type:'post',//HTTP请求类型
        timeout:10000,//超时时间设置为10秒；
        success:function(data){
            mui("#progressbar").progressbar().hide();
            addHtml(data);
        },
        error:function(xhr,type,errorThrown){
            //异常处理；
            mui("#progressbar").progressbar().hide();
            mui.toast('出错：'+errorThrown,{ duration:'long', type:'div' });
            console.log(type);
        }
    });
}
function getDesc(id){
    mui("#progressbar").progressbar().show();
    mui.ajax('/sou/desc',{
        data:{id:id, type:type},
        dataType:'json',
        type:'post',
        timeout:10000,
        success:function(data){
            mui("#progressbar").progressbar().hide();
            console.log(data);
            if(data.Url!==null&&data.Url!=="")chooseGo(data.Url);
            else mui.toast('居然没有获取到数据！',{ duration:'long', type:'div' });
        },
        error:function(xhr,type,errorThrown){
            //异常处理；
            mui("#progressbar").progressbar().hide();
            mui.toast('出错：'+errorThrown,{ duration:'long', type:'div' });
            console.log(type);
        }
    });
}
var page = 1;
var type="netease";
var search = document.getElementById("search_input");
search.addEventListener("keypress", function(e) { 
	if (e.keyCode == "13") {
            var str = search.value; 
		    if(str!==""){
		        page = 1;
		    	goSearch(str);
		    } else {
		    	mui.toast('请先输入关键词',{ duration:'short', type:'div' })
		    }
        }
});
mui.init();
