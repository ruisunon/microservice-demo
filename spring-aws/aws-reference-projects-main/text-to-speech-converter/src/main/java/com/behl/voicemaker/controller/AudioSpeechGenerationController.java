package com.behl.voicemaker.controller;

import org.springframework.core.io.InputStreamResource;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.ResponseStatus;
import org.springframework.web.bind.annotation.RestController;

import com.behl.voicemaker.dto.TextToSpeechConversionRequestDto;
import com.behl.voicemaker.service.TextToSpeechConvertionService;

import lombok.RequiredArgsConstructor;

@RestController
@RequiredArgsConstructor
public class AudioSpeechGenerationController {

	private final TextToSpeechConvertionService textToSpeechConvertionService;

	/**
	 * 
	 * Exposed API endpoint providing the capability to convert provided textual
	 * content into Audio speech
	 * 
	 * @param request JSON request-body containing the textual content intended to
	 *                be converted into Audio speech
	 * @return InputStreamResource signiying the generated Audio speech
	 */
	@PostMapping(value = "/v1/generate/audio", consumes = MediaType.APPLICATION_JSON_VALUE)
	@ResponseStatus(value = HttpStatus.OK)
	public ResponseEntity<InputStreamResource> audioGenerationHandler(
			@RequestBody(required = true) final TextToSpeechConversionRequestDto request) {
		final var generatedAudioSpeech = new InputStreamResource(
				textToSpeechConvertionService.convert(request.getText()));

		return ResponseEntity.status(HttpStatus.OK)
				.header(HttpHeaders.CONTENT_DISPOSITION,
						"attachment;filename=audio_speech_" + System.currentTimeMillis() + ".mp3")
				.body(generatedAudioSpeech);
	}

}
