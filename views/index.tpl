<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width,initial-scale=1,minimum-scale=1,maximum-scale=1,user-scalable=no" />
    <title></title>
    <script src="/static/js/mui.min.js"></script>
    <link href="/static/css/mui.min.css" rel="stylesheet"/>
    <link href="/static/css/index.css" rel="stylesheet"/>
    <link href="/static/css/audioplayer.css" rel="stylesheet"/>
</head>
<body>
<div class="mui-card">
    <div class="mui-card-header">个人音乐实验室（网易）</div>
    <div class="mui-card-content-inner">
        <div class="mui-input-row mui-search">
                <input type="search" class="mui-input-clear" placeholder="" id="search_input"/>
        </div>
        <div id="progressbar"></div>
        <ul class="mui-table-view mui-table-view-chevron" id="search_list">
        </ul>
    </div>
    <div class="mui-card-footer">
        <div id="pages" class="page-parent"></div>
    </div>
</div>
<div class="mui-card">
     <div class="mui-card-content-inner">
         <ul class="mui-table-view">
             <div class="mui-row">
                 <span class="mui-icon mui-icon-flag"></span>
                 <a class="mui-navigate-right" href="https://github.com/qiusunshine" target="_blank">源码下载</a>
             </div>
         </ul>
     </div>
</div>
<script type="text/javascript" src="/static/js/index.js" ></script>
<script type="text/javascript" src="/static/js/jquery.min.js" ></script>
<script type="text/javascript" src="/static/js/audioplayer.min.js" ></script>
</body>
</html>