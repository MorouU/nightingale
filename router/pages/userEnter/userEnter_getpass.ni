<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <title>NI - 密码找回</title>
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
                <p class="account_number on" style="font-size:30px">密码找回</p>
            </div>
			<form action="#" id="findAuthEmail" style="display:block">
            <div class="form1">
                <p class="p-input pos">
                    <label for="forgetAuthEmail">邮箱</label>
                    <input type="text" id="forgetAuthEmail" required>
                </p>
				<button id="forgetAuthEmailSend" class="lang-btn"  >获取邮箱验证码</button>
            </div>
		</form>
		<form action="#" id="changePass" style="display:none">
            <div class="form1">
				<p class="p-input pos">
                    <label for="newPass">输入新密码</label>
                    <input type="text" id="forgetNewPass" required>
                </p>
				<p class="p-input pos">
                    <label for="reNewPass">确认新密码</label>
                    <input type="text" id="forgetReNewPass" required>
                </p>
				<p class="p-input pos">
                    <label for="emailAuthCode">邮箱验证码</label>
                    <input type="text" id="forgetEmailAuthCode" required>
                </p>
				<button id="userChangePass" class="lang-btn" >修改密码</button>
            </div>
		</form>
            <p class="right_now">已有账号，<a href="/userEnter/login">马上登录</a></p>
             <p class="right">Powered by (c) 2021 (●'◡'●)Morouu~</p>
        </div>
    </div>
</div>
<script src="/static/userEnter/js/jquery.js"></script>
<script src="/static/userEnter/js/agree.js"></script>
</body>
</html>