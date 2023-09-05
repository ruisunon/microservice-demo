package com.behl.transcriber.controller;

import org.springframework.core.io.InputStreamResource;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestPart;
import org.springframework.web.bind.annotation.ResponseStatus;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.multipart.MultipartFile;

import com.behl.transcriber.service.SpeechToTextConvertionService;

import lombok.RequiredArgsConstructor;
import lombok.extern.log4j.Log4j2;

@Log4j2
@RestController
@RequiredArgsConstructor
public class AudioToSpeechConversionController {

	private final SpeechToTextConvertionService speechToTextConvertionService;

	/**
	 * Exposed API endpoint method capable of Automatic speech recognition to
	 * convert audio-to-text.
	 * 
	 * @param audioFile to convert into a textual representation
	 * @return JSON file containing the textual representation of the given
	 *         audiofile
	 */
	@PostMapping(value = "/audio/transcribe", consumes = MediaType.MULTIPART_FORM_DATA_VALUE)
	@ResponseStatus(value = HttpStatus.OK)
	public ResponseEntity<InputStreamResource> speechToTextConverter(
			@RequestPart(name = "file", required = true) final MultipartFile audioFile) {
		log.info("Received audio file {} to convert into text", audioFile.getOriginalFilename());

		final var resultJsonFile = speechToTextConvertionService.convert(audioFile);
		final var resultInputStream = new InputStreamResource(resultJsonFile);

		return ResponseEntity.status(HttpStatus.OK)
				.header(HttpHeaders.CONTENT_DISPOSITION, "attachment;filename=result.json").body(resultInputStream);
	}

}
