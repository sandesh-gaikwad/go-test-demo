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
                 emailext body: 'A Test EMail', recipientProviders: [[$class: 'DevelopersRecipientProvider'], [$class: 'RequesterRecipientProvider']], subject: 'Test'
            }
    }
}
