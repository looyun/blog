{{define "navbar"}}

    <div class="col-md-1"></div>
    <div class="container">
<div class="col-md-8">
    <a class="navbar-brand" href="/">My Blog</a>
    <ul class="nav navbar-nav">
        <li {{if .IsHome}}class="active"{{end}}><a href="/">Home</a></li>
        <li {{if .IsCategory}}class="active"{{end}}><a href="/category">Category</a></li>
        <li {{if .IsTopic}}class="active"{{end}}><a href="/topic">Topic</a></li>
    </ul>
</div>
    <div class="pull-right">
    	<ul class="nav navbar-nav">
    		{{if .IsLogin}}
    		<li><a href="/login?exit=true">log out</a></li>
    		{{else}}
    		<li><a href="/login">admin login</a></li>
    		{{end}}
    	</ul>
    </div>
</div>
{{end}}