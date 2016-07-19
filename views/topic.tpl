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
        <div class="col-md-10">
            <h1>Topic</h1>
            <a href="/topic/add" class="btn btn-default">add topic</a>
            <table class="table table striped">
                <thead>
                    <tr>
                        <th>#</th>
                        <th>topic name</th>
                        <th>last update</th>
                        <th>views</th>
                        <th>reply count</th>
                        <th>operation</th>
                    </tr>
                </thead>
                <tbody>
                    {{range .Topics}}
                    <tr>
                        <th>{{.Id}}</th>
                        <th><a href="/topic/{{.Id}}">{{.Title}}</a></th>
                        <th>{{dateformat .Updated "2006-01-02"}}</th>
                        <th>{{.Views}}</th>
                        <th>{{.ReplyCount}}</th>
                        <th><a href="/topic/modify/{{.Id}}">modify </a><a href="/topic/delete/{{.Id}}/{{.Category}}">delete</a></th>
                    </tr>
                    {{end}}
                </tbody>
            </table>
        </div>
    </div>
</body>

</html>
