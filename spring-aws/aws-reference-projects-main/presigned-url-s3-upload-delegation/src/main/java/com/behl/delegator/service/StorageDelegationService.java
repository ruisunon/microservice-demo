package com.behl.delegator.service;

import java.net.URISyntaxException;
import java.net.URL;

import org.joda.time.LocalDateTime;
import org.springframework.boot.context.properties.EnableConfigurationProperties;
import org.springframework.http.HttpStatus;
import org.springframework.stereotype.Service;
import org.springframework.web.server.ResponseStatusException;

import com.amazonaws.HttpMethod;
import com.amazonaws.SdkClientException;
import com.amazonaws.services.s3.AmazonS3;
import com.amazonaws.services.s3.model.GeneratePresignedUrlRequest;
import com.behl.delegator.configuration.AwsS3ConfigurationProperties;

import lombok.RequiredArgsConstructor;
import lombok.extern.log4j.Log4j2;

@Log4j2
@Service
@RequiredArgsConstructor
@EnableConfigurationProperties(value = AwsS3ConfigurationProperties.class)
public class StorageDelegationService {

	private final AmazonS3 amazonS3;
	private final AwsS3ConfigurationProperties awsS3ConfigurationProperties;

	/**
	 * 
	 * Method to generate S3 Presigned URL corresponding to the configured S3 Bucket
	 * and expiration time which can be used to delegate the responsibility of file
	 * uploadation to the client itself. The process would only be successful if the
	 * IAM user/role generating the presigned URL has the permission to upload
	 * objects to the configured bucket.
	 * 
	 * @param objectKey for which the presigned URL is to be generated. It is
	 *                  recommended to use the same name that the file is intended
	 *                  to be stored as.
	 * 
	 * @return Presigned URL to enable client to upload object to configured S3
	 *         bucket on the behalf of the backend server. The uploadation has to be
	 *         completed by the client prior to the passing of configured expiration
	 *         time.
	 * 
	 * @see com.behl.delegator.configuration.AwsS3ConfigurationProperties
	 * @see com.behl.delegator.configuration.AwsIAMConfigurationProperties
	 */
	public String generateObjectUploadationPresignedUrl(final String objectKey) {
		final var generatePresignedUrlRequest = new GeneratePresignedUrlRequest(
				awsS3ConfigurationProperties.getS3().getBucketName(), objectKey, HttpMethod.PUT);
		generatePresignedUrlRequest.setExpiration(new LocalDateTime()
				.plusMinutes(awsS3ConfigurationProperties.getS3().getPresignedUrl().getExpirationTime()).toDate());

		URL presignedUrl;
		try {
			log.info("Attempting to generate S3 presigned URL against request {}", generatePresignedUrlRequest);
			presignedUrl = amazonS3.generatePresignedUrl(generatePresignedUrlRequest);
		} catch (final SdkClientException exception) {
			log.error(
					"Exception occured while communicating with Amazon S3 service to generate presigned URL against request {}",
					generatePresignedUrlRequest, exception);
			throw new ResponseStatusException(HttpStatus.INTERNAL_SERVER_ERROR, "Unable to generate presigned URLs",
					exception);
		}

		log.info("Successfully generated S3 presigned URL against request {}", generatePresignedUrlRequest);
		try {
			return presignedUrl.toURI().toString();
		} catch (final URISyntaxException exception) {
			log.error("Generated Presigned URL could not be parsed as a URI reference", exception);
			throw new ResponseStatusException(HttpStatus.EXPECTATION_FAILED);
		}
	}

}
