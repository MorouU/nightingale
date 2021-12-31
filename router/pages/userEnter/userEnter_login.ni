<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <title>NI - 账号登录</title>
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
        <div class="form-data pos">
            <div class="change-login">
                <p class="account_number on" style="font-size:30px">账号登录</p>
            </div>
		
            <div class="form1">
                <p class="p-input pos">
                    <label for="loginAccount">用户名 / 邮箱</label>
                    <input type="text" id="loginAccount" required>
                </p>
                <p class="p-input pos">
                    <label for="loginPass">密码</label>
                    <input type="password" id="loginPass" autocomplete="new-password" required>
                </p>
				<p class="p-input pos">
                    <label for="loginAuthCode">验证码</label>
                    <input type="text" id="loginAuthCode" required>
                </p>
				<div style="text-align:center;width:280px,height:80px">
					<span  style="padding:0;border:none;cursor:pointer;">
						<a href="#">
							<img src="data:images/image.png;base64,{{ .authCode }}" id="authCodeImage" style="width:330px;height:80px;box-shadow: 0px 0px 4px 2px #66afe9;border-radius:15px"/>
						</a>
					</span>
				</div>
				<br>
            </div>
            <div class="r-forget cl">
                <a href="/userEnter/register" class="z">账号注册</a>
                <a href="/userEnter/forget" class="y">忘记密码</a>
            </div>
            <button id="userLoginButton" class="lang-btn">登录</button>
		
            <div class="third-party">
                <a href="/page/login/qhq-run#" class="log-qq icon-qq-round"></a>
                <a href="/page/login/qhq-run#" class="log-qq icon-weixin"></a>
                <a href="/page/login/qhq-run#" class="log-qq icon-sina1"></a>
            </div>
            <p class="right">Powered by (c) 2021 (●'◡'●)Morouu~</p>
        </div>
			
    </div>
	
</div>

<script src="/static/userEnter/js/jquery.js"></script>
<script src="/static/userEnter/js/agree.js"></script>
</body>
</html>