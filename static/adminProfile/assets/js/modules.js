
// 上一页
	$("#previous").click(function(){
		
			var postData = {
				
				authID : $("#authID").val(),
				step : "previous",
				currentPage : $("#currentPage").text(),

			};
			
			$.ajax({
				url: "/adminProfile/"+$("#authID").val()+"/"+$("#userID").val()+"/modules",
				data: postData,
				type: "POST",
				dataType: "json",
				success: function(data) {
					if(data.status == "ok"){
						
						$("#moduleTableRows").html(data.html);
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
				url: "/adminProfile/"+$("#authID").val()+"/"+$("#userID").val()+"/modules",
				data: postData,
				type: "POST",
				dataType: "json",
				success: function(data) {
					if(data.status == "ok"){
						
						$("#moduleTableRows").html(data.html);
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
				url: "/adminProfile/"+$("#authID").val()+"/"+$("#userID").val()+"/modules",
				data: postData,
				type: "POST",
				dataType: "json",
				success: function(data) {
					if(data.status == "ok"){
						
						$("#moduleTableRows").html(data.html);
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
				findModule : $("#findInputSearch").val(),
				
			};
			
			$.ajax({
				url: "/adminProfile/"+$("#authID").val()+"/"+$("#userID").val()+"/modules",
				data: postData,
				type: "POST",
				dataType: "json",
				success: function(data) {
					if(data.status == "ok"){
						
						$("#moduleTableRows").html(data.html);
						$("#currentPage").text(data.page);
						$("#totalPage").text(data.total);
						
					}
				},
				error: function(data){
					console.log(data)
				}
			});
	});
	
// 返回

	$("#moduleBack").click(function(){
		
			location.href = "/adminProfile/"+$("#authID").val()+"/userList";
	});
	
// 删除
function delModule(mID){
	
	$("#delete-module").modal("show");
	$("#deleteModuleID").val(mID);
	
}	

	$("#moduleDeleteConfirm").click(function(){
		
		var postData = {
				
			authID : $("#authID").val(),
			step : "delete",
			userID : $("#userID").val(),
			moduleID : $("#deleteModuleID").val(),
		};
			
		$.ajax({
			url: "/adminProfile/"+$("#authID").val()+"/"+$("#userID").val()+"/modules",
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
	
// 信息保存
function editModule(mID){
	
	var postData = {
				
			authID : $("#authID").val(),
			step : "view",
			userID : $("#userID").val(),
			moduleID : mID,
		};
			
		$.ajax({
			url: "/adminProfile/"+$("#authID").val()+"/"+$("#userID").val()+"/modules",
			data: postData,
			type: "POST",
			dataType: "json",
			success: function(data) {
					
				showAlert(data.status,"<h3>" + data.contents + "</h3>");
					
				if(data.status == "ok"){
						
						
						$("#editModuleID").val(mID);
						$("#editModuleTitle").text("模块ID -> "+mID);
						$("#editModuleName").val(data.moduleName);
						$("#editModuleData").val(data.moduleData);
						$("#editModuleCode").text(data.moduleCode);
						
						$("#edit-module").modal("show");
						
					
						
				}

			},
			error: function(data){
				console.log(data)
			}
		});
	
}

$("#editModuleSave").click(function(){
		
			var postData = {
				
				authID : $("#authID").val(),
				step : "save",
				moduleID : $("#editModuleID").val(),
				moduleName : $("#editModuleName").val(),
				moduleData : $("#editModuleData").val(),
				moduleCode : $("#editModuleCode").val(),
				
			};
			
			$.ajax({
				url: "/adminProfile/"+$("#authID").val()+"/"+$("#userID").val()+"/modules",
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