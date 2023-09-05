package com.behl.scheduler.utility;

import org.springframework.http.HttpStatus;
import org.springframework.stereotype.Component;
import org.springframework.web.server.ResponseStatusException;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;

import lombok.RequiredArgsConstructor;
import lombok.extern.log4j.Log4j2;

@Log4j2
@Component
@RequiredArgsConstructor
public class JsonUtility {

	private final ObjectMapper objectMapper;

	public String toJson(final Object object) {
		try {
			return objectMapper.writeValueAsString(object);
		} catch (final JsonProcessingException exception) {
			log.error("Exception occurred while converting object to JSON String", exception);
			throw new ResponseStatusException(HttpStatus.NOT_IMPLEMENTED);
		}
	}

}
