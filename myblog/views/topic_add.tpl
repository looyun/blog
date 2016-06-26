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
	<h1>Add Topic</h1>
	<form method="POST" action="/topic">
		  <div class="form-group">
		  	<label>title</label>
		  	<input type="text" name="title" class="form-control">
		  </div>
	<div class="form-group">
		<label>content</label>
		<textarea name="content" cols="30" rows="10" class="form-control"></textarea>
	</div>
	<button type="submit" class="btn btn-default">submit</button>
	</form>
</div>
</body>
</html>