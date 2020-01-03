var template1= `<style>
[v-cloak] {
    display: none;
}

.chatbox {
    position: fixed;
    bottom: 0;
    right: 5%;
    min-width: 255px;
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
    max-width: 170px;
    text-align: left;
}
</style>
<div class="chatbox" id="chatbox">
<!-- 提示框 -->
<div @click="show()" class="chatbox-notice" id="notice" v-cloak>
    <div class="chatbox-notice-font">
        聯絡客服
    </div>
    <img class="chatbox-notice-pic" src="http://cus.gcbodycheck.com/assets/img/chat.png" alt="">
</div>
<!-- 聊天框 -->
<div class="chatbox-room" id="box" style="display: none;" v-cloak>
    <div class="chatbox-room-head">
        大中華醫療服務
        <img @click="show()" src="http://cus.gcbodycheck.com/assets/img/min.png" alt="">
    </div>
    <div class="chatbox-room-box" id="boxindex">

        <template v-for="log in chatList">
            <div v-if="log.src_type == 'user'" class="chatbox-room-item-left">
                <img class="chatbox-room-ava" src="http://cus.gcbodycheck.com/assets/img/kefu.png" alt="">
                <div class="chatbox-room-msg">
                    {{log.content}}
                </div>
                <div class="chatbox-room-time">{{timestampToTime(log.create_at)}}</div>
            </div>

            <div v-if="log.src_type == 'ip'" class="chatbox-room-item-right">
                <div class="chatbox-room-msg">
                    {{log.content}}
                </div>
                <img class="chatbox-room-ava" src="http://cus.gcbodycheck.com/assets/img/user.png" alt="">
                <div class="chatbox-room-time">{{timestampToTime(log.create_at)}}</div>
            </div>
        </template>

    </div>
    <div class="chatbox-room-input">
        <textarea name="msg" placeholder="請輸入消息" v-model="msg" id="input"></textarea>
        <button @click="sendMsg()">發送</button>
    </div>
</div>
</div>`

var template2 = `
    <style>
        [v-cloak]{  
    display: none;  
}  

.chatbox {
    position: fixed;
    bottom: 100px;
    right: 80px;
}

.chatbox .chatbox-notice{
    width: 50px;
    height: 50px;
    cursor: pointer;
}
.chatbox .chatbox-notice img{
    width: 100%;
    height: 100%;
}
.chatbox .chatbox-notice-font{
    display: inline-block;
    margin-left: 30%;
}
.chatbox .chatbox-notice-pic{
    display: inline-block;
    width: 30px;
    vertical-align: middle;

}

.chatbox .chatbox-room{
    width: 400px;
    height: 250px;
    background-color: #fff;
    border-radius: 20px;
    box-shadow: 1px 1px 20px #9c9999;;
    position: fixed;
    bottom: 180px;
    right: 50px;
}
.chatbox .chatbox-room-head{
    height: 33px;
    line-height: 33px;
    padding-left: 50px;
    background-color: #fff;
    border-radius: 10px 10px 0 0;
}
.chatbox .chatbox-head-tit{
    display: inline-block;
    color: #7F7F7F;
}
.chatbox .chatbox-close{
    width: 25px;
    height: 25px;
    position: absolute;
    line-height: 22px;
    right: 5px;
    top: 5px;
    text-align: center;
    border-radius: 100px;
    background-color: #E4E4E4;
}
.chatbox .chatbox-room-head img{
    width: 20px;
    height: 20px;
    cursor: pointer;
    -ms-overflow-style: none;
}
.chatbox .chatbox-room-box{
    height: 160px;
    overflow-y: scroll;
}

.chatbox-room-box::-webkit-scrollbar {
    display: none;
}

.chatbox .chatbox-room-input{
    padding: 5px;
    height: 40px;
    display: none;
}
.chatbox .chatbox-room-input input{
    width: 85%;
    margin-top: 5px;
    height: 100%;
    border: 0;
    outline: none;
    resize:  none;
}
.chatbox .chatbox-room-input button{
    height: 35px;
    line-height: 35px;
    width: 50px;
    background-color: #50E3C2;
    text-align: center;
    border: none;
    color:#fff;
    border-radius: 4px;
    outline: none;
    margin-bottom: 5px;
    margin-right: 5px;
    cursor: pointer;
}
.chatbox .chatbox-room-item-left{
    text-align: left;
    width: 95%;
    padding: 10px 0 0 5%;

}
.chatbox .chatbox-room-item-right{
    text-align: right;
    width: 95%;
    padding: 10px 5% 0 0;
    
}
.chatbox .chatbox-room-ava{
    width: 30px;
    height: 30px;
    vertical-align: middle;
}
.chatbox .chatbox-room-time{
    font-size: 10px;
    color: #ccc;
}
.chatbox .chatbox-room-msg{
    padding: 5px 10px;
    background-color: #E4E4E4;
    display: inline-block;
    border-radius: 10px;
    vertical-align: middle;
    max-width: 300px;
    text-align: left;
}
.chatbox .chatbox-room-btn{
    display: inline-block;
    color: #7F7F7F;
    cursor: pointer;
}
.chatbox .chatbox-init{
    background-color: #D4A88C;
    height: 40px;
    line-height: 40px;
    text-align: center;
    border-radius: 10px;
    color: #fff;
    width: 50%;
    position: absolute;
    top: 130px;
    left: 25%;
    cursor: pointer;
}
@media (max-width:765px) {
    .chatbox .chatbox-room{
        width: 80% !important;
        bottom: 20px !important;
        right: 65px !important;
    }
    .chatbox .chatbox-room-msg{
        max-width: 220px !important;
    }
    .chatbox {
        bottom: 13px !important;
        right: 13px !important;
    }
}
    </style>
    
    <div class="chatbox" id="chatbox">
        <!-- 提示框 -->
        <div @click="show2()" class="chatbox-notice" id="notice" v-cloak>
            <img class="chatbox-notice-pic" src="http://cus.gcbodycheck.com/assets/img/btn1.png" alt="">
        </div>
        <!-- 聊天框 -->
        <div class="chatbox-room" id="box" style="display: none;" v-cloak>
            <div class="chatbox-room-head">
                <div class="chatbox-head-tit">
                    大中華醫療服務
                </div>
                <div class="chatbox-close">
                    <img @click="show2()" src="http://cus.gcbodycheck.com/assets/img/close.png" alt="">
                </div>
            </div>
            <div class="chatbox-room-box" id="boxcroll">

                <template v-for="log in chatList">
                    <div v-if="log.src_type == 'user'" class="chatbox-room-item-left">
                        <img class="chatbox-room-ava" src="http://cus.gcbodycheck.com/assets/img/company1.png" alt="">
                        <div class="chatbox-room-msg">
                            {{log.content}}
                        </div>
                        <div class="chatbox-room-time">{{timestampToTime(log.create_at)}}</div>
                    </div>

                    <div v-if="log.src_type == 'ip'" class="chatbox-room-item-right">
                        <div class="chatbox-room-msg">
                            {{log.content}}
                        </div>
                        <img class="chatbox-room-ava" src="http://cus.gcbodycheck.com/assets/img/user.png" alt="">
                        <div class="chatbox-room-time">{{timestampToTime(log.create_at)}}</div>
                    </div>
                </template>

            </div>
            <div class="chatbox-room-input" id="msginput">
                <input name="msg" placeholder="請輸入消息" v-model="msg" id="input">
                <div class="chatbox-room-btn" @click="sendMsg()">
                    發送
                </div>
            </div>
            <div class="chatbox-init" @click="beginChat()" id="beginchat">
                聯絡大中華醫療
            </div>
        </div>
    </div>
`

var template3 = `
    <style>
        [v-cloak]{  
    display: none;  
}  

.chatbox {
    position: fixed;
    bottom: 130px;
    right: 100px;
}

.chatbox .chatbox-notice{
    width: 100px;
    height: 100px;
    line-height: 100px;
    text-align: center;
    cursor: pointer;
}
.chatbox .chatbox-notice img{
    width: 89px;
    height: 89px;
}
.chatbox .chatbox-notice-font{
    display: inline-block;
    margin-left: 30%;
}
.chatbox .chatbox-notice-pic{
    display: inline-block;
    width: 30px;
    vertical-align: middle;

}

.chatbox .chatbox-room{
    width: 400px;
    height: 350px;
    background-color: #fff;
    border-radius: 20px;
    box-shadow: 1px 1px 20px #9c9999;;
    position: fixed;
    bottom: 240px;
    right: 50px;
}
.chatbox .chatbox-room-head{
    height: 49px;
    line-height: 49px;
    background-color: #177CB0;
    border-radius: 10px 10px 0 0;
}
.chatbox .chatbox-head-tit{
    display: inline-block;
    color: #7F7F7F;
}
.chatbox .chatbox-close{
    width: 25px;
    height: 25px;
    position: absolute;
    line-height: 24px;
    right: 20px;
    top: 14px;
    text-align: center;
    border-radius: 100px;
    background-color: #E4E4E4;
}
.chatbox .chatbox-room-head img{
    width: 20px;
    height: 20px;
    cursor: pointer;
    -ms-overflow-style: none;
}
.chatbox .chatbox-room-box{
    height: 244px;
    overflow-y: scroll;
}

.chatbox-room-box::-webkit-scrollbar {
    display: none;
}

.chatbox .chatbox-room-input{
    padding: 5px;
    height: 40px;
}
.chatbox .chatbox-room-input input{
    width: 85%;
    margin-top: 5px;
    height: 100%;
    border: 0;
    outline: none;
    resize:  none;
}
.chatbox .chatbox-room-input button{
    height: 35px;
    line-height: 35px;
    width: 50px;
    background-color: #50E3C2;
    text-align: center;
    border: none;
    color:#fff;
    border-radius: 4px;
    outline: none;
    margin-bottom: 5px;
    margin-right: 5px;
    cursor: pointer;
}
.chatbox .chatbox-room-item-left{
    text-align: left;
    width: 95%;
    padding: 10px 0 0 5%;

}
.chatbox .chatbox-room-item-right{
    text-align: right;
    width: 95%;
    padding: 10px 5% 0 0;
    
}
.chatbox .chatbox-room-ava{
    width: 30px;
    height: 30px;
    vertical-align: middle;
}
.chatbox .chatbox-room-time{
    font-size: 10px;
    color: #ccc;
}
.chatbox .chatbox-room-msg{
    padding: 5px 10px;
    background-color: #E4E4E4;
    display: inline-block;
    border-radius: 10px;
    vertical-align: middle;
    max-width: 300px;
    text-align: left;
}
.chatbox .chatbox-room-btn{
    display: inline-block;
    color: #7F7F7F;
    cursor: pointer;
}
.chatbox .chatbox-init{
    background-color: #177CB0;
    height: 40px;
    line-height: 40px;
    text-align: center;
    border-radius: 10px;
    color: #fff;
    width: 50%;
    position: absolute;
    top: 130px;
    left: 25%;
    cursor: pointer;
}
@media (max-width:765px) {
    .chatbox .chatbox-notice{
        width: 50px;
        height: 50px;
        line-height: 50px;
    }
    .chatbox .chatbox-notice-pic{
        width: 100% !important;
        height: 100% !important;
    }
    .chatbox .chatbox-room{
        width: 80% !important;
        bottom: 55px !important;
        right: 69px !important;
    }
    .chatbox .chatbox-room-msg{
        max-width: 188px !important;
    }
    .chatbox {
            bottom: 48px !important;
            right: 5% !important;
    }
}
    </style>
    
    <div class="chatbox" id="chatbox">
        <!-- 提示框 -->
        <div @click="show3()" class="chatbox-notice" id="notice" v-cloak>
            <img class="chatbox-notice-pic" src="http://cus.gcbodycheck.com/assets/img/btn2.png" alt="">
        </div>
        <!-- 聊天框 -->
        <div class="chatbox-room" id="box" style="display: none;" v-cloak>
            <div class="chatbox-room-head">
                <div class="chatbox-head-tit">
                   <img src="http://cus.gcbodycheck.com/assets/img/tit1.png" alt="" style="width:200px !important;height: 50px;">
                </div>
                <div class="chatbox-close">
                    <img @click="show3()" src="http://cus.gcbodycheck.com/assets/img/close.png" alt="">
                </div>
            </div>
            <div class="chatbox-room-box" id="boxcroll">

                <template v-for="log in chatList">
                    <div v-if="log.src_type == 'user'" class="chatbox-room-item-left">
                        <img class="chatbox-room-ava" src="http://cus.gcbodycheck.com/assets/img/kefu.png" alt="">
                        <div class="chatbox-room-msg">
                            {{log.content}}
                        </div>
                        <div class="chatbox-room-time">{{timestampToTime(log.create_at)}}</div>
                    </div>

                    <div v-if="log.src_type == 'ip'" class="chatbox-room-item-right">
                        <div class="chatbox-room-msg">
                            {{log.content}}
                        </div>
                        <img class="chatbox-room-ava" src="http://cus.gcbodycheck.com/assets/img/user.png" alt="">
                        <div class="chatbox-room-time">{{timestampToTime(log.create_at)}}</div>
                    </div>
                </template>

            </div>
            <div class="chatbox-room-input" id="msginput">
                <input name="msg" placeholder="請輸入消息" v-model="msg" id="input">
                <div class="chatbox-room-btn" @click="sendMsg()">
                    發送
                </div>
            </div>
        </div>
    </div>
`

var chat = new Vue({
    el: "#chatbox",
    data: {
        id: 0,
        uuid: "",
        showvalue: 0,
        webSocket: {},
        one: 0,
        chatList: [
            {"group_id": 0, "ip_id": 0, "src_type": "user", "cmd": "msg", "content": "你好！我們有什麼可以幫到您的？", "create_at": parseInt(((new Date()).valueOf().toString()).substring(0, ((new Date()).valueOf().toString()).length - 3))}
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
            var template = theRequest["template"]

            if (template == "" || template == null || template == undefined) {
                document.write(template1)
            } else if (template == 1) {
                document.write(template1)
            } else if (template == 2) {
                document.write((template2))
            } else if (template == 3) {
                document.write(template3)
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
        show2: function() {
            if (this.showvalue == 0) {
                this.showvalue = 1
                $("#box").fadeIn()
            } else {
                this.showvalue = 0
                $("#box").fadeOut()
            }
        },
        show3: function() {
            if (this.showvalue == 0) {
                this.showvalue = 1
                $("#box").fadeIn()
                if (this.one != 1) {
                    this.initWebSocket()
                    this.one = 1
                }
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
                this.scrollBottom()
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

            var id = window.localStorage.getItem("id")
            var time1 = (new Date()).valueOf().toString()
            var time2 = parseInt(time1.substring(0, time1.length - 3))

            var struct = {
                "group_id": parseInt(this.id),
                "ip_id": parseInt(0),
                "src_type": "ip",
                "cmd": "msg",
                "content": this.msg,
                "create_at": time2
            }
            this.chatList.push(struct)
            this.msg = ""
            this.webSocket.send(JSON.stringify(struct))
            this.scrollBottom()
        },
    },
    created() {
        this.getValue()
    },
    mounted() {
        $("#input").keypress(function (e) {
            if (e.which == 13) {
                chat.sendMsg()
            }
        });
    }
})