{{template "header"}}
    <title>Category - My Beego Blog</title>
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
		<h1>{{.Category}}</h1></div>
	<div class="page-header">
		{{range .Topic}}
		<h1><a href="/topic/{{.Id}}">{{.Title}}</a></h1>
		<h6 class="text-muted"><a class="label label-default" href="/category/{{.Category}}">{{.Category}}</a> Create in {{dateformat .Created "2006-01-02"}},{{.Views}} views,{{.ReplyCount}} replies.</h6>
		<p>{{.Content}}</p>
		{{end}}
	</div>
	</div>
	<div class="col-md-1"></div>
	<div class="col-md-2">
		<div class="page-header" data-spy="affix" data-offset-top="-20">
			<h3>Categories</h3>
			<ul>
				<li><a href="/">all</a></li>
			</ul>
			{{range .Categories}}
			<ul>
				<li><a href="/category/{{.Title}}">{{.Title}}</a></li>
			</ul>
			{{end}}
		</div>
	</div>
</div>
</body>
</html>