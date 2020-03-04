
pipeline {
    agent {
        label 'base-arm'
    }
    stages {
        stage('checkout') {
            steps {
                script {
                    library "alauda-cicd@feat/multi-arch-build"
                    checkout scm
                }
            }
        }
        stage('build') {
            steps {
                script {
                    container('tools') {
                        // withCredentials([usernamePassword(credentialsId: "cpaas-system-global-credentials-alaudak8s", usernameVariable: "USERNAME",  passwordVariable: "PASSWORD")]) {
                        //     sh "echo '${USERNAME} ${PASSWORD}' > hell  && cat hell"
                        //     sh "docker login index.alauda.cn -u ${USERNAME} -p ${PASSWORD}"
                        // }
                        // sh "docker buildx build --platform linux/amd64,linux/arm64 --tag index.alauda.cn/alaudak8s/hello:build-${env.BUILD_NUMBER} --push . "
                        def image = deploy.dockerBuildWithRegister(
                            dockerfile: "Dockerfile",
                            address: "alaudak8s/hello",
                            register: "index.alauda.cn",
                            credentialsId: "alauda-index-k8s",
                            tag: env.BUILD_NUMBER
                        ).setArg("commit_id", env.COMMIT_ID)
                        image.multiArch()
                        // try {
                            
                        // } catch (exc) {
                        //     input "wait? ${exc}"
                        // }
                        
                    }
                }
            }
        }
    }
}