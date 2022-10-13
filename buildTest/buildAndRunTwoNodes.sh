#/bin/bash

# delete data folders
rm -Rf ./konsta-non-signer
rm -Rf ./konsta-signer

# back to root
cd ..



# env GOOS=linux GOARCH=amd64

# delete binaries and rebuild
rm -Rf ./konsta-non-signer   
go build -o konsta-non-signer .

rm -Rf ./konsta-signer       
go build -o konsta-signer .


# copy keys
cp -Rf ./buildTest/validator-keys/us-1/ ./buildTest/konsta-signer/
cp -Rf ./buildTest/validator-keys/us-2/ ./buildTest/konsta-non-signer/

# copy binaries
cp ./konsta-signer ./buildTest/konsta-signer/
cp ./konsta-non-signer ./buildTest/konsta-non-signer/

# copy genesises
cp ./genesis.json ./buildTest/konsta-signer/
cp ./genesis.json ./buildTest/konsta-non-signer/

# cd ./buildTest/konsta-signer
#  ./konsta-signer server   --data-dir ./data-dir --chain genesis.json  --libp2p 0.0.0.0:2478  --grpc-address 0.0.0.0:1632  --jsonrpc 0.0.0.0:1541 --seal --sealAndSign

# ////////////////////
# cd ..

# cd ./konsta-non-signer
# ./konsta-non-signer server   --data-dir ./data-dir --chain genesis.json  --libp2p 0.0.0.0:2478  --grpc-address 0.0.0.0:2632  --jsonrpc 0.0.0.0:2541 --seal



# polygon-edge genesis --consensus ibft --ibft-validator=0x86B4371B5fF3201596A9835bf055Fca8c75009Ee --ibft-validator=0x09a6ff42E6bC0a8dbb0afe8ba1ce0BAD256fc89F --bootnode=/ip4/10.244.0.181/tcp/1478/p2p/16Uiu2HAm2ESK9DmHMNSouZUAJtMx7vzf1oJXtkGk7ZCHBAHWBwPr   --bootnode=/ip4/10.244.1.76/tcp/1478/p2p/16Uiu2HAmSGr8fC6wwHn8JoEqdhiNAAdv7uexmvQr1NhciHKvNgWr    --premine=0x86B4371B5fF3201596A9835bf055Fca8c75009Ee:1000000000000000000000 --premine=0x09a6ff42E6bC0a8dbb0afe8ba1ce0BAD256fc89F:1000000000000000000000 
