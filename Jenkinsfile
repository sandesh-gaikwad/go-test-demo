node {
    stage('Checkout Code') {
        checkout scm
    }
    stage('Build'){
        sh "go version"
    }
    stage('Test'){
        sh "go test -v"
    }
}
