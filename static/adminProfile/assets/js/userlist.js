
// 上一页
	$("#previous").click(function(){
		
			var postData = {
				
				authID : $("#authID").val(),
				step : "previous",
				currentPage : $("#currentPage").text(),

			};
			
			$.ajax({
				url: "/adminProfile/"+$("#authID").val()+"/userList",
				data: postData,
				type: "POST",
				dataType: "json",
				success: function(data) {
					if(data.status == "ok"){
						
						$("#userTableRows").html(data.html);
						$("#currentPage").text(data.page);
						$("#totalPage").text(data.total);
						
					}
				},
				error: function(data){
					console.log(data)
				}
			});
	});
	
// 下一页
	$("#next").click(function(){
		
			var postData = {
				
				authID : $("#authID").val(),
				step : "next",
				currentPage : $("#currentPage").text(),
				
			};
			
			$.ajax({
				url: "/adminProfile/"+$("#authID").val()+"/userList",
				data: postData,
				type: "POST",
				dataType: "json",
				success: function(data) {
					if(data.status == "ok"){
						
						$("#userTableRows").html(data.html);
						$("#currentPage").text(data.page);
						$("#totalPage").text(data.total);
						
					}
				},
				error: function(data){
					console.log(data)
				}
			});
	});

// 重置
	$("#findBack").click(function(){
		
			var postData = {
				
				authID : $("#authID").val(),
				step : "back",

			};
			
			$.ajax({
				url: "/adminProfile/"+$("#authID").val()+"/userList",
				data: postData,
				type: "POST",
				dataType: "json",
				success: function(data) {
					if(data.status == "ok"){
						
						$("#userTableRows").html(data.html);
						$("#currentPage").text(data.page);
						$("#totalPage").text(data.total);
						
					}
				},
				error: function(data){
					console.log(data)
				}
			});
	});

// 查找
	$("#findGo").click(function(){
		
			var postData = {
				
				authID : $("#authID").val(),
				step : "find",
				findAccount : $("#findInputSearch").val(),
				
			};
			
			$.ajax({
				url: "/adminProfile/"+$("#authID").val()+"/userList",
				data: postData,
				type: "POST",
				dataType: "json",
				success: function(data) {
					if(data.status == "ok"){
						
						$("#userTableRows").html(data.html);
						$("#currentPage").text(data.page);
						$("#totalPage").text(data.total);
						
					}
				},
				error: function(data){
					console.log(data)
				}
			});
	});
	
// 删除
function delUser(uID){
	$("#delete-user").modal("show");
	$("#deleteUserID").val(uID);
}	
	$("#userDeleteConfirm").click(function(){
		
		var postData = {
				
			authID : $("#authID").val(),
			step : "delete",
			userID : $("#deleteUserID").val(),
		};
			
		$.ajax({
			url: "/adminProfile/"+$("#authID").val()+"/userList",
			data: postData,
			type: "POST",
			dataType: "json",
			success: function(data) {
					
				showAlert(data.status,"<h3>" + data.contents + "</h3>");
					
				if(data.status == "refresh"){
						
					location.reload();
						
				}
			},
			error: function(data){
				console.log(data)
			}
		});
	});
	
// 编辑
function editUser(uID){
	
	var postData = {
				
		authID : $("#authID").val(),
		step : "view",
		userID : uID,
	};
	
	$.ajax({
			url: "/adminProfile/"+$("#authID").val()+"/userList",
			data: postData,
			type: "POST",
			dataType: "json",
			success: function(data) {
					
				if(data.status == "ok"){
					
					$("#editUserID").val(uID);
					$("#editUserTitle").text("用户ID -> "+uID);
					$("#editUserName").val(data.userName);
					$("#editUserEmail").val(data.userEmail);
					$("#editUserPhone").val(data.userPhone);
					$("#editUserPass").val("");
					
					if(data.userStatus != "1"){
						$("#userStatus_1").removeAttr("selected");
						$("#userStatus_0").attr("selected","selected");
					}else{
						$("#userStatus_0").removeAttr("selected");
						$("#userStatus_1").attr("selected","selected");
					}
					
					if(data.userLevel == "0"){
						$("#userLevel_1").removeAttr("selected");
						$("#userLevel_2").removeAttr("selected");
						$("#userLevel_0").attr("selected","selected");
					}else if(data.userLevel == "1"){
						$("#userLevel_0").removeAttr("selected");
						$("#userLevel_2").removeAttr("selected");
						$("#userLevel_1").attr("selected","selected");
					}else{
						$("#userLevel_0").removeAttr("selected");
						$("#userLevel_1").removeAttr("selected");
						$("#userLevel_2").attr("selected","selected");
					}
					
					$("#edit-user").modal("show");
						
				}
			},
			error: function(data){
				console.log(data)
			}
		});
		
}

	$("#editUserSave").click(function(){
		
		var postData = {
				
			authID : $("#authID").val(),
			step : "edit",
			userID : $("#editUserID").val(),
			userName : $("#editUserName").val(),
			userEmail : $("#editUserEmail").val(),
			userPhone : $("#editUserPhone").val(),
			userPass : $("#editUserPass").val(),
			userStatus : $("#editUserStatus").val(),
			userLevel : $("#editUserLevel").val(),
			
		};
			
		$.ajax({
			url: "/adminProfile/"+$("#authID").val()+"/userList",
			data: postData,
			type: "POST",
			dataType: "json",
			success: function(data) {
					
				showAlert(data.status,"<h3>" + data.contents + "</h3>");
					
				if(data.status == "refresh"){
						
					location.reload();
						
				}
			},
			error: function(data){
				console.log(data)
			}
		});
	});
	
// 新建用户

	$("#newUserOK").click(function(){
		
		var postData = {
				
			authID : $("#authID").val(),
			step : "new",
			userName : $("#newUserName").val(),
			userEmail : $("#newUserEmail").val(),
			userPhone : $("#newUserPhone").val(),
			userPass : $("#newUserPass").val(),
			userStatus : $("#newUserStatus").val(),
			userLevel : $("#newUserLevel").val(),
			
		};
			
		$.ajax({
			url: "/adminProfile/"+$("#authID").val()+"/userList",
			data: postData,
			type: "POST",
			dataType: "json",
			success: function(data) {
					
				showAlert(data.status,"<h3>" + data.contents + "</h3>");
					
				if(data.status == "refresh"){
						
					location.reload();
						
				}
			},
			error: function(data){
				console.log(data)
			}
		});
	});

// 用户任务

function tasks(uID){
	
	location.href = "/adminProfile/" + $("#authID").val() + "/" + uID + "/tasks" 
		
}

// 用户模块

function modules(uID){
	
	location.href = "/adminProfile/" + $("#authID").val() + "/" + uID + "/modules" 
		
}