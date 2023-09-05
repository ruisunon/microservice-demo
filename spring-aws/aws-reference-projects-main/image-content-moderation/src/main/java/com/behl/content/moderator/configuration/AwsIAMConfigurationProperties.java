package com.behl.content.moderator.configuration;

import org.springframework.boot.context.properties.ConfigurationProperties;

import lombok.Data;

/**
 * Maps AWS IAM user security credentials configured in active .properties file
 * to the below mentioned variables. The permissions attached to the configured
 * IAM user will be used when making any API call to AWS.
 * 
 * AmazonRekognition.class bean object is created using the below configured
 * credentials. @see AwsRekognitionConfiguration#amazonRekognition()
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