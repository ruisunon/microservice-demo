package com.rollingstone.filters;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.cloud.gateway.filter.GlobalFilter;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.http.HttpHeaders;
import reactor.core.publisher.Mono;

@Configuration
public class RollingStoneResponseFilter {
 
    final Logger logger =LoggerFactory.getLogger(RollingStoneResponseFilter.class);
    
    @Autowired
	FilterUtils rollingStoneFilterUtils;
 
    @Bean
    public GlobalFilter postGlobalFilter() {
        return (exchange, chain) -> {
            return chain.filter(exchange).then(Mono.fromRunnable(() -> {
            	  HttpHeaders requestHeaders = exchange.getRequest().getHeaders();
            	  String rollingStoneCorrelationId = rollingStoneFilterUtils.getRollingStoneCorrelationId(requestHeaders);
            	  logger.debug("Adding the rollingstone correlation id to the egress headers. {}", rollingStoneCorrelationId);
                  exchange.getResponse().getHeaders().add(FilterUtils.ROLLINGSTONE_CORRELATION_ID, rollingStoneCorrelationId);
                  logger.debug("Completing egress request for {}.", exchange.getRequest().getURI());
              }));
        };
    }
}
