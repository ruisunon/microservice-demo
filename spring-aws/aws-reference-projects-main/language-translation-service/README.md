### Language Translation Service | Amazon Translate

Language translation service provides an opportunity to target an international audience along with fulfilling various use-cases of localizing website or application content for diverse set of users, translating large volumes of text for analysis, and enabling cross-lingual communication between users.

This proof-of-concept exposes REST API endpoints capable of translating text written in a source language to a specified target language. The inital cURL command given below can be used to list all supported languages that can be used for language-translation process and the latter command to perform language translation.


```
curl --location --request GET 'http://localhost:8080/translate/language'
```
```
curl --location --request POST 'http://localhost:8080/translate' \
--header 'Content-Type: application/json' \
--data-raw '{
    "text" : "How are you?",
    "sourceLanguageCode" : "en",
    "targetLanguageCode" : "hi"
}'
```

----

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
   "Version":"2012-10-17",
   "Id":"language-translation-service-policy",
   "Statement":[
      {
         "Sid":"textTranslationPermission",
         "Effect":"Allow",
         "Action":[
            "translate:TranslateText",
            "translate:ListLanguages"
         ],
         "Resource":"*"
      }
   ]
}
```
* Go to `application.properties` file and replace below values with the above recieved security credentials
```
com.behl.aws.access-key=<AWS-IAM-ACCESS-KEY-ID>
com.behl.aws.secret-access-key=<AWS-IAM-SECRET-ACCESS-KEY>
```
* Run the application

```
mvn spring-boot:run &
```

* Sample cURL to list all supported languages that can be used for language-translation

```
curl --location --request GET 'http://localhost:8080/translate/language'
```

* Sample cURL to perform language translation

```
curl --location --request POST 'http://localhost:8080/translate' \
--header 'Content-Type: application/json' \
--data-raw '{
    "text" : "How are you?",
    "sourceLanguageCode" : "en",
    "targetLanguageCode" : "hi"
}'
```
#### Alternative Setup Guide through Docker
* Create an image out of the Dockerfile

```
docker build -t language-translation-service .                        
```
* Start a container from the above created image

```
docker run -d -p 8080:8080 language-translation-service
```
* To view the logs of the started container
```
docker container logs -f <container-id>
```
* To stop the container
```
docker container stop <container-id>
```
