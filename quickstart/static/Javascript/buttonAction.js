/**
 * Created by Administrator on 2018/7/24.
 */

// Document Ready and the fun begin :)


/* 判断输入数字是否合法，0-255范围*/
function limitInput(o){
    if(!isNaN(o.value)){
        var value=o.value;
        var min=0;
        var max=255;
        if(parseInt(value)<min||parseInt(value)>max){
            alert('请输入0-255范围内数字！');
            o.value ='';
        }
    }else{
        alert("请输入0-255范围内数字");
        o.value ='';
    }
}

$(function () {


    /*管理页面  确定 按钮*/
    $('#magApplybtn').click(function () {
        var url = "aa";   //地址
        $.ajax({
            url: url,
            type: 'post',
            data: 'XXX',
            beforeSend: function (XMLHttpRequest) {
                $("#loading").css("display", "block");
            },
            success: function (data, textStatus) {
                alert('开始回调，状态文本值：返回数据：');
                $("#loading").css("display", "none");
                window.location.href = "connection.html#table_wTabs-5";   //停留在原页面
            },
            error: function (XMLHttpRequest, textStatus, errorThrown) {
                alert('数据保存失败！');
                $("#loading").css("display", "none");
                window.location.href = "connection.html#table_wTabs-5";   //停留在原页面
            }
        });
    });







});
