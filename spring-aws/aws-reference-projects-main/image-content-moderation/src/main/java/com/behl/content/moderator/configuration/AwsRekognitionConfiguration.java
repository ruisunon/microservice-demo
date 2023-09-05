package com.behl.content.moderator.configuration;

import org.springframework.boot.context.properties.EnableConfigurationProperties;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

import com.amazonaws.auth.AWSStaticCredentialsProvider;
import com.amazonaws.auth.BasicAWSCredentials;
import com.amazonaws.regions.Regions;
import com.amazonaws.services.rekognition.AmazonRekognition;
import com.amazonaws.services.rekognition.AmazonRekognitionClientBuilder;

import lombok.RequiredArgsConstructor;

@Configuration
@EnableConfigurationProperties(value = AwsIAMConfigurationProperties.class)
@RequiredArgsConstructor
public class AwsRekognitionConfiguration {

	private final AwsIAMConfigurationProperties awsIAMConfigurationProperties;

	/**
	 * Registers AmazonRecognition.class bean object in the Spring IOC container
	 * which then can be autowired directly into the consuming class. The bean is
	 * constructed using the IAM user security credentials defined in the active
	 * .properties file @see AwsIAMConfigurationProperties.class
	 * 
	 * If the application is being hosted in an EC2 Instance or ECS and IAM
	 * Role/Instance profile is being used for authentication then
	 * '.withCredentials(new DefaultAWSCredentialsProviderChain())' can be used
	 * instead of the 'AWSStaticCredentialsProvider' defined below.
	 * AwsIAMConfigurationProperties.class can be discarded in this scenario as well
	 */
	@Bean
	public AmazonRekognition amazonRekognition() {
		BasicAWSCredentials credentials = new BasicAWSCredentials(awsIAMConfigurationProperties.getAccessKey(),
				awsIAMConfigurationProperties.getSecretAccessKey());

		return AmazonRekognitionClientBuilder.standard().withRegion(Regions.AP_SOUTH_1)
				.withCredentials(new AWSStaticCredentialsProvider(credentials)).build();
	}

}