package com.behl.delegator.controller;

import java.util.Map;

import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseStatus;
import org.springframework.web.bind.annotation.RestController;

import com.behl.delegator.service.StorageDelegationService;

import lombok.RequiredArgsConstructor;
import lombok.extern.log4j.Log4j2;

@Log4j2
@RestController
@RequiredArgsConstructor
@RequestMapping(value = "/upload")
public class ObjectStorageController {

	private final StorageDelegationService storageDelegationService;

	/**
	 * Exposed API endpoint method with the ability to generate Presigned URL to
	 * delegate the repsonsibility of file uploadation to the client.
	 * 
	 * @param objectKey which will be corresponded to the object being uploaded
	 *                  through generated presigned URL.
	 * 
	 * @return Presigned URL to enable client to upload object to configured S3
	 *         bucket on the behalf of the backend server. The uploadation has to be
	 *         completed by the client prior to the passing of configured expiration
	 *         time.
	 */
	@GetMapping(value = "/{objectKey}", produces = MediaType.APPLICATION_JSON_VALUE)
	@ResponseStatus(value = HttpStatus.OK)
	public ResponseEntity<Map<String, String>> generateObjectUploadationPresignedUrl(
			@PathVariable(name = "objectKey", required = true) final String objectKey) {
		log.info("Recieved request to generate presigned URL corresponding to object-key {}", objectKey);
		final var presignedUri = storageDelegationService.generateObjectUploadationPresignedUrl(objectKey);
		return ResponseEntity.status(HttpStatus.OK).body(Map.of("URI", presignedUri));
	}

}
