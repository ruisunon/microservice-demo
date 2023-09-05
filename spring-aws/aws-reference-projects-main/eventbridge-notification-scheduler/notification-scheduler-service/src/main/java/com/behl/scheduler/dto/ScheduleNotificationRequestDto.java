package com.behl.scheduler.dto;

import java.time.LocalDateTime;

import lombok.Builder;
import lombok.Getter;
import lombok.extern.jackson.Jacksonized;

@Getter
@Builder
@Jacksonized
public class ScheduleNotificationRequestDto {

	private final LocalDateTime timestamp;
	private final String emailId;
	private final String subject;
	private final String body;
	private final String ruleName;
	private final String eventBusName;
	private final String targetId;

	public static ScheduleNotificationRequestDto of(final ScheduleNotificationRequestDto scheduleNotificationRequestDto,
			final String ruleName, final String eventBusName, final String targetId) {
		return ScheduleNotificationRequestDto.builder().timestamp(scheduleNotificationRequestDto.getTimestamp())
				.emailId(scheduleNotificationRequestDto.getEmailId())
				.subject(scheduleNotificationRequestDto.getSubject()).body(scheduleNotificationRequestDto.getBody())
				.ruleName(ruleName).eventBusName(eventBusName).targetId(targetId).build();
	}

}