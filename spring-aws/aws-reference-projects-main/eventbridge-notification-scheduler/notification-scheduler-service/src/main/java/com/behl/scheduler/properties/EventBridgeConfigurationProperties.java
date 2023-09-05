package com.behl.scheduler.properties;

import org.springframework.boot.context.properties.ConfigurationProperties;

import lombok.Data;

@Data
@ConfigurationProperties(prefix = "com.behl.scheduler.aws.eventbridge")
public class EventBridgeConfigurationProperties {

	private String targetArn;
	private String eventBusName;
	private String rulePrefix;

}