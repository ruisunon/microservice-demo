package com.behl.scheduler.utility;

import org.springframework.boot.context.properties.EnableConfigurationProperties;
import org.springframework.stereotype.Component;

import com.amazonaws.services.eventbridge.model.Target;
import com.behl.scheduler.properties.EventBridgeConfigurationProperties;

import lombok.RequiredArgsConstructor;

@Component
@RequiredArgsConstructor
@EnableConfigurationProperties(value = EventBridgeConfigurationProperties.class)
public class TargetBuilder {

	private final EventBridgeConfigurationProperties eventBridgeConfigurationProperties;

	/**
	 * Returns an AWS Target object to be associated with an EventBridge Rule
	 * 
	 * @param targetId: unique identifier of the created target
	 * @param input:    JSON string to be sent to the created target
	 */
	public Target build(final String targetId, final String input) {
		return new Target().withArn(eventBridgeConfigurationProperties.getTargetArn()).withId(targetId)
				.withInput(input);
	}

}