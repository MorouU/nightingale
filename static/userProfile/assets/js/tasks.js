
// 上一页
	$("#previous").click(function(){
		
			var postData = {
				
				authID : $("#authID").val(),
				step : "previous",
				currentPage : $("#currentPage").text(),

			};
			
			$.ajax({
				url: "/userProfile/"+$("#authID").val()+"/tasks",
				data: postData,
				type: "POST",
				dataType: "json",
				success: function(data) {
					if(data.status == "ok"){
						
						$("#taskTableRows").html(data.html);
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
				url: "/userProfile/"+$("#authID").val()+"/tasks",
				data: postData,
				type: "POST",
				dataType: "json",
				success: function(data) {
					if(data.status == "ok"){
						
						$("#taskTableRows").html(data.html);
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
				url: "/userProfile/"+$("#authID").val()+"/tasks",
				data: postData,
				type: "POST",
				dataType: "json",
				success: function(data) {
					if(data.status == "ok"){
						
						$("#taskTableRows").html(data.html);
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
				findTaskID : $("#findInputSearch").val(),
				
			};
			
			$.ajax({
				url: "/userProfile/"+$("#authID").val()+"/tasks",
				data: postData,
				type: "POST",
				dataType: "json",
				success: function(data) {
					if(data.status == "ok"){
						
						$("#taskTableRows").html(data.html);
						$("#currentPage").text(data.page);
						$("#totalPage").text(data.total);
						
					}
				},
				error: function(data){
					console.log(data)
				}
			});
	});

// 启用/停用
function controlTask(tID){
	
	var statusText = "";
	
	if ($("#taskStatus_"+tID).text() == "启用"){
		statusText = "你确定停用这个任务吗 ?";
	}else{
		statusText = "你确定启用这个任务吗 ?";
	}

	$("#statusTask").text(statusText);
	$("#controlTaskID").val(tID);
	$("#control-task").modal("show");
	
	
}
	$("#taskControlConfirm").click(function(){
		
		var postData = {
				
			authID : $("#authID").val(),
			step : "control",
			taskID : $("#controlTaskID").val(),
			
		};
			
		$.ajax({
			url: "/userProfile/"+$("#authID").val()+"/tasks",
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


// 删除
function delTasks(tID){
	$("#delete-task").modal("show");
	$("#deleteTaskID").val(tID);
}	
	$("#taskDeleteConfirm").click(function(){
		
		var postData = {
				
			authID : $("#authID").val(),
			step : "delete",
			taskID : $("#deleteTaskID").val(),
		};
			
		$.ajax({
			url: "/userProfile/"+$("#authID").val()+"/tasks",
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
	
// 获取API
function getTaskApi(tID){
	
	var postData = {
				
		authID : $("#authID").val(),
		step : "getApi",
		taskID : tID,
	};
	
	$.ajax({
			url: "/userProfile/"+$("#authID").val()+"/tasks",
			data: postData,
			type: "POST",
			dataType: "json",
			success: function(data) {
					
				if(data.status == "ok"){
						
					$("#taskApiName").text(data.taskName);
					$("#taskApiData").text(data.taskApi);
					$("#getApi-task").modal("show");
						
				}
			},
			error: function(data){
				console.log(data)
			}
		});
		
}

	
	$("#taskApiData").select(function(){
		
		document.execCommand('copy');
		showAlert("ok","<h3>已复制选中内容</h3>");
		
	});


// 新建任务

	$("#newTaskNext").click(function(){
		
		var postData = {
				
			authID : $("#authID").val(),
			step : "new",
			taskName : $("#newTaskName").val(),
			taskModule : $("#newTaskModule").val(),
			
		};
			
		$.ajax({
			url: "/userProfile/"+$("#authID").val()+"/tasks",
			data: postData,
			type: "POST",
			dataType: "json",
			success: function(data) {
					
				if(data.status == "nextCustom"){
						
					$("#customTaskName").text(data.taskName);
					$("#customCode").text(data.taskCode);
					$("#customModuleID").val(data.moduleID);
					$("#new-custom-task").modal("show");
						
				}
				
				if(data.status == "nextDefault"){
						
					$("#defaultTaskName").text(data.taskName);
					$("#defaultModuleID").val(data.moduleID);
					$("#default-task").modal("show");
						
				}
				
				if(data.status == "nextRequest"){
						
					$("#requestTaskName").text(data.taskName);
					$("#requestModuleID").val(data.moduleID);
					$("#request-task").modal("show");
						
				}
				
				if(data.status == "nextFish"){
						
					$("#fishTaskName").text(data.taskName);
					$("#fishModuleID").val(data.moduleID);
					$("#fish-task").modal("show");
						
				}
				
				if(data.status == "nextPort"){
					
					$("#portTaskName").text(data.taskName);
					$("#portModuleID").val(data.moduleID);
					$("#port-task").modal("show");
					
				}
				
				if(data.status == "error"){
					showAlert(data.status,"<h3>" + data.contents + "</h3>");
				}
			},
			error: function(data){
				console.log(data)
			}
		});
	});

// 自定义模块

	$("#customOK").click(function(){
		
		var postData = {
				
			authID : $("#authID").val(),
			step : "customAdd",
			taskName : $("#newTaskName").val(),
			taskData : $("#newTaskData").val(),
			taskStatus : $("#newTaskStatus").val(),
			taskCode : $("#customCode").val(),
			moduleID : $("#customModuleID").val(),
			
		};
			
		$.ajax({
			url: "/userProfile/"+$("#authID").val()+"/tasks",
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
	
	$("#customBack").click(function(){
		
		$("#new-task").modal("show");
		
	})
	
// 默认模块

	$("#defaultOK").click(function(){
		
		var postData = {
				
			authID : $("#authID").val(),
			step : "defaultAdd",
			taskName : $("#newTaskName").val(),
			taskData : $("#newTaskData").val(),
			taskStatus : $("#newTaskStatus").val(),
			moduleID : $("#defaultModuleID").val(),
			defaultCurrentUrl : document.getElementById("defaultCurrentUrl").checked.toString(),
			defaultRefererUrl : document.getElementById("defaultRefererUrl").checked.toString(),
			defaultCookie : document.getElementById("defaultCookie").checked.toString(),
			defaultOS : document.getElementById("defaultOS").checked.toString(),
			defaultBrowser : document.getElementById("defaultBrowser").checked.toString(),
			defaultScreenResolution : document.getElementById("defaultScreenResolution").checked.toString(),
			defaultWebPage : document.getElementById("defaultWebPage").checked.toString(),
			defaultScreenShot : document.getElementById("defaultScreenShot").checked.toString(),
			
		};
			
		$.ajax({
			url: "/userProfile/"+$("#authID").val()+"/tasks",
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
	
	$("#defaultBack").click(function(){
		
		$("#new-task").modal("show");
		
	})

// 域内请求

$("#requestGET").click(

	function(){
			
			if($("#requestGET").attr("checked") != "checked"){
				
				$("#requestGET").attr("checked","checked");
				$("#requestGETS").attr("style","display:block");
				
			}else{
				
				$("#requestGET").removeAttr("checked");
				$("#requestGETS").attr("style","display:none");
				
			}
			
});

$("#requestPOST").click(

	function(){
			
			if($("#requestPOST").attr("checked") != "checked"){
				
				$("#requestPOST").attr("checked","checked");
				$("#requestPOSTS").attr("style","display:block");
				
			}else{
				
				$("#requestPOST").removeAttr("checked");
				$("#requestPOSTS").attr("style","display:none");
				
			}
			
});

$("#requestFILE").click(

	function(){
			
			if($("#requestFILE").attr("checked") != "checked"){
				
				let fileFormatS = "";
				$("#requestFILE").attr("checked","checked");
				$("#requestFILES").attr("style","display:block");
				
				
				
				
			}else{
				
				$("#requestFILE").removeAttr("checked");
				$("#requestFILES").attr("style","display:none");
				$("#fileSpace").html("");
				
			}
			
});


$("#requestOK").click(function(){
	
	
		let formData = new FormData();
		formData.append("authID",$("#authID").val());
		formData.append("step","requestAdd");
		formData.append("taskName",$("#newTaskName").val());
		formData.append("taskData",$("#newTaskData").val());
		formData.append("taskStatus",$("#newTaskStatus").val());
		formData.append("moduleID",$("#requestModuleID").val());
		formData.append("requestUrl",$("#requestUrl").val());
		
		if($("#requestGET").attr("checked") == "checked"){
			formData.append("requestGET",$("#requestGETV").val());
		}
		
		if($("#requestPOST").attr("checked") == "checked"){
			formData.append("requestPOST",$("#requestPOSTV").val());
		}
		
		if($("#requestFILE").attr("checked") == "checked"){
			if($("#requestFILEV")[0].files.length > 0){
				for(var i=0;i<$("#requestFILEV")[0].files.length;i++){
					formData.append("requestFILE",$("#requestFILEV")[0].files[i]);
				}
			}
			
		}
			
		$.ajax({
			url: "/userProfile/"+$("#authID").val()+"/tasks",
			data: formData,
			type: "POST",
			contentType: false,
			processData: false,
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
	
	$("#requestBack").click(function(){
		
		$("#new-task").modal("show");
		
	})

// 钓鱼

$("#fishCustom").click(

	function(){
			
			if($("#fishCustom").attr("checked") != "checked"){
				
				$("#fishCustom").attr("checked","checked");
				$("#fishURLS").attr("style","display:block");
				
				$("#fishThis").removeAttr("checked");
				$("#fishFILES").attr("style","display:none");
				
			}else{
				
				$("#fishCustom").removeAttr("checked");
				$("#fishURLS").attr("style","display:none");
				
			}
			
});
	
$("#fishThis").click(

	function(){
			
			if($("#fishThis").attr("checked") != "checked"){
				
				$("#fishThis").attr("checked","checked");
				$("#fishFILES").attr("style","display:block");
				
				$("#fishCustom").removeAttr("checked");
				$("#fishURLS").attr("style","display:none");
				
			}else{
				
				$("#fishThis").removeAttr("checked");
				$("#fishFILES").attr("style","display:none");
				
			}
			
});


$("#fishOK").click(function(){
	
	
		let formData = new FormData();
		formData.append("authID",$("#authID").val());
		formData.append("step","fishAdd");
		formData.append("taskName",$("#newTaskName").val());
		formData.append("taskData",$("#newTaskData").val());
		formData.append("taskStatus",$("#newTaskStatus").val());
		formData.append("moduleID",$("#fishModuleID").val());
		
		if($("#fishCustom").attr("checked") == "checked"){
			formData.append("fishUrl",$("#fishUrl").val());
		}
		
		if($("#fishThis").attr("checked") == "checked"){
			formData.append("fishFILE",$("#fishFILEV")[0].files[0]);
		}
			
		$.ajax({
			url: "/userProfile/"+$("#authID").val()+"/tasks",
			data: formData,
			type: "POST",
			contentType: false,
			processData: false,
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

$("#fishBack").click(function(){
		
	$("#new-task").modal("show");
		
})

// 端口探测

$("#portOK").click(function(){
	
		var postData = {
			
			authID : $("#authID").val(),
			step : "portAdd",
			taskName : $("#newTaskName").val(),
			taskData : $("#newTaskData").val(),
			taskStatus : $("#newTaskStatus").val(),
			moduleID : $("#portModuleID").val(),
			portIP:$("#portUrl").val(),
			portScan:$("#portScan").val(),
			
		};
			
		$.ajax({
			url: "/userProfile/"+$("#authID").val()+"/tasks",
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

$("#portBack").click(function(){
		
	$("#new-task").modal("show");
		
})
	
// 查看详情

function viewTask(tID){
	
	var postData = {
				
			authID : $("#authID").val(),
			step : "viewTask",
			taskID : tID,
			
		};
	
	$.ajax({
			url: "/userProfile/"+$("#authID").val()+"/tasks",
			data: postData,
			type: "POST",
			dataType: "json",
			success: function(data) {
				
				if(data.status == "redirect"){
						
					location.href=data.url;
						
				}
			},
			error: function(data){
				console.log(data)
			}
		});
	
	
}