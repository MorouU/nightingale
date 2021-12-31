<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <title>NI - 管理员登录</title>
    <meta name="description" content="">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    <meta name="robots" content="all,follow">
	<link href="/static/index/assets/img/apple-touch-icon.png" rel="icon">
    <link rel="stylesheet" href="/static/adminEnter/css/bootstrap.min.css">
    <link rel="stylesheet" href="/static/adminEnter/css/style.default.css" id="theme-stylesheet">
  </head>
  <body>
    <div class="page login-page">
      <div class="container d-flex align-items-center">
        <div class="form-holder has-shadow">
          <div class="row">
            <!-- Logo & Information Panel-->
            <div class="col-lg-6">
              <div class="info d-flex align-items-center">
                <div class="content">
                  <div class="logo">
                    <h1>欢迎登录</h1>
                  </div>
                  <p>—— 管理后台 ——</p>
                </div>
              </div>
            </div>
            <!-- Form Panel    -->
            <div class="col-lg-6 bg-white">
              <div class="form d-flex align-items-center">
                <div class="content">
                  
                    <div class="form-group">
                      <input  type="text" id="loginAccount" required data-msg="请输入用户名" placeholder="用户名 / 邮箱" class="input-material">
                    </div>
                    <div class="form-group">
                      <input  type="password" id="loginPass" required data-msg="请输入密码" placeholder="密码" class="input-material">
                    </div>
					<div class="form-group">
                      <input  type="text" id="loginAuthCode" required data-msg="请输入验证码" placeholder="验证码" class="input-material">
					  <br>
					  
                    </div>
					<div class="form-group">
						<div style="text-align:center;width:380px,height:80px">
						<span  style="padding:0;border:none;cursor:pointer;">
							<a href="#">
								<img src="data:images/image.png;base64,{{ .authCode }}" id="authCodeImage" style="width:380px;height:80px;box-shadow: 0px 0px 4px 2px #66afe9;border-radius:15px"/>
							</a>
						</span>
						</div>
					</div>
                    <button type="button" class="btn btn-primary" id="adminLoginButton">登录</button>
                  
                  <br />
                  <small>没有账号?</small><a href="#" class="signup">&nbsp;不给注册</a>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <!-- JavaScript files-->
    <!-- Main File-->
    <script src="/static/adminEnter/js/jquery.js"></script>
	<script src="/static/adminEnter/js/agree.js"></script>
  </body>
</html>