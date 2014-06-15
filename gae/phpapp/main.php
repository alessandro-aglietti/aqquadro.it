<?php
require 'vendor/autoload.php';

$loader = new Twig_Loader_Filesystem('./');
$twig = new Twig_Environment($loader);

$container = array(
	"twig" => $twig
);

$app = new Tonic\Application(array(
    'load' => 'example.php'
));
$request = new Tonic\Request();

$resource = $app->getResource($request);

$resource->container = $container;

$response = $resource->exec();
$response->output();

?>