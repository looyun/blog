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
            <h1>Add Topic</h1>
            <form method="POST" action="/topic" enctype="multipart/form-data">
                <div class="form-group">
                    <label>title</label>
                    <input type="text" name="title" class="form-control">
                </div>
                <div class="form-group">
                    <label>category</label>
                    <select class="form-control" name="category">
                        {{range .Category}}
                        <option>{{.Title}}</option>
                        {{end}}
                    </select>
                </div>
                <div class="form-group">
                    <label>content</label>
                    <textarea name="content" cols="30" rows="10" class="form-control"></textarea>
                </div>
                <div class="form-group">
                    <label>attachment</label>
                    <input type="file" name="attachment" class="form-control">
                </div>
                <button type="submit" class="btn btn-default">submit</button>
            </form>
        </div>
    </div>
    <script src="//cdn.bootcss.com/jquery/3.0.0/jquery.min.js"></script>
    <script src="/static/js/bootstrap.min.js"></script>
    <script src="/static/js/bootstrap.js"></script>
</body>

</html>
