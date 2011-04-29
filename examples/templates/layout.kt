<!DOCTYPE HTML PUBLIC '-//W3C//DTD HTML 4.01 Transitional//EN'>
<html>
    <head>
        <meta http-equiv='Content-type' content='text/html; charset=utf-8'>
        <link href='style.css' type='text/css' rel='stylesheet'>
        <title>$Title</title>
    </head>
    <body>
        <div id='Top'>Title of the page</div>
        <div id='Menu'>$menu.Render(Menu)</div>
        <div id='Container0'>
            <div id='Container1'>
                <div id='Left'>$left.Render(Left)</div>
                <div id='Right'>$right.Render(Right)</div>
            </div>
        </div>
        <div id='Bottom'>
            Started: $Started; Hits: $Hits; Last client: $LastCliAddr
        </div>
    </body>
</html>
