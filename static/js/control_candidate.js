$(function() {

	init_candidate_page();

	$("#btn_candidate_back").click(function() {
		window.history.back(-1);
	});

	$('#btn_candidate_save_comment').click(function() {
		save_candidate_comment();
	});



});