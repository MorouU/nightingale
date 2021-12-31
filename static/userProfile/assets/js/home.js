
// 上一页
	$("#previous").click(function(){
		
			var postData = {
				
				authID : $("#authID").val(),
				step : "previous",
				currentPage : $("#currentPage").text(),

			};
			
			$.ajax({
				url: "/userProfile/"+$("#authID").val()+"/home",
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
				url: "/userProfile/"+$("#authID").val()+"/home",
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
				url: "/userProfile/"+$("#authID").val()+"/home",
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
				url: "/userProfile/"+$("#authID").val()+"/home",
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

