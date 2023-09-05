

sudo apt update && \
sudo apt install apt-transport-https ca-certificates curl software-properties-common -y && \
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add - && \
sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu focal stable" && \
apt-cache policy docker-ce && \
sudo apt install docker-ce -y && \
sudo apt-get install openjdk-11-jdk -y


curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl" && \
        sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl && \
        curl -LO https://get.helm.sh/helm-v3.7.1-linux-amd64.tar.gz && \
        tar -zxvf  helm-v3.7.1-linux-amd64.tar.gz  && mv ./linux-amd64/helm /usr/local/bin/helm


sudo systemctl status docker

