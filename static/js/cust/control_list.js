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
            data: 6
        }, {
            data: 7
        },{
            data: 8
        },{
            data: 9
        },{
            "data": null,
            "defaultContent": "<button class='btn btn-link'>More</button>"
        }]

    });

    $('#candidatelist tbody').on('click', 'button', function() {
        var pdata = candidatetable.row($(this).parents('tr')).data();
        $.cookie("id", pdata[0]);
        window.location.href = "/candidate";

    });



});