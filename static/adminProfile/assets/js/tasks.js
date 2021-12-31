
// 上一页
	$("#previous").click(function(){
		
			var postData = {
				
				authID : $("#authID").val(),
				step : "previous",
				currentPage : $("#currentPage").text(),

			};
			
			$.ajax({
				url: "/adminProfile/"+$("#authID").val()+"/"+$("#userID").val()+"/tasks",
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
				url: "/adminProfile/"+$("#authID").val()+"/"+$("#userID").val()+"/tasks",
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
				url: "/adminProfile/"+$("#authID").val()+"/"+$("#userID").val()+"/tasks",
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
				findTask : $("#findInputSearch").val(),
				
			};
			
			$.ajax({
				url: "/adminProfile/"+$("#authID").val()+"/"+$("#userID").val()+"/tasks",
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
	
// 返回

	$("#taskBack").click(function(){
		
			location.href = "/adminProfile/"+$("#authID").val()+"/userList";
	});
	
// 删除
function delTask(tID){
	
	$("#delete-task").modal("show");
	$("#deleteTaskID").val(tID);
	
}	

	$("#taskDeleteConfirm").click(function(){
		
		var postData = {
				
			authID : $("#authID").val(),
			step : "delete",
			userID : $("#userID").val(),
			taskID : $("#deleteTaskID").val(),
		};
			
		$.ajax({
			url: "/adminProfile/"+$("#authID").val()+"/"+$("#userID").val()+"/tasks",
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
function editTask(tID){
	
	var postData = {
				
			authID : $("#authID").val(),
			step : "view",
			userID : $("#userID").val(),
			taskID : tID,
		};
			
		$.ajax({
			url: "/adminProfile/"+$("#authID").val()+"/"+$("#userID").val()+"/tasks",
			data: postData,
			type: "POST",
			dataType: "json",
			success: function(data) {
					
				showAlert(data.status,"<h3>" + data.contents + "</h3>");
					
				if(data.status == "ok"){
						
						
						$("#editTaskID").val(tID);
						$("#editTaskTitle").text("任务ID -> "+tID);
						$("#editTaskName").val(data.taskName);
						$("#editTaskData").val(data.taskData);
						
						if(data.taskStatus != "1"){
							$("#taskStatus_1").removeAttr("selected");
							$("#taskStatus_0").attr("selected","selected");
						}else{
							$("#taskStatus_0").removeAttr("selected");
							$("#taskStatus_1").attr("selected","selected");
						}
						
						$("#edit-task").modal("show");
						
					
						
				}
			},
			error: function(data){
				console.log(data)
			}
		});
	
}

$("#editTaskSave").click(function(){
		
			var postData = {
				
				authID : $("#authID").val(),
				step : "save",
				taskID : $("#editTaskID").val(),
				taskName : $("#editTaskName").val(),
				taskData : $("#editTaskData").val(),
				taskStatus : $("#editTaskStatus").val(),
				
			};
			
			$.ajax({
				url: "/adminProfile/"+$("#authID").val()+"/"+$("#userID").val()+"/tasks",
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