package com.behl.scheduler.utility;

import java.time.LocalDateTime;

import lombok.extern.log4j.Log4j2;

@Log4j2
public class CronExpressionGenerator {

	public static String generate(final LocalDateTime timestamp) {
		String cronExpression = "cron(MINUTE HOUR DAY_OF_MONTH MONTH ? YEAR)"
				.replaceAll("MINUTE", String.valueOf(timestamp.getMinute()))
				.replaceAll("HOUR", String.valueOf(timestamp.getHour()))
				.replaceAll("DAY_OF_MONTH", String.valueOf(timestamp.getDayOfMonth()))
				.replaceAll("MONTH", String.valueOf(timestamp.getMonthValue()))
				.replaceAll("YEAR", String.valueOf(timestamp.getYear()));
		log.info("Cron expression '{}' generated against timestamp {}", cronExpression, timestamp);
		return cronExpression;
	}

}