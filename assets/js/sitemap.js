var sitemap = new Vue({
    delimiters: ['${', '}'],
    el: "#sitemap",
    data: {
        appList: [],
        page_size: 10,
        page_index: 0,
        pagelist: [],
        search: "",
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
        //加载app记录
        LoadAppList: function() {
            var id = window.localStorage.getItem("id")
            var token = window.localStorage.getItem("token")
            axios({
                method: "post",
                url: "/api/app/list",
                data: {
                    "id": parseInt(id),
                    "token": token,
                    "page_size": parseInt(this.page_size),
                    "page_index": parseInt(this.page_index),
                    "search": this.search,
                    "total": 0,
                },
                headers: {
                    "Content-Type": "application/json"
                }
            }).then(res => {
                if (res.data.code == 0) {
                    this.appList = res.data.rows
                    this.total = res.data.total
                    this.pageparse(res.data.total)
                } else if (res.data.code == -1) {
                    alert(res.data.msg)
                    if (res.data.path != null) {
                        window.location.href = res.data.path
                    }
                }
            })
        },
        //重置UUID
        ResetUUID: function(aid) {
            if (confirm("确认重置UUID嘛")) {
                var id = window.localStorage.getItem("id")
                var token = window.localStorage.getItem("token")
                axios({
                    method: "post",
                    url: "/api/app/resetuuid",
                    data: {
                        "id": parseInt(id),
                        "token": token,
                        "aid": aid
                    },
                    headers: {
                        "Content-Type": "application/json"
                    }
                }).then(res => {
                    console.log(res)
                    if (res.data.code == 0) {
                        alert(res.data.msg)
                        this.LoadAppList()
                    } else if (res.data.code == -1) {
                        alert(res.data.msg)
                        if (res.data.path != null) {
                            window.location.href = res.data.path
                        }
                    }
                })
            }

        },
        //删除站点
        DeleteSite: function(aid) {
            if (confirm("确认删除站点嘛")) {
                var id = window.localStorage.getItem("id")
                var token = window.localStorage.getItem("token")
                axios({
                    method: "post",
                    url: "/api/app/delete",
                    data: {
                        "id": parseInt(id),
                        "token": token,
                        "aid": aid
                    },
                    headers: {
                        "Content-Type": "application/json"
                    }
                }).then(res => {
                    if (res.data.code == 0) {
                        alert(res.data.msg)
                        this.LoadAppList()
                    } else if (res.data.code == -1) {
                        alert(res.data.msg)
                        if (res.data.path != null) {
                            window.location.href = res.data.path
                        }
                    }
                })
            }
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
        //解析分页按钮
        pageparse: function(all) {
            //求出总共的页数
            var pageNum = Math.ceil(all / this.page_size)
            //求出当前的页数
            var pageNow = Math.ceil(this.page_index / this.page_size) + 1
            //记录显示的分页数据
            var list  = new Array()

            if (pageNow <= 3) {
                var loop_start = 1
            } else {
                var loop_start = pageNow - 2
            }
            if ((pageNow + 2) >  pageNum) {
                var loop_end = pageNum
            } else {
                if (loop_start == 1) {
                    if (pageNum < 5) {
                        var loop_end = pageNum
                    } else {
                        var loop_end = 5
                    }
                } else {
                    var loop_end = pageNow + 2
                }
            }
            for (loop_start; loop_start <= loop_end; loop_start++) {
                if (loop_start == pageNow) {
                    list.push({"num": loop_start, "class":"am-active"})
                } else {
                    list.push({"num": loop_start, "class":""})
                }
            }
            this.pagelist = list

            //渲染快进和快退按钮
            document.getElementById("uPage").className = ""
            document.getElementById("downPage").className = ""
            if (pageNow == 1) {
                document.getElementById("uPage").classList.add("am-disabled")
            } else if (pageNow == pageNum) {
                document.getElementById("downPage").classList.add("am-disabled")
            }
        },
        //页面跳转
        jumPage: function(num) {
            this.page_index = ((num - 1) * this.page_size) - 1
            if (this.page_index < 0) {
                this.page_index = 0
            }
            //重新获取数据
            this.LoadAppList()
        },
        //快进
        upPage: function() {
            this.page_index = this.page_index - (this.page_size * 5)
            if (this.page_index < 0) {
                this.page_index = 0
            }
            this.LoadAppList()
        },
        //快退
        downPage: function() {
            this.page_index = this.page_index + (this.page_size * 5)
            if (this.page_index > (this.total - this.page_size)) {
                this.page_index = this.total - this.page_size
            }
            this.LoadAppList()
        },
        //搜索
        searchSite: function() {
            this.LoadAppList()
        }
    },
    created() {
        this.checkToken()
        this.LoadAppList()
    }
})