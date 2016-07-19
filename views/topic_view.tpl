{{template "header"}}
<title>Topic - My Beego Blog</title>
</head>

<body>
    <nav class="navbar navbar-default">
        <div class="container">
            {{template "navbar" .}}
        </div>
    </nav>
    <div class="container">
        <div class="col-md-1"></div>
        <div class="col-md-8">
            <div class="page-header">
                {{with .Topic}}
                <h1>{{.Title}}<a href="/topic/modify/{{.Id}}" class="btn btn-primary btn-sm">modify</a></h1>
                <p>Attachment: <a href="/attachment/{{.Id}}/{{.Attachment}}">{{.Attachment}}</a></p>
                <h5 class="text-muted"><a class="label label-default" href="/category/{{.Category}}">{{.Category}}</a> Create in {{dateformat .Created "2006-01-02"}},{{.Views}} views,{{.ReplyCount}} replies.</h5>
                <p>{{.Content}}</p>
                {{end}}
            </div>
            <div>
                <h3>Comment:</h3>
                <form method="POST" action="/comment">
                    <input type="hidden" name="tid" value={{.Topic.Id}}>
                    <div class="form-group">
                        <label>name</label>
                        <input type="text" name="name" class="form-control" id="name">
                    </div>
                    <div class="form-group">
                        <label>content</label>
                        <textarea name="content" id="content" cols="20" rows="4" class="form-control"></textarea>
                    </div>
                    <button type="submit" class="btn btn-default" onclick="return checkInput();">comment</button>
                    <div class="pull-right">
                        <button class="btn btn-default" onclick="return backToHome();">Back</button>
                    </div>
                </form>
                <div class="page-header" id="comment">
                    <h3>Comment Area:</h3> {{$t:=.Topic.Id}} {{range .Comment}}
                    <div class="panel panel-default">
                        <div class="panel-heading">
                            <h3 class="panel-title">
                            {{.Name}}:
                            	<div class="pull-right">
                            		<a type="submit" class="btn btn-danger btn-xs" href="/comment?id={{.Id}}&tid={{$t}}"><span class="glyphicon glyphicon-remove" aria-hidden="true"></span>
                            		</a>
                            	</div>
        					</h3>
                        </div>
                        <div class="panel-body">
                            {{.Content}}
                        </div>
                    </div>
                    {{end}}
                </div>
            </div>
        </div>
    </div>
    <script type="text/javascript">
    function checkInput() {
        var name = document.getElementById("name");
        if (name.value.length == 0) {
            alert("please input your name!");
            return false
        }

        var content = document.getElementById("content");
        if (content.value.length == 0) {
            alert("content can't be empty!");
            return false
        }
    }

    function backToHome() {
        window.location.href = "/topic";
        return false;
    }
    </script>
</body>

</html>
