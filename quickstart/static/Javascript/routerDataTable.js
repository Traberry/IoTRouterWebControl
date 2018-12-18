/**
 * Created by Administrator on 2018/7/24.
 */


var data = [];
var titles = ['目标网络地址', '子网掩码', '网关', '接口'];
$(function () {
    var table = $('#staticRouter').DataTable({
        "ajax": {
            type: "GET",
            url: "/connection/staticRouteTable",
            dataSrc: "Result"
        },
        columns: [
            {data: "dstAddr"},
            {data: "dstMask"},
            {data: "gateway"},
            {data: "port"}
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
            "sEmptyTable": "未添加任何静态路由",
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
        "columnDefs": [{
            "searchable": false,
            "orderable": true,
            "targets": 0
        }],
        "order": [[1, 'asc']]
    });
    table.on('order.dt search.dt', function() {
        table.column(0, {
            search: 'applied',
            order: 'applied'
        }).nodes().each(function(cell, i) {
            cell.innerHTML = i + 1;
        });
    }).draw();
    $('#staticRouter tbody').on('click', 'tr', function () {
        if ( $(this).hasClass('selected') ) {
            $(this).removeClass('selected');
        }
        else {
            table.$('tr.selected').removeClass('selected');
            $(this).addClass('selected');
        }
    });
    $("#cancelAdd").on('click', function() {
        console.log('cancelAdd');
        $("#addRouterModal").find('input').val('')
    })
    $("#addRoutersInfo").on('click', function() {
        console.log('addRoutersInfo');
        /*if (data.length) {*/
        if ($("#addRouterModal").find('input').val() == '') {

            alert('请输入内容')
            /*	}*/
        } else {
            var targetIP = $("#targetIP_1").val() + "." + $("#targetIP_2").val() + "." + $("#targetIP_3").val() + "." + $("#targetIP_4").val()
            var mask = $("#mask_1").val() + "." + $("#mask_2").val() + "." + $("#mask_3").val() + "." + $("#mask_4").val()
            var gateway = $("#gateway_1").val() + "." + $("#gateway_2").val() + "." + $("#gateway_3").val() + "." + $("#gateway_4").val()
            var interfaceValue = $("#interface").val()
            var addRouterInfos = [].concat(targetIP, mask, gateway, interfaceValue);
            for (var i = 0; i < addRouterInfos.length; i++) {
                if (addRouterInfos[i] === '') {
                    alert(titles[i] + '内容不能为空')
                }
            }
            table.row.add([targetIP, mask, gateway, interfaceValue]).draw();
            $("#addRouterModal").find('input').val('')
        }
    })
    /*	$("#addBooksInfo").click();*/
    $("#btn_add").click(function(){
        console.log('add');
        $("#addRouter").modal()
    });
    $('#btn_edit').click(function () {
        console.log('edit');
        if (table.rows('.selected').data().length) {
            $("#editRouterInfo").modal()
            var rowData = table.rows('.selected').data()[0];

            var eachValue = new Array();
            eachValue[0] = rowData[1];
            var eachIP = rowData[2].split('.');
            for (var i = 0;i < 4; i ++) {
                eachValue[1 + i] = eachIP[i];
            }
            var eachMask = rowData[3].split('.');
            for (var i = 0; i < 4; i ++) {
                eachValue[5 + i] = eachMask[i];
            }
            var eachGateway = rowData[4].split('.');
            for (var i = 0; i < 4; i ++) {
                eachValue[9 + i] = eachGateway[i];
            }

            var inputs = $("#editRouterModal").find('input')

            for (var i = 0; i < inputs.length; i++) {
                $(inputs[i]).val(eachValue[i])
            }
        } else {
            alert('请选择一条静态路由信息');
        }
    });
    $("#saveEdit").click(function() {
        var editRouterName = $("#editRouterName").val()
        var editTargetIP = $("#editTargetIP_1").val() + "." + $("#editTargetIP_2").val() + "." + $("#editTargetIP_3").val() + "." + $("#editTargetIP_4").val()
        var editMask = $("#editMask_1").val() + "." + $("#editMask_2").val() + "." + $("#editMask_3").val() + "." + $("#editMask_4").val()
        var editGateway = $("#editGateway_1").val() + "." + $("#editGateway_2").val() + "." + $("#editGateway_3").val() + "." + $("#editGateway_4").val()
        var editInterface = $("#editInterface").val()
        var newRowData = [].concat(editRouterName, editTargetIP, editMask, editGateway, editInterface);
        var tds = Array.prototype.slice.call($('.selected td'))
        for (var i = 0; i < newRowData.length; i++) {
            if (newRowData[i] !== '') {
                tds[i + 1].innerHTML = newRowData[i];
            } else {
                alert(titles[i] + '内容不能为空')
            }
        }
    })
    $("#cancelEdit").click(function() {
        console.log('cancelAdd');
        $("#editBookModal").find('input').val('')
    })
    $('#btn_delete').click(function () {
        if (table.rows('.selected').data().length) {
            $("#deleteRouter").modal()
        } else {
            alert('请选择一条静态路由信息');
        }
    });
    $('#delete').click(function () {
        table.row('.selected').remove().draw(false);
    });
});

/* ------ 想办法知道 当前的登录的用户信息，知道后，设置$("#btn_add").attr("disabled",true);  ------------*/
$.ajax({
    type:'get',
    url:'/user/judgeUser',
    dataType:'json',
    success:function(data){
        if (data == false){
            $("#btn_add").attr("disabled",true);
            $("#btn_edit").attr("disabled",true);
            $("#btn_delete").attr("disabled",true);
        }
    }
});