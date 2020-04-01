
pipeline {
    agent {
        label 'base'
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
                    container('tools') {
                        sh "docker ps -a" 
                    }
                }
            }
        }
    }
}