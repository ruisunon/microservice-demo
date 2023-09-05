package com.behl.content.moderator.controller;

import java.util.Map;

import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestPart;
import org.springframework.web.bind.annotation.ResponseStatus;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.multipart.MultipartFile;

import com.amazonaws.util.CollectionUtils;
import com.behl.content.moderator.service.ImageContentModerationService;
import com.behl.content.moderator.utils.FileUtil;
import com.behl.content.moderator.utils.ResponseUtil;

import lombok.RequiredArgsConstructor;
import lombok.extern.log4j.Log4j2;

@Log4j2
@RestController
@RequiredArgsConstructor
public class UserProfileController {

	private final ImageContentModerationService imageContentModerationService;

	/**
	 * API endpoint method to set/update logged-in users profile picture. Image
	 * content moderation is executed prior to proccessing
	 * 
	 * @param profileImage : MultipartFile instance to set as current profile
	 *                     picture of the logged-in user
	 * @return <code>HttpStatus.OK</code> If given image has successfully passed
	 *         content moderation execution and no harmful content is found
	 *         <code>HttpStatus.NOT_ACCEPTABLE</code> If a file with content-type of
	 *         image/png or image/jpeg is not provided
	 *         <code>HttpStatus.NOT_ACCEPTABLE</code> If given image has failed
	 *         content moderation execution
	 */
	@ResponseStatus(value = HttpStatus.OK)
	@PostMapping(value = "/users/profile/image", consumes = MediaType.MULTIPART_FORM_DATA_VALUE, produces = MediaType.APPLICATION_JSON_VALUE)
	public ResponseEntity<?> updateProfilePicture(
			@RequestPart(name = "file", required = true) final MultipartFile profileImage) {
		if (Boolean.FALSE.equals(FileUtil.isJpeg(profileImage)) && Boolean.FALSE.equals(FileUtil.isPng(profileImage)))
			return ResponseEntity.status(HttpStatus.NOT_ACCEPTABLE)
					.body(Map.of("error", "Only image/png or image/jpeg are allowed"));

		final var moderationLabels = imageContentModerationService.moderate(profileImage);

		if (!CollectionUtils.isNullOrEmpty(moderationLabels)) {
			log.error("{} Failed content moderation process. Stopping further processing",
					profileImage.getOriginalFilename());
			final String errorResponse = ResponseUtil.failedContentModerationResponse(moderationLabels);
			return ResponseEntity.status(HttpStatus.NOT_ACCEPTABLE).body(Map.of("error", errorResponse));
		}

		log.info("Provided image {} successfully passed content moderation process. Proceeding with further processing",
				profileImage.getOriginalFilename());
		return ResponseEntity.status(HttpStatus.OK).build();
	}

}