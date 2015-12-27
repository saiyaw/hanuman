$(function() {
	$("#iattachment").fileinput();

    $("#btn_new_save").click(function() {
        save_candidate_info();
        window.location.href="/candidate";
    });

    $('#btn_new_save_comment').click(function(){
    	save_candidate_comment();

    });

});