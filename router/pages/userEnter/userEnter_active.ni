<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <title>NI - 账号激活</title>
    <meta http-equiv="content-type" content="text/html; charset=utf-8" />
	<link href="/static/index/assets/img/apple-touch-icon.png" rel="icon">
    <link rel="stylesheet" href="/static/userEnter/css/base.css">
    <link rel="stylesheet" href="/static/userEnter/css/iconfont.css">
    <link rel="stylesheet" href="/static/userEnter/css/reg.css">
</head>
<body>
<div id="ajax-hook"></div>
<div class="wrap">
    <div class="wpn">
        <div class="form-data find_password">
            <div class="change-login">
                <p class="account_number on" style="font-size:30px">{{ .content }}</p>
				
            </div>
			<div class="change-login">
				<p class="account_number on" style="font-size:15px">若 3 秒后未跳转，可以 <a href="/userEnter/login"><span style="font-size:15px">点击此链接</span></a> 手动跳转。</p>
			</div>
			
            <p class="right_now">已有账号，<a href="/userEnter/login">马上登录</a></p>
             <p class="right">Powered by (c) 2021 (●'◡'●)Morouu~</p>
        </div>
    </div>
</div>
<script src="/static/userEnter/js/jquery.js"></script>
<script>
	window.setTimeout("location.href = '/userEnter/login'", 3000);
</script>
</body>
</html>