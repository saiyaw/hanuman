$(function() {

	alert($.cookie("md5"));

	$("#btn_candidate_back").click(function() {
		window.history.back(-1);
	});



});