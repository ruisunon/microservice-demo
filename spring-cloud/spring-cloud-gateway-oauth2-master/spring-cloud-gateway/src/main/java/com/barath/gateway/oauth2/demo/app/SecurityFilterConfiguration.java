package com.barath.gateway.oauth2.demo.app;

import java.lang.invoke.MethodHandle;
import java.lang.invoke.MethodHandles;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.cloud.gateway.filter.GlobalFilter;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.core.annotation.Order;
import org.springframework.http.HttpHeaders;
import org.springframework.web.reactive.function.client.WebClient;

import reactor.core.publisher.Mono;

@Configuration
public class SecurityFilterConfiguration {
	
	private static final Logger logger = LoggerFactory.getLogger(MethodHandles.lookup().lookupClass());
	
	@Bean
	@Order(-1)
	public GlobalFilter tokenValidationFilter(WebClient webClient) {
	    return (exchange, chain) -> {
	    	logger.info("token validation filter");
	    	HttpHeaders headers = exchange.getRequest().getHeaders();
	         webClient.post()
	            .header("Authorization", headers.get("Authorization").get(0))
	            .retrieve()
	            .bodyToMono(Object.class)
	            .log()
	            .onErrorReturn(Mono.empty());
	         return chain.filter(exchange); 
	      
	    };
	}

}
