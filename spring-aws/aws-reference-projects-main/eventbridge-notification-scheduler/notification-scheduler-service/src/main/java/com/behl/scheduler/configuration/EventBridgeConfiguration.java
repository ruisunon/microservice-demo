package com.behl.scheduler.configuration;

import java.util.UUID;

import org.springframework.boot.context.properties.EnableConfigurationProperties;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.context.annotation.Primary;

import com.amazonaws.auth.AWSStaticCredentialsProvider;
import com.amazonaws.auth.BasicAWSCredentials;
import com.amazonaws.regions.Regions;
import com.amazonaws.services.eventbridge.AmazonEventBridge;
import com.amazonaws.services.eventbridge.AmazonEventBridgeClientBuilder;
import com.amazonaws.services.eventbridge.model.Target;
import com.behl.scheduler.properties.AwsIAMConfigurationProperties;
import com.behl.scheduler.properties.EventBridgeConfigurationProperties;

import lombok.RequiredArgsConstructor;

@Configuration
@RequiredArgsConstructor
@EnableConfigurationProperties(value = { EventBridgeConfigurationProperties.class,
		AwsIAMConfigurationProperties.class })
public class EventBridgeConfiguration {

	private final EventBridgeConfigurationProperties eventBridgeConfigurationProperties;
	private final AwsIAMConfigurationProperties awsIAMConfigurationProperties;

	@Bean
	@Primary
	public AmazonEventBridge amazonEventBridgeClient() {
		BasicAWSCredentials credentials = new BasicAWSCredentials(awsIAMConfigurationProperties.getAccessKey(),
				awsIAMConfigurationProperties.getSecretAccessKey());
		return AmazonEventBridgeClientBuilder.standard().withRegion(Regions.AP_SOUTH_1)
				.withCredentials(new AWSStaticCredentialsProvider(credentials)).build();
	}

	@Bean
	@Primary
	public Target target() {
		return new Target().withArn(eventBridgeConfigurationProperties.getTargetArn())
				.withId(UUID.randomUUID().toString());
	}

}