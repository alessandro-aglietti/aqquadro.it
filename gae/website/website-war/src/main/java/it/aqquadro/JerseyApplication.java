package it.aqquadro;

import it.aqquadro.webapp.IndexResource;

import javax.ws.rs.ApplicationPath;

import org.glassfish.jersey.filter.LoggingFilter;
import org.glassfish.jersey.server.ResourceConfig;
import org.glassfish.jersey.server.mvc.beanvalidation.MvcBeanValidationFeature;
import org.glassfish.jersey.server.mvc.mustache.MustacheMvcFeature;

@ApplicationPath("/")
public class JerseyApplication extends ResourceConfig {

    public JerseyApplication() {
        // Resources.
        packages(IndexResource.class.getPackage().getName());

        // Features.
        register(MvcBeanValidationFeature.class);

        // Providers.
        register(LoggingFilter.class);
        register(MustacheMvcFeature.class);

        // Properties.
        property(MustacheMvcFeature.TEMPLATE_BASE_PATH, "/mustache");
    }
}
