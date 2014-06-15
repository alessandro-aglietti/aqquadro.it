<?php

/**
 * This class defines an example resource that is wired into the URI /example
 * @uri /
 */
class ExampleResource extends Tonic\Resource {

    /**
     * @method GET
     */
    function exampleMethod() {
		$sleepThis = rand(0,3000);
		usleep($sleepThis * 1000);
		
		$model = array(
			'titolo1' => 'Titolo 1',
			'titolo3' => 'Sleep for ' . $sleepThis . 'ms'
		);
		
        return new Tonic\Response(Tonic\Response::OK, $this->container["twig"]->render('hello.html', $model));
    }

}

?>