function add_comment_into_list(createtime, content) {
	var component = '<li class="list-group-item"><h4 class="list-group-item-heading">' + createtime + '</h4><p class="list-group-item-heading">' + content + '</p></li>';
	$("#ul_candidate_comment_list").append(component);

}

function add_attachment_into_list(filename) {
	var component = '<li class="list-group-item"><a class="list-group-item" href="/static/file/' + filename + '">' + filename + '</a></li>';
	$("#ul_candidate_attachment_list").append(component);

}

function get_candidate_info() {
	var result = null;
	var id = $.cookie("id");
	$.ajax({
		type: "POST",
		async: false,
		url: "/getcandidatebyid",
		data: {
			"id": id
		},
		success: function(output) {
			result = output;
		}
	});

	return result;
}

function get_comment_list() {
	var result = null;
	var id = $.cookie("id");
	$.ajax({
		type: "POST",
		async: false,
		url: "/getcommentlist",
		data: {
			"candidateid": id
		},
		success: function(output) {
			result = output;
		}
	});

	return result;

}

function get_attachment_list() {
	var result = null;
	var id = $.cookie("id");
	$.ajax({
		type: "POST",
		async: false,
		url: "/getattachments",
		data: {
			"candidateid": id
		},
		success: function(output) {
			result = output;
		}
	});

	return result;

}

function fill_candidate_label(){
    var labels = get_candidate_label_list();
    var labeldata = [];
    $.each(labels, function(index, value){
        content = value['labelid'];
        labeldata.push(content);
    }); 

    $('#select_candidate_label').val(labeldata).trigger("change");
}

function fill_candidate_info(info) {
	$('#iname').val(info.Fullname);
	$('#iage').val(info.Age);
	$('#sgender').val(info.Gender);
	$('#iworkyear').val(info.Workyear);
	$('#iemail').val(info.Email);
	$('#imobile').val(info.Mobile);
	$('#icity').val(info.City);
	$('#ipost').val(info.Post);
	$('#icompany').val(info.Company);

	fill_candidate_label();
}

function update_candidate_info() {
	var result = null;
	$.ajax({
		type: "POST",
		async: false,
		url: "/updatecandidate",
		data: {
			"id": $.cookie("id"),
			"fullname": $("#iname").val(),
			"age": $("#iage").val(),
			"gender": $("#sgender").val(),
			"mobile": $("#imobile").val(),
			"email": $("#iemail").val(),
			"city": $("#icity").val(),
			"post": $("#ipost").val(),
			"workyear": $("#iworkyear").val(),
			"company": $("#icompany").val(),
		},
		success: function(r) {
			result = r;
		}
	});
	return result;
}

function update_candidate_label() {
	remove_candidate_label();
	insert_candidate_label();
}

function remove_candidate_label() {
	$.ajax({
		type: "POST",
		async: false,
		url: "/deletecandidatelabel",
		data: {
			"candidateid": $.cookie("id"),
		},
		success: function(r) {
			result = r;
		}
	});
}




function init_comment_list() {
	var list = get_comment_list();
	$.each(list, function(index, value) {
		add_comment_into_list(ParseTimetoStr(value[3]), value[2]);
	});
}

function init_attachment_list() {
	var list = get_attachment_list();
	$.each(list, function(index, value) {
		add_attachment_into_list(value[2]);
	});
}


function init_candidate_page() {
	// init basic information
	var candidateinfo = get_candidate_info();
	fill_candidate_info(candidateinfo);

	// init comment
	init_comment_list();

	// init attachment
	init_attachment_list();
}