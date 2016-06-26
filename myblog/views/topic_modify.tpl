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
	<h1>Modify Topic</h1>
	<form method="POST" action="/topic">
	<input type="hidden" name="id" value="{{.Id}}">
		  <div class="form-group">
		  	<label>title</label>
		  	<input type="text" name="title" class="form-control" value="{{.Topic.Title}}">
		  </div>
	<div class="form-group">
		<label>content</label>
		<textarea name="content" cols="30" rows="10" class="form-control">{{.Topic.Content}}</textarea>
	</div>
	<button type="submit" class="btn btn-default">Edit</button>
	</form>
</div>
</body>
</html>