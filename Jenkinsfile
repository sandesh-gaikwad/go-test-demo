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
        }
    post {
            always {
                 emailext attachLog: true,
                 body: "${currentBuild.currentResult}: Job ${env.JOB_NAME} build ${env.BUILD_NUMBER}\n More info at: ${env.BUILD_URL}",
                 //recipientProviders: [developers(), requestor()],
                 to: "sandesh.gaikwad@afourtech.com"    
                 subject: "Jenkins Build ${currentBuild.currentResult}: Job ${env.JOB_NAME}"
            }
    }
}
