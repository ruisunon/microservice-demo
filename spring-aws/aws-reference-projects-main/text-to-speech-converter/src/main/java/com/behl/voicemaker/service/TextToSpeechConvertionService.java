package com.behl.voicemaker.service;

import java.io.InputStream;
import java.util.List;
import java.util.Random;

import org.springframework.http.HttpStatus;
import org.springframework.stereotype.Service;
import org.springframework.web.server.ResponseStatusException;

import com.amazonaws.SdkClientException;
import com.amazonaws.services.polly.AmazonPolly;
import com.amazonaws.services.polly.model.DescribeVoicesRequest;
import com.amazonaws.services.polly.model.DescribeVoicesResult;
import com.amazonaws.services.polly.model.OutputFormat;
import com.amazonaws.services.polly.model.SynthesizeSpeechRequest;
import com.amazonaws.services.polly.model.SynthesizeSpeechResult;
import com.amazonaws.services.polly.model.Voice;

import lombok.RequiredArgsConstructor;
import lombok.extern.log4j.Log4j2;

@Log4j2
@Service
@RequiredArgsConstructor
public class TextToSpeechConvertionService {

	private final AmazonPolly amazonPollyClient;

	/**
	 * Generates Audio Speech from provided textual content
	 * 
	 * @param <code>text</code> to convert into an audio speech.
	 * @return InputStream signifying the generated audio speech
	 */
	public InputStream convert(final String text) {
		log.info("Beginning convertion of text '{}' into MP3 Audio Speech", text);
		final var randomVoice = getRandomVoice();
		final var speechGenerationRequest = new SynthesizeSpeechRequest().withText(text)
				.withVoiceId(randomVoice.getId()).withOutputFormat(OutputFormat.Mp3);

		SynthesizeSpeechResult speechGenerationResult;
		try {
			speechGenerationResult = amazonPollyClient.synthesizeSpeech(speechGenerationRequest);
		} catch (SdkClientException exception) {
			log.error(
					"Exception occured while communicating with Amazon Polly service against request to generate speech from provided textual content");
			throw new ResponseStatusException(HttpStatus.INTERNAL_SERVER_ERROR, "Unable to generate Audio speech",
					exception);
		}
		log.info("Successfully generated Audio speech against provided textual content");
		return speechGenerationResult.getAudioStream();
	}

	/**
	 * Method to select a Random Voice.class Instance from the available
	 * List<Voice>. The logic can be changed to return a hard-coded Voice.class
	 * instance for each API call based on business requirements.
	 * 
	 * @return Voice.class Instance using which the process of speech generation is
	 *         to be undertaken
	 */
	private Voice getRandomVoice() {
		log.info("Received request to return a random Voice.class instance to be used in speech generation");
		DescribeVoicesResult describeVoicesResult;
		try {
			describeVoicesResult = amazonPollyClient.describeVoices(new DescribeVoicesRequest());
		} catch (SdkClientException exception) {
			log.error(
					"Exception occured while communicating with Amazon Polly service against request to describe configured voices");
			throw new ResponseStatusException(HttpStatus.INTERNAL_SERVER_ERROR, "Unable to generate Audio speech",
					exception);
		}

		List<Voice> voices = describeVoicesResult.getVoices();
		Voice selectedVoice = voices.get(new Random().nextInt(voices.size()));

		log.info("From {} total available choises, Returning Voice {}", voices.size(), selectedVoice);
		return selectedVoice;
	}

}
