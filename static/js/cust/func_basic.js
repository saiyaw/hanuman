function insert_candidate_info() {
    var result = null;
    $.ajax({
        type: "POST",
        async: false,
        url: "/insertcandidate",
        data: {
            "fullname": $("#input_candidate_name").val(),
            "age": $("#input_candidate_age").val(),
            "gender": $("#select_candidate_gender").val(),
            "mobile": $("#input_candidate_mobile").val(),
            "email": $("#input_candidate_email").val(),
            "post": $("#input_candidate_post").val(),
            "city": $("#input_candidate_city").val(),
            "workyear": $("#input_candidate_workyear").val(),
            "company" : $("#input_candidate_company").val(),
        },
        success: function(r) {
            result = r;
            $.cookie("id", result);
        }
    });
    return result;
}

function update_candidate_info() {
    var result = null;
    $.ajax({
        type: "POST",
        async: false,
        url: "/updatecandidate",
        data: {
            "candidateid" : $.cookie("id"),
            "fullname": $("#input_candidate_name").val(),
            "age": $("#input_candidate_age").val(),
            "gender": $("#select_candidate_gender").val(),
            "mobile": $("#input_candidate_mobile").val(),
            "email": $("#input_candidate_email").val(),
            "post": $("#input_candidate_post").val(),
            "city": $("#input_candidate_city").val(),
            "workyear": $("#input_candidate_workyear").val(),
            "company" : $("#input_candidate_company").val(),
        },
        success: function(r) {
            result = r;
     //       $.cookie("id", result);
        }
    });
    return result;
}

function save_candidate_label() {
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


function insert_candidate_label() {
    var labels = $("#select_candidate_label").val();
    var labeldata = "";
    $.each(labels, function(index, value) {
        labeldata += value + "|";
    })
    if (labeldata.length == 0){
        return;
    }
    $.ajax({
        type: "POST",
        async: false,
        url: "/insertcandidatelabel",
        data: {
            "candidateid": $.cookie("id"),
            "labels": labeldata,
        },
        success: function(r) {
            result = r;
        }
    });
}