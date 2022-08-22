# /bin/bash


# compile 
go build -o node1

go build -o node2

# copy and replace into respected folders

cp -Rf ./node1 blockchain1/

cp -Rf ./node2 blockchain2/

# run both 

