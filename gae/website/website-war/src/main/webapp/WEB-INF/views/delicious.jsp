<%@ page contentType="text/html;charset=UTF-8" language="java" %>
<%@ taglib prefix="fn" uri="http://java.sun.com/jsp/jstl/functions" %>


<?php
	header("Expires: Thu, 04 May 1989 14:00:00 GMT");
	header("Last-Modified: " . gmdate("D, d M Y H:i:s") . " GMT");
	header("Cache-Control: no-store, no-cache, must-revalidate");
	header("Cache-Control: post-check=0, pre-check=0", false);
	header("Pragma: no-cache");
	
	$ch = curl_init("http://feeds.delicious.com/v2/json/aqquadro?count=10");
	curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);

	$response = curl_exec($ch);
	curl_close($ch);
	if ( $response !== FALSE) {
		header('Content-type: text/html');
		
		require 'Mustache/Autoloader.php';
		Mustache_Autoloader::register();

		$m = new Mustache_Engine;
		
		$headers = getallheaders();
		
		if( strrpos(strtolower($headers["User-Agent"]), 'lynx') !== FALSE ) {
			
			$echoes = $m->render('{{#links}}<p>{{{n}}}<br /><a target="_blank" href="{{{u}}}">{{{d}}}</a></p>{{/links}}', array('links' => json_decode($response)));

			$echoes = $echoes . '<p>more del.icio.us<br /><a target="_blank" href="https://delicious.com/aqquadro">aqquadro @ del.icio.us</a></p>';

		} else {
		
			$echoes = $m->render('{{#links}}<blockquote><p>{{{n}}}</p><small><a target="_blank" href="{{{u}}}">{{{d}}}</a></small></blockquote>{{/links}}', array('links' => json_decode($response)));

			$echoes = $echoes . '<blockquote><p>more del.icio.us</p><small><a target="_blank" href="https://delicious.com/aqquadro">aqquadro @ del.icio.us</a></small></blockquote>';		
		}
		
		echo $echoes;
	} else {
		header($_SERVER['SERVER_PROTOCOL'] . ' 500 Internal Server Error', true, 500);
	}
?>