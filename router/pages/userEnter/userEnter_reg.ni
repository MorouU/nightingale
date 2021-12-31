<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <title>NI - 账号注册</title>
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
                <p class="account_number on" style="font-size:30px">账号注册</p>
            </div>
                <form action="#">
                    <p class="p-input pos" >
                        <label for="regUser">用户名</label>
                        <input type="text" id="regUser" autocomplete="off" required>
                        <span class="tel-warn tel-err hide"><em></em><i class="icon-warn"></i></span>
                    </p>
					<p class="p-input pos">
                        <label for="regEmail">邮箱</label>
                        <input type="text" id="regEmail" autocomplete="off" required>
                        <span class="tel-warn tel-err hide"><em></em><i class="icon-warn"></i></span>
                    </p>
                    <p class="p-input pos" id="pwd">
                        <label for="regPass">输入密码</label>
                        <input type="password" id="regPass" required>
                        <span class="tel-warn pwd-err hide"><em></em><i class="icon-warn" style="margin-left: 5px"></i></span>
                    </p>
                    <p class="p-input pos" id="confirmpwd">
                        <label for="regRePass">确认密码</label>
                        <input type="password" id="regRePass" required>
                        <span class="tel-warn confirmpwd-err hide"><em></em><i class="icon-warn" style="margin-left: 5px"></i></span>
                    </p>
                <div class="reg_checkboxline pos">
                    <span class="z" id="regAgreeTrigger"><i class="icon-ok-sign boxcol" nullmsg="请同意!"></i></span>
                    <input type="hidden" id="regAgree" value="1">
                    <div class="Validform_checktip"></div>
                    <p>我已阅读并接受 <a href="/page/reg/qhq-run#">《使用协议》</a></p>
                </div>
                <button id="userRegisterButton" class="lang-btn">注册</button>
                <div class="bottom-info">已有账号，<a href="/userEnter/login">马上登录</a></div>
			</form>
                <p class="right">Powered by (c) 2021 (●'◡'●)Morouu~</p>
            </div>
        </div>
    </div>
    <script src="/static/userEnter/js/jquery.js"></script>
    <script src="/static/userEnter/js/agree.js"></script>
</body>
</html>
