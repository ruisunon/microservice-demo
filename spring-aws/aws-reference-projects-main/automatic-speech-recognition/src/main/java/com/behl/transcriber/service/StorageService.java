package com.behl.transcriber.service;

import java.io.IOException;

import org.springframework.boot.context.properties.EnableConfigurationProperties;
import org.springframework.http.HttpStatus;
import org.springframework.stereotype.Service;
import org.springframework.web.multipart.MultipartFile;
import org.springframework.web.server.ResponseStatusException;

import com.amazonaws.SdkClientException;
import com.amazonaws.services.s3.AmazonS3;
import com.amazonaws.services.s3.model.AmazonS3Exception;
import com.amazonaws.services.s3.model.ObjectMetadata;
import com.amazonaws.services.s3.model.S3Object;
import com.behl.transcriber.configuration.AwsS3ConfigurationProperties;

import lombok.RequiredArgsConstructor;
import lombok.extern.log4j.Log4j2;

/**
 * Service class acting as an abstraction layer to communicate with AWS Simple
 * Storage Service.
 */

@Log4j2
@Service
@RequiredArgsConstructor
@EnableConfigurationProperties(value = AwsS3ConfigurationProperties.class)
public class StorageService {

	private final AmazonS3 amazonS3;

	/**
	 * Method which communicates with AWS S3 service to store the provided
	 * Multipartfile object into desired S3 Bucket.
	 * 
	 * @param file signifying the object to be stored in the S3 Bucket corresponding
	 *             to the provided bucket-name
	 * @throws ResponseStatusException if communication with AWS S3 service fails to
	 *                                 store the given object
	 * @return HttpStatus.OK if object is stored successfully in S3 Bucket
	 */
	public HttpStatus save(final MultipartFile file, final String bucketName) {
		log.info("Received request to store {} in S3 bucket {}", file.getOriginalFilename(), bucketName);
		final var metadata = constructMetadata(file);
		try {
			amazonS3.putObject(bucketName, file.getOriginalFilename(), file.getInputStream(), metadata);
		} catch (final SdkClientException | IOException exception) {
			log.error("Exception occured while communicating with Amazon S3 service to store object {} in S3 Bucket {}",
					file.getOriginalFilename(), bucketName);
			throw new ResponseStatusException(HttpStatus.INTERNAL_SERVER_ERROR, "Unable to store object in AWS S3",
					exception);
		}
		log.info("Successfully saved {} in S3 Bucket", file.getOriginalFilename(), bucketName);
		return HttpStatus.OK;
	}

	/**
	 * @param objectKey  to retrieve from AWS S3 Bucket
	 * @param bucketName from which the object is to be retrieved
	 * @param isPolling  if set to true does not throw any exception, can be used
	 *                   when an object that is contracted to be stored in the S3
	 *                   bucket at near future timestamp
	 * @return S3Object.class instance representing the retrieved object
	 */
	public S3Object retrieve(final String objectKey, final String bucketName, final boolean isPolling) {
		S3Object retrievedObject;
		try {
			retrievedObject = amazonS3.getObject(bucketName, objectKey);
		} catch (final AmazonS3Exception exception) {
			if (Boolean.TRUE.equals(isPolling))
				return null;
			log.error(
					"Exception occured while communicating with Amazon S3 service to retrieve object {} from S3 Bucket {}",
					objectKey, bucketName, exception);
			throw new ResponseStatusException(HttpStatus.INTERNAL_SERVER_ERROR,
					"Unable to retrieve object from AWS S3 Bucket", exception);
		}
		log.info("Successfully retrieved object {} from S3 Bucket", objectKey, bucketName);
		return retrievedObject;
	}

	/**
	 * @param objectKey  corresponding to the object to retrieve
	 * @param bucketName corresponding to the bucket in which the object to retrieve
	 *                   is contracted to be stored
	 * @return S3Object.class instance representing the retrieved object
	 */
	public S3Object pollForObject(final String objectKey, final String bucketName) {
		log.info("Beginning polling for object {} in S3 Bucket {}", objectKey, bucketName);
		S3Object generatedJsonTextResponse = null;
		var resultRetrieved = false;
		int attempt = 0;
		while (!resultRetrieved) {
			attempt++;
			generatedJsonTextResponse = retrieve(objectKey, bucketName, Boolean.TRUE);
			if (generatedJsonTextResponse != null)
				resultRetrieved = true;
		}
		log.info("Retrieved object with key {} from S3 Bucket in attempt number {}", objectKey, bucketName, attempt);
		return generatedJsonTextResponse;
	}

	/**
	 * @param objectKey  corresponding to the stored object in S3
	 * @param bucketName corresponding to the bucket in which the object is stored
	 * 
	 * @return S3 Object URI String corresponding to the given parameters
	 */
	public String constructS3Uri(final String objectKey, final String bucketName) {
		return "s3://" + bucketName + "/" + objectKey;
	}

	private ObjectMetadata constructMetadata(final MultipartFile file) {
		log.info("Received request to generate object Metadata for {}", file.getOriginalFilename());
		final ObjectMetadata metadata = new ObjectMetadata();
		metadata.setContentLength(file.getSize());
		metadata.setContentType(file.getContentType());
		metadata.setContentDisposition(file.getOriginalFilename());
		return metadata;
	}

}