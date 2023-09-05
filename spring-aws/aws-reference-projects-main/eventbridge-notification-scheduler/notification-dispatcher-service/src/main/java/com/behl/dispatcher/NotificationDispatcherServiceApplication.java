package com.behl.dispatcher;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cloud.aws.autoconfigure.context.ContextRegionProviderAutoConfiguration;
import org.springframework.cloud.aws.autoconfigure.context.ContextStackAutoConfiguration;

@SpringBootApplication(exclude = { ContextRegionProviderAutoConfiguration.class, ContextStackAutoConfiguration.class })
public class NotificationDispatcherServiceApplication {

	public static void main(String[] args) {
		SpringApplication.run(NotificationDispatcherServiceApplication.class, args);
	}

}
