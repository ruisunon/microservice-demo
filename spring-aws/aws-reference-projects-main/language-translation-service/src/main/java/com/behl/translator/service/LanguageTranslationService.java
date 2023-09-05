package com.behl.translator.service;

import java.util.List;

import org.springframework.http.HttpStatus;
import org.springframework.stereotype.Service;
import org.springframework.web.server.ResponseStatusException;

import com.amazonaws.SdkClientException;
import com.amazonaws.services.translate.AmazonTranslate;
import com.amazonaws.services.translate.model.Language;
import com.amazonaws.services.translate.model.ListLanguagesRequest;
import com.amazonaws.services.translate.model.ListLanguagesResult;
import com.amazonaws.services.translate.model.TranslateTextRequest;
import com.amazonaws.services.translate.model.TranslateTextResult;
import com.behl.translator.dto.TranslationRequestDto;
import com.behl.translator.dto.TranslationResponseDto;

import lombok.RequiredArgsConstructor;
import lombok.extern.log4j.Log4j2;

@Log4j2
@Service
@RequiredArgsConstructor
public class LanguageTranslationService {

	private final AmazonTranslate amazonTranslate;

	/**
	 * Method to translate texts in a source language to a specified target
	 * language. The language codes being specified in source and target can be
	 * retrieved from {@link #getSupportedLanguages()}
	 * 
	 * @throws {@link org.springframework.web.server.ResponseStatusException} when
	 *                communication with Amazon Translate service fails
	 * @param translationRequestDto {@link com.behl.translator.dto.TranslationRequestDto}
	 * @return {@link com.behl.translator.dto.TranslationResponseDto}
	 */
	public TranslationResponseDto translate(final TranslationRequestDto translationRequestDto) {
		final TranslateTextRequest translateTextRequest = new TranslateTextRequest()
				.withText(translationRequestDto.getText())
				.withSourceLanguageCode(translationRequestDto.getSourceLanguageCode())
				.withTargetLanguageCode(translationRequestDto.getTargetLanguageCode());

		TranslateTextResult translateTextResult;
		try {
			log.info("Attempting to call TranslateText API againt request {}", translateTextRequest);
			translateTextResult = amazonTranslate.translateText(translateTextRequest);
		} catch (SdkClientException exception) {
			log.error("Exception occured while communicating with Amazon Translate service against request {}",
					translateTextRequest);
			throw new ResponseStatusException(HttpStatus.INTERNAL_SERVER_ERROR,
					"Communication with Amazon Translate service failed", exception);
		}

		log.info("Successfully recieved translated text result corresponding to request {}", translateTextRequest);
		return TranslationResponseDto.builder().translatedText(translateTextResult.getTranslatedText())
				.originalText(translationRequestDto.getText())
				.sourceLanguageCode(translateTextResult.getSourceLanguageCode())
				.targetLanguageCode(translateTextResult.getTargetLanguageCode()).build();
	}

	/**
	 * @return List of {@link com.amazonaws.services.translate.model.Language} which
	 *         are supported for the process of language-translation by Amazon
	 *         Translate Service. The returned object contains both the language
	 *         name and language code, the latter of which is to be used when
	 *         <a href=
	 *         "URL#https://docs.aws.amazon.com/translate/latest/APIReference/API_TranslateText.html">TranslateText
	 *         API</a> is called
	 */
	public List<Language> getSupportedLanguages() {
		log.info("Recieved request to return list of all supported languages for language-translation");
		ListLanguagesResult listLanguagesResult;
		try {
			listLanguagesResult = amazonTranslate.listLanguages(new ListLanguagesRequest());
		} catch (final SdkClientException exception) {
			log.error(
					"Exception occured while communicating with Amazon Translate service against request to retrieve list of all supported languages");
			throw new ResponseStatusException(HttpStatus.INTERNAL_SERVER_ERROR,
					"Communication with Amazon Translate service failed", exception);
		}
		final var supportedLanguages = listLanguagesResult.getLanguages();
		log.info("Successfully recieved list of {} supported languages for the process of language-translation",
				supportedLanguages.size());
		return supportedLanguages;
	}

}