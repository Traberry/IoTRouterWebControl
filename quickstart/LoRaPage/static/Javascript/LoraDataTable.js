/**
 * Created by Administrator on 2018/8/30.
 */

/* ------组织Table  Begin  -------*/
var data = [];
var titles = ['ID', '名称', '显示名称', '是否有网关'];
$(function () {
    var table = $('#organizationTable').DataTable({
        "ajax": {
            url: "/lora/organizationTable",
            dataSrc: "Result"
        },
        columns: [
            {data: "ID"},
            {data: "Name"},
            {data: "DisplayName"},
            {data: "CanHaveGateways"}
        ],
        "pagingType": "full_numbers",
        "bSort": true,
        "language": {
            "sProcessing": "处理中...",
            "sLengthMenu": "显示 _MENU_ 项结果",
            "sZeroRecords": "没有匹配结果",
            "sInfo": "显示第 _START_ 至 _END_ 项结果，共 _TOTAL_ 项",
            "sInfoEmpty": " 显示第 0 至 0 项结果，共 0 项",
            "sInfoFiltered": "(由 _MAX_ 项结果过滤)",
            "sInfoPostFix": "",
            "sSearch": "搜索:",
            "sUrl": "",
            "sEmptyTable": "未添加组织",
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
            },

        },
        "aLengthMenu": [5, 10, 25],//也可以直接用一维数组设置数量
        "order": [[0, 'asc']]
    });
    $('#organizationTable tbody').on('click', 'tr', function () {
        if ( $(this).hasClass('selected') ) {
            $(this).removeClass('selected');
        }
        else {
            table.$('tr.selected').removeClass('selected');
            $(this).addClass('selected');
        }
    });
})

/* ------用户Table  Begin  -------*/
var userData = [];
$(function () {
    var table = $('#userTable').DataTable({
        "ajax": {
            url: "/lora/userTable",
            dataSrc: "Result"
        },
        columns: [
            {data: "UserName"},
            {data: "IsActive"},
            {data: "IsAdmin"}
        ],
        "pagingType": "full_numbers",
        "bSort": true,
        "language": {
            "sProcessing": "处理中...",
            "sLengthMenu": "显示 _MENU_ 项结果",
            "sZeroRecords": "没有匹配结果",
            "sInfo": "显示第 _START_ 至 _END_ 项结果，共 _TOTAL_ 项",
            "sInfoEmpty": " 显示第 0 至 0 项结果，共 0 项",
            "sInfoFiltered": "(由 _MAX_ 项结果过滤)",
            "sInfoPostFix": "",
            "sSearch": "搜索:",
            "sUrl": "",
            "sEmptyTable": "未添加数据",
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
            },

        },
        "columnDefs": [{
            "searchable": false,
            "orderable": true,
            "targets": 0
        }],
        "aLengthMenu": [5, 10, 25],//也可以直接用一维数组设置数量
        "order": [[0, 'asc']]
    });
    $('#userTable tbody').on('click', 'tr', function () {
        if ($(this).hasClass('selected')) {
            $(this).removeClass('selected');
        }
        else {
            table.$('tr.selected').removeClass('selected');
            $(this).addClass('selected');
        }
    });
});

/* ------服务器Table  Begin  -------*/
var serverData = [];
$(function () {
    var table = $('#serverTable').DataTable({
        "ajax": {
            url: "/lora/serverTable",
            dataSrc: "Result"
        },
        columns: [
            {data: "Name"},
            {data: "Server"}
        ],
        "pagingType": "full_numbers",
        "bSort": true,
        "language": {
            "sProcessing": "处理中...",
            "sLengthMenu": "显示 _MENU_ 项结果",
            "sZeroRecords": "没有匹配结果",
            "sInfo": "显示第 _START_ 至 _END_ 项结果，共 _TOTAL_ 项",
            "sInfoEmpty": " 显示第 0 至 0 项结果，共 0 项",
            "sInfoFiltered": "(由 _MAX_ 项结果过滤)",
            "sInfoPostFix": "",
            "sSearch": "搜索:",
            "sUrl": "",
            "sEmptyTable": "未添加数据",
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
            },

        },
        "columnDefs": [{
            "searchable": false,
            "orderable": true,
            "targets": 0
        }],
        "aLengthMenu": [5, 10, 25],//也可以直接用一维数组设置数量
        "order": [[0, 'asc']]
    });
    $('#serverTable tbody').on('click', 'tr', function () {
        if ($(this).hasClass('selected')) {
            $(this).removeClass('selected');
        }
        else {
            table.$('tr.selected').removeClass('selected');
            $(this).addClass('selected');
        }
    });
});

/* ------应用Table  Begin  -------*/
var appData = [];
$(function () {
    var table = $('#appTable').DataTable({
        "ajax": {
            url: "/lora/applicationTable",
            dataSrc: "Result"
        },
        columns: [
            {data: "ID"},
            {data: "OrganizationID"},
            {data: "Name"},
            {data: "ServiceProfileName"},
            {data: "Description"}
        ],
        "pagingType": "full_numbers",
        "bSort": true,
        "language": {
            "sProcessing": "处理中...",
            "sLengthMenu": "显示 _MENU_ 项结果",
            "sZeroRecords": "没有匹配结果",
            "sInfo": "显示第 _START_ 至 _END_ 项结果，共 _TOTAL_ 项",
            "sInfoEmpty": " 显示第 0 至 0 项结果，共 0 项",
            "sInfoFiltered": "(由 _MAX_ 项结果过滤)",
            "sInfoPostFix": "",
            "sSearch": "搜索:",
            "sUrl": "",
            "sEmptyTable": "未添加数据",
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
            },

        },
        "columnDefs": [{
            "searchable": false,
            "orderable": true,
            //"targets": 0
        }],
        "aLengthMenu": [10, 25, 50],//也可以直接用一维数组设置数量
        "order": [[0, 'asc']]
    });
    $('#appTable tbody').on('click', 'tr', function () {
        if ($(this).hasClass('selected')) {
            $(this).removeClass('selected');
        }
        else {
            table.$('tr.selected').removeClass('selected');
            $(this).addClass('selected');
        }
    });
});


/* ------网关Table  Begin  -------*/
var gatewayData = [];
$(function () {
    var table = $('#gatewayTable').DataTable({
        "ajax": "/lora/gatewayTable",
        "pagingType": "full_numbers",
        "bSort": true,
        "language": {
            "sProcessing": "处理中...",
            "sLengthMenu": "显示 _MENU_ 项结果",
            "sZeroRecords": "没有匹配结果",
            "sInfo": "显示第 _START_ 至 _END_ 项结果，共 _TOTAL_ 项",
            "sInfoEmpty": " 显示第 0 至 0 项结果，共 0 项",
            "sInfoFiltered": "(由 _MAX_ 项结果过滤)",
            "sInfoPostFix": "",
            "sSearch": "搜索:",
            "sUrl": "",
            "sEmptyTable": "未添加数据",
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
            },

        },
        "columnDefs": [{
            orderable: false,
            "targets":2,
            "render": function(data, type, row, meta) {
                data = data.split(",");
                var lihtml = "";
                for (var i = 0; i < data.length; i ++) {
                    lihtml += "<li style='left:"+ 10 * (i+1)+"px;height:"+data[i]*3+"px'></li>";   //每个柱柱位置信息
                }
                return  '<div id="vert"><ul>'+lihtml+'</ul></div>';

            }
        },{
            orderable: false,
            defaultContent: "<button id='viewrow' class='btn btn-primary btn-table' type='button' title='查看'><i class='fa fa-eye'></i></button>",
            targets:   -1
        }],
        "aLengthMenu": [10, 25, 50],//也可以直接用一维数组设置数量
    });
    $('#gatewayTable tbody').on('click', 'tr', function () {
        if ($(this).hasClass('selected')) {
            $(this).removeClass('selected');
        }
        else {
            table.$('tr.selected').removeClass('selected');
            $(this).addClass('selected');
        }
    });

    /* 每一行查看按钮点击*/
    $("#gatewayTable tbody").on("click","#viewrow", function(){


            var data = table.row($(this).parents('tr')).data();
            $("#viewGatewayInfo").modal()
    })
});


/* ------网关日志Table  Begin  -------*/
var gatewayLogData = [];
$(function () {
    var table = $('#gatewayLogTable').DataTable({
        data: gatewayLogData,
        "pagingType": "full_numbers",
        "bSort": true,
        "language": {
            "sProcessing": "处理中...",
            "sLengthMenu": "显示 _MENU_ 项结果",
            "sZeroRecords": "没有匹配结果",
            "sInfo": "显示第 _START_ 至 _END_ 项结果，共 _TOTAL_ 项",
            "sInfoEmpty": " 显示第 0 至 0 项结果，共 0 项",
            "sInfoFiltered": "(由 _MAX_ 项结果过滤)",
            "sInfoPostFix": "",
            "sSearch": "搜索:",
            "sUrl": "",
            "sEmptyTable": "未添加数据",
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
            },

        },
        "aLengthMenu": [5, 10, 25],//也可以直接用一维数组设置数量
        "order": [[0, 'asc']]
    });
    $('#gatewayLogTable tbody').on('click', 'tr', function () {
        if ($(this).hasClass('selected')) {
            $(this).removeClass('selected');
        }
        else {
            table.$('tr.selected').removeClass('selected');
            $(this).addClass('selected');
        }
    });
});
