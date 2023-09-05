package com.barath.gateway.oauth2.demo.app;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.web.reactive.function.client.WebClient;

@Configuration
public class WebClientConfiguration {
	
	@Bean
	public WebClient webClient() {
		return WebClient.create("http://localhost:8080/auth/realms/master/protocol/openid-connect/userinfo");
	}

}
