pipeline {
    agent {label 'Node-1'}
    tools {
         go 'go-1.14.3'
     }
    stages {
        stage ('checkout code') {
            steps {
                checkout scm
            }
        }
        stage('Test') {
            steps {
                        withEnv(["PATH+EXTRA=${HOME}/go/bin"]){
                        sh 'go test -v'
                    }
                }
            }

        stage('Test Coverage') {
            steps {
                        withEnv(["PATH+EXTRA=${HOME}/go/bin"]){
                        sh 'go test -coverprofile cover_repport.out'
                    }
                }
            }
        }
    post {
            always {
                 emailext attachLog: true,
                 body: "${currentBuild.currentResult}: Job ${env.JOB_NAME} build ${env.BUILD_NUMBER}\n More info at: ${env.BUILD_URL}",
                 to: "sandesh.gaikwad@afourtech.com", 
                 subject: "Jenkins Build ${currentBuild.currentResult}: Job ${env.JOB_NAME} - ${env.BUILD_NUMBER}"
            }
    }
}
