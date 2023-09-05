package com.behl.content.moderator.utils;

import java.util.List;
import java.util.stream.Collectors;

import com.amazonaws.services.rekognition.model.ModerationLabel;

public class ResponseUtil {

	private ResponseUtil() {
	}

	public static String failedContentModerationResponse(final List<ModerationLabel> moderationLabels) {
		return "Our Systems Detected that the image may contain " + String.join(", ", moderationLabels.parallelStream()
				.map(moderationLabel -> moderationLabel.getName()).collect(Collectors.toList())
				+ ". Therefore we are unable to update your profile image at this time, Kindly raise a help desk ticket if you feel this was shown in error");
	}

}