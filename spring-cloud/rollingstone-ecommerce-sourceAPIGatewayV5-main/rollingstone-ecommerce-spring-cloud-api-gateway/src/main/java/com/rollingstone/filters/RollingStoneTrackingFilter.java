package com.rollingstone.filters;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.cloud.gateway.filter.GatewayFilterChain;
import org.springframework.cloud.gateway.filter.GlobalFilter;
import org.springframework.core.annotation.Order;
import org.springframework.http.HttpHeaders;
import org.springframework.stereotype.Component;
import org.springframework.web.server.ServerWebExchange;
import reactor.core.publisher.Mono;

@Order(1)
@Component
public class RollingStoneTrackingFilter implements GlobalFilter {

	private static final Logger logger = LoggerFactory.getLogger(RollingStoneTrackingFilter.class);

	@Autowired
	FilterUtils rollingStoneFilterUtils;

	@Override
	public Mono<Void> filter(ServerWebExchange exchange, GatewayFilterChain filterChain) {
		HttpHeaders requestHeaders = exchange.getRequest().getHeaders();
		if (isRollingStoneCorrelationIdPresent(requestHeaders)) {
			logger.debug("tmx-correlation-id found in tracking filter: {}. ",
					rollingStoneFilterUtils.getRollingStoneCorrelationId(requestHeaders));
		} else {
			String rollingStoneCorrelationID = generateRollingStoneCorrelationId();
			exchange = rollingStoneFilterUtils.setRollingStoneCorrelationId(exchange, rollingStoneCorrelationID);
			logger.debug("rollingstone-correlation-id generated in tracking filter: {}.", rollingStoneCorrelationID);
		}
		
		return filterChain.filter(exchange);
	}


	private boolean isRollingStoneCorrelationIdPresent(HttpHeaders requestHeaders) {
		if (rollingStoneFilterUtils.getRollingStoneCorrelationId(requestHeaders) != null) {
			return true;
		} else {
			return false;
		}
	}

	private String generateRollingStoneCorrelationId() {
		return java.util.UUID.randomUUID().toString();
	}

}