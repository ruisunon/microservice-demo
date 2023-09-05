package com.behl.voicemaker.dto;

import lombok.Builder;
import lombok.Getter;
import lombok.extern.jackson.Jacksonized;

@Getter
@Builder
@Jacksonized
public class TextToSpeechConversionRequestDto {

	/**
	 * The textual content which is required to be processed and converted into an
	 * Audio speech
	 */
	private final String text;

}
