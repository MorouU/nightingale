   $(function(){
   	$(".icon-ok-sign").click(function(){
	    verify = $(this).hasClass('boxcol') ? 0 : 1;
	    $(this).toggleClass('boxcol');
	    ob = $('input[name=agree]');
	    ob.val(ob.val()==0?1:0);
	    ob.val(ob.val() !== verify ? verify : ob.val());

	    ob.val() == 0 ? $('#errormsg').removeClass('hide').addClass('show') : $('#errormsg').removeClass('show').addClass('hide');
    });
   	//输入框输入时模拟placeholder效果
   	var oInput = $(".form-data input");
   	oInput.focus(function () {
		$(this).siblings("label").hide();
    });
   	oInput.blur(function () {
		if($(this).val()==""){
			$(this).siblings("label").show();
		}
	});
   	// 输入框内容变化按钮颜色发生变化
   	oInput.keyup(function () {
		if($(this).val()!="jquery.js"){
			$(".log-btn").removeClass("off")
		}else{
            $(".log-btn").addClass("off")
		}
    });
});
  
  	// 验证码刷新
	$("#authCodeImage").click(function(){
			$.ajax({
				url: "/userEnter/authCodeRefresh",
				data: {refresh: '1'},
				type: "POST",
				dataType: "json",
				success: function(data) {
					if(data.status == "ok"){
						$("#authCodeImage").attr("src","data:images/image.png;base64," + data.contents);
					}
				},
				error: function(data){
					console.log(data)
				}
			});
	});
	
	// 用户登录
	$("#userLoginButton").click(function(){
		
			var postData = {
				
				userAccount : $("#loginAccount").val(),
				loginPass : $("#loginPass").val(),
				authCode : $("#loginAuthCode").val(),

			};
			
			$.ajax({
				url: "/userEnter/login",
				data: postData,
				type: "POST",
				dataType: "json",
				success: function(data) {
					if(data.status == "ok"){
						alert(data.contents);
						location.href="/userProfile/" + data.uuid + "/home";
					}else{
						alert(data.contents);
						$("#authCodeImage").click();
					}
				},
				error: function(data){
					console.log(data)
				}
			});
	});
	
	// 用户注册
	$("#userRegisterButton").click(function(){
		
			var postData = {
				
				userName : $("#regUser").val(),
				email : $("#regEmail").val(),
				pass : $("#regPass").val(),
				repass : $("#regRePass").val(),
				agree : $("#regAgree").val(),

			};
			
			$.ajax({
				url: "/userEnter/register",
				data: postData,
				type: "POST",
				dataType: "json",
				success: function(data) {
					if(data.status == "ok"){
						alert(data.contents);
						location.href="/userEnter/login";
					}else{
						alert(data.contents);
					}
				},
				error: function(data){
					console.log(data)
				}
			});
	});
	
	$("#regAgreeTrigger").click(function(){
		if($("#regAgree").val() != "1"){
			$("#regAgree").val("1");
		}else{
			$("#regAgree").val("0");
		}
	});
	
	// 密码找回
	$("#forgetAuthEmailSend").click(function(){
		
			var postData = {
				
				authEmail : $("#forgetAuthEmail").val(),
				step : "1",

			};
			
			$.ajax({
				url: "/userEnter/forget",
				data: postData,
				type: "POST",
				dataType: "json",
				success: function(data) {
					if(data.status == "ok"){
						alert(data.contents);
						$("#findAuthEmail").attr("style","display:none");
						$("#changePass").attr("style","display:block");
					}else{
						alert(data.contents);
					}
				},
				error: function(data){
					console.log(data)
				}
			});
	});

	// 密码更变
	$("#userChangePass").click(function(){
		
			var postData = {
				
				newPass : $("#forgetNewPass").val(),
				newRePass : $("#forgetReNewPass").val(),
				authCode : $("#forgetEmailAuthCode").val(),
				step : "2",

			};
			
			$.ajax({
				url: "/userEnter/forget",
				data: postData,
				type: "POST",
				dataType: "json",
				success: function(data) {
					if(data.status == "ok"){
						alert(data.contents);
						location.href="/userEnter/login";
					}else{
						alert(data.contents);
					}
				},
				error: function(data){
					console.log(data)
				}
			});
	});

	
	