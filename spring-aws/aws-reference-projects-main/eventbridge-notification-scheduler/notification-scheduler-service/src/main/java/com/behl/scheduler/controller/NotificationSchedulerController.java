package com.behl.scheduler.controller;

import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.ResponseStatus;
import org.springframework.web.bind.annotation.RestController;

import com.behl.scheduler.dto.ScheduleNotificationRequestDto;
import com.behl.scheduler.service.NotificationSchedulerService;

import lombok.RequiredArgsConstructor;

@RestController
@RequiredArgsConstructor
public class NotificationSchedulerController {

	private final NotificationSchedulerService notificationSchedulerService;

	@PostMapping(value = "/v1/schedule", consumes = MediaType.APPLICATION_JSON_VALUE)
	@ResponseStatus(value = HttpStatus.OK)
	public ResponseEntity<HttpStatus> schedule(
			@RequestBody(required = true) final ScheduleNotificationRequestDto scheduleNotificationRequestDto) {
		notificationSchedulerService.schedule(scheduleNotificationRequestDto);
		return ResponseEntity.ok().build();
	}

}