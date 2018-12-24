
$(function () {
    $('#userLink').click( function () {
        $.ajax({
            type: 'post',
            url: '/jump',
            success: function (data) {
                console.log("go to another website");
            }
        });

    } );
});