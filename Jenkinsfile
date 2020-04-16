pipeline {
  agent {
        label "go"
      }
    stages {
      stage('Test') {
        steps {
            sh 'go test'
        }
      }
    }
}
