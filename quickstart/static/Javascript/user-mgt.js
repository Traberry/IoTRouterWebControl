/**
 * Created by Administrator on 2018/7/26.
 */




var data = [];
var titles = ['用户信息'];
$(function () {
    var table = $('#userTable').DataTable({
        "ajax": {
            url: "/user/userList",
            dataSrc: "Result"
        },
        columns: [
            {data: "UserName"},
            {data: "Email"},
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
            "sEmptyTable": "未添加任何用户信息",
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
            "searchable": true,
            "orderable": true,
            "targets": 0
        }],
        "order": [[0, 'asc']]
    });


    $('#userTable tbody').on('click', 'tr', function () {
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
        $("#addUserModal").find('input').val('')
    })

    $("#addUserInfo").on('click', function() {
        console.log('addUserInfo');
        /*if (data.length) {*/
        if ( $("#UserName").val() == '' || $("#UserEmail").val() == '' ||  $("#UserPassword").val() == '') {
            alert('请输入内容')
            /*	}*/
        } else {
            var userName = $("#UserName").val();
            var userEmail = $("#UserEmail").val();
            var isActive = $("#isActive").val();
            var isOrgAdmin = $("#isOrgAdmin").val();


            table.row.add([userName, userEmail,isActive,isOrgAdmin] ).draw();  //新增用户信息 （这里没有添加密码 和 系统管理员）
            $("#addUserModal").find('input').val('')
            $.ajax({
                url: '/user/addUser',
                type:'post',
                data: {"userName": userName, "userEmail": userEmail},
                success: function (Res) {
                    alert(Res);
                },
                error:function(e){
                    alert("发送失败");  //当前为测试，正式时请改为“发送失败"
                }
            })
        }
    })
    /*	$("#addUserInfo").click();*/
    $("#btn_add").click(function(){
        console.log('add');
        $("#addUser").modal()
    });

    $('#btn_edit').click(function () {
        console.log('edit');
        if (table.rows('.selected').data().length) {
            $("#editUserInfo").modal()
            var rowData = table.rows('.selected').data()[0];

            var eachValue = new Array();
            eachValue[0] = rowData[0];
            $("#editUserName").val(rowData[0]);
            $("#editUserEmail").val(rowData[1]);

            $("#isActiveEdit").val(rowData[2]);
            $("#isOrgAdminEdit").val(rowData[3]);
        } else {
            alert('请选择一条用户信息');
        }
    });
    $("#saveEdit").click(function() {

        var editUserName = $("#editUserName").val();
        var editUserEmail =  $("#editUserEmail").val();
        var isActiveEdit = $("#isActiveEdit").val();
        var isOrgAdminEdit = $("#isOrgAdminEdit").val();
        var newRowData = [].concat(editUserName, editUserEmail,isActiveEdit, isOrgAdminEdit);
        console.log(newRowData);

        var tds = Array.prototype.slice.call($('.selected td'))
        for (var i = 0; i < newRowData.length; i++) {
            if (newRowData[i] !== '') {
                tds[i].innerHTML = newRowData[i];
            } else {
                alert(titles[i] + '内容不能为空')
            }
        }
    })
    $("#cancelEdit").click(function() {
        console.log('cancelAdd');
        $("#editUserModal").find('input').val('')
    })
    $('#btn_delete').click(function () {
        if (table.rows('.selected').data().length) {
            $("#deleteUser").modal()
        } else {
            alert('请选择一条用户信息');
        }
    });
    $('#delete').click(function () {
        table.row('.selected').remove().draw(false);
    });


})