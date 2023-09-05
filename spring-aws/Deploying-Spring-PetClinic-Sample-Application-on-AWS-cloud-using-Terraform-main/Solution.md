[![Screenshot-2022-12-16-at-21-46-06.png](https://i.postimg.cc/SQ78s3Jm/Screenshot-2022-12-16-at-21-46-06.png)](https://postimg.cc/nMrMSd4W)

# Solution for Subtask I 

We will deploy Spring PetClinic  Application on AWS cloud using Terraform in this project. We will create following AWS resources via Terraform

    * AWS VPC
    * AWS Subnet
    * AWS Internet Gateway
    * AWS Security Groups
    * AWS EC2 Instance

We will provision our EC2 instances via Cloud Init (user-data) which is recommended way of doing it when we compare to remote provisioners.

Firstly of all ,we will need to initiate terraform on our main directory ,terraform will download AWS provider to the local directory.A provider in Terraform is a plugin that enables interaction with an API. This includes Cloud providers and Software-as-a-service providers. The providers are specified in the Terraform configuration code like below

```
provider "aws" {
  # No need to add since we have added all configuration including region,access_key and secret_access_key via aws configure list command 
  # it is located at ~/.aws/credentials
}
```
We will need also to provide our AWS credentials which will authenticate and start to provision AWS resources. I am also using terraform variables ,which help us to reuse it multiple times in our code.

Snippet from Terraform variables.

```
vpc_cidr_block = "10.20.0.0/16"
subnet_cidr_block = "10.20.15.0/24"
availibility_zone = "us-east-1c"
my_source_ip ="0.0.0.0/0"
env_prefix = "petclinic"
instance_type = "t3.micro"
my_public_key_location = "~/.ssh/id_rsa.pub"
ssh_private_key = "~/.ssh/id_rsa"
user_name = "ubuntu"

```
When we run terraform apply command to create an infrastructure on cloud, Terraform creates a state file called “terraform.tfstate”. This State File contains full details of resources in our terraform code. When you modify something on your code and apply it on cloud, terraform will look into the state file, and compare the changes made in the code from that state file and the changes to the infrastructure based on the state file.  

The terraform plan command lets you to preview the actions Terraform would take to modify your infrastructure, or save a speculative plan which you can apply later. The function of terraform plan is speculative: you cannot apply it unless you save its contents and pass them to a terraform apply command.

```
zhajili$ terraform plan
data.aws_ami.ubuntu18: Reading...
data.aws_ami.ubuntu22: Reading...
data.aws_ami.ubuntu22: Read complete after 1s [id=ami-06878d265978313ca]
data.aws_ami.ubuntu18: Read complete after 2s [id=ami-08fdec01f5df9998f]

Terraform used the selected providers to generate the following execution plan. Resource actions are indicated with the following symbols:
  + create

Terraform will perform the following actions:

  # aws_default_route_table.main-route-table will be created
  + resource "aws_default_route_table" "main-route-table" {
      + arn                    = (known after apply)
      + default_route_table_id = (known after apply)
      + id                     = (known after apply)
      + owner_id               = (known after apply)
      + route                  = [
          + {
              + cidr_block                 = "0.0.0.0/0"
              + core_network_arn           = ""
              + destination_prefix_list_id = ""
              + egress_only_gateway_id     = ""
              + gateway_id                 = (known after apply)
              + instance_id                = ""
              + ipv6_cidr_block            = ""
              + nat_gateway_id             = ""
              + network_interface_id       = ""
              + transit_gateway_id         = ""
              + vpc_endpoint_id            = ""
              + vpc_peering_connection_id  = ""
            },
        ]
      + tags                   = {
          + "Name" = "petclinic-main-route-table"
        }
      + tags_all               = {
          + "Name" = "petclinic-main-route-table"
        }
      + vpc_id                 = (known after apply)
    }

  # aws_instance.petclinic_application will be created
  + resource "aws_instance" "petclinic_application" {
      + ami                                  = "ami-06878d265978313ca"
      + arn                                  = (known after apply)
      + associate_public_ip_address          = true
      + availability_zone                    = "us-east-1c"
      + cpu_core_count                       = (known after apply)
      + cpu_threads_per_core                 = (known after apply)
      + disable_api_stop                     = (known after apply)
      + disable_api_termination              = (known after apply)
      + ebs_optimized                        = (known after apply)
      + get_password_data                    = false
      + host_id                              = (known after apply)
      + host_resource_group_arn              = (known after apply)
      + iam_instance_profile                 = (known after apply)
      + id                                   = (known after apply)
      + instance_initiated_shutdown_behavior = (known after apply)
      + instance_state                       = (known after apply)
      + instance_type                        = "t3.micro"
      + ipv6_address_count                   = (known after apply)
      + ipv6_addresses                       = (known after apply)
      + key_name                             = "AWS_KEY_PAIR"
      + monitoring                           = (known after apply)
      + outpost_arn                          = (known after apply)
      + password_data                        = (known after apply)
      + placement_group                      = (known after apply)
      + placement_partition_number           = (known after apply)
      + primary_network_interface_id         = (known after apply)
      + private_dns                          = (known after apply)
      + private_ip                           = (known after apply)
      + public_dns                           = (known after apply)
      + public_ip                            = (known after apply)
      + secondary_private_ips                = (known after apply)
      + security_groups                      = (known after apply)
      + source_dest_check                    = true
      + subnet_id                            = (known after apply)
      + tags                                 = {
          + "Name" = "petclinic-app"
        }
      + tags_all                             = {
          + "Name" = "petclinic-app"
        }
      + tenancy                              = (known after apply)
      + user_data                            = "d0442a9d4e134e82552fc1d72c9cec7b15b5e324"
      + user_data_base64                     = (known after apply)
      + user_data_replace_on_change          = false
      + vpc_security_group_ids               = (known after apply)

      + capacity_reservation_specification {
          + capacity_reservation_preference = (known after apply)

          + capacity_reservation_target {
              + capacity_reservation_id                 = (known after apply)
              + capacity_reservation_resource_group_arn = (known after apply)
            }
        }

      + ebs_block_device {
          + delete_on_termination = (known after apply)
          + device_name           = (known after apply)
          + encrypted             = (known after apply)
          + iops                  = (known after apply)
          + kms_key_id            = (known after apply)
          + snapshot_id           = (known after apply)
          + tags                  = (known after apply)
          + throughput            = (known after apply)
          + volume_id             = (known after apply)
          + volume_size           = (known after apply)
          + volume_type           = (known after apply)
        }

      + enclave_options {
          + enabled = (known after apply)
        }

      + ephemeral_block_device {
          + device_name  = (known after apply)
          + no_device    = (known after apply)
          + virtual_name = (known after apply)
        }

      + maintenance_options {
          + auto_recovery = (known after apply)
        }

      + metadata_options {
          + http_endpoint               = (known after apply)
          + http_put_response_hop_limit = (known after apply)
          + http_tokens                 = (known after apply)
          + instance_metadata_tags      = (known after apply)
        }

      + network_interface {
          + delete_on_termination = (known after apply)
          + device_index          = (known after apply)
          + network_card_index    = (known after apply)
          + network_interface_id  = (known after apply)
        }

      + private_dns_name_options {
          + enable_resource_name_dns_a_record    = (known after apply)
          + enable_resource_name_dns_aaaa_record = (known after apply)
          + hostname_type                        = (known after apply)
        }

      + root_block_device {
          + delete_on_termination = (known after apply)
          + device_name           = (known after apply)
          + encrypted             = (known after apply)
          + iops                  = (known after apply)
          + kms_key_id            = (known after apply)
          + tags                  = (known after apply)
          + throughput            = (known after apply)
          + volume_id             = (known after apply)
          + volume_size           = (known after apply)
          + volume_type           = (known after apply)
        }
    }

  # aws_instance.petclinic_mysql will be created
  + resource "aws_instance" "petclinic_mysql" {
      + ami                                  = "ami-08fdec01f5df9998f"
      + arn                                  = (known after apply)
      + associate_public_ip_address          = true
      + availability_zone                    = "us-east-1c"
      + cpu_core_count                       = (known after apply)
      + cpu_threads_per_core                 = (known after apply)
      + disable_api_stop                     = (known after apply)
      + disable_api_termination              = (known after apply)
      + ebs_optimized                        = (known after apply)
      + get_password_data                    = false
      + host_id                              = (known after apply)
      + host_resource_group_arn              = (known after apply)
      + iam_instance_profile                 = (known after apply)
      + id                                   = (known after apply)
      + instance_initiated_shutdown_behavior = (known after apply)
      + instance_state                       = (known after apply)
      + instance_type                        = "t3.micro"
      + ipv6_address_count                   = (known after apply)
      + ipv6_addresses                       = (known after apply)
      + key_name                             = "AWS_KEY_PAIR"
      + monitoring                           = (known after apply)
      + outpost_arn                          = (known after apply)
      + password_data                        = (known after apply)
      + placement_group                      = (known after apply)
      + placement_partition_number           = (known after apply)
      + primary_network_interface_id         = (known after apply)
      + private_dns                          = (known after apply)
      + private_ip                           = "10.20.15.200"
      + public_dns                           = (known after apply)
      + public_ip                            = (known after apply)
      + secondary_private_ips                = (known after apply)
      + security_groups                      = (known after apply)
      + source_dest_check                    = true
      + subnet_id                            = (known after apply)
      + tags                                 = {
          + "Name" = "petclinic-db"
        }
      + tags_all                             = {
          + "Name" = "petclinic-db"
        }
      + tenancy                              = (known after apply)
      + user_data                            = "c0b7bc898cfdcfe4591e649b92a02e658b1729c7"
      + user_data_base64                     = (known after apply)
      + user_data_replace_on_change          = false
      + vpc_security_group_ids               = (known after apply)

      + capacity_reservation_specification {
          + capacity_reservation_preference = (known after apply)

          + capacity_reservation_target {
              + capacity_reservation_id                 = (known after apply)
              + capacity_reservation_resource_group_arn = (known after apply)
            }
        }

      + ebs_block_device {
          + delete_on_termination = (known after apply)
          + device_name           = (known after apply)
          + encrypted             = (known after apply)
          + iops                  = (known after apply)
          + kms_key_id            = (known after apply)
          + snapshot_id           = (known after apply)
          + tags                  = (known after apply)
          + throughput            = (known after apply)
          + volume_id             = (known after apply)
          + volume_size           = (known after apply)
          + volume_type           = (known after apply)
        }

      + enclave_options {
          + enabled = (known after apply)
        }

      + ephemeral_block_device {
          + device_name  = (known after apply)
          + no_device    = (known after apply)
          + virtual_name = (known after apply)
        }

      + maintenance_options {
          + auto_recovery = (known after apply)
        }

      + metadata_options {
          + http_endpoint               = (known after apply)
          + http_put_response_hop_limit = (known after apply)
          + http_tokens                 = (known after apply)
          + instance_metadata_tags      = (known after apply)
        }

      + network_interface {
          + delete_on_termination = (known after apply)
          + device_index          = (known after apply)
          + network_card_index    = (known after apply)
          + network_interface_id  = (known after apply)
        }

      + private_dns_name_options {
          + enable_resource_name_dns_a_record    = (known after apply)
          + enable_resource_name_dns_aaaa_record = (known after apply)
          + hostname_type                        = (known after apply)
        }

      + root_block_device {
          + delete_on_termination = (known after apply)
          + device_name           = (known after apply)
          + encrypted             = (known after apply)
          + iops                  = (known after apply)
          + kms_key_id            = (known after apply)
          + tags                  = (known after apply)
          + throughput            = (known after apply)
          + volume_id             = (known after apply)
          + volume_size           = (known after apply)
          + volume_type           = (known after apply)
        }
    }

  # aws_internet_gateway.petclinic_vpc_internet_gateway will be created
  + resource "aws_internet_gateway" "petclinic_vpc_internet_gateway" {
      + arn      = (known after apply)
      + id       = (known after apply)
      + owner_id = (known after apply)
      + tags     = {
          + "Name" = "petclinic-internet-gateway"
        }
      + tags_all = {
          + "Name" = "petclinic-internet-gateway"
        }
      + vpc_id   = (known after apply)
    }

  # aws_security_group.petclinic_sg will be created
  + resource "aws_security_group" "petclinic_sg" {
      + arn                    = (known after apply)
      + description            = "Allow Inbound SSH and TCP port 8080 traffic"
      + egress                 = [
          + {
              + cidr_blocks      = [
                  + "0.0.0.0/0",
                ]
              + description      = ""
              + from_port        = 0
              + ipv6_cidr_blocks = []
              + prefix_list_ids  = []
              + protocol         = "-1"
              + security_groups  = []
              + self             = false
              + to_port          = 0
            },
        ]
      + id                     = (known after apply)
      + ingress                = [
          + {
              + cidr_blocks      = [
                  + "0.0.0.0/0",
                ]
              + description      = "HTTP from VPC"
              + from_port        = 8080
              + ipv6_cidr_blocks = []
              + prefix_list_ids  = []
              + protocol         = "tcp"
              + security_groups  = []
              + self             = false
              + to_port          = 8080
            },
          + {
              + cidr_blocks      = [
                  + "0.0.0.0/0",
                ]
              + description      = "MYSQL port"
              + from_port        = 3306
              + ipv6_cidr_blocks = []
              + prefix_list_ids  = []
              + protocol         = "tcp"
              + security_groups  = []
              + self             = false
              + to_port          = 3306
            },
          + {
              + cidr_blocks      = [
                  + "0.0.0.0/0",
                ]
              + description      = "SSH from VPC"
              + from_port        = 22
              + ipv6_cidr_blocks = []
              + prefix_list_ids  = []
              + protocol         = "tcp"
              + security_groups  = []
              + self             = false
              + to_port          = 22
            },
        ]
      + name                   = "PETCLINIC_SG"
      + name_prefix            = (known after apply)
      + owner_id               = (known after apply)
      + revoke_rules_on_delete = false
      + tags                   = {
          + "Name" = "petclinic-sg"
        }
      + tags_all               = {
          + "Name" = "petclinic-sg"
        }
      + vpc_id                 = (known after apply)
    }

  # aws_subnet.my_app_subnet-1 will be created
  + resource "aws_subnet" "my_app_subnet-1" {
      + arn                                            = (known after apply)
      + assign_ipv6_address_on_creation                = false
      + availability_zone                              = "us-east-1c"
      + availability_zone_id                           = (known after apply)
      + cidr_block                                     = "10.20.15.0/24"
      + enable_dns64                                   = false
      + enable_resource_name_dns_a_record_on_launch    = false
      + enable_resource_name_dns_aaaa_record_on_launch = false
      + id                                             = (known after apply)
      + ipv6_cidr_block_association_id                 = (known after apply)
      + ipv6_native                                    = false
      + map_public_ip_on_launch                        = false
      + owner_id                                       = (known after apply)
      + private_dns_hostname_type_on_launch            = (known after apply)
      + tags                                           = {
          + "Name" = "petclinic"
        }
      + tags_all                                       = {
          + "Name" = "petclinic"
        }
      + vpc_id                                         = (known after apply)
    }

  # aws_vpc.petclinic_vpc will be created
  + resource "aws_vpc" "petclinic_vpc" {
      + arn                                  = (known after apply)
      + cidr_block                           = "10.20.0.0/16"
      + default_network_acl_id               = (known after apply)
      + default_route_table_id               = (known after apply)
      + default_security_group_id            = (known after apply)
      + dhcp_options_id                      = (known after apply)
      + enable_classiclink                   = (known after apply)
      + enable_classiclink_dns_support       = (known after apply)
      + enable_dns_hostnames                 = true
      + enable_dns_support                   = true
      + enable_network_address_usage_metrics = (known after apply)
      + id                                   = (known after apply)
      + instance_tenancy                     = "default"
      + ipv6_association_id                  = (known after apply)
      + ipv6_cidr_block                      = (known after apply)
      + ipv6_cidr_block_network_border_group = (known after apply)
      + main_route_table_id                  = (known after apply)
      + owner_id                             = (known after apply)
      + tags                                 = {
          + "Name" = "petclinic-vpc"
        }
      + tags_all                             = {
          + "Name" = "petclinic-vpc"
        }
    }

Plan: 7 to add, 0 to change, 0 to destroy.

─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────


```

Output shows that which resource will be created when when apply terraform code. Let's apply our code and check the output.

```

Plan: 7 to add, 0 to change, 0 to destroy.

Do you want to perform these actions?
  Terraform will perform the actions described above.
  Only 'yes' will be accepted to approve.

  Enter a value: yes

aws_vpc.petclinic_vpc: Creating...
aws_vpc.petclinic_vpc: Still creating... [10s elapsed]
aws_vpc.petclinic_vpc: Creation complete after 15s [id=vpc-00d7a3453547733a8]
aws_internet_gateway.petclinic_vpc_internet_gateway: Creating...
aws_subnet.my_app_subnet-1: Creating...
aws_security_group.petclinic_sg: Creating...
aws_subnet.my_app_subnet-1: Creation complete after 2s [id=subnet-095271cacb7ccf4b1]
aws_internet_gateway.petclinic_vpc_internet_gateway: Creation complete after 2s [id=igw-09a31402e6275f2b5]
aws_default_route_table.main-route-table: Creating...
aws_default_route_table.main-route-table: Creation complete after 2s [id=rtb-042d039135b87bc7d]
aws_security_group.petclinic_sg: Creation complete after 6s [id=sg-0e71b043d0bb9cae9]
aws_instance.petclinic_application: Creating...
aws_instance.petclinic_mysql: Creating...
aws_instance.petclinic_mysql: Still creating... [10s elapsed]
aws_instance.petclinic_application: Still creating... [10s elapsed]
aws_instance.petclinic_application: Creation complete after 15s [id=i-03082cfd11418ca8b]
aws_instance.petclinic_mysql: Creation complete after 15s [id=i-011d089f60d685f55]

Apply complete! Resources: 7 added, 0 changed, 0 destroyed.

```
From output we can see that 7 resources have been created ,let's login to our AWS and check EC2 instances ,then we will ssh to our app server and check status.

[![Screenshot-2022-12-16-at-20-54-08.png](https://i.postimg.cc/wv5nswYS/Screenshot-2022-12-16-at-20-54-08.png)](https://postimg.cc/gXJgCqJD)

As we can see that petclinic-app is still Initializing ,the reason for that is mvn is still downloading dependencies and at the end application will be started with MySQL profile.

```
ubuntu@APPLICATION:~$ ps ax | grep java
   9028 ?        Sl     1:27 /usr/bin/java -classpath /usr/share/maven/boot/plexus-classworlds-2.x.jar -Dclassworlds.conf=/usr/share/maven/bin/m2.conf -Dmaven.home=/usr/share/maven -Dlibrary.jansi.path=/usr/share/maven/lib/jansi-native -Dmaven.multiModuleProjectDirectory=/home/ubuntu/spring-petclinic org.codehaus.plexus.classworlds.launcher.Launcher spring-boot:run -Dspring-boot.run.profiles=mysql
   9192 ?        Sl     0:09 /usr/lib/jvm/java-11-openjdk-amd64/bin/java -Xverify:none -XX:TieredStopAtLevel=1 -cp /home/ubuntu/spring-petclinic/target/classes:/root/.m2/repository/ch/qos/logback/spring-boot-devtools/2.7.3/spring-boot-devtools-2.7.3.jar:/root/.m2/repository/org/springframework/boot/spring-boot/2.7.3/spring-boot-2.7.3.jar:/root/.m2/repository/org/springframework/boot/spring-boot-autoconfigure/2.7.3/spring-boot-autoconfigure-2.7.3.jar org.springframework.samples.petclinic.PetClinicApplication --spring.profiles.active=mysql
   9354 pts/0    R+     0:00 grep --color=auto java

ubuntu@APPLICATION:~$ netstat -tulnp
(Not all processes could be identified, non-owned process info
 will not be shown, you would have to be root to see it all.)
Active Internet connections (only servers)
Proto Recv-Q Send-Q Local Address           Foreign Address         State       PID/Program name    
tcp        0      0 127.0.0.53:53           0.0.0.0:*               LISTEN      -                   
tcp        0      0 0.0.0.0:22              0.0.0.0:*               LISTEN      -                   
tcp6       0      0 :::22                   :::*                    LISTEN      -                   
tcp6       0      0 :::35729                :::*                    LISTEN      -                   
tcp6       0      0 :::8080                 :::*                    LISTEN      -                   
udp        0      0 10.20.15.240:68         0.0.0.0:*                           -                   
udp        0      0 127.0.0.1:323           0.0.0.0:*                           -                   
udp        0      0 127.0.0.53:53           0.0.0.0:*                           -                   
udp6       0      0 ::1:323                 :::*                                -                  


```

We can check EC2 Cloud Init logs in following directory.

```
/var/log/cloud-init-output.log

```

```
19:54:19.904 [Thread-0] DEBUG org.springframework.boot.devtools.restart.classloader.RestartClassLoader - Created RestartClassLoader org.springframework.boot.devtools.restart.classloader.RestartClassLoader@7f760151


              |\      _,,,--,,_
             /,`.-'`'   ._  \-;;,_
  _______ __|,4-  ) )_   .;.(__`'-'__     ___ __    _ ___ _______
 |       | '---''(_/._)-'(_\_)   |   |   |   |  |  | |   |       |
 |    _  |    ___|_     _|       |   |   |   |   |_| |   |       | __ _ _
 |   |_| |   |___  |   | |       |   |   |   |       |   |       | \ \ \ \
 |    ___|    ___| |   | |      _|   |___|   |  _    |   |      _|  \ \ \ \
 |   |   |   |___  |   | |     |_|       |   | | |   |   |     |_    ) ) ) )
 |___|   |_______| |___| |_______|_______|___|_|  |__|___|_______|  / / / /
 ==================================================================/_/_/_/

:: Built with Spring Boot :: 2.7.3


2022-12-16 19:54:20.421  INFO 9192 --- [  restartedMain] o.s.s.petclinic.PetClinicApplication     : Starting PetClinicApplication using Java 11.0.17 on APPLICATION with PID 9192 (/home/ubuntu/spring-petclinic/target/classes started by root in /home/ubuntu/spring-petclinic)
2022-12-16 19:54:20.427  INFO 9192 --- [  restartedMain] o.s.s.petclinic.PetClinicApplication     : The following 1 profile is active: "mysql"
2022-12-16 19:54:20.543  INFO 9192 --- [  restartedMain] .e.DevToolsPropertyDefaultsPostProcessor : Devtools property defaults active! Set 'spring.devtools.add-properties' to 'false' to disable
2022-12-16 19:54:20.543  INFO 9192 --- [  restartedMain] .e.DevToolsPropertyDefaultsPostProcessor : For additional web related logging consider setting the 'logging.level.web' property to 'DEBUG'
2022-12-16 19:54:21.744  INFO 9192 --- [  restartedMain] .s.d.r.c.RepositoryConfigurationDelegate : Bootstrapping Spring Data JPA repositories in DEFAULT mode.
2022-12-16 19:54:21.802  INFO 9192 --- [  restartedMain] .s.d.r.c.RepositoryConfigurationDelegate : Finished Spring Data repository scanning in 43 ms. Found 2 JPA repository interfaces.
2022-12-16 19:54:22.494  INFO 9192 --- [  restartedMain] o.s.b.w.embedded.tomcat.TomcatWebServer  : Tomcat initialized with port(s): 8080 (http)
2022-12-16 19:54:22.509  INFO 9192 --- [  restartedMain] o.apache.catalina.core.StandardService   : Starting service [Tomcat]
2022-12-16 19:54:22.512  INFO 9192 --- [  restartedMain] org.apache.catalina.core.StandardEngine  : Starting Servlet engine: [Apache Tomcat/9.0.65]
2022-12-16 19:54:22.591  INFO 9192 --- [  restartedMain] o.a.c.c.C.[Tomcat].[localhost].[/]       : Initializing Spring embedded WebApplicationContext
2022-12-16 19:54:22.593  INFO 9192 --- [  restartedMain] w.s.c.ServletWebServerApplicationContext : Root WebApplicationContext: initialization completed in 2041 ms
2022-12-16 19:54:22.751  INFO 9192 --- [  restartedMain] com.zaxxer.hikari.HikariDataSource       : HikariPool-1 - Starting...
2022-12-16 19:54:23.138  INFO 9192 --- [  restartedMain] com.zaxxer.hikari.HikariDataSource       : HikariPool-1 - Start completed.
2022-12-16 19:54:23.149  INFO 9192 --- [  restartedMain] o.s.b.a.h2.H2ConsoleAutoConfiguration    : H2 console available at '/h2-console'. Database available at 'jdbc:mysql://10.20.15.200:3306/petclinic'
2022-12-16 19:54:23.512  INFO 9192 --- [  restartedMain] org.ehcache.core.EhcacheManager          : Cache 'vets' created in EhcacheManager.
2022-12-16 19:54:23.521  INFO 9192 --- [  restartedMain] org.ehcache.jsr107.Eh107CacheManager     : Registering Ehcache MBean javax.cache:type=CacheStatistics,CacheManager=urn.X-ehcache.jsr107-default-config,Cache=vets
2022-12-16 19:54:23.644  INFO 9192 --- [  restartedMain] o.hibernate.jpa.internal.util.LogHelper  : HHH000204: Processing PersistenceUnitInfo [name: default]
2022-12-16 19:54:23.707  INFO 9192 --- [  restartedMain] org.hibernate.Version                    : HHH000412: Hibernate ORM core version 5.6.10.Final
2022-12-16 19:54:23.823  INFO 9192 --- [  restartedMain] o.hibernate.annotations.common.Version   : HCANN000001: Hibernate Commons Annotations {5.1.2.Final}
2022-12-16 19:54:23.932  INFO 9192 --- [  restartedMain] org.hibernate.dialect.Dialect            : HHH000400: Using dialect: org.hibernate.dialect.MySQL57Dialect
2022-12-16 19:54:24.589  INFO 9192 --- [  restartedMain] o.h.e.t.j.p.i.JtaPlatformInitiator       : HHH000490: Using JtaPlatform implementation: [org.hibernate.engine.transaction.jta.platform.internal.NoJtaPlatform]
2022-12-16 19:54:24.598  INFO 9192 --- [  restartedMain] j.LocalContainerEntityManagerFactoryBean : Initialized JPA EntityManagerFactory for persistence unit 'default'
2022-12-16 19:54:25.558  INFO 9192 --- [  restartedMain] o.s.b.d.a.OptionalLiveReloadServer       : LiveReload server is running on port 35729
2022-12-16 19:54:25.569  INFO 9192 --- [  restartedMain] o.s.b.a.e.web.EndpointLinksResolver      : Exposing 13 endpoint(s) beneath base path '/actuator'
2022-12-16 19:54:25.625  INFO 9192 --- [  restartedMain] o.s.b.w.embedded.tomcat.TomcatWebServer  : Tomcat started on port(s): 8080 (http) with context path ''
2022-12-16 19:54:25.645  INFO 9192 --- [  restartedMain] o.s.s.petclinic.PetClinicApplication     : Started PetClinicApplication in 5.725 seconds (JVM running for 6.369)
2022-12-16 19:56:13.815  INFO 9192 --- [nio-8080-exec-1] o.a.c.c.C.[Tomcat].[localhost].[/]       : Initializing Spring DispatcherServlet 'dispatcherServlet'
2022-12-16 19:56:13.816  INFO 9192 --- [nio-8080-exec-1] o.s.web.servlet.DispatcherServlet        : Initializing Servlet 'dispatcherServlet'
2022-12-16 19:56:13.832  INFO 9192 --- [nio-8080-exec-1] o.s.web.servlet.DispatcherServlet        : Completed initialization in 13 ms
Cloud-init v. 22.4.2-0ubuntu0~22.04.1 running 'modules:final' at Fri, 16 Dec 2022 19:57:47 +0000. Up 454.61 seconds.
Cloud-init v. 22.4.2-0ubuntu0~22.04.1 finished at Fri, 16 Dec 2022 19:57:47 +0000. Datasource DataSourceEc2Local.  Up 454.79 seconds

```

# Solution for Subtask II

In this task we are implementing first task via Jenkins ,our Jenkinsfile will contain like that

```
def COLOR_MAP = [
    'SUCCESS': 'good', 
    'FAILURE': 'danger',
]
pipeline {
agent any
tools {
  terraform 'terraform'
}

 stages { 
  stage ('CHECKOUT GIT ') { 
     steps { 
       cleanWs()
       sh  'git clone https://github.com/hacizeynal/Deploying-Spring-PetClinic-Sample-Application-on-AWS-cloud-using-Terraform.git'
      }
      } 
  
  stage ('TERRAFORM INIT') { 
    steps {
    sh '''
    cd Deploying-Spring-PetClinic-Sample-Application-on-AWS-cloud-using-Terraform/
    terraform init
    ''' 
    }
   }
   
  stage ('TERRAFORM APPLY') { 
    steps {
    sh '''
    cd Deploying-Spring-PetClinic-Sample-Application-on-AWS-cloud-using-Terraform/
    terraform apply --auto-approve
    ''' 
    }
   }

  stage ("WAIT TIME TILL DEPLOYMENT") {
    steps{
      sleep time: 300, unit: 'SECONDS'
      echo "Waiting 5 minutes for deployment to complete prior starting health check testing"
    }  
    }

  stage ('CHECK HEALTH STATUS') {
    environment {
      PUBLIC_DYNAMIC_URL = "${sh(script:'cd Deploying-Spring-PetClinic-Sample-Application-on-AWS-cloud-using-Terraform/ && terraform output -raw application_public_public_dns', returnStdout: true).trim()}"
    } 
    steps {
      sh "curl -X GET http://${env.PUBLIC_DYNAMIC_URL}:8080/actuator/health/custom"

      echo "Application is UP running successfully ! :) "
        }
      }
    }

  post {
    always {
        echo 'Slack Notifications.'
        slackSend channel: '#jenkins',
            color: COLOR_MAP[currentBuild.currentResult],
            message: "*${currentBuild.currentResult}:* Job ${env.JOB_NAME} build ${env.BUILD_NUMBER} \n More info at: ${env.BUILD_URL}"
            // message: "Application is running on ${env.PUBLIC_DYNAMIC_URL}"
                }
            }  
  }

```

Since application deployment takes some time I have added 5 minutes wait time before doing health check ,if healh check is successful Slack Notification will be send about Success ,otherwise Failure message will be send to Slack Channel.

```
SUCCESS: Job Petclinic via Terraform build 53
More info at: http://ec2-18-208-199-196.compute-1.amazonaws.com:8080/job/Petclinic%20via%20Terraform/53/

```
Another challenging point is preparing dynamic URL for application ,since it is dynamic we should take URL via terraform output and update Jenkins ENV variable 

```

stage ('CHECK HEALTH STATUS') {
    environment {
      PUBLIC_DYNAMIC_URL = "${sh(script:'cd Deploying-Spring-PetClinic-Sample-Application-on-AWS-cloud-using-Terraform/ && terraform output -raw application_public_public_dns', returnStdout: true).trim()}"
    } 
    steps {
      sh "curl -X GET http://${env.PUBLIC_DYNAMIC_URL}:8080/actuator/health/custom"

      echo "Application is UP running successfully ! :) "

```



