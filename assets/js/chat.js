var chat = new Vue({
    el: "#chat",
    data: {
        webSocket:{},
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
        //初始化websocket
        initwebsocket: function() {
            var id = window.localStorage.getItem("id")
            var token = window.localStorage.getItem("token")
            var uuid = window.localStorage.getItem("uuid")
            var url="ws://"+location.host+"/api/chat?id="+id+"&token=" + token + "&uuid=" + uuid + "&role=admin"
            // var url2="ws://"+location.host+"/api/chat?id="+ id + "&uuid=e1a0390f-911b-401c-8f20-2efcb8a62aaa&role=ip"
            this.webSocket = new WebSocket(url)
            //消息处理
            
            this.webSocket.onmessage = function(evt){
                if(evt.data.indexOf("}")>-1){
                    console.log(evt.data)
                }else{
                    console.log("recv<=="+evt.data)
                }
            }
            //关闭回调
            this.webSocket.onclose=function (evt) {
                console.log(evt.data)
            }
            //出错回调
            this.webSocket.onerror=function (evt) {
                console.log(evt.data)
            }
        },
        sendMsg: function() {
            this.webSocket.send("你好")
        }
    },
    created() {
        this.checkToken()
        this.initwebsocket()
    }
})