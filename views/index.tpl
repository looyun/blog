{{template "header"}}
    <title>Home - My Beego Blog</title>
</head>
<body>
<nav class="navbar navbar-default">
  <div class="container">
    {{template "navbar" .}}
  </div>
</nav>
<div class="container">
		{{range .Topics}}
	<div class="page-header">
		<h1><a href="/topic/view/{{.Id}}">{{.Title}}</a></h1>
		<h6 class="text-muted">Create in{{.Created}},{{.Views}} views,{{.ReplyCount}} replies.</h6>
		<p>{{.Content}}</p>
		{{end}}
	</div>
</div>
</body>
</html>