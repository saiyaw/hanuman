function delete_one_label(){
	var label_data = $("#select_admin_label").select2("data")[0];
	$.ajax({
        type: "POST",
        async: false,
        url: "/deletelabel",
        data: {
            "labelid": label_data.id
        },
        success: function(r) {
            result = r;
            location.reload();
        }
    });
}