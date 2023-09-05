package com.behl.delegator.configuration;

import org.springframework.boot.context.properties.ConfigurationProperties;

import lombok.Data;

/**
 * Maps AWS S3 configuration values defined in active .properties file to the
 * below mentioned instance variables which will be used when generating a
 * presigned-URL to delegate the responsibility to upload an object to the
 * client.
 * 
 * @see com.behl.delegator.service.StorageDelegationService
 */
@Data
@ConfigurationProperties(prefix = "com.behl.aws")
public class AwsS3ConfigurationProperties {

	private S3 s3 = new S3();

	@Data
	public class S3 {

		/**
		 * Name of the S3 Bucket to be considered as the destination for uploading the
		 * object for which the presigned-url is being generated.
		 * 
		 * @see https://docs.aws.amazon.com/AmazonS3/latest/userguide/creating-bucket.html
		 */
		private String bucketName;

		private PresignedUrl presignedUrl = new PresignedUrl();

		@Data
		public class PresignedUrl {

			/**
			 * Signifies the time in minutes post which the generated presigned-url will
			 * become invalid and un-useable
			 */
			private Integer expirationTime;

		}

	}

}