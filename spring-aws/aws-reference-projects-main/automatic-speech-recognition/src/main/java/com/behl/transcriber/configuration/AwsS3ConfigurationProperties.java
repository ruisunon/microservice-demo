package com.behl.transcriber.configuration;

import org.springframework.boot.context.properties.ConfigurationProperties;

import lombok.Data;

/**
 * Maps AWS S3 Input and Output Bucket information configured in active
 * .properties file to the below mentioned instance variables. The S3 Buckets
 * corresponding to the configured names will be used during the process of
 * automatic speech recognition and subsequent textual result generation when
 * communicating with Amazon Transcribe service.
 */
@Data
@ConfigurationProperties(prefix = "com.behl.aws")
public class AwsS3ConfigurationProperties {

	private S3 s3 = new S3();

	@Data
	public class S3 {

		/**
		 * <code>Name</code> of the S3 Bucket to be configured to store Input audio
		 * files prior to beginning the transcribe process, to be defined in the active
		 * .properties file corresponding to the key of
		 * 'com.behl.aws.s3.input-bucket-name'
		 * 
		 * Do not include the <code>S3://</code> prefix of the specified bucket.
		 */
		private String inputBucketName;

		/**
		 * <code>Name</code> of the S3 Bucket to be configured to store Output JSON
		 * files containing the textual result post transcribe process, to be defined in
		 * the active .properties file corresponding to the key of
		 * 'com.behl.aws.s3.output-bucket-name'
		 * 
		 * Do not include the <code>S3://</code> prefix of the specified bucket.
		 */
		private String outputBucketName;

	}

}