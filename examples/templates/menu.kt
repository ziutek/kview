<ul>
$for i, v in Content:
    <li><a href='$v.Url'$if i==Selected: id='CURRENT'$end>$v.Name</a></li>
$end
</ul>
