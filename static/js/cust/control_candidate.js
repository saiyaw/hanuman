$(function() {

	init_candidate_page();

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

	/*
		$("#btn_candidate_back").click(function() {
			window.history.back(-1);
		});
	*/
	$('#btn_candidate_save_comment').click(function() {
		save_candidate_comment();
	});

	$("#btn_candidate_save").click(function() {
		update_candidate_info();
	});
});