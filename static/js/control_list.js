$(function() {
    var candidatetable = $('#candidatelist').DataTable({
        stateSave: true,
        "ajax": {
            "url": "/getcandidatelist",
            "dataSrc": ''
        },
        columns: [{
            data: 0
        }, {
            data: 1
        }, {
            data: 2
        }, {
            data: 3
        }, {
            data: 4
        }, {
            data: 5
        }, {
            "data": null,
            "defaultContent": "<button  class='btn btn-link' data-toggle='modal' data-target='#patientmodal'>More</button>"
        }]

    });

    $('#patient tbody').on('click', 'button', function() {
        var pdata = candidatetable.row($(this).parents('tr')).data();
        $.ajax({
            url: '/getcandidatebyid',
            type: 'POST',
            async: false,
            data: {
                "id": pdata[0]
            },
        }).done(function(output) {
            var result = JSON.stringify(output);
            //           alert(result);
        }).fail(function() {
            console.log("error");
            alert("error");
        }).always(function() {
            console.log("complete");
        });
    });



});