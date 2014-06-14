package it.aqquadro.webapp;

import java.util.HashMap;
import java.util.Map;
import java.util.Random;
import java.util.logging.Logger;

import javax.validation.constraints.NotNull;
import javax.ws.rs.GET;
import javax.ws.rs.Path;
import javax.ws.rs.Produces;
import javax.ws.rs.core.Context;
import javax.ws.rs.core.UriInfo;

import org.glassfish.jersey.server.mvc.Template;

@Path("/noscript")
public class IndexResource {

	private final static Logger LOGGER = Logger.getLogger(IndexResource.class
			.getName());

	@NotNull
	@Context
	private UriInfo uriInfo;

	@GET
	@Produces("text/html")
	@Template(name = "/index")
	public Map<String, Object> get() throws InterruptedException {
		LOGGER.info("IndexResource.get()");

		Map<String, Object> model = new HashMap<String, Object>();

		int sleepThisMillisec = new Random().nextInt(3000);
		
		Thread.sleep(sleepThisMillisec);

		model.put("msg", "Great! " + new Random().nextGaussian());
		model.put("sleep", "Sleep for " + sleepThisMillisec + "ms");

		return model;
	}
}
