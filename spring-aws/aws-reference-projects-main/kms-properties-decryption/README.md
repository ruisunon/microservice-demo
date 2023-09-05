### Encrypting secret-keys in .properties file using KMS

A POC Spring-boot web application that stores protected keys in application.properties file in an encrypted format. [spring-cloud-config-aws-kms](https://github.com/zalando/spring-cloud-config-aws-kms) is used to decrypt the encrypted-values during application startup by defining the key-id in `bootstrap.yaml` file and the credentials are resolved using the [Default Credential Provider Chain](https://docs.aws.amazon.com/sdk-for-java/v1/developer-guide/credentials.html#credentials-default). Alternatively an IAM Role can be used with the required permission if using an EC2 Instance to run the web app [(Guide)](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/iam-roles-for-amazon-ec2.html).

#### Local Setup and Steps

* Install Java 11+ (recommended to use [SdkMan](https://sdkman.io))

```
sdk install java 11.0.2-open
```
* Install Maven (recommended to use [SdkMan](https://sdkman.io))

```
sdk install maven
```
* Create an IAM user with security credentials (programmatic-access) enabled | [Guide](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_users_create.html)
* Create a KMS Symmetric Key or use an existing one with the above created IAM user registered as a key-user to inherit access to basic encryption actions | [Guide](https://docs.aws.amazon.com/kms/latest/developerguide/create-keys.html)
* Clone the repository or this [individual subdirectory](https://github.com/hardikSinghBehl/aws-java-reference-pocs/blob/main/INDIVIDUAL_FOLDER_CLONE.md)
* Configure the security-credentials of the above created user using [Default Credential Provider Chain](https://docs.aws.amazon.com/sdk-for-java/v1/developer-guide/credentials.html#credentials-default)
* Configure the key-id and the region in bootstrap.properties 

```
aws:
    kms:
        keyId: <Key-id-Goes-Here>
        region: ap-south-1
```
* Generate a JWT Secret key and base64 encode it

```
echo -n 093617ebfa4b9af9700db274ac204ffa34195494d97b9c26c23ad561de817922 | openssl base64
```
* Encrypt the base64 encoded secret key using KMS 

```
aws kms encrypt --key-id <Key-id-Goes-Here> --plaintext <Base-64-encoded-output>
```
* Store the encrypted ciphertext-blob in the application.properties file with the prefix of `{cipher}`

```
com.behl.decryptor.jwt.secret-key={cipher}<Cipher-text-blob-goes-here>
```
* Run the spring-boot application and access the APIs as configured in [AuthenticationController.class](https://github.com/hardikSinghBehl/aws-java-reference-pocs/blob/main/kms-properties-decryption/src/main/java/com/behl/decryptor/controller/AuthenticationController.java)

```
mvn spring-boot:run
```

#### Demonstration Video

https://user-images.githubusercontent.com/69693621/168938415-dc32999f-7731-422e-aa94-95ca3e7ed168.mov



