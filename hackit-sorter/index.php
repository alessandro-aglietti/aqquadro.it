<?php
$con=mysqli_connect("127.0.0.1","aqquadro","2Jw4ZqRFZ6b6HSeC","aqquadro");

// Check connection
if (mysqli_connect_errno()) {
	echo "Failed to connect to MySQL: " . mysqli_connect_error();
}
$rand = rand ( 0, 1 );

if ( $rand == 0 ) {
	$rand = "Bologna";
} else {
	$rand = "Napoli";
}

$stamp = time();

$ip = $_SERVER['REMOTE_ADDR'];

$insert = "INSERT INTO hackit (ip, stamp, value) VALUES ('" . $ip . "', " . $stamp . ", '" . $rand . "')";

$hitsQuery = "SELECT COUNT(*) hits FROM hackit";

$result = mysql_query($hitsQuery);
$row = mysql_fetch_assoc($result);
$hits = $row['hits'];

$hitsQuery = "SELECT COUNT(*) hits FROM hackit WHERE value = 'Napoli'";

$result = mysql_query($hitsQuery);
$row = mysql_fetch_assoc($result);
$naphits = $row['hits'];

$hitsQuery = "SELECT COUNT(*) hits FROM hackit WHERE value = 'Bologna'";

$result = mysql_query($hitsQuery);
$row = mysql_fetch_assoc($result);
$bohits = $row['hits'];

mysqli_close($con);
?>
<html>
    <head>
    </head>
    <body style="padding-top: 50px;">
		<p>Hits: <?php echo $hits?></p>
		<p>Napoli: <?php echo $naphits?></p>
		<p>Bologna: <?php echo $bohits?></p>
    </body>
</html>