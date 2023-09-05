package com.behl.transcriber.configuration;

import org.springframework.boot.context.properties.EnableConfigurationProperties;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

import com.amazonaws.auth.AWSCredentials;
import com.amazonaws.auth.AWSStaticCredentialsProvider;
import com.amazonaws.auth.BasicAWSCredentials;
import com.amazonaws.regions.Regions;
import com.amazonaws.services.s3.AmazonS3;
import com.amazonaws.services.s3.AmazonS3ClientBuilder;

import lombok.RequiredArgsConstructor;

@Configuration
@RequiredArgsConstructor
@EnableConfigurationProperties(value = AwsIAMConfigurationProperties.class)
public class AwsS3Configuration {

	private final AwsIAMConfigurationProperties awsIAMConfigurationProperties;

	/**
	 * Registers AmazonS3.class bean object in the Spring IOC container which then
	 * can be autowired directly into the consuming class to communicate with Amazon
	 * Simple Storage Service. The bean is constructed using the IAM user security
	 * credentials defined in the active .properties file @see
	 * AwsIAMConfigurationProperties.class
	 * 
	 * If the application is being hosted in an EC2 Instance or ECS and an IAM
	 * Role/Instance profile is being used for authentication then
	 * '.withCredentials(new DefaultAWSCredentialsProviderChain())' can be used
	 * instead of the 'AWSStaticCredentialsProvider' defined below.
	 * AwsIAMConfigurationProperties.class can be discarded in this scenario as well
	 */
	@Bean
	public AmazonS3 amazonS3() {
		AWSCredentials awsCredentials = new BasicAWSCredentials(awsIAMConfigurationProperties.getAccessKey(),
				awsIAMConfigurationProperties.getSecretAccessKey());

		return AmazonS3ClientBuilder.standard().withRegion(Regions.AP_SOUTH_1)
				.withCredentials(new AWSStaticCredentialsProvider(awsCredentials)).build();
	}

}