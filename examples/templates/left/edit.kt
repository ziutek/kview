<h4>Edit page</h4>

$for _, p in @[1]:
    <p>$:p</p>
$end
<p style='font-size: 75%'>
    (number of paragraphs printed above: $printf("%04d", len(@[1])))
</p>
