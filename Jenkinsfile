node {
    stage('Checkout Code') {
        checkout scm
    }
    stage('Build'){
        sh "go build ."
    }
    stage('Test'){
        sh "go test -v"
    }
}