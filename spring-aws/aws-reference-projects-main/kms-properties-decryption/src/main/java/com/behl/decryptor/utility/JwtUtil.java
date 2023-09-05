package com.behl.decryptor.utility;

import java.nio.charset.Charset;
import java.time.LocalDateTime;
import java.util.Date;
import java.util.HashMap;
import java.util.Map;
import java.util.UUID;
import java.util.concurrent.TimeUnit;
import java.util.function.Function;

import org.springframework.boot.context.properties.EnableConfigurationProperties;
import org.springframework.http.HttpStatus;
import org.springframework.stereotype.Component;
import org.springframework.web.server.ResponseStatusException;

import com.behl.decryptor.properties.JwtConfigurationProperties;

import io.jsonwebtoken.Claims;
import io.jsonwebtoken.Jwts;
import io.jsonwebtoken.security.Keys;
import io.jsonwebtoken.security.SignatureException;
import lombok.AllArgsConstructor;
import lombok.extern.slf4j.Slf4j;

@Slf4j
@Component
@EnableConfigurationProperties(JwtConfigurationProperties.class)
@AllArgsConstructor
public class JwtUtil {

	private final JwtConfigurationProperties jwtConfigurationProperties;

	public String extractUserName(final String token) {
		return extractClaim(token, Claims::getSubject);
	}

	public UUID extractUserId(final String token) {
		return UUID.fromString((String) extractAllClaims(token).get("user_id"));
	}

	private Claims extractAllClaims(String token) {
		try {
			return Jwts.parserBuilder()
					.setSigningKey(
							jwtConfigurationProperties.getJwt().getSecretKey().getBytes(Charset.forName("UTF-8")))
					.build().parseClaimsJws(token.replace("Bearer ", "")).getBody();
		} catch (final SignatureException exception) {
			log.error("Untrusted JWT signature found", exception);
			throw new ResponseStatusException(HttpStatus.UNAUTHORIZED);
		}
	}

	private <T> T extractClaim(final String token, final Function<Claims, T> claimsResolver) {
		final Claims claims = extractAllClaims(token);
		return claimsResolver.apply(claims);
	}

	public String generateToken(final String userName) {
		Map<String, Object> claims = new HashMap<>();
		claims.put("user_id", UUID.randomUUID());
		claims.put("user_name", userName);
		claims.put("created_at", LocalDateTime.now());
		return createToken(claims, userName, TimeUnit.MINUTES.toMillis(90));
	}

	private String createToken(final Map<String, Object> claims, final String subject, final Long expiration) {
		final var key = Keys
				.hmacShaKeyFor(jwtConfigurationProperties.getJwt().getSecretKey().getBytes(Charset.forName("UTF-8")));
		return Jwts.builder().setClaims(claims).setSubject(subject).setIssuedAt(new Date(System.currentTimeMillis()))
				.setExpiration(new Date(System.currentTimeMillis() + expiration)).signWith(key).compact();
	}

}