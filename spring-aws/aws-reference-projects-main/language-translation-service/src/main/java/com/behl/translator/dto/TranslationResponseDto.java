package com.behl.translator.dto;

import javax.validation.constraints.NotBlank;

import lombok.Builder;
import lombok.Getter;
import lombok.extern.jackson.Jacksonized;

@Getter
@Builder
@Jacksonized
public class TranslationResponseDto {

	@NotBlank
	private final String translatedText;

	@NotBlank
	private final String originalText;

	@NotBlank
	private final String sourceLanguageCode;

	@NotBlank
	private final String targetLanguageCode;

}