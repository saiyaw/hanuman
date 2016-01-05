$(function() {
    var candidatetable = $('#table_candidate_list').DataTable({
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
        }, {
            data: 8
        }, {
            data: 9
        }, {
            data: 10
        }, {
            data: 11
        }, {
            "data": null,
            "defaultContent": "<button class='btn btn-link'>More</button>"
        }]

    });
    // hide the column :id
    candidatetable.column(0).visible(false);
    candidatetable.column(8).visible(false);
    candidatetable.column(9).visible(false);
    candidatetable.column(10).visible(false);

    $('#table_candidate_list tbody').on('click', 'button', function() {
        var pdata = candidatetable.row($(this).parents('tr')).data();
        $.cookie("id", pdata[0]);
        window.location.href = "/candidate";

    });

    $('#table_candidate_list tfoot th').each(function() {
        var title = $(this).text();
        if (title != "") {
            $(this).html('<input type="text" placeholder="搜索 ' + title + '" />');
        }
    });

    candidatetable.columns().every(function() {
        var that = this;

        $('input', this.footer()).on('keyup change', function() {
            if (that.search() !== this.value) {
                that.search(this.value).draw();
            }
        });
    });



});