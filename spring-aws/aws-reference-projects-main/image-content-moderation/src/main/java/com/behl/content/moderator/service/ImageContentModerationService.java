package com.behl.content.moderator.service;

import java.io.IOException;
import java.nio.ByteBuffer;
import java.util.List;

import org.springframework.http.HttpStatus;
import org.springframework.stereotype.Service;
import org.springframework.web.multipart.MultipartFile;
import org.springframework.web.server.ResponseStatusException;

import com.amazonaws.SdkClientException;
import com.amazonaws.services.rekognition.AmazonRekognition;
import com.amazonaws.services.rekognition.model.DetectModerationLabelsRequest;
import com.amazonaws.services.rekognition.model.DetectModerationLabelsResult;
import com.amazonaws.services.rekognition.model.Image;
import com.amazonaws.services.rekognition.model.ModerationLabel;

import lombok.RequiredArgsConstructor;
import lombok.extern.log4j.Log4j2;

@Log4j2
@Service
@RequiredArgsConstructor
public class ImageContentModerationService {

	private final AmazonRekognition amazonRekognition;

	/**
	 * 
	 * Performs Image content moderation against the provided MultipartFile.class
	 * object using Amazon Rekognition service.
	 * 
	 * @param file : The image file for which content moderation has to be
	 *             performed. only .png and .jpeg file extensions are supported
	 * @return List<ModerationLabel> obtained after content moderation detection is
	 *         performed. If the returned list contains any elements, then the
	 *         provided image file has failed content moderation and includes
	 *         inappropriate, unwanted, suggestive or offensive material. When an
	 *         empty list is returned, the provided image has passed content
	 *         moderation and can be processed further without harm
	 */
	public List<ModerationLabel> moderate(final MultipartFile file) {
		log.info("Content moderation execution request received for {}", file.getOriginalFilename());
		DetectModerationLabelsRequest request;
		try {
			request = new DetectModerationLabelsRequest()
					.withImage(new Image().withBytes(ByteBuffer.wrap(file.getBytes())));
		} catch (final IOException exception) {
			log.error("Exception occured while moderating file: {}", file.getOriginalFilename());
			throw new ResponseStatusException(HttpStatus.INTERNAL_SERVER_ERROR, "Unable to moderate provided file",
					exception);
		}

		DetectModerationLabelsResult moderationDetectionResult;
		try {
			moderationDetectionResult = amazonRekognition.detectModerationLabels(request);
		} catch (SdkClientException exception) {
			log.error("Exception occured while communicating with Amazon Rekognition service against request : {}",
					request);
			throw new ResponseStatusException(HttpStatus.INTERNAL_SERVER_ERROR, "Unable to moderate provided file",
					exception);
		}

		List<ModerationLabel> moderationLabels = moderationDetectionResult.getModerationLabels();
		log.info("Successfully performed content moderation against {} and found {} moderation labels",
				file.getOriginalFilename(), moderationLabels.size());
		return moderationLabels;
	}

}