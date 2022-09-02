sudo curl -fsSL https://get.docker.com | bash -s docker --mirror Aliyun
sudo yum install -y yum-utils \
  device-mapper-persistent-data \
  lvm2
sudo yum-config-manager \
    --add-repo \
    http://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo
sudo yum install docker-ce docker-ce-cli containerd.io
sudo systemctl start docker
docker version

# shellcheck disable=SC2046
sudo curl -L https://get.daocloud.io/docker/compose/releases/download/v2.4.1/docker-compose-`uname -s`-`uname -m` > /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
sudo ln -s /usr/local/bin/docker-compose /usr/bin/docker-compose
docker-compose --version
sudo systemctl enable docker.service
echo "添加docker国内镜像源"
sudo cp ./docker_daemon.json /etc/docker/daemon.json
sudo systemctl restart docker.service
echo