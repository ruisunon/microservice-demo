package com.behl.dispatcher.dto;

import java.time.LocalDateTime;

import lombok.Builder;
import lombok.Getter;
import lombok.ToString;
import lombok.extern.jackson.Jacksonized;

@Getter
@Builder
@Jacksonized
@ToString
public class NotificationDetailDto {

	private final LocalDateTime timestamp;
	private final String emailId;
	private final String subject;
	private final String body;
	private final String ruleName;
	private final String eventBusName;
	private final String targetId;

}