package com.behl.decryptor.properties;

import org.springframework.boot.context.properties.ConfigurationProperties;

import lombok.Data;

@ConfigurationProperties(prefix = "com.behl.decryptor")
@Data
public class JwtConfigurationProperties {

	private JWT jwt = new JWT();

	@Data
	public class JWT {
		private String secretKey;
	}

}