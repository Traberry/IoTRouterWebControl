/**
 * Created by Administrator on 2018/7/25.
 */


function submitFun() {

    /* 检测前，去除抖动及提示语 */
    $("#loginForm").removeClass('animated shake');
    $("#failTips").text("");

    /* 未输入用户名，显示提示语，并抖动 */
    if (loginForm.username.value == "")
    {
        $("#failTips").text("请填写用户名！");
        loginForm.username.focus();
        $("#loginForm").addClass('animated shake');
        return false;
    }

    /* 未输入密码，显示提示语，并抖动 */
    if (loginForm.password.value == "")
    {
        $("#failTips").text("请填写密码！");
        loginForm.password.focus();
        $("#loginForm").addClass('animated shake');
        return false;
    }

    return true;
}