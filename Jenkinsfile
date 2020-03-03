pipeline {
    agent {
        label 'base'
    }
    stages {
        stage('checkout') {
            steps {
                script {
                    checkout scm
                }
            }
        }
        stage('build') {
            steps {
                script {
                    container('tools') {
                        withCredentials([usernamePassword(credentialsId: "alaudak8s", usernameVariable: "USERNAME",  passwordVariable: "PASSWORD")]) {
                            sh "docker login index.alauda.cn -u ${USERNAME} -p ${PASSWORD}"
                        }
                        sh "docker buildx build --platform linux/amd64,linux/arm64 --tag index.alauda.cn/alaudak8s/hello:build-${env.BUILD_NUMBER} --push . "
                    }
                }
            }
        }
    }
}