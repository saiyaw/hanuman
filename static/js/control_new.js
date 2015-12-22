$(document).ready(function() {
    $('button[name=btn_new_next]').click(function() {
        var btnid = $(this).attr('id');
        var id = btnid.substr(btnid.lastIndexOf("_") + 1);
        display_screening_section(id);
        $(this).hide();
    });

    $("#btn_new_submit").click(function() {
        var result = get_all_answers();
        var s = "";
        $.each(result, function(index, value) {
            var questionid = index;
            var answer = value;
            s += questionid + ":" + answer + "|";
        });

        save_patient_answers(1, s);
        $('#divscreeningresult').show();
    });

});