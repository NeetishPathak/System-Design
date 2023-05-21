clear
. $(dirname ${BASH_SOURCE})/util.sh

desc "Examine client dir structure..."
run "tree ../client"
clear

desc "Initialize Client and send requests"
run "go run ../client/main.go"
clear
