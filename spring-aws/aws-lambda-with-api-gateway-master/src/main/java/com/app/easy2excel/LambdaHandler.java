package com.app.easy2excel;

import com.amazonaws.services.lambda.runtime.Context;
import com.amazonaws.services.lambda.runtime.RequestHandler;
import com.amazonaws.services.lambda.runtime.events.APIGatewayProxyRequestEvent;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;

import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.stream.Collectors;
import java.util.stream.Stream;

public class LambdaHandler implements RequestHandler<APIGatewayProxyRequestEvent,APIResponse> {

    public APIResponse handleRequest(APIGatewayProxyRequestEvent input, Context context) {
        return buildResponse(context);
    }

    private APIResponse buildResponse(Context context)  {
        Map<String,String> headers = new HashMap<>();
        headers.put("X-amazon-author","Lipsa");
        headers.put("X-amazon-apiVersion","v1");

        String body = null;
        try {
            body = new ObjectMapper().writeValueAsString(buildEmployee());
        } catch (JsonProcessingException e) {
            context.getLogger().log("Error:::" + e.getMessage());
        }

        APIResponse response = new APIResponse(body,200,headers);
        context.getLogger().log("APIResponse:::" + response);
        return  response;
    }


    private List<Employee> buildEmployee(){
        return Stream.of(
                new Employee(100, "Lipsa Patra", "lipsa@gmail.com"),
                new Employee(101, "Robert Browne", "robert@gmail.com"),
                new Employee(102, "Julia Robert", "julia@gmail.com"),
                new Employee(103, "Daisy Dsouza ", "daisy@gmail.com"),
                new Employee(104, "Jasmine Gupta", "jasmine@gmail.com")
                ).collect(Collectors.toList());
    }
}
