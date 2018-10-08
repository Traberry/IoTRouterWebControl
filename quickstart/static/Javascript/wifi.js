/**
 * Created by Administrator on 2018/7/26.
 */


$(function () {

    /* 管理页面  UPnP 启用 checkbox===================================== */

    $("#macfliter_cb").change(function () {

        if ($(this).attr("checked")) {

            $("#forbiddenFilter").attr("disabled", false);
            $("#allowFilter").attr("disabled", false);
            $("#forbiddenFilter").parent(".checked").css("opacity", "1");
            $("#allowFilter").parent(".checked").css("opacity", "1");

        }
        if (!$(this).attr("checked")) {
            $("#forbiddenFilter").attr("disabled", true);
            $("#allowFilter").attr("disabled", true);
            $("#forbiddenFilter").parent(".checked").css("opacity", "0.2");
            $("#allowFilter").parent(".checked").css("opacity", "0.2");
        }
    });
});



var data = [];
var titles = ['MAC地址'];
$(function () {
    var table = $('#MACTable').DataTable({
        data: data,
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


    $('#MACTable tbody').on('click', 'tr', function () {
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
        $("#addMACModal").find('input').val('')
    })

    $("#addMACInfo").on('click', function() {
        console.log('addMACInfo');
        /*if (data.length) {*/
        if ($("#addMACModal").find('input').val() == '') {

            alert('请输入内容')
            /*	}*/
        } else {
            var MACName = $("#MACName").val()
            var addMACInfos = [].concat(MACName);
            for (var i = 0; i < addMACInfos.length; i++) {
                if (addMACInfos[i] === '') {
                    alert(titles[i] + '内容不能为空')
                }
            }
            table.row.add(['', MACName]).draw();
            $("#addMACModal").find('input').val('')
        }
    })
    /*	$("#addMACInfo").click();*/
    $("#btn_add").click(function(){
        console.log('add');
        $("#addMAC").modal()
    });

    $('#btn_edit').click(function () {
        console.log('edit');
        if (table.rows('.selected').data().length) {
            $("#editMACInfo").modal()
            var rowData = table.rows('.selected').data()[0];

            var eachValue = new Array();
            eachValue[0] = rowData[1];

            var inputs = $("#editMACModal").find('input')

            for (var i = 0; i < inputs.length; i++) {
                $(inputs[i]).val(eachValue[i])
            }
        } else {
            alert('请选择一条MAC地址信息');
        }
    });
    $("#saveEdit").click(function() {
        var editMACName = $("#editMACName").val()
        var newRowData = [].concat(editMACName);
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
        $("#editMACModal").find('input').val('')
    })
    $('#btn_delete').click(function () {
        if (table.rows('.selected').data().length) {
            $("#deleteMAC").modal()
        } else {
            alert('请选择一条MAC地址信息');
        }
    });
    $('#delete').click(function () {
        table.row('.selected').remove().draw(false);
    });


    /* Wifi 保护设置页面上 按钮 点击*/
    $("#setBtn").click(function () {
        $("#btnSetModal").modal();
    });

    /* Wifi 保护设置页面上 PIN注册 点击*/
    $("#registerBtn").click(function () {
        $("#registerModal").modal();
    });
})