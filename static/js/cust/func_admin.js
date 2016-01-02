function delete_one_label(labelid) {
    $.ajax({
        type: "POST",
        async: false,
        url: "/deletelabel",
        data: {
            "labelid": labelid
        },
        success: function(r) {
            result = r;
        }
    });
}

function delete_candidate_label(labelid) {
    $.ajax({
        type: "POST",
        async: false,
        url: "/deletecandidatelabelbylabel",
        data: {
            "labelid": labelid
        },
        success: function(r) {
            result = r;

        }
    });

}