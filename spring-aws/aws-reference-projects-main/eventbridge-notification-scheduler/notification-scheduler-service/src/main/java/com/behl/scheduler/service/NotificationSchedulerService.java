package com.behl.scheduler.service;

import java.util.List;
import java.util.UUID;

import org.apache.commons.lang3.RandomStringUtils;
import org.springframework.boot.context.properties.EnableConfigurationProperties;
import org.springframework.http.HttpStatus;
import org.springframework.stereotype.Service;
import org.springframework.web.server.ResponseStatusException;

import com.amazonaws.SdkClientException;
import com.amazonaws.services.eventbridge.AmazonEventBridge;
import com.amazonaws.services.eventbridge.model.DeleteRuleRequest;
import com.amazonaws.services.eventbridge.model.PutRuleRequest;
import com.amazonaws.services.eventbridge.model.PutTargetsRequest;
import com.amazonaws.services.eventbridge.model.RuleState;
import com.behl.scheduler.dto.ScheduleNotificationRequestDto;
import com.behl.scheduler.properties.EventBridgeConfigurationProperties;
import com.behl.scheduler.utility.CronExpressionGenerator;
import com.behl.scheduler.utility.JsonUtility;
import com.behl.scheduler.utility.TargetBuilder;

import lombok.RequiredArgsConstructor;
import lombok.extern.log4j.Log4j2;

@Log4j2
@Service
@RequiredArgsConstructor
@EnableConfigurationProperties(value = EventBridgeConfigurationProperties.class)
public class NotificationSchedulerService {

	private final AmazonEventBridge amazonEventBridge;
	private final TargetBuilder targetBuilder;
	private final JsonUtility jsonUtility;
	private final EventBridgeConfigurationProperties eventBridgeConfigurationProperties;

	public void schedule(final ScheduleNotificationRequestDto scheduleNotificationRequestDto) {
		final String ruleName = eventBridgeConfigurationProperties.getRulePrefix()
				+ RandomStringUtils.randomAlphabetic(10);
		final String cronExpression = CronExpressionGenerator.generate(scheduleNotificationRequestDto.getTimestamp());

		var putRuleRequest = new PutRuleRequest();
		putRuleRequest.setName(ruleName);
		putRuleRequest.setScheduleExpression(cronExpression);
		putRuleRequest.setState(RuleState.ENABLED.name());
		putRuleRequest.setEventBusName(eventBridgeConfigurationProperties.getEventBusName());

		createRule(putRuleRequest);

		final String targetId = UUID.randomUUID().toString();
		final String input = jsonUtility.toJson(ScheduleNotificationRequestDto.of(scheduleNotificationRequestDto,
				ruleName, eventBridgeConfigurationProperties.getEventBusName(), targetId));
		final var target = targetBuilder.build(targetId, input);

		var targetRequest = new PutTargetsRequest();
		targetRequest.setEventBusName(eventBridgeConfigurationProperties.getEventBusName());
		targetRequest.setRule(ruleName);
		targetRequest.setTargets(List.of(target));

		createTarget(targetRequest);
		log.info("Successfully attached Eventbridge target with rule {} and input {}", ruleName, input);
	}

	private void createRule(PutRuleRequest putRuleRequest) {
		try {
			var putRuleResult = amazonEventBridge.putRule(putRuleRequest);
			log.info("Successfully created rule with ARN : {}", putRuleResult.getRuleArn());
		} catch (final SdkClientException exception) {
			log.error("Exception occurred while creating EventBridge Rule", exception);
			throw new ResponseStatusException(HttpStatus.EXPECTATION_FAILED);
		}
	}

	private void createTarget(final PutTargetsRequest targetRequest) {
		try {
			amazonEventBridge.putTargets(targetRequest);
		} catch (final SdkClientException exception) {
			log.error("Exception occurred while creating EventBridge Target", exception);
			deleteRule(targetRequest.getRule());
			throw new ResponseStatusException(HttpStatus.EXPECTATION_FAILED);
		}
	}

	private void deleteRule(final String ruleName) {
		log.info("Attempting to delete rule {} after unsuccessfull target creation", ruleName);
		final var deleteRuleRequest = new DeleteRuleRequest();
		deleteRuleRequest.setEventBusName(eventBridgeConfigurationProperties.getEventBusName());
		deleteRuleRequest.setForce(true);
		deleteRuleRequest.setName(ruleName);
		try {
			amazonEventBridge.deleteRule(deleteRuleRequest);
		} catch (final SdkClientException exception) {
			log.error("Unable to delete EventBridge rule {}", ruleName, exception);
		}
		log.info("Successfully deleted EventBridge rule {}", ruleName);
	}

}