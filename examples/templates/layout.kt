<!doctype html>
<html lang=en>
    <head>
		<meta charset='utf-8'>
        <link href='/static/style.css' type='text/css' rel='stylesheet'>
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
