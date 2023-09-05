package com.behl.transcriber.service;

import java.io.InputStream;

import org.apache.commons.lang3.RandomStringUtils;
import org.springframework.boot.context.properties.EnableConfigurationProperties;
import org.springframework.http.HttpStatus;
import org.springframework.stereotype.Service;
import org.springframework.web.multipart.MultipartFile;
import org.springframework.web.server.ResponseStatusException;

import com.amazonaws.SdkClientException;
import com.amazonaws.services.s3.model.S3Object;
import com.amazonaws.services.transcribe.AmazonTranscribe;
import com.amazonaws.services.transcribe.model.Media;
import com.amazonaws.services.transcribe.model.StartTranscriptionJobRequest;
import com.behl.transcriber.configuration.AwsS3ConfigurationProperties;

import lombok.RequiredArgsConstructor;
import lombok.extern.log4j.Log4j2;

@Log4j2
@Service
@RequiredArgsConstructor
@EnableConfigurationProperties(AwsS3ConfigurationProperties.class)
public class SpeechToTextConvertionService {

	private final StorageService storageService;
	private final AmazonTranscribe amazonTranscribe;
	private final AwsS3ConfigurationProperties awsS3ConfigurationProperties;

	/**
	 * Converts a given Audio file into a textual representation. AWS S3 and Amazon
	 * Transcribe services are used in the mentioned process. To see the allowed
	 * audio formats, Refer this <a href=
	 * "URL#https://docs.aws.amazon.com/transcribe/latest/dg/how-input.html">Article</a>
	 * 
	 * @param audioFile to convert into a textual representation
	 * @return InputStream representing a JSON file containing the textual
	 *         representation of the given audiofile
	 */
	public InputStream convert(final MultipartFile audioFile) {
		final var s3Configuration = awsS3ConfigurationProperties.getS3();
		storageService.save(audioFile, s3Configuration.getInputBucketName());
		log.info("Successfully stored input audio file {} in input S3 Bucket {}. Proceeding with further processing",
				audioFile.getOriginalFilename(), s3Configuration.getInputBucketName());

		final var media = new Media();
		final var audioFileS3Uri = storageService.constructS3Uri(audioFile.getOriginalFilename(),
				s3Configuration.getInputBucketName());
		media.setMediaFileUri(audioFileS3Uri);

		final String transcriptionJobName = RandomStringUtils.randomAlphabetic(8);
		StartTranscriptionJobRequest startTranscriptionJobRequest = new StartTranscriptionJobRequest();
		startTranscriptionJobRequest.setIdentifyLanguage(Boolean.TRUE);
		startTranscriptionJobRequest.setOutputBucketName(s3Configuration.getOutputBucketName());
		startTranscriptionJobRequest.setTranscriptionJobName(transcriptionJobName);
		startTranscriptionJobRequest.setMedia(media);

		try {
			log.info("Starting transcription job against request {}", startTranscriptionJobRequest);
			amazonTranscribe.startTranscriptionJob(startTranscriptionJobRequest);
		} catch (final SdkClientException exception) {
			log.error("Exception occured while communicating with Amazon Transcribe service against request {}",
					startTranscriptionJobRequest);
			throw new ResponseStatusException(HttpStatus.INTERNAL_SERVER_ERROR, "Unable to transcribe audio file",
					exception);
		}
		log.info("Successfully started transcription job against request {}", startTranscriptionJobRequest);

		log.info("Beginning polling for JSON file result against transaction Job request {}",
				startTranscriptionJobRequest);
		final S3Object generatedJsonTextResponse = storageService.pollForObject(transcriptionJobName + ".json",
				s3Configuration.getOutputBucketName());

		log.info("Successfully retrieved generated JSON file response against transaction job request {}",
				startTranscriptionJobRequest);
		return generatedJsonTextResponse.getObjectContent();
	}

}