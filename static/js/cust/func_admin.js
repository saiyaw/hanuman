function delete_one_keyword(){
	var keyword_data = $("#select_admin_label").select2("data")[0];
	$.ajax({
        type: "POST",
        async: false,
        url: "/deletekeyword",
        data: {
            "id": keyword_data.id
        },
        success: function(r) {
            result = r;
            location.reload();
        }
    });
}