$(function() {

	$("#select_candidate_label").select2({
		theme: "bootstrap",
		language: "zh-CN",
		tags: "true",
		data: get_label_list()
	});	

	$("#iattachment").fileinput({
		language: "zh",
		showCaption: true,
		uploadUrl: "/insertattachment",
		maxFileCount: 1,
		autoReplace: true,
		maxFileSize: 10240,
		showPreview: false,
		uploadExtraData: {
			"candidateid": $.cookie("id")
		}
	}).on('fileuploaded', function(event, data, previewId, index) {
		location.reload();
	});

	init_candidate_page();

	$('#btn_candidate_save_comment').click(function() {
		save_candidate_comment();
	});

	$("#btn_candidate_save").click(function() {
		update_candidate_info();
		update_candidate_label();
	});
});