package com.behl.translator.controller;

import java.util.List;

import javax.validation.Valid;

import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseStatus;
import org.springframework.web.bind.annotation.RestController;

import com.amazonaws.services.translate.model.Language;
import com.behl.translator.dto.TranslationRequestDto;
import com.behl.translator.dto.TranslationResponseDto;
import com.behl.translator.service.LanguageTranslationService;

import lombok.RequiredArgsConstructor;

@RestController
@RequiredArgsConstructor
@RequestMapping(value = "/translate")
public class LanguageTranslationController {

	private final LanguageTranslationService languageTranslationService;

	/**
	 * Exposed API Endpoint that returns the list of all available/supported
	 * {@link com.amazonaws.services.translate.model.Language} that can be used for
	 * language translation when calling endpoint API method
	 * {@link #translateText(TranslationRequestDto)}
	 * 
	 * @return List of {@link com.amazonaws.services.translate.model.Language}
	 */
	@GetMapping(value = "/language", produces = MediaType.APPLICATION_JSON_VALUE)
	@ResponseStatus(value = HttpStatus.OK)
	public ResponseEntity<List<Language>> retrieveSupportedLanguages() {
		return ResponseEntity.ok(languageTranslationService.getSupportedLanguages());
	}

	/**
	 * 
	 * Method to translate texts in a source language to a specified target
	 * language. The language codes being specified in source and target can be
	 * retrieved from {@link #retrieveSupportedLanguages()}
	 * 
	 * @param translationRequestDto {@link com.behl.translator.dto.TranslationRequestDto}
	 *                              in API request-body
	 * @return {@link com.behl.translator.dto.TranslationResponseDto} containing the
	 *         translated text result
	 */
	@PostMapping(produces = MediaType.APPLICATION_JSON_VALUE)
	@ResponseStatus(value = HttpStatus.OK)
	public ResponseEntity<TranslationResponseDto> translateText(
			@Valid @RequestBody(required = true) final TranslationRequestDto translationRequestDto) {
		return ResponseEntity.ok(languageTranslationService.translate(translationRequestDto));
	}

}