var template1= `<style>
[v-cloak] {
    display: none;
}

.chatbox {
    position: fixed;
    bottom: 0;
    right: 5%;
    min-width: 250px;
    width: 10%;
}

.chatbox .chatbox-notice {
    width: 78%;
    height: 50px;
    line-height: 50px;
    background-color: #fff;
    font-size: 20px;
    font-weight: 600;
    border-radius: 10px 10px 0 0;
    box-shadow: 1px 1px 3px #aa9999;
    padding: 0 30px;
    cursor: pointer;
}

.chatbox .chatbox-notice-font {
    display: inline-block;
    margin-left: 15%;
}

.chatbox .chatbox-notice-pic {
    display: inline-block;
    width: 30px;
    vertical-align: middle;

}

.chatbox .chatbox-room {
    width: 100%;
    height: 400px;
    background-color: #fff;
    border: 1px solid #ccc;
    border-radius: 10px 10px 0 0;
    min-width: 250px;
}

.chatbox .chatbox-room-head {
    width: 100%;
    height: 50px;
    line-height: 50px;
    text-indent: 10px;
    font-size: 20px;
    font-weight: 600;
    color: #fff;
    background-color: #50E3C2;
    border-radius: 10px 10px 0 0;
}

.chatbox .chatbox-room-head img {
    width: 30px;
    height: 30px;
    vertical-align: middle;
    margin-left: 20%;
    cursor: pointer;
}

.chatbox .chatbox-room-box {
    height: 240px;
    background-color: #efefef;
    border-bottom: 1px solid #ccc;
    overflow-y: scroll;
}

.chatbox .chatbox-room-input {
    height: 100px;
    text-align: right;
}

.chatbox .chatbox-room-input textarea {
    width: 98%;
    margin-top: 5px;
    height: 55px;
    border: 0;
    outline: none;
    resize: none;
}

.chatbox .chatbox-room-input button {
    height: 35px;
    line-height: 35px;
    width: 50px;
    background-color: #50E3C2;
    text-align: center;
    border: none;
    color: #fff;
    border-radius: 4px;
    outline: none;
    margin-bottom: 5px;
    margin-right: 5px;
    cursor: pointer;
}

.chatbox .chatbox-room-item-left {
    text-align: left;
    width: 95%;
    padding: 5% 0 0 5%;

}

.chatbox .chatbox-room-item-right {
    text-align: right;
    width: 95%;
    padding: 5% 0 5% 5%;

}

.chatbox .chatbox-room-ava {
    width: 30px;
    height: 30px;
    vertical-align: top;
}

.chatbox .chatbox-room-time {
    font-size: 10px;
    color: #ccc;
}

.chatbox .chatbox-room-msg {
    padding: 5px 10px;
    background-color: #fff;
    display: inline-block;
    border-radius: 10px;
    vertical-align: top;
    max-width: 220px;
    text-align: left;
}
</style>
<div class="chatbox" id="chatbox">
<!-- 提示框 -->
<div @click="show()" class="chatbox-notice" id="notice" v-cloak>
    <div class="chatbox-notice-font">
        聯絡客服
    </div>
    <img class="chatbox-notice-pic" src="http://www.awgo.top/assets/img/chat.png" alt="">
</div>
<!-- 聊天框 -->
<div class="chatbox-room" id="box" style="display: none;" v-cloak>
    <div class="chatbox-room-head">
        大中華醫療服務
        <img @click="show()" src="http://www.awgo.top/assets/img/min.png" alt="">
    </div>
    <div class="chatbox-room-box" id="boxindex">

        <template v-for="log in chatList">
            <div v-if="log.src_type == 'user'" class="chatbox-room-item-left">
                <img class="chatbox-room-ava" src="http://www.awgo.top/assets/img/kefu.png" alt="">
                <div class="chatbox-room-msg">
                    {{log.content}}
                </div>
                <div class="chatbox-room-time">{{timestampToTime(log.create_at)}}</div>
            </div>

            <div v-if="log.src_type == 'ip'" class="chatbox-room-item-right">
                <div class="chatbox-room-msg">
                    {{log.content}}
                </div>
                <img class="chatbox-room-ava" src="http://www.awgo.top/assets/img/user.png" alt="">
                <div class="chatbox-room-time">{{timestampToTime(log.create_at)}}</div>
            </div>
        </template>

    </div>
    <div class="chatbox-room-input">
        <textarea name="msg" placeholder="請輸入消息" v-model="msg"></textarea>
        <button @click="sendMsg()">發送</button>
    </div>
</div>
</div>`

var chat = new Vue({
    el: "#chatbox",
    data: {
        id: 0,
        uuid: "",
        showvalue: 0,
        webSocket: {},
        chatList: [],
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
            var template = theRequest["template"]

            if (template == "" || template == null || template == undefined) {
                document.write(template1)
            } else if (template == 1) {
                document.write(template1)
            }
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
        initWebSocket: function() {
            var url="wss://"+ "www.awgo.top" +"/api/chat?id=" + this.id + "&uuid=" + this.uuid + "&role=ip"
            this.webSocket = new WebSocket(url)
            //消息处理
            this.webSocket.onmessage = function(evt){
                if(evt.data.indexOf("}")>-1){
                    this.chatList.push(JSON.parse(evt.data))
                    console.log("display:", document.getElementById("box").style.display)
                    if (document.getElementById("box").style.display == "none") {
                        alert("你有一条新的消息")
                    }
                    this.scrollBottom()
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
            var ele = document.getElementById('boxindex');
            ele.scrollTop = ele.scrollHeight;
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
            this.scrollBottom()
        },
        //滚动到底部

    },
    created() {
        this.getValue()
        this.initWebSocket()
        this.scrollBottom()
    },
    mounted() {
        //this.scrollBottom()
    }
})