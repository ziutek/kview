<ul>
$for i, v in content:
    <li><a href='$v.url'$if i==selected: id='CURRENT'$end>$v.name</a></li>
$end
</ul>
