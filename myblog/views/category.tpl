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
	<h1>Category</h1>
	<form method="GET" action="/category">
		  <div class="form-group">
		  <label>Category</label>
		  <input id="name" class="form-control" placeholder="name" name="name">
		  </div>
		  <input type="hidden" name="op" value="add">
		  <button type="submit" class="btn btn-default" onclick="return checkInput();">Add
		  </button>
	</form>

		<script type="text/javascript">
			function checkInput() {
				var name=document.getElementById("name");
				if(name.value.length==0){
					alert("please input category name!");
					return false
				}
			}
		</script>
		<table class="table table-striped">
			<thead>
				<tr>
					<th>#</th>
					<th>name</th>
					<th>category number</th>
					<th>operation</th>
				</tr>
			</thead>
			<thead>
			{{range .Categories}}
			<tr>
				<th>{{.Id}}</th>
				<th>{{.Title}}</th>
				<th>{{.TopicCount}}</th>
				<th>
					<a href="/category?op=del&id={{.Id}}">delete</a>
				</th>
			<tr>
			{{end}}
			</thead>
		</table>
</div>
</body>
</html>