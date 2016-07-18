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
	<div class="col-md-1"></div>
	<div class="col-md-8">
		{{range .Topics}}
		<div class="page-header">
			<h1><a href="/topic/{{.Id}}">{{.Title}}</a></h1>
			<h6 class="text-muted"><a class="label label-default" href="/category/{{.Category}}">{{.Category}}</a> Create in {{dateformat .Created "2006-01-02"}},{{.Views}} views,{{.ReplyCount}} replies.</h6>
			<p>{{.Content}}</p>
		</div>
		{{end}}

		<div class="copyright"><p>© 2015 - 2016 <a href="https://github.com/looyun">looyun</a>, unless otherwise noted.</p></div>
	</div>
	<div class="col-md-1"></div>
	<div class="col-md-2">
		<div class="page-header" data-spy="affix" data-offset-top="-20">
			<h3>Categories</h3>
			<ul>
				<li><a href="/">all</a></li>
			</ul>
			{{range .Category}}
			<ul>
				<li><a href="/category/{{.Title}}">{{.Title}}</a></li>
			</ul>

			{{end}}
			<ul><a name="BackToTop" href="/#top">Top</a></ul>
		</div>
	</div>
</div>
</body>
</html>