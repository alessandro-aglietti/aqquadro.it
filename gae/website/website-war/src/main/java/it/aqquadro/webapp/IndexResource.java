package it.aqquadro.webapp;

import java.util.HashMap;
import java.util.Map;
import java.util.logging.Logger;

import javax.validation.constraints.NotNull;
import javax.ws.rs.GET;
import javax.ws.rs.Path;
import javax.ws.rs.Produces;
import javax.ws.rs.core.Context;
import javax.ws.rs.core.UriInfo;

import org.glassfish.jersey.server.mvc.Template;

@Path("/")
public class IndexResource {

	private final static Logger LOGGER = Logger.getLogger(IndexResource.class
			.getName());

	@NotNull
	@Context
	private UriInfo uriInfo;

	@GET
	@Produces("text/html")
	@Template(name = "/index")
	public Map<String, Object> get() {
		LOGGER.info("IndexResource.get()");

		Map<String, Object> model = new HashMap<String, Object>();

		model.put("msg", "Great!");

		return model;
	}
}
