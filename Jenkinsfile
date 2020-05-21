pipeline {
    agent {label 'Node-1'}
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
