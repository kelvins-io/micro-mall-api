wget https://dl.google.com/go/go1.16.15.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.16.15.linux-amd64.tar.gz

cd ~ || exit
mkdir go
# shellcheck disable=SC1012
echo export GOROOT=/usr/local/go\nexport GOBIN="$GOROOT"/bin\nexport GOPROXY=https://goproxy.cn,https://goproxy.io,direct\nexport GOPATH=~/go\nexport PATH="$PATH":"$HOME"/bin:"$GOBIN":"$GOROOT" >>..bash_profile

go version
