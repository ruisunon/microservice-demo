package com.behl.content.moderator.utils;

import org.springframework.util.MimeTypeUtils;
import org.springframework.web.multipart.MultipartFile;

public class FileUtil {

	private FileUtil() {
	}

	/**
	 * @param file : MultipartFile instance to evaluate
	 * @return <code>true</code> if content-type of file is image/png,
	 *         <code>false</code> if not.
	 */
	public static boolean isPng(final MultipartFile file) {
		return file.getContentType().equalsIgnoreCase(MimeTypeUtils.IMAGE_PNG_VALUE);
	}

	/**
	 * @param file : MultipartFile instance to evaluate
	 * @return <code>true</code> if content-type of file is image/jpeg,
	 *         <code>false</code> if not.
	 */
	public static boolean isJpeg(final MultipartFile file) {
		return file.getContentType().equalsIgnoreCase(MimeTypeUtils.IMAGE_JPEG_VALUE);
	}

}
