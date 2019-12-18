
var chat = new Vue({
    el: "#chatbox",
    data: {
        id: 0,
        uuid: "",
        showvalue: 0,
        webSocket: {},
        chatList: [
            {"group_id": 0, "ip_id": 0, "src_type": "user", "cmd": "msg", "content": "你好！我們有什麼可以幫到您的？", "create_at": ""}
        ],
        msg: "",
    },
    methods: {
        getValue: function() {
            var test = document.getElementById("chatboxjs");
            var src = test.getAttribute("src");
            var theRequest = new Object();
            if (src.indexOf("?") != -1) {
                var str = src.substr(src.indexOf('?') + 1);
                var strs = str.split("&");
                for (var i = 0; i < strs.length; i++) {
                    theRequest[strs[i].split("=")[0]] = unescape(strs[i].split("=")[1]);
                }
            }
            this.uuid = theRequest['uuid']
            this.id  = theRequest['id']
            // var template = theRequest["template"]

            // if (template == "" || template == null || template == undefined) {
            //     document.write(template1)
            // } else if (template == 1) {
            //     document.write(template1)
            // }
        },
        show: function() {
            if (this.showvalue == 0) {
                this.showvalue = 1
                $("#box").fadeIn()
                document.getElementById("notice").style.display = "none"
            } else {
                this.showvalue = 0
                $("#box").fadeOut()
                setTimeout("document.getElementById('notice').style.display = 'block'",450)  
            }
        },
        show2: function() {
            if (this.showvalue == 0) {
                this.showvalue = 1
                $("#box").fadeIn()
            } else {
                this.showvalue = 0
                $("#box").fadeOut() 
            }
        },
        initWebSocket: function() {
            var url="wss://"+ "cus.gcbodycheck.com" +"/api/chat?id=" + this.id + "&uuid=" + this.uuid + "&role=ip"
            this.webSocket = new WebSocket(url)
            //消息处理
            this.webSocket.onmessage = function(evt){
                if(evt.data.indexOf("}")>-1){
                    this.chatList.push(JSON.parse(evt.data))
                    console.log("display:", document.getElementById("box").style.display)
                    if (document.getElementById("box").style.display == "none") {
                        alert("你有一条新的消息")
                    }
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
        scrollBottom: function() {
            // document.getElementById("boxcroll").scrollTop = document.getElementById("boxcroll").scrollHeight
            console.log(document.getElementById("boxcroll").innerHTML)
            $("#boxcroll").scrollTop($("#boxcroll")[0].scrollHeight);
        },
        //初始哈websocket,開始聊天
        beginChat: function() {
            document.getElementById("beginchat").style.display = 'none'
            document.getElementById("msginput").style.display = 'block'
            this.initWebSocket()
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
        //发送消息
        sendMsg: function() {
            if (this.msg == "" || this.msg == null || this.msg == undefined) {
                return false
            }

            var struct = {
                "group_id": parseInt(this.id), 
                "ip_id": parseInt(0), 
                "src_type": "ip", 
                "cmd": "msg", 
                "content": this.msg,
                "create_at": (new Date()).valueOf()
            }
            this.chatList.push(struct)
            this.msg = ""
            this.webSocket.send(JSON.stringify(struct))
        },
    },
    created() {
        this.getValue()
    },
})