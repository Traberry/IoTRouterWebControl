/**
 * Created by Administrator on 2018/8/30.
 * Modified by Administrator on 2018/10/17.
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
                render : function (obj){
                    return "<button id='viewApp' class='btn btn-primary btn-table' type='button' title='查看'><i class='fa fa-eye'></i></button>";
                },
                targets:  5
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

        //var data = table.row($(this).parents('tr')).data();

        var Row = $(this).parents('tr')[0];//通过获取该td所在的tr，即td的父级元素，取出第一列序号元素
        var data = $("#appTable").dataTable().fnGetData(Row);
        alert (data.ID);

        if (data)
        {
            var test = data.ID;   // 如果需要变量，请修改这里
            appDevicetable1.fnReloadAjax("/lora/applicationTable/deviceOfAppID/"+test);
            $("#viewAppInfo").modal();
        }
    })

    var appDeviceData=[];
    var appDevicetable1 = $('#appDeviceTable').dataTable({
        //   data: appDeviceData,
        "ajax": {
            url: "/lora/applicationTable/deviceOfAppID/?:id",
            dataSrc: "Result"
        },
        "columns": [
            {"data": "Name"},
            {"data": "DeviceStatusBattery"},
            {"data": "DeviceStatusMargin"}
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

/* ------应用下的设备Table  Begin  -------*/
$(function () {


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
                    lihtml += "<li style='left:"+ 10 * (i+1)+"px;height:"+data[i]*0.01+"px'></li>";   //每个柱柱位置信息
                }
                return  '<div id="vert"><ul>'+lihtml+'</ul></div>';

            }
        },{
            orderable: false,
           // defaultContent: "<button id='viewrow' class='btn btn-primary btn-table' type='button' title='查看'><i class='fa fa-eye'></i></button>",
            render : function (obj){
                return "<button id='viewrow' class='btn btn-primary btn-table' type='button' title='查看'><i class='fa fa-eye'></i></button>";
            },
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

        //var data = table.row($(this).parents('tr')).data();
        var Row = $(this).parents('tr')[0];//通过获取该td所在的tr，即td的父级元素，取出第一列序号元素
        var data = $("#gatewayTable").dataTable().fnGetData(Row);
        alert (data.MAC);

        if (data)
        {
            $.ajax({
                url: '/lora/gatewayTable',//请求后台加载数据的方法
                type:'post',
                //data: "mac=" + data[1] ,
                data: "mac=" + data.MAC,
                success: function (jsonResult) {
                    var obj = JSON.parse(jsonResult);
                    var str = obj.Longtitude + "," + obj.Latitude;
                    document.getElementById("gatewayMacAddress").innerHTML = obj.MAC;
                    document.getElementById("gatewayAltitude").innerHTML = obj.Altitude;
                    document.getElementById("gatewayCoordinates").innerHTML = str;
                    document.getElementById("gatewayLastSeenAt").innerHTML = obj.LastSeenAt;
                },
                error:function(e){
                    alert("发送失败");  //当前为测试，正式时请改为“发送失败"
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
var deviceTableArray = [];
//$(function () {
function  initDeviceTable (tableID, val) {
    console.error('yyy', tableID);
    var appID = tableID-1;

    deviceTableArray[appID] = $('#deviceTable'+ tableID).DataTable({
        "ajax": {
            type: "GET",
            url: "/lora/deviceTableDevices/" + val,
            dataSrc: "Result"
        },
        "columns": [
            {"data": "Name"},
            {"data": "DevEUI"},
            {"data": "DeviceProfileName"},
            {"data": "DeviceStatusMargin"},
            {"data": "DeviceStatusBattery"}
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
            render : function (obj){
                return "<button id='viewBtn"+ tableID +"' class='btn btn-primary btn-table' type='button' title='查看'><i class='fa fa-eye'></i></button><button id='activeBtn"+ tableID + "' class='btn btn-primary btn-table margin-left-10' type='button' title='激活'><i class='fa fa-flash'></i></button><button id='closeBtn' class='btn btn-primary btn-table margin-left-10' type='button' title='关闭'><i class='fa fa-close'></i></button><button id='resetBtn'+tableID class='btn btn-primary btn-table margin-left-10' type='button' title='重置'><i class='fa fa-refresh'></i></button>";
            },
            targets: 5
        }],

        "aLengthMenu": [10, 25, 50],//也可以直接用一维数组设置数量
    });

    $('#deviceTable' + tableID +' tbody').on('click', 'tr', function () {
        if ($(this).hasClass('selected')) {
            $(this).removeClass('selected');
        }
        else {
            deviceTableArray[appID].$('tr.selected').removeClass('selected');
            $(this).addClass('selected');
        }
    });

    /* 每一行查看按钮点击*/
    $("#deviceTable" + tableID +" tbody").on("click", "#viewBtn"+tableID, function () {

        var Row = $(this).parents('tr')[0];//通过获取该td所在的tr，即td的父级元素，取出第一列序号元素
        var data =  $("#deviceTable" + tableID).dataTable().fnGetData(Row);
        alert (data.Name);

        if (data)
        {
            $.ajax({
                url: '/lora/deviceActivation',//请求后台加载数据的方法
                type:'post',
                data: "devEUI=" + data.DevEUI,
                success: function (jsonResult) {
                    var obj = JSON.parse(jsonResult);
                    document.getElementById("deviceAddress").innerHTML = obj.DevAddr;
                    document.getElementById("deviceNetworkSKey").innerHTML = obj.AppSKey;
                    document.getElementById("deviceAppSKey").innerHTML = obj.NwkSkey;
                    document.getElementById("deviceUpCount").innerHTML = obj.FCntUp;
                    document.getElementById("deviceDownCount").innerHTML = obj.FCntDown;
                },
                error:function(e){
                    alert("发送失败");  //当前为测试，正式时请改为“发送失败"
                }
            })
        }

        $("#viewDeviceInfo").modal()
    })
}



function setDevice(){

    var appArrarys = [];
    var totalNum;
    var tableName = new Array();

    $.ajax({
        type: "GET",
        url: "/lora/deviceTableApps",
        dataType: "json",
        success: function (data) {
            console.error('xxx', data);
            console.error('no',data[0].number);
            totalNum = data.length;
            for (var i = 0; i < totalNum; i++){
                tableName.push( data[i].name)
            }
            setPagination (totalNum, tableName);

            for (var i = 0; i < totalNum; i++) {
                console.info("item: " + i);
                var val = data[i].number;
                console.error('vvvv', i);
                initDeviceTable(i, val);
            }

        }
    });

};



function setPagination(totalNum, tableName) {
    (function( name) {
        console.info(totalNum + "  " + name);
        var container = $('#pagination-' + name);
        var sources = function () {
            var result = [];

            for (var i = 1; i <= totalNum; i++) {
                result.push(i);
            }

            return result;
        }();

        var options = {
            dataSource: sources,
            callback: function (response, pagination) {
                window.console && console.log(response, pagination);

                var dataHtml = '<ul>';

                $.each(response, function (index, item) {

                    dataHtml += '<li>' + '<div class="config_title app_name"><span class="label content_title config_title_txt">应用'+ tableName[item-1] + '</span></div><table id="deviceTable'+ item + '" class="table table-striped table-bordered row-border hover order-column" cellspacing="0" width="100%" style="text-align: center;">'
                        + '<thead><tr><th>设备名称</th><th>设备EUI</th><th>设备配置文件</th><th>链路边际</th><th>电量</th><th>操作</th></tr></thead>'
                        + '<tbody></tbody></table></li>'
                });

                dataHtml += '</ul>';
                container.prev().html(dataHtml);
            }
        };

        //$.pagination(container, options);

        container.addHook('beforeInit', function () {
            window.console && console.log('beforeInit...');
        });
        container.pagination(options);

        container.addHook('beforePageOnClick', function () {
            window.console && console.log('beforePageOnClick...');
            //return false
        });
    })('table');
}