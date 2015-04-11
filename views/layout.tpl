<!DOCTYPE html>
<html>
<head>
    <title>Lin Li</title>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <link rel="stylesheet" href="static/bootstrap/dist/css/bootstrap.min.css">
    <link rel="stylesheet" href="static/bootstrap/dist/css/bootstrap-theme.min.css">
    {{.HtmlHead}}
</head>
<body>

    <div class="container">
        {{.FormTodo}}
        {{.LayoutContent}}
        <div class="sidebar">
            {{.SideBar}}
        </div>    
    </div>

    {{.Scripts}}
</body>
</html>