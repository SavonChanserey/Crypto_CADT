# Crypto_CADT

# from Week3/lab1

cd /Users/savonchanserey/path/to/Week3/lab1

# change module name to include Week3/lab1

go mod edit -module=github.com/SavonChanserey/Crypto/Week3/lab1
go mod edit -replace=github.com/SavonChanserey/Crypto=./

# tidy
go mod tidy

# How to run these code

go run main.go -mode md5 -hash 6a85dfd77d9cb35770c9dc6728d73d3f -wordlist nord_vpn.txt -out verbose_md5.txt

go run main.go -mode sha1 -hash aa1c7d931cf140bb35a5a16adeb83a551649c3b9 -wordlist nord_vpn.txt -out verbose_sha1.txt

go run main.go -mode sha512 -hash 485f5c36c6f8474f53a3b0e361369ee3e32ee0de2f368b87b847dd23cb284b316bb0f026ada27df76c31ae8fa8696708d14b4d8fa352dbd8a31991b90ca5dd38 -wordlist nord_vpn.txt -out verbose_sha512.txt


# Push to github

cd Crypto_CADT
git init 
git add .
git commit -m "add"
git remote add origin https://github.com/SavonChanserey/Crypto_CADT.git
git branch -M main
git pull origin main --rebase
git push -u origin main