package com.app.easy2excel;

import lombok.AllArgsConstructor;
import lombok.Data;

import java.util.Map;

@Data
@AllArgsConstructor
public class APIResponse {
    private String body;
    private Integer statusCode;
    private Map<String, String> headers;
}
