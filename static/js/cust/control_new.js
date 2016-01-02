$(function() {

	$("#sgender").select2({
		theme: "bootstrap",
		minimumResultsForSearch: Infinity
	});
	$("#select_new_label").select2({
		theme: "bootstrap",
		language: "zh-CN",
		tags: "true",
		data: get_label_list(),
	});



	$("#btn_new_save").click(function() {
		var result = save_candidate_info();
		if (parseInt(result) > 0) {
			insert_candidate_label();
			window.location.href = "/candidate";
		}

	});

	$('#btn_new_save_comment').click(function() {
		save_candidate_comment();


	});



});