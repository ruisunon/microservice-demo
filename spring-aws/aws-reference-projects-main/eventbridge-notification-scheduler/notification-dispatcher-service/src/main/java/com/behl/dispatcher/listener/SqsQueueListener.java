package com.behl.dispatcher.listener;

import org.springframework.cloud.aws.messaging.listener.SqsMessageDeletionPolicy;
import org.springframework.cloud.aws.messaging.listener.annotation.SqsListener;
import org.springframework.stereotype.Component;

import com.behl.dispatcher.dto.NotificationDetailDto;
import com.behl.dispatcher.service.EventBridgeService;
import com.behl.dispatcher.utility.JsonUtility;

import lombok.RequiredArgsConstructor;
import lombok.extern.log4j.Log4j2;

@Log4j2
@Component
@RequiredArgsConstructor
public class SqsQueueListener {

	private final EventBridgeService eventBridgeService;

	@SqsListener(value = "${com.behl.aws.sqs.name}", deletionPolicy = SqsMessageDeletionPolicy.ON_SUCCESS)
	public void consumer(final String message) {
		log.info("Input received from SQS Queue {}", message);
		NotificationDetailDto notificationDetailDto = JsonUtility.parseJson(message, NotificationDetailDto.class);
		eventBridgeService.deleteRule(notificationDetailDto.getRuleName(), notificationDetailDto.getEventBusName(),
				notificationDetailDto.getTargetId());
	}

}