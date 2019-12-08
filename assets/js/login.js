var login = new Vue({
    el: "#login",
    data:{
        user: "",
        pass: "",
    },
    methods: {
        //登陆api
        login: function() {
            if (this.user == "" || this.user == undefined) {
                console.log(this.user)
                alert("请输入登陆账号")
                return false
            } else if (this.pass == "" || this.pass == undefined) {
                alert("请输入登陆密码")
                return false
            }

            axios({
                method: "post",
                url: "/api/login",
                data: {
                    "user": this.user,
                    "pass": this.pass
                },
                headers: {
                    "Content-Type": "application/json"
                }
            }).then(res => {
                if (res.data.code == 0) {
                    window.localStorage.setItem("id", res.data.data.id)
                    window.localStorage.setItem("token", res.data.data.token)
                    window.localStorage.setItem("uuid", res.data.data.uuid)
                    alert("登陆成功")
                    window.location.href = "/index"
                } else if (res.data.code == -1) {
                    alert(res.data.msg)
                }
            })
        },
        checkToken: function() {
            var id = window.localStorage.getItem("id")
            var token = window.localStorage.getItem("token")
            if (id == "" || id == null || id == undefined) {
                return false
            } else if (token == "" || token == null || token == undefined) {
                return false
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
                if (res.data.code == 0) {
                    alert("你已经登陆，即将自动跳转到主页")
                    window.location.href = "/index"
                }
            })
        }
    },
    created() {
        this.checkToken()
    }
})