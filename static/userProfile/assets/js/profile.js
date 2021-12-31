
// 用户信息修改
	$("#editSubmit").click(function(){
		
			var postData = {
				
				authID : $("#authID").val(),
				step : "info",
				editUserName : $("#editUserName").val(),
				editUserEmail : $("#editUserEmail").val(),
				editUserPhone : $("#editUserPhone").val(),

			};
			
			$.ajax({
				url: "/userProfile/"+$("#authID").val()+"/profile",
				data: postData,
				type: "POST",
				dataType: "json",
				success: function(data) {
					showAlert(data.status,"<h3>" + data.contents + "</h3>");
					if(data.status == "ok"){
						$("#editUserName").val(data.userName);
						$("#editUserEmail").val(data.userEmail);
						$("#editUserPhone").val(data.userPhone);
					}
				},
				error: function(data){
					console.log(data)
				}
			});
	});
	
// 用户密码修改
	$("#changeSubmit").click(function(){
		
			var postData = {
				
				authID : $("#authID").val(),
				step : "pass",
				changeUserCPass : $("#changeUserCPass").val(),
				changeUserNPass : $("#changeUserNPass").val(),
				changeUserVPass : $("#changeUserVPass").val(),

			};
			
			$.ajax({
				url: "/userProfile/"+$("#authID").val()+"/profile",
				data: postData,
				type: "POST",
				dataType: "json",
				success: function(data) {
					showAlert(data.status,"<h3>" + data.contents + "</h3>");
				},
				error: function(data){
					console.log(data)
				}
			});
	});