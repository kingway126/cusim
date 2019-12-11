var chat = new Vue({
    el: "#chat",
    delimiters: ['${', '}'],
    data: {
        webSocket:{},
        iplist: [],
        showroom: {
            "ip": "",
            "app_name": "",
            "name": {"String": ""},
            "email": {"String": ""},

        },
        chatlog: [],
        msg: "",
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
        //初始化用户列表
        initChatList: function() {
            var id = window.localStorage.getItem("id")
            var token = window.localStorage.getItem("token")

            axios({
                method: "post",
                url: "/api/ipuser/list",
                data: {
                    "id": parseInt(id),
                    "token": token
                },
                headers: {
                    "Content-Type": "application/json"
                }
            }).then(res => {
                if (res.data.code == -1) {
                    alert(res.data.msg)
                    if (res.data.url != null) {
                        window.location.href = res.data.url
                    }
                } else if (res.data.code == 0) {
                    this.iplist = res.data.rows
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
        //发送消息
        sendMsg: function() {
            if (this.msg == "" || this.msg == null || this.msg == undefined) {
                return false
            }

            var id = window.localStorage.getItem("id")
            //生成結構體
            var struct = {
                "group_id": parseInt(id),
                "ip_id": this.showroom.id,
                "src_type": "user",
                "cmd": "msg",
                "content": this.msg,
                "date": Date.parse(new Date()),
            }
            //將消息加載到chatlog裡面
            this.chatlog.push(struct)
            //並清除輸入框的消息
            this.msg = ""
            //發送消息
            this.webSocket.send(JSON.stringify(struct))
        },
        //聊天
        chat: function(id) {
            this.showroom = this.iplist[id]
            this.getChatlog(this.iplist[id].id)
            document.getElementById("chatroom").style.display = "inline-block"

        },
        //获取聊天记录
        getChatlog: function(iid) {
            var id = window.localStorage.getItem("id")
            var token = window.localStorage.getItem("token")

            axios({
                method: "post",
                url: "/api/chat/list",
                data: {
                    "id": parseInt(id),
                    "token": token,
                    "iid": parseInt(iid)
                },
                headers: {
                    "Content-Type": "application/json"
                }
            }).then(res => {
                if (res.data.code == -1) {
                    alert(res.data.msg)
                    if (res.data.url != null) {
                        window.location.href = res.data.url
                    }
                } else if (res.data.code == 0) {
                    this.showroom.no_read = 0
                    this.chatlog = res.data.rows
                }
            })
        },
        //将时间戳转换成日期时间
        timestampToTime: function(timestamp) {
            var date = new Date(timestamp * 1000);//时间戳为10位需*1000，时间戳为13位的话不需乘1000
            var Y = date.getFullYear() + '-';
            var M = (date.getMonth()+1 < 10 ? '0'+(date.getMonth()+1) : date.getMonth()+1) + '-';
            var D = date.getDate() + ' ';
            var h = date.getHours() + ':';
            var m = date.getMinutes() + ':';
            var s = date.getSeconds();
            return Y+M+D+h+m+s;
        },
    },
    created() {
        this.checkToken()
        this.initChatList()
        this.initwebsocket()
    }
})