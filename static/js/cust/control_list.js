$(function() {
    var candidatetable = $('#table_candidate_list').DataTable({
        stateSave: true,
        language: {

            "sProcessing": "处理中...",
            "sLengthMenu": "显示 _MENU_ 项结果",
            "sZeroRecords": "没有匹配结果",
            "sInfo": "显示第 _START_ 至 _END_ 项结果，共 _TOTAL_ 项",
            "sInfoEmpty": "显示第 0 至 0 项结果，共 0 项",
            "sInfoFiltered": "(由 _MAX_ 项结果过滤)",
            "sInfoPostFix": "",
            "sSearch": "搜索:",
            "sUrl": "",
            "sEmptyTable": "表中数据为空",
            "sLoadingRecords": "载入中...",
            "sInfoThousands": ",",
            "oPaginate": {
                "sFirst": "首页",
                "sPrevious": "上页",
                "sNext": "下页",
                "sLast": "末页"
            },
            "oAria": {
                "sSortAscending": ": 以升序排列此列",
                "sSortDescending": ": 以降序排列此列"
            }

        },
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