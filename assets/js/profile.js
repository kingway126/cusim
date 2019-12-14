var profile = new Vue({
    el: "#profile",
    data: {
        id: 0,
        uuid: "",
        email: "",
    },
    methods: {
        //检验token
        checkToken: function () {
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
        //加载用户信息
        getUserInfo: function () {
            var id = window.localStorage.getItem("id")
            var token = window.localStorage.getItem("token")
            axios({
                method: "post",
                url: "/api/user",
                data: {
                    "id": parseInt(id),
                    "token": token,
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
                    this.uuid = res.data.data.uuid
                    this.email = res.data.data.email
                }
            })
        },
        //更新表单信息
        updateUser: function () {
            if (this.email == null || this.email == "" || this.email == undefined) {
                alert("请输入邮箱")
                return
            }

            var id = window.localStorage.getItem("id")
            var token = window.localStorage.getItem("token")
            axios({
                method: "post",
                url: "/api/user/email",
                data: {
                    "id": parseInt(id),
                    "token": token,
                    "email": this.email
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
                }
            })
        }
    },
    created() {
        this.id = window.localStorage.getItem("id")
        this.checkToken()
        this.getUserInfo()
    }

})