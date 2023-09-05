package com.alexander.bryksin.microservive.springwebfluxgrpc.dto;

import javax.validation.constraints.DecimalMin;
import java.math.BigDecimal;

public record WithdrawBalanceDto(@DecimalMin(value = "0.0") BigDecimal amount) {
}
