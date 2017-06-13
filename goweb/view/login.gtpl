<html>
    <head>
        <title>Login Page</title>
    </head>
    <body>
        <form action="/login" method="post">
            Username:<input type="text" name="username">
            Password:<input type="password" name="password">
            <input type="hidden" name="token" value="{{.}}">
            <input type="submit" value="Login">
        </form>

        <p>Current List of Users are:</p>
        {{range .}}
        <li>{{.Emailaddress}}</li>
        {{end}}
    </body>
</html>
