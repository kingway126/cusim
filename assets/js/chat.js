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
            var url="wss://"+location.host+"/api/chat?id="+id+"&token=" + token + "&uuid=" + uuid + "&role=admin"
            // var url2="ws://"+location.host+"/api/chat?id="+ id + "&uuid=e1a0390f-911b-401c-8f20-2efcb8a62aaa&role=ip"
            this.webSocket = new WebSocket(url)
            //消息处理
            
            this.webSocket.onmessage = function(evt){
                if(evt.data.indexOf("}")>-1){
                    this.do(JSON.parse(evt.data))
                }else{
                    console.log("recv<=="+evt.data)
                }
            }.bind(this)
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
            var time1 = (new Date()).valueOf().toString()
            var time2 = parseInt(time1.substring(0, time1.length - 3))
            //生成結構體
            var struct = {
                "group_id": parseInt(id),
                "ip_id": this.showroom.id,
                "src_type": "user",
                "cmd": "msg",
                "content": this.msg,
                "create_at": time2,
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
        //处理消息
        do: function(data) {
            //判断消息类型
            if (data.cmd == "notice") {
                //如果是notic的话，给目标的最新一条消息设置未发送的状态
                var len = this.chatlog.length
                var msg = document.getElementsByClassName("chat-notice")
                msg[msg.length -1].innerText = "xx"

            } else if (data.cmd == "msg") {
                //如果是msg的话，先解析是哪个ipuser的
                var index = this.findIpuser(data.iid)
                console.log("index:", index)

                if (index != null) {
                    //存儲iid
                    var iid = this.iplist[index]
                    console.log("存在")
                    //將發送消息的ip推到頂部
                    this.toFirst(this.iplist, index)

                    //如果存在该ipuser的话
                    console.log("showroomid:"+this.showroom.id)
                    if (this.showroom.id != undefined || this.showroom.id == iid) {
                        //并且就是当前展示的页面，就消息push到chatlog
                        //將記錄添加到聊天框
                        this.chatlog.push(data)
                        //并将消息设置成已读
                        this.setRead(this.showroom)
                    } else {
                        //不是当前展示的页面，就给iplist的ipuser增加一个未读消息的状态值
                        this.iplist[index].no_read.Int32 = this.iplist[index].no_read.Int32 + 1
                    }

                } else {
                    //如果不存在ipuser的话，获取ipuser的信息，并push到iplist表的顶部
                    this.getIpuser(data.iid)
                }
            }

        },
        //将消息设置成已读
        setRead: function(iid) {
            var id = window.localStorage.getItem("id")
            var token = window.localStorage.getItem("token")
            axios({
                method: "post",
                url: "/api/chats/read",
                data: {
                    "id": parseInt(id),
                    "token": token,
                    "iid": iid
                },
                headers: {
                    "Content-Type": "application/json"
                }
            }).then(res => {
                if (res.data.code == -1) {
                    if (res.data.path != null) {
                        window.location.href = res.data.path
                    }
                }
            })
        },
        //获取接收到的消息对应的ipuser
        findIpuser: function(iid) {
            for (var i = 0; i < this.iplist.length; i++) {
                console.log("compare: iplist.id:" + this.iplist[i].id + ", iid:", iid)
                if (this.iplist[i].id == iid) {
                    return i
                }
            }

            return null
        },
        //將某一個元素移動到第一位
        toFirst: function(fieldData,index) {

            if(index!=0){
                fieldData.unshift(fieldData.splice(index , 1)[0]);
            }
        },
        getIpuser: function(iid) {
            var id = window.localStorage.getItem("id")
            var token = window.localStorage.getItem("token")

            axios({
                method: "post",
                url: "/api/ipuser",
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
                    console.log(res)
                    this.iplist.unshift(res.data.data)
                }
            })
        }
    },
    created() {
        this.checkToken()
        this.initChatList()
        this.initwebsocket()
    }
})