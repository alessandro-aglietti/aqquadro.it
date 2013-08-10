<?php
	$headers = getallheaders();
	if( strrpos(strtolower($headers["User-Agent"]), 'lynx') !== FALSE ) {
		header("HTTP/1.1 301 Moved Permanently");
		header("Location: text.php");
		exit();
	}
	$imageID = rand ( 0, 93 );
	header("Expires: Thu, 04 May 1989 14:00:00 GMT");
	header("Last-Modified: " . gmdate("D, d M Y H:i:s") . " GMT");
	header("Cache-Control: no-store, no-cache, must-revalidate");
	header("Cache-Control: post-check=0, pre-check=0", false);
	header("Pragma: no-cache");
?>
<html lang="en">
    <head>
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<title>Alessandro Aglietti</title>
		<script src="//code.jquery.com/jquery-1.10.1.min.js"></script>
		<script src="//code.jquery.com/jquery-migrate-1.2.1.min.js"></script>
		<link rel="stylesheet" href="//netdna.bootstrapcdn.com/bootstrap/3.0.0-rc1/css/bootstrap.min.css">
		<link href="//netdna.bootstrapcdn.com/font-awesome/3.2.1/css/font-awesome.min.css" rel="stylesheet">
		<script src="//netdna.bootstrapcdn.com/bootstrap/3.0.0-rc1/js/bootstrap.min.js"></script>
        <script type="text/javascript">
			$(function(){
				$.ajax({
					url: "deliciuos.php?r=<?php echo $imageID ?>"
				}).done(function ( data ) {
					$("#links").children().last().remove();
					$("#links").append(data);
				});
			});
		</script>
		<style type="text/css">
			#footer {
				height: 60px;
				background-color: #f5f5f5;
				margin-top: 20px;
			}
			.credit {
				margin: 20px 0;
				text-align: center;
			}
		</style>
    </head>
    <body style="padding-top: 50px;">
    	<div class="navbar navbar-inverse navbar-fixed-top">
			<div class="navbar-inner">
				<div class="container">
					<a class="navbar-brand">aqquadro</a>
				</div>
			</div>
		</div>
		<div class="container">
			<div class="row">
				<div class="col-lg-6">
					<h2><a href="http://aqquadro.it/?r=<?php echo $imageID ?>">.about() <i class="icon-refresh"></i></a></h2>
					<address>
						<strong>Alessandro Aglietti</strong><br />
						<abbr title="curriculum vitae"><a target="_blank" href="http://cv.alessandroaglietti.com">CV</a></abbr>
					</address>
					<div style="text-align:center;"><img src="https://googledrive.com/host/0B-udBnWnmH6JSGFEemdZbl9WX2c/<?php echo $imageID ?>" class="img-responsive img-thumbnail" alt="rand"></div>
				</div>
				<div class="col-lg-6" id="links">
					<h2><a target="_blank" href="http://feeds.delicious.com/v2/rss/aqquadro?count=50">.delicious() <i class="icon-rss"></i></a></h2>
					<p>loading...</p>
				</div>
			</div>
		</div>
		<div id="footer">
			<div class="container">
				<p class="text-muted credit">Fork on <a target="_blank" href="https://github.com/alessandro-aglietti/aqquadro.it">GitHub</a></p>
			</div>
		</div>
    </body>
</html>