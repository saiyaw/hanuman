function get_candidate_labels() {
    var result = null;
    $.ajax({
        type: 'GET',
        url: '/getlabellist',
        async: false,
        success: function(data) {
            result = eval(data);
        }
    });
    return result;

}


function init_list_page() {
    var labels = get_candidate_labels();
    if (labels != null) {
        $.each(labels, function(index, value) {

            str = "<span class='button-checkbox'>   <button type='button' class='btn' data-color='primary' id=label_" + value[0] + ">" + value[1] + "</button>  <input type='checkbox' class='hidden'/> </span>";

            $("#div_label_list").append(str);
        });
    }

}