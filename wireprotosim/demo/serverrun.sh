clear
. $(dirname ${BASH_SOURCE})/util.sh

desc "Examine server dir structure..."
run "tree ../server"
clear

desc "Initialize Server"
run "go run ../server/main.go"
clear