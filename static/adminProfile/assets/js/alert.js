function showAlert(title,contents){
	if(contents != "<h3>undefined</h3>" && contents.length > 0){
		console.log(contents);
		$("#AlertTitle").text(title);
		$("#AlertBody").html(contents);
		$("#ThisAlert").modal('show');
		
	}
}