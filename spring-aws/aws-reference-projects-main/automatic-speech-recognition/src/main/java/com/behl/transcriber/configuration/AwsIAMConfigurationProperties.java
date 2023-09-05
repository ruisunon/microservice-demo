package com.behl.transcriber.configuration;

import org.springframework.boot.context.properties.ConfigurationProperties;

import lombok.Data;

/**
 * Maps AWS IAM user security credentials configured in active .properties file
 * to the below mentioned instance variables. The permissions attached to the
 * configured IAM user will be used for policy evaluation logic when making any
 * API call to AWS.
 * 
 * AmazonTranscribe.class and AmazonS3.class bean objects are created using the
 * below configured credentials.
 * 
 * @see AmazonTranscribeConfiguration#amazonTranscribe()
 * @see AwsS3Configuration#amazonS3()
 */

@Data
@ConfigurationProperties(prefix = "com.behl.aws")
public class AwsIAMConfigurationProperties {

	/**
	 * Access key ID of the IAM user to be defined corresponding to the key
	 * `com.behl.aws.access-key` in active .properties file
	 */
	private String accessKey;

	/**
	 * Secret Access key of the IAM user to be defined corresponding to the key
	 * `com.behl.aws.secret-access-key` in active .properties file
	 */
	private String secretAccessKey;

}