package com.behl.translator.dto;

import javax.validation.constraints.NotBlank;
import javax.validation.constraints.Size;

import lombok.Builder;
import lombok.Getter;
import lombok.extern.jackson.Jacksonized;

@Getter
@Builder
@Jacksonized
public class TranslationRequestDto {

	@NotBlank
	@Size(max = 500)
	private final String text;

	@NotBlank
	private final String sourceLanguageCode;

	@NotBlank
	private final String targetLanguageCode;

}