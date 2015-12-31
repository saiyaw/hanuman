function save_candidate_info() {
    var result = null;
    $.ajax({
        type: "POST",
        async: false,
        url: "/insertcandidate",
        data: {
            "fullname": $("#iname").val(),
            "age": $("#iage").val(),
            "gender": $("#sgender").val(),
            "mobile": $("#imobile").val(),
            "email": $("#iemail").val(),
            "post": $("#ipost").val(),
            "city": $("#icity").val(),
            "workyear": $("#iworkyear").val(),
        },
        success: function(r) {
            result = r;
            $.cookie("id", result);
        }
    });
    return result;
}