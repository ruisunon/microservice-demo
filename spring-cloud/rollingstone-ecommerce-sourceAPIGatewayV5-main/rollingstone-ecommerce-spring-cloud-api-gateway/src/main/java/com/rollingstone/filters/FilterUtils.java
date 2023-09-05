package com.rollingstone.filters;

import org.springframework.http.HttpHeaders;
import org.springframework.stereotype.Component;
import org.springframework.web.server.ServerWebExchange;

import java.util.List;


@Component
public class FilterUtils {

	public static final String ROLLINGSTONE_CORRELATION_ID = "rollingstone-correlation-id";


	public String getRollingStoneCorrelationId(HttpHeaders requestHeaders){
		if (requestHeaders.get(ROLLINGSTONE_CORRELATION_ID) !=null) {
			List<String> header = requestHeaders.get(ROLLINGSTONE_CORRELATION_ID);
			return header.stream().findFirst().get();
		}
		else{
			return null;
		}
	}

	public ServerWebExchange setRollingStoneRequestHeader(ServerWebExchange exchange, String name, String value) {
		return exchange.mutate().request(
							exchange.getRequest().mutate()
							.header(name, value)
							.build())
						.build();	
	}
	
	public ServerWebExchange setRollingStoneCorrelationId(ServerWebExchange exchange, String rollingStoneCorrelationId) {
		return this.setRollingStoneRequestHeader(exchange, ROLLINGSTONE_CORRELATION_ID, rollingStoneCorrelationId);
	}

}
