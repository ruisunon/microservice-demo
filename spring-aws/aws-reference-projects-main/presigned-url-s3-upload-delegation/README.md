### S3 Presigned URL to delegate the object upload responsibility to the client application

This architecture can be used to offload the backend server from the task of large object uploadation to S3 by generating presigned URL with embedded temporary credentials and delegating this responsibility to the consuming client, hence reducing the CPU Utilization and RAM usage. Using the Presigned URLs the client can directly upload the objects to S3 without sending the file to the backend server, the client inherits all the permissions attached to the IAM user/role that generated the presigned URL.

[Reference Article](https://docs.aws.amazon.com/AmazonS3/latest/userguide/PresignedUrlUploadObject.html)

---

#### Architecture Diagram

![presigned-url-s3-upload-delegation-architecture-diagram](https://user-images.githubusercontent.com/69693621/198973959-e714d483-a770-470e-9649-784b32eff65f.jpg)

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

* Create or use an existing S3 Bucket which is to be used as a destination to upload objects to.

* Create an IAM user with programmatic access enabled with the below mentioned policy attached

```
{
   "Version":"2012-10-17",
   "Id":"S3PresignedObjectUploadDelegation",
   "Statement":[
      {
         "Sid":"S3ObjectUploadOperation",
         "Effect":"Allow",
         "Action":[
            "s3:PutObject"
         ],
         "Resource":"arn:aws:s3:::<bucket-name-goes-here>/*"
      }
   ]
}
```
* Go to `application.properties` file and replace below values with the above received S3 bucket name and security credentials
```
com.behl.aws.access-key=<AWS-IAM-ACCESS-KEY-ID>
com.behl.aws.secret-access-key=<AWS-IAM-SECRET-ACCESS-KEY>
```
```
com.behl.aws.s3.bucket-name=<S3-BUCKET-NAME>
com.behl.aws.s3.presigned-url.expiration-time=2
```

* Run the application

```
mvn spring-boot:run &
```

* Sample cURL to access the sole exposed API endpoint

```
curl --location --request GET 'http://localhost:8080/upload/{objectKey}'
```

#### Alternative Setup Guide through Docker
* Create an image out of the Dockerfile

```
docker build -t presigned-url-s3-upload-delegation-application .                        
```
* Start a container from the above created image

```
docker run -d -p 8080:8080 presigned-url-s3-upload-delegation-application
```
* To view the logs of the started container
```
docker container logs -f <container-id>
```
* To stop the container
```
docker container stop <container-id>
```
