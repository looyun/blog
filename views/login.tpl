{{template "header"}}
<title>Login - My Beego Blog</title>
</head>

<body>
    <nav class="navbar">
    </nav>
    <div class="container" style="width: 500px">
        <form method="post" action="/login">
            <div class="form-group">
                <label>User</label>
                <input id="uname" class="form-control" placeholder="admin" name="uname">
            </div>
            <div class="form-group">
                <label>Password</label>
                <input id="pwd" type="password" class="form-control" placeholder="password" name="pwd">
            </div>
            <div class="checkbox">
                <label>
                    <input type="checkbox" name="auotologin"> Auto login
                </label>
            </div>
            <button type="submit" class="btn btn-default" onclick="return checkInput();">Login</button>
            <div class="pull-right">
                <button class="btn btn-default" onclick="return backToHome();">Back</button>
            </div>
        </form>
        <script type="text/javascript">
        function checkInput() {
            var uname = document.getElementById("uname");
            if (uname.value.length == 0) {
                alert("please input account!");
                return false
            }

            var pwd = document.getElementById("pwd");
            if (pwd.value.length == 0) {
                alert("please input password!");
                return false
            }
        }

        function backToHome() {
            window.location.href = "/";
            return false;
        }
        </script>
    </div>
</body>

</html>
