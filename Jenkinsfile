node {
    
    def root = tool name: 'go-1.14.3', type:'go'
    
    stage('Checkout Code') {
         checkout scm
    }
    
    stage('Build'){
        withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
                sh "go version"
         }
    }
    stage('Test'){
       withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
                sh "go test -v"
         }
    }
}
