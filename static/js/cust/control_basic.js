$(function() {

	$("#select_candidate_gender").select2({
		theme: "bootstrap",
		minimumResultsForSearch: Infinity
	});

	$("#select_candidate_label").select2({
		theme: "bootstrap",
		language: "zh-CN",
		tags: "true",
		data: get_label_list(),
	});

	$("#btn_candidate_save").click(function() {
		var current_url = $(location).attr('pathname');
		if (current_url == "/new") {
			var result = insert_candidate_info();
			if (parseInt(result) > 0) {
				save_candidate_label();
				window.location.href = "/candidate";
			}
		} else if (current_url == "/candidate") {
			update_candidate_info();
			save_candidate_label();
		} else {
			// do nothing.
		}

	});

});


