<?php
	$headers = getallheaders();
	if( strrpos(strtolower($headers["User-Agent"]), 'lynx') !== FALSE ) {
		header("HTTP/1.1 301 Moved Permanently");
		header("Location: http://aqquadro.it/text.php");
		exit();
	}
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
		<meta http-equiv="content-type" content="text/html; charset=utf-8" />
		<meta name="description" content="Alessandro Aglietti, from Firenze, is aqquadro: since 2004 face-to-face to a monitor. Not to write poems! Be RSS: ban social network from life for life." />
		<title>aqquadro rovers</title>
		<script src="//code.jquery.com/jquery-1.10.1.min.js"></script>
		<script src="//code.jquery.com/jquery-migrate-1.2.1.min.js"></script>
		<link rel="stylesheet" href="//netdna.bootstrapcdn.com/bootstrap/3.0.0-rc1/css/bootstrap.min.css">
		<link href="//netdna.bootstrapcdn.com/font-awesome/3.2.1/css/font-awesome.min.css" rel="stylesheet">
		<script src="//netdna.bootstrapcdn.com/bootstrap/3.0.0-rc1/js/bootstrap.min.js"></script>
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
			
			#footer div.container {
				background-image: url('http://www.kopimi.com/kopimi/k/mkopimi_gay.gif');
				background-repeat: no-repeat;
				background-position-y: 6px;
			}
		</style>
		<script type="text/javascript">

		  var _gaq = _gaq || [];
		  _gaq.push(['_setAccount', 'UA-11408016-2']);
		  _gaq.push(['_trackPageview']);

		  (function() {
		    var ga = document.createElement('script'); ga.type = 'text/javascript'; ga.async = true;
		    ga.src = ('https:' == document.location.protocol ? 'https://ssl' : 'http://www') + '.google-analytics.com/ga.js';
		    var s = document.getElementsByTagName('script')[0]; s.parentNode.insertBefore(ga, s);
		  })();

		</script>
    </head>
    <body style="padding-top: 50px;">
    	<div class="navbar navbar-inverse navbar-fixed-top">
			<div class="navbar-inner">
				<div class="container">
					<button type="button" class="navbar-toggle" data-toggle="collapse" data-target=".nav-collapse">
						<span class="icon-bar"></span>
						<span class="icon-bar"></span>
						<span class="icon-bar"></span>
					</button>
					<a class="navbar-brand" href="http://aqquadro.it/alessandro.aglietti/">aqquadro</a>
					<div class="nav-collapse collapse">
						<ul class="nav navbar-nav">
							<li class="active"><a href="http://aqquadro.it/rovers/">rovers</a></li>
							<li><a href="http://cloud.feedly.com/v3/opml?feedlyToken=AQAAIH17InYiOiJwcm9kdWN0aW9uIiwicCI6MSwiYSI6ImZlZWRseSIsInQiOjEsIngiOiJzdGFuZGFyZCIsImkiOiIxZTBmZDBjMS1hNmFiLTQ4YTEtOGRjNi0wNWFiMDQ3MjQ5YTIiLCJlIjoxMzgxODQwNTAxMzA1fQ">my rss</a></li>
							<li><a href="http://aqquadro.it/wall">wall</a></li>
						</ul>
					</div>
				</div>
			</div>
		</div>
		<div class="container">
			<div class="row">
				<div class="col-lg-4">
					<div style="text-align:center;">
						<h2><a target="_blank" href="http://it.wikipedia.org/wiki/Aaron_Swartz">Aaron Swartz</a></h2>
						<img src="https://googledrive.com/host/0B-udBnWnmH6JSGFEemdZbl9WX2c/98" class="img-responsive img-thumbnail" alt="rand">
					</div>
				</div>
				<div class="col-lg-4">
					<div style="text-align:center;">
						<h2><a target="_blank" href="http://it.wikipedia.org/wiki/Fabrizio_De_Andr%C3%A9">Fabrizio De Andr&eacute;</a></h2>
						<img src="https://googledrive.com/host/0B-udBnWnmH6JSGFEemdZbl9WX2c/99" class="img-responsive img-thumbnail" alt="rand">
					</div>
				</div>
				<div class="col-lg-4">
					<div style="text-align:center;">
						<h2><a target="_blank" href="http://it.wikipedia.org/wiki/Daft_Punk">Daft Punk</a></h2>
						<img src="https://googledrive.com/host/0B-udBnWnmH6JSGFEemdZbl9WX2c/100" class="img-responsive img-thumbnail" alt="rand">
					</div>
				</div>
			</div>
			<div class="row">
				<div class="col-lg-4">
					<div style="text-align:center;">
						<h2><a target="_blank" href="http://it.wikipedia.org/wiki/Giorgio_Ambrosoli">Giorgio Ambrosoli</a></h2>
						<img src="https://googledrive.com/host/0B-udBnWnmH6JSGFEemdZbl9WX2c/104" class="img-responsive img-thumbnail" alt="rand">
					</div>
				</div>
				<div class="col-lg-4">
					<div style="text-align:center;">
						<h2><a target="_blank" href="http://it.wikipedia.org/wiki/Salvador_Allende">Salvador Allende</a></h2>
						<img src="https://googledrive.com/host/0B-udBnWnmH6JSGFEemdZbl9WX2c/110" class="img-responsive img-thumbnail" alt="rand">
					</div>
				</div>
				<div class="col-lg-4">
					<div style="text-align:center;">
						<h2><a target="_blank" href="http://memoriaaudiovisual.cl/~interac/victimas/?p=1028">Lagos Rios Oscar Reinaldo</a></h2>
						<img src="https://googledrive.com/host/0B-udBnWnmH6JSGFEemdZbl9WX2c/109" class="img-responsive img-thumbnail" alt="rand">
					</div>
				</div>
			</div>
			<div class="row">
				<div class="col-lg-4">
					<div style="text-align:center;">
						<h2><a target="_blank" href="http://it.wikipedia.org/wiki/Luis_Sep%C3%BAlveda">Luis Sep&uacute;lveda</a></h2>
						<img src="https://googledrive.com/host/0B-udBnWnmH6JSGFEemdZbl9WX2c/111" class="img-responsive img-thumbnail" alt="rand">
					</div>
				</div>
				<div class="col-lg-4"></div>
				<div class="col-lg-4"></div>
			</div>
		</div>
		<div id="footer">
			<div class="container">
				<p class="text-muted credit"><abbr title="http://www.kopimi.com/kopimi/">kopimi</abbr> from <a target="_blank" href="https://github.com/alessandro-aglietti/aqquadro.it">GitHub</a></p>
			</div>
		</div>
    </body>
</html>