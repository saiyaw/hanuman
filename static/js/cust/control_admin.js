$(function() {

	$("#select_admin_label").select2({
		theme: "bootstrap",
		language: "zh-CN",
		tags: "true",
		//		tokenSeparators: [','],
		data: get_keyword_list()
	});

	$('#btn_admin_add_label').click(function() {
		add_new_keyword();
	});

	$('#btn_admin_delete_label').click(function() {
		delete_one_keyword();
	});

});


