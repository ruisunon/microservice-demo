### Automatic Speech Recognition | Speech to Text Conversion | Amazon Transcribe

This proof-of-concept exposes a single REST API endpoint with the ability to perform **Automatic speech recognition** and convert the provided audio file into a **textual representation**. **Amazon Transcribe** along with **AWS S3** is used to achieve the mentioned functionality. The provided audio file is first stored in the source S3 bucket, the stored objects URI is provided when calling the [StartTranscriptionJob API](https://docs.aws.amazon.com/transcribe/latest/APIReference/API_StartTranscriptionJob.html) exposed by the Amazon Transcribe service. Post successfull execution of the job, the generated result is stored in the configured destination S3 bucket in JSON format which is polled by the application and returned to the user. 

The use cases of subtitle generation and language identification can be fulfilled by using this implementation. 

### Setup Guide

* Install `Java 17` or higher and `Maven`, recommended to use [SdkMan](https://sdkman.io)

```
sdk install java 17-open
```

```
sdk install maven
```
* Clone the repository or [clone this subdirectory](https://stackoverflow.com/questions/600079/how-do-i-clone-a-subdirectory-only-of-a-git-repository)

* Create 2 or use existing S3 Buckets to be used for the purpose of storing the input audio file and the latter to store generated textual content result. Alternatively a single S3 Bucket can also be used to store both the input and ouput files, if following this approach it is recommended to seperate both the type of files thorugh different prefixes. [OutputKey field](https://docs.aws.amazon.com/transcribe/latest/APIReference/API_StartTranscriptionJob.html#transcribe-StartTranscriptionJob-request-OutputKey) can be used to achieve this functionality when calling `startTranscriptionJob` API

* Create an IAM user with programmatic access with the below mentioned policy
```
{
    "Version": "2012-10-17",
    "Id": "aws-transcribe-integration-poc-java-spring-boot",
    "Statement": [
        {
            "Sid": "s3Permissions",
            "Effect": "Allow",
            "Action": [
                "s3:PutObject",
                "s3:GetObject"
            ],
            "Resource": [
                "arn:aws:s3:::<INPUT-BUCKET-NAME>/*",
                "arn:aws:s3:::<OUTPUT-BUCKET-NAME>/*"
            ]
        },
        {
            "Sid": "amazonTranscribePermissions",
            "Effect": "Allow",
            "Action": [
                "transcribe:GetTranscriptionJob",
                "transcribe:StartTranscriptionJob"
            ],
            "Resource": "*"
        }
    ]
}
```

* Go to `application.properties` file and replace below values with the above received security credentials and S3 bucket names
```
com.behl.aws.access-key=<AWS-IAM-ACCESS-KEY-ID>
com.behl.aws.secret-access-key=<AWS-IAM-SECRET-ACCESS-KEY>
```
```
com.behl.aws.s3.input-bucket-name=<INPUT-BUCKET-NAME>
com.behl.aws.s3.output-bucket-name=<OUTPUT-BUCKET-NAME>
```
* Run the application

```
mvn spring-boot:run &
```

* Sample cURL to test the sole exposed API endpoint

```
curl --location --request POST 'http://localhost:8080/audio/transcribe' \
--form 'file=@"path-to-file"'
```
#### Alternative Setup Guide through Docker
* Create an image out of the Dockerfile

```
docker build -t automatic-speech-recognition .                        
```
* Start a container from the above created image

```
docker run -d -p 8080:8080 automatic-speech-recognition
```
* To view the logs of the started container
```
docker container logs -f <container-id>
```
* To stop the container
```
docker container stop <container-id>
```
