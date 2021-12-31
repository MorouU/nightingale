
// 上一页
	$("#previous").click(function(){
		
			var postData = {
				
				authID : $("#authID").val(),
				step : "previous",
				currentPage : $("#currentPage").text(),

			};
			
			$.ajax({
				url: "/userProfile/"+$("#authID").val()+"/"+$("#taskID").val()+"/view",
				data: postData,
				type: "POST",
				dataType: "json",
				success: function(data) {
					if(data.status == "ok"){
						
						$("#viewTableRows").html(data.html);
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
				url: "/userProfile/"+$("#authID").val()+"/"+$("#taskID").val()+"/view",
				data: postData,
				type: "POST",
				dataType: "json",
				success: function(data) {
					if(data.status == "ok"){
						
						$("#viewTableRows").html(data.html);
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
				url: "/userProfile/"+$("#authID").val()+"/"+$("#taskID").val()+"/view",
				data: postData,
				type: "POST",
				dataType: "json",
				success: function(data) {
					if(data.status == "ok"){
						
						$("#viewTableRows").html(data.html);
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
				findRecordID : $("#recordSearch").val(),
				
			};
			
			$.ajax({
				url: "/userProfile/"+$("#authID").val()+"/"+$("#taskID").val()+"/view",
				data: postData,
				type: "POST",
				dataType: "json",
				success: function(data) {
					if(data.status == "ok"){
						
						$("#viewTableRows").html(data.html);
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

	$("#viewBack").click(function(){
		
			location.href = "/userProfile/"+$("#authID").val()+"/tasks";
	});
	
// 删除
function delRecord(rID){
	$("#delete-record").modal("show");
	$("#deleteRecordID").val(rID);
}	
	$("#recordDeleteConfirm").click(function(){
		
		var postData = {
				
			authID : $("#authID").val(),
			step : "delete",
			recordID : $("#deleteRecordID").val(),
		};
			
		$.ajax({
			url: "/userProfile/"+$("#authID").val()+"/"+$("#taskID").val()+"/view",
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
	$("#taskSave").click(function(){
		
			var postData = {
				
				authID : $("#authID").val(),
				step : "save",
				taskID : $("#taskID").val(),
				taskName : $("#taskName").val(),
				taskData : $("#taskData").val(),
				
			};
			
			$.ajax({
				url: "/userProfile/"+$("#authID").val()+"/"+$("#taskID").val()+"/view",
				data: postData,
				type: "POST",
				dataType: "json",
				success: function(data) {
					
					showAlert(data.status,"<h3>" + data.contents + "</h3>");
					
					if(data.status == "ok"){
						
						$("#taskTitle").val(data.taskName);
						$("#taskName").val(data.taskName);
						$("#taskData").text(data.taskData);
						
					}
				},
				error: function(data){
					console.log(data)
				}
			});
	});

// 详情

function viewRecord(rID){
	
	var postData = {
				
		authID : $("#authID").val(),
		taskID : $("#taskID").val(),
		step : "view",
		recordID : rID,
				
	};
			
			$.ajax({
				url: "/userProfile/"+$("#authID").val()+"/"+$("#taskID").val()+"/view",
				data: postData,
				type: "POST",
				dataType: "json",
				success: function(data) {
					
					if(data.status == "ok"){
						
						if(data.returnType == "custom"){
							
							let jsonData = JSON.parse(data.reContents);
							
							let method = jsonData["method"];
							let ip = jsonData["ip"];
							let time = jsonData["time"];
							let agent = jsonData["agent"];
							let headers = JSON.stringify(jsonData["headers"]);
							
							$("#customRecordID").text(rID);
							$("#customMethod").text(method);
							$("#customIP").text(ip);
							$("#customTime").text(time);
							$("#customAgent").text(agent);
							$("#customHeaders").text(headers);
							
							$("#custom-record").modal("show");
							
						}
						
						if(data.returnType == "public_default"){
							
							let jsonData = JSON.parse(data.reContents);
							
							let method = jsonData["method"];
							let ip = jsonData["ip"];
							let time = jsonData["time"];
							let agent = jsonData["agent"];
							let headers = JSON.stringify(jsonData["headers"]);
							
							$("#defaultRecordID").text(rID);
							$("#defaultMethod").text(method);
							$("#defaultIP").text(ip);
							$("#defaultTime").text(time);
							$("#defaultAgent").text(agent);
							$("#defaultHeaders").text(headers);
							
							
							$("#defaultExtra").html(data.extraContents);
							$("#default-record").modal("show");
							
						}
						
						if(data.returnType == "public_request"){
							
							let jsonData = JSON.parse(data.reContents);
							
							let method = jsonData["method"];
							let ip = jsonData["ip"];
							let time = jsonData["time"];
							let agent = jsonData["agent"];
							let headers = JSON.stringify(jsonData["headers"]);
							
							$("#requestRecordID").text(rID);
							$("#requestMethod").text(method);
							$("#requestIP").text(ip);
							$("#requestTime").text(time);
							$("#requestAgent").text(agent);
							$("#requestHeaders").text(headers);

							$("#request-record").modal("show");
							
						}
						
						if(data.returnType == "public_fish"){
							
							let jsonData = JSON.parse(data.reContents);
							
							let method = jsonData["method"];
							let ip = jsonData["ip"];
							let time = jsonData["time"];
							let agent = jsonData["agent"];
							let headers = JSON.stringify(jsonData["headers"]);
							
							$("#fishRecordID").text(rID);
							$("#fishMethod").text(method);
							$("#fishIP").text(ip);
							$("#fishTime").text(time);
							$("#fishAgent").text(agent);
							$("#fishHeaders").text(headers);

							$("#fish-record").modal("show");
							
						}
						
						if(data.returnType == "public_port"){
							
							let jsonData = JSON.parse(data.reContents);
							
							let method = jsonData["method"];
							let ip = jsonData["ip"];
							let time = jsonData["time"];
							let agent = jsonData["agent"];
							let headers = JSON.stringify(jsonData["headers"]);
							let portResult = JSON.parse(data.extraContents);
							
							$("#portRecordID").text(rID);
							$("#portMethod").text(method);
							$("#portIP").text(ip);
							$("#portTime").text(time);
							$("#portAgent").text(agent);
							$("#portHeaders").text(headers);
							$("#portOpen").text(portResult.opens);
							$("#portClosed").text(portResult.closeds);

							$("#port-record").modal("show");
							
						}
						
					}
				},
				error: function(data){
					console.log(data)
				}
			});
	
}
