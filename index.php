<?php
	$headers = getallheaders();
	if( strrpos(strtolower($headers["User-Agent"]), 'lynx') !== FALSE ) {
		header("HTTP/1.1 301 Moved Permanently");
		header("Location: text.php");
		exit();
	}
	$imageID = rand ( 0, 165 );
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
		<title>Alessandro Aglietti</title>
		<link rel="shortcut icon" type="image/ico" href="http://piratbyran.org/favicon.ico" />
		<script src="//code.jquery.com/jquery-1.10.1.min.js"></script>
		<script src="//code.jquery.com/jquery-migrate-1.2.1.min.js"></script>
		<link rel="stylesheet" href="//netdna.bootstrapcdn.com/bootstrap/3.0.0-rc1/css/bootstrap.min.css">
		<link href="//netdna.bootstrapcdn.com/font-awesome/3.2.1/css/font-awesome.min.css" rel="stylesheet">
		<script src="//netdna.bootstrapcdn.com/bootstrap/3.0.0-rc1/js/bootstrap.min.js"></script>
        <script type="text/javascript">
			$(function(){
				$.ajax({
					url: "http://aqquadro.it/deliciuos.php?r=<?php echo $imageID ?>"
				}).done(function ( data ) {
					$("#links").children().last().remove();
					$("#links").append(data);
				});
				
				$("#clickthis").click(function(eventObject){
					var randd = Math.floor((Math.random() * 165));
					
					$(".progress").removeClass("hide").addClass("in");
					var memfucker = window.setInterval(function(){
						var rand = Math.floor((Math.random() * 98));
						$(".progress").css("width", rand + "%");
					}, 750);
					
					
					$("#tochangeit").attr("src", "https://googledrive.com/host/0B-udBnWnmH6JSGFEemdZbl9WX2c/" + randd).load(function(){
						clearInterval(memfucker);
						$(".progress").addClass("hide").removeClass("in");
					});
					$(".alert").removeClass("hide").addClass('in');
					return false;
				})
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
			
			#footer div.container {
				background-image: url('http://aqquadro.it/kopimi-gay.png');
				background-repeat: no-repeat;
				background-position-y: 6px;
				background-size: 100px;
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
							<li><a href="http://aqquadro.it/rovers/">rovers</a></li>
							<li><a href="http://aqquadro.it/aqquadro-rss-subscriptions.xml">my rss</a></li>
							<li><a href="http://aqquadro.it/wall">wall</a></li>
						</ul>
					</div>
				</div>
			</div>
		</div>
		<div class="container">
			<div class="row">
				<div class="col-lg-6">
					<h2><a id="clickthis" href="http://aqquadro.it/alessandro.aglietti/?r=<?php echo $imageID ?>">.about() <i class="icon-refresh"></i></a></h2>
					<address>
						<strong>Alessandro Aglietti</strong>, from Firenze, <strong>is aqquadro</strong>: since 2004 <strong>face-to-face to a monitor</strong>. Not to write poems! Be RSS: ban social network from life for life.<br />
						<abbr title="curriculum vitae"><a target="_blank" href="http://cv.alessandroaglietti.com">CV</a></abbr>
						<br />
						<abbr title="dal database al web"><a target="_blank" href="https://github.com/alessandro-aglietti/itis-leonardo-da-vinci">Dal database al web</a></abbr>
					</address>
					<p style="text-align: center;">(a-b) + ((a/b)(b/a)(c-d))</p>
					<iframe width="100%" height="166" scrolling="no" frameborder="no" src="https://w.soundcloud.com/player/?url=https%3A//api.soundcloud.com/playlists/27796585&amp;color=ff5500&amp;auto_play=true&amp;hide_related=false&amp;show_artwork=true"></iframe>
					<br />
					<br />
					<p style="text-align:center;"><abbr title="Benjamin Franklin"><i>Chi e' pronto a rinunciare alle proprie liberta' fondamentali per comprarsi briciole di temporanea sicurezza non merita ne' la liberta' ne' la sicurezza</i></abbr></p>
					<div style="text-align:center;">
						<div class="progress progress-striped progress-info active hide fade">
						  <div class="bar" style="width: 40%;"></div>
						</div>
						<img id="tochangeit" src="https://googledrive.com/host/0B-udBnWnmH6JSGFEemdZbl9WX2c/<?php echo $imageID ?>" class="img-responsive img-thumbnail" alt="rand"></div>
				</div>
				<div class="col-lg-6" id="links">
					<h2><a target="_blank" href="http://feeds.delicious.com/v2/rss/aqquadro?count=50">.delicious() <i class="icon-rss"></i></a></h2>
					<p>loading...</p>
				</div>
			</div>
		</div>
		<div id="footer">
			<div class="container">
				<p class="text-muted credit"><abbr title="https://en.wikipedia.org/wiki/Piratbyr%C3%A5n#Kopimi">kopimi</abbr> from <a target="_blank" href="https://github.com/alessandro-aglietti/aqquadro.it">GitHub</a></p>
			</div>
		</div>
		<div style="position: absolute; top: 22px; left: 145px; z-index: 9999;" class="alert alert-info fade hide"><button type="button" class="close" data-dismiss="alert">×</button><strong>Holy guacamole!</strong> Clickjacking to show random pics!</div>
    </body>
</html>