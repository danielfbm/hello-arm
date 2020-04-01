
pipeline {
    agent {
        label 'golang'
    }
    stages {
        stage('checkout') {
            steps {
                script {
                    def scmVars = checkout scm
                    echo "scmVars: ${scmVars}"
                }
            }
        }
        stage('build') {
            steps {
                script {
                    container('golang') {
                        sh "go build -o hello ." 
                        sh "ls -lah"
                    }
                }
            }
        }
    }
}