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