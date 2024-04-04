go
wget https://golang.org/dl/go1.17.6.linux-arm64.tar.gz 
sudo tar -C /usr/local -xzf go1.17.6.linux-arm64.tar.gz 

# wget https://golang.org/dl/go1.18.2.linux-amd64.tar.gz 
# sudo tar -C /usr/local -xzf go1.18.2.linux-amd64.tar.gz

echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc 
source ~/.bashrc
go version
go

cd sib-seal-ec2-test
go mod init sib-seal-ec2-test 
go mod tidy