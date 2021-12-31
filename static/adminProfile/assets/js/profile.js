
// 用户信息修改
	$("#editSubmit").click(function(){
		
			var postData = {
				
				authID : $("#authID").val(),
				step : "info",
				editAdminName : $("#editAdminName").val(),
				editAdminEmail : $("#editAdminEmail").val(),
				editAdminPhone : $("#editAdminPhone").val(),

			};
			
			$.ajax({
				url: "/adminProfile/"+$("#authID").val()+"/profile",
				data: postData,
				type: "POST",
				dataType: "json",
				success: function(data) {
					showAlert(data.status,"<h3>" + data.contents + "</h3>");
					if(data.status == "ok"){
						$("#editAdminName").val(data.adminName);
						$("#editAdminEmail").val(data.adminEmail);
						$("#editAdminPhone").val(data.adminPhone);
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
				changeAdminCPass : $("#changeAdminCPass").val(),
				changeAdminNPass : $("#changeAdminNPass").val(),
				changeAdminVPass : $("#changeAdminVPass").val(),

			};
			
			$.ajax({
				url: "/adminProfile/"+$("#authID").val()+"/profile",
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