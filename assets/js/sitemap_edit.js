var edit = new Vue({
    el: "#edit",
    data: {
        aid: 0,
        api: "/api/app/add",
        name: "",
        url: "",
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
        //获取url中的参数
        GetQueryParam: function(name) {
            var reg = new RegExp("(^|&)"+ name +"=([^&]*)(&|$)");
            var r = window.location.search.substr(1).match(reg);
            if(r!=null)return  unescape(r[2]); return null;
        },
        //编辑站点信息
        editapp: function() {
            if (this.name == "" || this.name == null || this.name == undefined) {
                alert("请输入站点名称")
                return false
            }
            var id = window.localStorage.getItem("id")
            var token = window.localStorage.getItem("token")
            axios({
                method: "post",
                url: this.api,
                data: {
                    "id": parseInt(id),
                    "token": token,
                    "aid": this.aid,
                    "name": this.name,
                    "url": this.url
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


        },
        //获取站点的信息
        getapp: function() {
            var id = window.localStorage.getItem("id")
            var token = window.localStorage.getItem("token")
            axios({
                method: "post",
                url: "/api/app",
                data: {
                    "id": parseInt(id),
                    "token": token,
                    "aid": parseInt(this.aid),
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
                    this.name = res.data.data.name
                    this.url = res.data.data.url
                }
            })
        }
    },
    created() {
        this.checkToken()
        //获取请求的参数
        var param = this.GetQueryParam("aid")
        if (param != null) {
            this.api = "/api/app/update"
            this.aid = param
            this.getapp()
        }
    }
})