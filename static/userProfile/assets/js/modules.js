
// 添加自定义模块
	$("#moduleAddSave").click(function(){
		
			var postData = {
				
				authID : $("#authID").val(),
				step : "add",
				moduleAddName : $("#moduleAddName").val(),
				moduleAddData : $("#moduleAddData").val(),
				moduleAddCode  : $("#moduleAddCode").val(),

			};
			
			$.ajax({
				url: "/userProfile/"+$("#authID").val()+"/modules",
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
	
// 修改自定义模块
	$("#moduleModifySave").click(function(){
		
			var postData = {
				
				authID : $("#authID").val(),
				step : "modify",
				moduleID : $("#moduleID").val(),
				moduleModifyName : $("#moduleModifyName").val(),
				moduleModifyData : $("#moduleModifyData").val(),
				moduleModifyCode : $("#moduleModifyCode").val(),

			};
			
			$.ajax({
				url: "/userProfile/"+$("#authID").val()+"/modules",
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

// 删除自定义模块
	$("#moduleDeleteConfirm").click(function(){
		
			var postData = {
				
				authID : $("#authID").val(),
				step : "delete",
				moduleID : $("#moduleID").val(),

			};
			
			$.ajax({
				url: "/userProfile/"+$("#authID").val()+"/modules",
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

// 查看自定义模块详情
function viewModuleCustom(mID){
	$("#moduleID").val(mID);
	$("#moduleModifyName").val($("#moduleCustomTitle_" + mID).text());
	$("#moduleModifyData").text($("#moduleCustomData_" + mID).text());
	
	var postData = {
				
				authID : $("#authID").val(),
				step : "view",
				moduleID : mID,

			};
	
	$.ajax({
		url: "/userProfile/"+$("#authID").val()+"/modules",
		data: postData,
		type: "POST",
		dataType: "json",
		success: function(data) {
			if(data.status == "ok"){
				$("#moduleModifyCode").text(data.code);
				
			}else{
				$("#moduleModifyCode").text("加载失败");
			}
			$("#moduleCustomView").modal('show');
		},
		error: function(data){
			console.log(data)
		}
	});
	
	
	
}

// 查看公共模块详情
function viewModule(mID){
	
	$("#moduleName").text($("#moduleTitle_" + mID).text());
	$("#moduleData").text($("#moduleData_" + mID).text());
	
	$("#moduleView").modal('show');
	
}

