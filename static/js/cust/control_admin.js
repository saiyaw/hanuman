$(function() {

	$("#select_admin_label").select2({
		theme: "bootstrap",
		language: "zh-CN",
		tags: "true",
		tokenSeparators: [','],
		data: get_label_list()
	});

	$('#btn_admin_add_label').click(function() {
		add_new_label();
	});

	$('#btn_admin_delete_label').click(function() {
		delete_one_label();
	});

});