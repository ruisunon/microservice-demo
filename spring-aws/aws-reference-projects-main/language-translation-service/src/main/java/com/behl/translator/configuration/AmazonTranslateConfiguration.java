package com.behl.translator.configuration;

import org.springframework.boot.context.properties.EnableConfigurationProperties;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

import com.amazonaws.auth.AWSStaticCredentialsProvider;
import com.amazonaws.auth.BasicAWSCredentials;
import com.amazonaws.regions.Regions;
import com.amazonaws.services.translate.AmazonTranslate;
import com.amazonaws.services.translate.AmazonTranslateClientBuilder;

import lombok.RequiredArgsConstructor;

@Configuration
@EnableConfigurationProperties(AwsIAMConfigurationProperties.class)
@RequiredArgsConstructor
public class AmazonTranslateConfiguration {

	private final AwsIAMConfigurationProperties awsIAMConfigurationProperties;

	/**
	 * Registers AmazonTranslate.class bean object in the Spring IOC container which
	 * then can be autowired directly into the consuming class(es). The bean is
	 * constructed using the IAM user's security credentials defined in the active
	 * .properties file mapped to AwsIAMConfigurationProperties.class
	 * 
	 * If the application is being hosted in an EC2 Instance or ECS and IAM
	 * Role/Instance profile is being used for authentication then
	 * '.withCredentials(new DefaultAWSCredentialsProviderChain())' can be used
	 * instead of the 'AWSStaticCredentialsProvider' defined below.
	 * AwsIAMConfigurationProperties.class can be discarded in this scenario as well
	 * 
	 * @see AwsIAMConfigurationProperties
	 * @see com.amazonaws.services.translate.AmazonTranslate
	 */
	@Bean
	public AmazonTranslate amazonTranslate() {
		BasicAWSCredentials credentials = new BasicAWSCredentials(awsIAMConfigurationProperties.getAccessKey(),
				awsIAMConfigurationProperties.getSecretAccessKey());

		return AmazonTranslateClientBuilder.standard().withRegion(Regions.AP_SOUTH_1)
				.withCredentials(new AWSStaticCredentialsProvider(credentials)).build();
	}

}