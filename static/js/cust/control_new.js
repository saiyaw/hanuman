$(function() {

	$("#sgender").select2({
		theme: "bootstrap",
		minimumResultsForSearch: Infinity
	});
	$("#ilabel").select2({
			theme: "bootstrap",
			language: "zh-CN",
			tags: "true",
		//	allowClear: true,
			tokenSeparators: [',', ' '],
			data: get_keyword_list(),
			/*
					ajax: {
						url: "/getkeywordlist",
						dataType: "json",
						delay: 250,
						data: function(params) {
							return params;
						},
						processResults: function(data, params) {
							params.page = params.page || 1;
							return {
								results: data[1],
								pagination: {
									more: (params.page * 30) < data.total_count
								}
							}
						},
			cache: true
		}*/
	});



$("#btn_new_save").click(function() {
	var result = save_candidate_info();
	if (parseInt(result) > 0) {
		window.location.href = "/candidate";
	}

});

$('#btn_new_save_comment').click(function() {
	save_candidate_comment();

});


$('#btn_new_add_label').click(function() {

	var keyword_data = JSON.stringify($("#ilabel").select2("data"));
//	
	add_new_keyword();

});

});