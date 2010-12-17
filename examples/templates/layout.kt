<!DOCTYPE HTML PUBLIC '-//W3C//DTD HTML 4.01 Transitional//EN'>
<html>
    <head>
        <meta http-equiv='Content-type' content='text/html; charset=utf-8'>
        <link href='style.css' type='text/css' rel='stylesheet'>
        <title>$title</title>
    </head>
    <body>
        <div id='Top'>Title of the page</div>
        <div id='Menu'>$Menu.Render(ctx.menu)</div>
        <div id='Container0'>
            <div id='Container1'>
                <div id='Left'>$Left.Render(ctx.left)</div>
                <div id='Right'>$Right.Render(ctx.right)</div>
            </div>
        </div>
        <div id='Bottom'>Service started: $started, hits: $hits</div>
    </body>
</html>
