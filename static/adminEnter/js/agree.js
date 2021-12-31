
	// 验证码刷新
	$("#authCodeImage").click(function(){
			$.ajax({
				url: "/adminEnter/authCodeRefresh",
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
	
	// 管理员登录
	$("#adminLoginButton").click(function(){
		
			var postData = {
				
				userAccount : $("#loginAccount").val(),
				loginPass : $("#loginPass").val(),
				authCode : $("#loginAuthCode").val(),

			};
			
			$.ajax({
				url: "/adminEnter/login",
				data: postData,
				type: "POST",
				dataType: "json",
				success: function(data) {
					if(data.status == "ok"){
						alert(data.contents);
						location.href="/adminProfile/" + data.uuid + "/profile";
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
