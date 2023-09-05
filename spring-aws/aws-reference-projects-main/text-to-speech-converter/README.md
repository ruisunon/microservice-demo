### Text To Speech Converter | Amazon Polly

Text to speech convertion is an important requirement in the domains of **content creation** through blogs, **E-learning portals**, and **automated telephonic responses**. Amazon Polly service makes use of deep learning technologies to generate natural-sounding human audio speeches from given textual content in different available voices.

This proof-of-concept built with Java Spring-boot exposes a single **REST API endpoint** with the ability to generate an Audio-speech in **MP3 format** from provided Textual content.

---

https://user-images.githubusercontent.com/69693621/156869860-979608e2-0621-4a46-aeda-cbbeb4bcacde.mov

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
    "Id": "text-to-speech-converter-permissions",
    "Statement": [
        {
            "Sid": "TextToSpeechConversion",
            "Effect": "Allow",
            "Action": [
                "polly:SynthesizeSpeech",
                "polly:StartSpeechSynthesisTask",
                "polly:DescribeVoices"
            ],
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
curl --location --request POST 'http://localhost:8080/v1/generate/audio' \
--header 'Content-Type: application/json' \
--data-raw '{
  "text": "Text to convert"
}'
```
#### Alternative Setup Guide through Docker
* Create an image out of the Dockerfile

```
docker build -t text-to-speech-converter .                        
```
* Start a container from the above created image

```
docker run -d -p 8080:8080 text-to-speech-converter
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
### References
* [Documentation](https://docs.aws.amazon.com/polly/latest/dg/how-text-to-speech-works.html)
* [Available voices in Amazon Polly](https://docs.aws.amazon.com/polly/latest/dg/voicelist.html)
* [Supported Languages in Amazon Polly](https://docs.aws.amazon.com/polly/latest/dg/SupportedLanguage.html)
