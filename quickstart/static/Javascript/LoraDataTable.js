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
            "searchable": true,
            "orderable": true,
            //"targets": 0
        },
            {
                orderable: false,
                //defaultContent: "<button id='viewApp' class='btn btn-primary btn-table' type='button' title='查看'><i class='fa fa-eye'></i></button>",
                targets:   5,
                render : function (obj){
                    return "<button id='viewApp' class='btn btn-primary btn-table' type='button' title='查看'><i class='fa fa-eye'></i><button>";
                }
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
    /* 每一行查看按钮点击*/
    $("#appTable tbody").on("click","#viewApp", function(){

        var data = table.row($(this).parents('tr')).data();

        if (data)
        {
            $.ajax({
                url: '/lora/applicationTable/deviceOfAppID',//请求后台加载数据的方法
                type:'post',
                data: "id=" + data[0] ,
                success: function (data) {
                    alert("发送成功");
                },
                error:function(e){
                    alert("id=" + data[0]);  //当前为测试，正式时请改为“发送失败"
                }
            })
        }
        $("#viewAppInfo").modal();
    })
});

/* ------应用下的设备Table  Begin  -------*/
$(function () {
    var table = $('#appDeviceTable').DataTable({
        "ajax": {
            url: "/lora/applicationTable/deviceOfAppID",
            dataSrc: "Result"
        },
        columns: [
            {data: "Name"},
            {data: "DeviceStatusMargin"},
            {data: "DeviceStatusBattery"}
        ],
        /*   "ajax": {                        //带参数的ajax请求
         url: "apparrays.txt",
         dataType:'json',
         data: {id:"1"}, //请求的附加参数，用json对象
         method:'POST',
         success: function(data){
         $.messager.alert('消息',data.id,'');
         },
         },*/
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
        "aLengthMenu": [10, 25, 50],//也可以直接用一维数组设置数量
        "order": [[0, 'asc']]
    });
    $('#appDeviceTable tbody').on('click', 'tr', function () {
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
        "ajax": {
            url: "/lora/gatewayTable",
            dataSrc: "Result"
        },
        columns: [
            {data: "Name"},
            {data: "MAC"},
            {data: "Activity"}
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
            orderable: false,
            "targets":2,
            "render": function(data, type, row, meta) {
                data = data.split(",");
                var lihtml = "";
                for (var i = 0; i < data.length; i ++) {
                    lihtml += "<li style='left:"+ 10 * (i+1)+"px;height:"+data[i]*0.03+"px'></li>";   //每个柱柱位置信息
                }
                return  '<div id="vert"><ul>'+lihtml+'</ul></div>';

            }
        },{
            orderable: false,
            render : function (obj){
                return "<button id='viewrow' class='btn btn-primary btn-table' type='button' title='查看'><i class='fa fa-eye'></i></NOtton>";
            },
            //defaultContent: "<button id='viewrow' class='btn btn-primary btn-table' type='button' title='查看'><i class='fa fa-eye'></i></button>",
            targets:   3
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

        if (data)
        {
            $.ajax({
                url: '/lora/gatewayTable',//请求后台加载数据的方法
                type:'post',
                data: "mac=" + data[1] ,
                success: function (data) {
                    alert("发送成功");
                },
                error:function(e){
                    alert(data[1]);  //当前为测试，正式时请改为“发送失败"
                }
            })
        }
        $("#viewGatewayInfo").modal();
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


/* ------设备Table  Begin  -------*/
var deviceData = [];
$(function () {
    var table = $('#deviceTable').DataTable({
        "ajax": {
            url: "/lora/deviceTable",
            dataSrc: "Result"
        },
        columns: [
            {data: "Name"},
            {data: "DevEUI"},
            {data: "DeviceProfileName"},
            {data: "DeviceStatusMargin"},
            {data: "DeviceStatusBattery"}
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
            orderable: false,
            //defaultContent: "<button id='viewBtn' class='btn btn-primary btn-table' type='button' title='查看'><i class='fa fa-eye'></i></button><button id='activeBtn' class='btn btn-primary btn-table margin-left-10' type='button' title='激活'><i class='fa fa-flash'></i></button><button id='closeBtn' class='btn btn-primary btn-table margin-left-10' type='button' title='关闭'><i class='fa fa-close'></i></button><button id='resetBtn' class='btn btn-primary btn-table margin-left-10' type='button' title='重置'><i class='fa fa-refresh'></i></button>",
            targets:   5,
            render : function (obj){
                return "<button id='viewBtn' class='btn btn-primary btn-table' type='button' title='查看'><i class='fa fa-eye'></i><button><button id='activeBtn' class='btn btn-primary btn-table margin-left-10' type='button' title='激活'><i class='fa fa-flash'></i><button><button id='closeBtn' class='btn btn-primary btn-table margin-left-10' type='button' title='关闭'><i class='fa fa-close'></i><button><button id='resetBtn' class='btn btn-primary btn-table margin-left-10' type='button' title='重置'><i class='fa fa-refresh'></i><button>";
            }
        }],

        "aLengthMenu": [10, 25, 50],//也可以直接用一维数组设置数量
    });
    $('#deviceTable tbody').on('click', 'tr', function () {
        if ($(this).hasClass('selected')) {
            $(this).removeClass('selected');
        }
        else {
            table.$('tr.selected').removeClass('selected');
            $(this).addClass('selected');
        }
    });

    /* 每一行查看按钮点击*/
    $("#deviceTable tbody").on("click","#viewBtn", function(){

        var data = table.row($(this).parents('tr')).data();

        if (data)
        {
            $.ajax({
                url: '/postarrays.txt',//请求后台加载数据的方法
                type:'post',
                data: "deviceName=" + data[0] ,
                success: function (data) {
                    alert("发送成功");
                },
                error:function(e){
                    alert(data[1]);  //当前为测试，正式时请改为“发送失败"
                }
            })
        }
        $("#viewDeviceInfo").modal()
    })
});


