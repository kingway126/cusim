var pwd = new Vue({
    el: "#pwd",
    data: {
        pwd: "",
        repwd: "",
    },
    methods: {
        //检验token
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
        },
        //更新密码
        updataPwd: function () {
            if (this.pwd == "" || this.pwd == null || this.pwd == undefined) {
                alert("请输出新密码")
                return false
            }

            if (this.pwd != this.repwd) {
                alert("两次输入的密码不一致")
                return false
            }

            var id = window.localStorage.getItem("id")
            var token = window.localStorage.getItem("token")
            axios({
                method: "post",
                url: "/api/user/pwd",
                data: {
                    "id": parseInt(id),
                    "token": token,
                    "pwd": this.pwd
                },
                headers: {
                    "Content-Type": "application/json"
                }
            }).then(res => {
                if (res.data.code == -1) {
                    alert(res.data.msg)
                    if (res.data.path != null) {
                        window.location.href = res.data.path
                    }
                } else if (res.data.code == 0) {
                    alert(res.data.msg)
                    if (res.data.path != null) {
                        window.location.href = res.data.path
                    }
                }
            })
        }
    },
    created() {
        this.checkToken()
    }
})