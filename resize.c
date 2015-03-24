/*
 ============================================================================
 Name        : resize.c
 Author      : Centny
 ============================================================================
 */

#include <stdio.h>
#include <stdlib.h>
#include <opencv/cv.h>
#include <opencv/highgui.h>
/*
Args
 fp:the target image file path.
 sp:the save path for resized image.
 w:target width.
 h:target height.
Return
 0:success
 -1:the target file not found.
 -2:can save the file to sp.
*/
int resize(char* fp, char* sp, int w, int h) {
	IplImage* src = NULL;
	IplImage* dst = NULL;
	src = cvLoadImage(fp, CV_LOAD_IMAGE_UNCHANGED);
	if (src == NULL) {
		return -1;
	}
	dst = cvCreateImage(cvSize(w, h), src->depth, src->nChannels);
	cvResize(src, dst, CV_INTER_AREA);
	int ret = cvSaveImage(sp, dst, 0);
	cvFree_(src);
	cvFree_(dst);
	if (ret) {
		return 0;
	} else {
		return -2;
	}
}

