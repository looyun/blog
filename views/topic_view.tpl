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
	{{with .Topic}}
	<h1>{{.Title}}
	<a href="/topic/modify/{{.Id}}" class="btn btn-default">modify topic</a></h1>
	{{.Content}}
	{{end}}
</div>
</body>
</html>