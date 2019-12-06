var index = new Vue({
    el: "#index",
    data: {
        uuid: ""
    },
    methods: {

        checkToken: function() {
            var id = window.localStorage.getItem("id")
            var token = window.localStorage.getItem("token")
            if (id == "" || id == null || id == undefined) {
                alert("非法登入")
                window.location.href = "/user/login"
            } else if (token == "" || token == null || token == undefined) {
                alert("非法登入")
                window.location.href = "/user/login"
            }

            axios({
                method: "post",
                url: "/api/user/check",
                data: {
                    "id": parseInt(id),
                    "token": token
                },
                headers: {
                    "Content-Type": "application/json"
                }
            }).then(res => {
                if (res.data.code == -1) {
                    alert("非法登入")
                    window.location.href = "/user/login"
                } else if (res.data.code == 0) {
                    this.uuid = res.data.data
                }
            })
        }
    },
    created() {
        this.checkToken()
    },
})