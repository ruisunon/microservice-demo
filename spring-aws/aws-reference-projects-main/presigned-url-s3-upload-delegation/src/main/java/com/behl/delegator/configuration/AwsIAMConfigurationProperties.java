package com.behl.delegator.configuration;

import org.springframework.boot.context.properties.ConfigurationProperties;

import lombok.Data;

/**
 * <p>
 * Maps AWS IAM user security credentials configured in active .properties file
 * to the below mentioned instance variables. The permissions attached to the
 * configured IAM user will be used in the policy evaluation logic when making
 * any API call to AWS.
 * </p>
 * 
 * <p>
 * AmazonS3.class bean object is created using the below configured credentials
 * which can be autowired in the consuming classes to communicate with Amazon
 * Simple Storage Service
 * </p>
 * 
 * @see AwsS3Configuration#amazonS3()
 * @see com.amazonaws.services.s3.AmazonS3
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