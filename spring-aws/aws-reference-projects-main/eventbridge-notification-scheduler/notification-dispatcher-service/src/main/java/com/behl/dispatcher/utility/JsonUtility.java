package com.behl.dispatcher.utility;

import org.springframework.http.HttpStatus;
import org.springframework.web.server.ResponseStatusException;

import com.behl.dispatcher.dto.NotificationDetailDto;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.datatype.jsr310.JavaTimeModule;

import lombok.extern.log4j.Log4j2;

@Log4j2
public class JsonUtility {

	private JsonUtility() {
	}

	public static NotificationDetailDto parseJson(final String message, final Class<NotificationDetailDto> target) {
		NotificationDetailDto notificationDetailDto;
		try {
			notificationDetailDto = new ObjectMapper().registerModule(new JavaTimeModule()).readValue(message,
					NotificationDetailDto.class);
		} catch (JsonProcessingException exception) {
			log.error("Unable to read notification message", exception);
			throw new ResponseStatusException(HttpStatus.INTERNAL_SERVER_ERROR);
		}
		return notificationDetailDto;
	}

}
