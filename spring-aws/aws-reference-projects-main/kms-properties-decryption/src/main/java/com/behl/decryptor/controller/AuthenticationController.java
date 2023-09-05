package com.behl.decryptor.controller;

import java.util.Map;

import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import com.behl.decryptor.utility.JwtUtil;

import lombok.RequiredArgsConstructor;

@RestController
@RequiredArgsConstructor
@RequestMapping(value = "/v1/auth/user")
public class AuthenticationController {

	/**
	 * Controller class that simulates the sign-up and login functionalities. No
	 * actual implementation available other than creation and parsing of JWT token
	 */

	private final JwtUtil jwtUtils;

	@PostMapping(value = "/sign-up/{userName}", produces = MediaType.APPLICATION_JSON_VALUE)
	public ResponseEntity<Map<String, String>> signInHandler(
			@PathVariable(name = "userName", required = true) final String userName) {
		return ResponseEntity.ok(Map.of("accessToken", jwtUtils.generateToken(userName)));
	}

	@PostMapping(value = "/login", produces = MediaType.APPLICATION_JSON_VALUE)
	public ResponseEntity<Map<String, Object>> loginHandler(@RequestBody(required = true) final String accessToken) {
		return ResponseEntity.ok(Map.of("userId", jwtUtils.extractUserId(accessToken), "userName",
				jwtUtils.extractUserName(accessToken)));
	}

}
