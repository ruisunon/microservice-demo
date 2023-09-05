### Image Content Moderation | Amazon Rekognition

Content moderation is an important requirement in social media, broadcast media, advertising, and e-commerce applications wherein inappropriate, unwanted, suggestive, or offensive content is required to be identified prior to processing and taken care of as per business requirements.

This proof-of-concept exposes a single REST API endpoint to set/update user's current profile picture. **Amazon Rekognition** has been integrated in the service layer to detect any inappropriate, suggestive, unwanted or offensive content in the image being uploaded, If any `ModerationLabels` are detected the API returns `HttpStatus.NOT_ACCEPTABLE`. On successful content moderation evaluation, `HttpStatus.OK` is returned by the endpoint.

----

https://user-images.githubusercontent.com/69693621/198328377-def1e1b8-af53-461b-b251-a489d9988ba6.mov

---

### Setup Guide

* Install `Java 17` or higher and `Maven`, recommended to use [SdkMan](https://sdkman.io)

```
sdk install java 17-open
```

```
sdk install maven
```

* Clone the repository or [clone this subdirectory](https://stackoverflow.com/questions/600079/how-do-i-clone-a-subdirectory-only-of-a-git-repository)

* Create an IAM user with programmatic access with the below mentioned policy

```
{
    "Version": "2012-10-17",
    "Id": "rekognition-detect-moderation-labels-permission",
    "Statement": [
        {
            "Sid": "detectModerationLabels",
            "Effect": "Allow",
            "Action": "rekognition:DetectModerationLabels",
            "Resource": "*"
        }
    ]
}
```
* Go to `application.properties` file and replace below values with the above received security credentials
```
com.behl.aws.access-key=<AWS-IAM-ACCESS-KEY-ID>
com.behl.aws.secret-access-key=<AWS-IAM-SECRET-ACCESS-KEY>
```
* Run the application

```
mvn spring-boot:run &
```

* Sample cURL to test the sole exposed API endpoint

```
curl --location --request POST 'http://localhost:8080/users/profile/image' \
--form 'file=@"/path/to/file"'
```
#### Alternative Setup Guide through Docker
* Create an image out of the Dockerfile

```
docker build -t image-content-moderation-application .                        
```
* Start a container from the above created image

```
docker run -d -p 8080:8080 image-content-moderation-application
```
* To view the logs of the started container
```
docker container logs -f <container-id>
```
* To stop the container
```
docker container stop <container-id>
```
---

#### References
* [Documentation](https://docs.aws.amazon.com/rekognition/latest/dg/moderation.html)
* [Explanation video](https://www.youtube.com/watch?v=U3nsR1yyxKk)
