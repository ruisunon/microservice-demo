### Notification Scheduler Service using [AWS EventBridge](https://aws.amazon.com/eventbridge/) and [SQS](https://aws.amazon.com/sqs/)
POC to implement a cloud based scheduler architecture which enables the user to specify a future timestamp at which the custom message being provided is delivered to the destination application. AWS EventBridge and SQS are being used to achieve this functionality. This architecture provides fault tolerance and works even if the backend application(s) are temporarily down, an advantage over a [Java based Quartz implementation](https://github.com/hardikSinghBehl/quartz-scheduler-daily-mail-subscription-spring-boot).

**NOTE:** [more AWS service(s)](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-targets.html) besides SQS can be used as an target for the EventBridge rule. Simply specify the ARN of the resource to use in the below mentioned key in [application.properties](https://github.com/hardikSinghBehl/aws-java-reference-pocs/blob/main/eventbridge-notification-scheduler/notification-scheduler-service/src/main/resources/application.properties) file of `notification-scheduler-service`
```
com.behl.scheduler.aws.eventbridge.target-arn=<ARN-Goes-Here>
```
#### Services
* **notification-scheduler-service** : Backend application that creates EventBridge rules and registers a SQS target with custom input and cron expression
* **notification-dispatcher-service** : Backend application listening to the configured SQS queue, the application receives the custom input provided when the timestamp specified in the cron expression is reached
#### Sample API Request
Sample cURL to access sole notification-scheduler-service API, the timestamp is to be provided in UTC and would be converted to a cron expression and registered to an EventBridge Rule by the application.
```
curl --location --request POST 'http://localhost:8080/v1/schedule' \
--header 'Content-Type: application/json' \
--data-raw '{
    "timestamp": "2022-08-18T04:34:00.000000Z",
    "emailId": "all@gmail.com",
    "subject": "Clearence Sale",
    "body": "Get upto 50% off on all items"
}'
```
#### Demonstration screen-recording

https://user-images.githubusercontent.com/69693621/185271511-27ca30fb-705c-4558-9322-9bd7766809a5.mp4

----

#### AWS Services and Environment Setup
* Install Java 17 (recommended to use [SdkMan](https://sdkman.io))

```
sdk install java 17-open
```
* Install Maven (recommended to use [SdkMan](https://sdkman.io))

```
sdk install maven
```
* Create the below mentioned AWS Services
  * AWS Event-bus under EventBridge, the default event-bus can also be used | [Reference](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-create-event-bus.html)
  * Standard SQS Queue | [Reference](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-configure-create-queue.html)
* Attach the below statement on the access-policy of the target being used **(SQS Queue created above in this context)**. If using any other AWS Services like lambda, modify the Resource and Action accordingly. The value of `<Rule-Name-Prefix>` below should match the one configured in key `com.behl.scheduler.aws.eventbridge.rule-prefix` in [.properties file](https://github.com/hardikSinghBehl/aws-java-reference-pocs/blob/main/eventbridge-notification-scheduler/notification-scheduler-service/src/main/resources/application.properties)
```
  {
      "Sid": "f852fca8-3f5b-405b-a804-23f848a42076",
      "Effect": "Allow",
      "Principal": {
        "Service": "events.amazonaws.com"
      },
      "Action": "sqs:SendMessage",
      "Resource": "<ARN-of-SQS-Queue>",
      "Condition": {
        "ArnEquals": {
          "aws:SourceArn": "arn:aws:events:<Region>:<Account-ID>:rule/<Rule-Name-Prefix>*"
        }
      }
    }
```
* Create an IAM user with programmatic access enabled and attach the below mentioned policy, download the .csv file containing the user credentials for future reference
```
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "NotificationSchedulerPermissions",
            "Effect": "Allow",
            "Action": [
                "events:DeleteRule",
                "events:PutTargets",
                "events:EnableRule",
                "events:PutRule",
                "events:RemoveTargets",
                "events:DisableRule",
                "sqs:DeleteMessage",
                "sqs:GetQueueUrl",
                "sqs:ReceiveMessage",
                "sqs:GetQueueAttributes",
                "sqs:ListQueueTags"
            ],
            "Resource": [
                "arn:aws:events:<aws-region>:<aws-account-id>:rule/<rule-name-prefix>*",
                "arn:aws:events:<aws-region>:<aws-account-id>:event-bus/<event-bus-name>",
                "arn:aws:sqs:<aws-region>:<aws-account-id>:<sqs-queue-name>"
            ]
        }
    ]
}
```
* Configure appropriate values in both service's `application.properties` file
```
### notification-scheduler-service ###

# AWS IAM Configuration
com.behl.aws.access-key=<Access-key-Goes-Here>
com.behl.aws.secret-access-key=<Secret-access-key-Goes-Here>

# EventBridge Configuration
com.behl.scheduler.aws.eventbridge.event-bus-name=<EventBridge-bus-name>
com.behl.scheduler.aws.eventbridge.target-arn=<target-SQS-Queue-ARN>
com.behl.scheduler.aws.eventbridge.rule-prefix=scheduler_
```
```
### notification-dispatcher-service ###

# AWS IAM Configuration
com.behl.aws.access-key={AWS_ACCESS_KEY:<Access-key-Goes-Here>}
com.behl.aws.secret-access-key={AWS_SECRET_ACCESS_KEY:<Secret-access-key-Goes-Here>}

# SQS Configuration
com.behl.aws.sqs.name=<SQS-Queue-Name-Goes-Here>
```
* Run the below command to start the application in the root directory of both the services
```
mvn spring-boot:run
```
