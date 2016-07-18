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

		  <div class="row">
		  	<div class="col-xs-2">
		  <label>category</label>
		  <select class="form-control" name="category">
		  {{range .Category}}
		    <option>{{.Title}}</option>
		    {{end}}
		  </select>
		  	</div>
		  </div>

	<div class="form-group">
		<label>content</label>
		<textarea name="content" cols="30" rows="10" class="form-control"></textarea>
	</div>
	<button type="submit" class="btn btn-default">submit</button>
	</form>
</div>

    <script src="//cdn.bootcss.com/jquery/3.0.0/jquery.min.js"></script>
	<script src="/static/js/bootstrap.min.js"></script>
	<script src="/static/js/bootstrap.js"></script>
</body>
</html>