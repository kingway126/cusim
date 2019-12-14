var index = new Vue({
    delimiters: ['${', '}'],
    el: "#index",
    data: {
        uuid: "",
        app: 0,
        user: 0,
        noread: 0,
        read: 0,
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
        },
        //获取首页的信息
        getIndexNum:  function () {
            var id = window.localStorage.getItem("id")
            var token = window.localStorage.getItem("token")
            axios({
                method: "post",
                url: "/api/index",
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
                    console.log(res)
                    this.app = res.data.data.app
                    this.user = res.data.data.user
                    this.noread = res.data.data.noread
                    this.read = res.data.data.read
                }
            })
        }
    },
    created() {
        this.checkToken()
        this.getIndexNum()
    },
})