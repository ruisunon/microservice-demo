package com.behl.dispatcher.service;

import org.springframework.stereotype.Service;

import com.amazonaws.SdkClientException;
import com.amazonaws.services.eventbridge.AmazonEventBridge;
import com.amazonaws.services.eventbridge.model.DeleteRuleRequest;
import com.amazonaws.services.eventbridge.model.DisableRuleRequest;
import com.amazonaws.services.eventbridge.model.RemoveTargetsRequest;

import lombok.RequiredArgsConstructor;
import lombok.extern.log4j.Log4j2;

@Log4j2
@Service
@RequiredArgsConstructor
public class EventBridgeService {

	private final AmazonEventBridge amazonEventBridge;

	public void deleteRule(final String ruleName, final String eventBusName, final String targetId) {
		final var disableRuleRequest = new DisableRuleRequest().withEventBusName(eventBusName).withName(ruleName);

		try {
			amazonEventBridge.disableRule(disableRuleRequest);
		} catch (SdkClientException exception) {
			log.error("Exception occurred while attempting to disable rule {} in event-bus {}", ruleName, eventBusName);
			throw new RuntimeException("Unable to disable EventBridge Rule", exception);
		}
		log.info("Successfully disabled EventBridge rule {} in event-bus {}", ruleName, eventBusName);

		var removeTargetsRequest = new RemoveTargetsRequest().withEventBusName(eventBusName).withForce(true)
				.withIds(targetId).withRule(ruleName);

		try {
			amazonEventBridge.removeTargets(removeTargetsRequest);
		} catch (SdkClientException exception) {
			log.error("Exception occurred while attempting to delete target {} associated with rule {}", targetId,
					ruleName, exception);
			throw new RuntimeException("Unable to delete target associated with EventBridge Rule", exception);
		}
		log.info("Successfully deleted target {} associated with EventBridge rule {} in event-bus {}", targetId,
				ruleName, eventBusName);

		var deleteRuleRequest = new DeleteRuleRequest().withEventBusName(eventBusName).withForce(true)
				.withName(ruleName);

		try {
			amazonEventBridge.deleteRule(deleteRuleRequest);
		} catch (SdkClientException exception) {
			log.error("Exception occurred while attempting to delete rule {} in event-bus {}", ruleName, eventBusName);
			throw new RuntimeException("Unable to delete EventBridge Rule", exception);
		}
		log.info("Successfully deleted EventBridge rule {} in event-bus {}", ruleName, eventBusName);
	}

}