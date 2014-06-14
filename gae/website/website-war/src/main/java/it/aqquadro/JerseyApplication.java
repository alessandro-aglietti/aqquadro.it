package it.aqquadro;

import it.aqquadro.webapp.IndexResource;

import javax.ws.rs.ApplicationPath;

import org.glassfish.jersey.CommonProperties;
import org.glassfish.jersey.filter.LoggingFilter;
import org.glassfish.jersey.server.ResourceConfig;
import org.glassfish.jersey.server.ServerProperties;
import org.glassfish.jersey.server.mvc.beanvalidation.MvcBeanValidationFeature;
import org.glassfish.jersey.server.mvc.mustache.MustacheMvcFeature;

@ApplicationPath("/")
public class JerseyApplication extends ResourceConfig {

    public JerseyApplication() {
        // Resources.
//        packages(IndexResource.class.getPackage().getName());
    	register(IndexResource.class);

        // Features.
        register(MvcBeanValidationFeature.class);

        // Providers.
        register(LoggingFilter.class);
        register(MustacheMvcFeature.class);

        // Properties.
        property(MustacheMvcFeature.TEMPLATE_BASE_PATH, "/mustache");
        
        // https://jersey.java.net/documentation/latest/deployment.html#deployment.autodiscovery.config
        property(CommonProperties.FEATURE_AUTO_DISCOVERY_DISABLE, "false");
        property(CommonProperties.JSON_PROCESSING_FEATURE_DISABLE, "false");
        property(CommonProperties.MOXY_JSON_FEATURE_DISABLE, "false");
        
        // https://jersey.java.net/documentation/latest/deployment.html#deployment.classpath-scanning
//        register(org.glassfish.jersey.server.filter.UriConnegFilter.class);
//        register(org.glassfish.jersey.server.validation.ValidationFeature.class);
//        register(org.glassfish.jersey.server.spring.SpringComponentProvider.class);
//        register(org.glassfish.jersey.grizzly2.httpserver.GrizzlyHttpContainerProvider.class);
//        property(ServerProperties.METAINF_SERVICES_LOOKUP_DISABLE, true);
        
    }
}
