pipeline {
    agent any

    // add timestamps to the console log
    options {
        timestamps()
    }


    stages {

        stage('Clone repo') {
            steps {
                sh "git clone https://github.com/spring-petclinic/spring-petclinic-angular.git"
                sh "git clone https://github.com/spring-petclinic/spring-petclinic-rest.git"
            }
        }

        stage('Build frontend') {
            steps {
                sh "npm uninstall -g angular-cli @angular/cli"
                sh "npm cache clean"
                sh "npm install -g @angular/cli@latest"
                sh "cd spring-petclinic-angular"
                sh "npm install --save-dev @angular/cli@latest"
                sh "npm install"
                sh "docker build -t spring-pet-clinic-angular:latest"
                sh "cd ~"

        }

        stage('Deploy') {
            steps {
                sh "docker run --rm -p 8080:8080 spring-pet-clinic-angular:latest"
                sh "docker run -p 9966:9966 springcommunity/spring-petclinic-rest"
                //Then clean the workspace after deployment ignoring node_modules directory
                cleanWs notFailBuild: true, patterns: [[pattern: 'node_modules', type: 'EXCLUDE']]
            }
        }

    }

    //post {
        //always {
            //junit '**/*.xml'
            //cobertura coberturaReportFile: 'frontend/coverage.xml', failNoReports: false
            //cobertura coberturaReportFile: 'blackcards/coverage.xml', failNoReports: false
            //cobertura coberturaReportFile: 'whitecards/coverage.xml', failNoReports: false
            //cobertura coberturaReportFile: 'magicmaker/coverage.xml', failNoReports: false
            //sh "docker-compose down || true"
        //}
    //}
}