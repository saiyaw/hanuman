$(function() {

	$("#btn_new_save").click(function() {
		var result = save_candidate_info();
		if (parseInt(result) > 0) {
			window.location.href = "/candidate";
		}

	});

	$('#btn_new_save_comment').click(function() {
		save_candidate_comment();

	});

});