# read latest version number

rm ./macBuilds/konsta

go build -o ./macBuilds/konsta

# "${baseCommand[@]}""${currentNumber[@]}" 
# clear old data 
cd ./buildTest/us-1/
rm -Rf ./data-dir/blockchain ./data-dir/consensus/metadata ./data-dir/consensus/snapshots ./data-dir/trie
rm -Rf ./konsta
cd ../..
cd ./buildTest/us-2/
rm -Rf ./data-dir/blockchain ./data-dir/consensus/metadata ./data-dir/consensus/snapshots ./data-dir/trie
rm -Rf ./konsta

cd ../..
# cpy new binaries 
sleep 5
cp ./macBuilds/konsta ./buildTest/us-1/
cp ./macBuilds/konsta ./buildTest/us-2/

cd ./buildTest/us-1/

./konsta server   --data-dir ./data-dir --chain genesis.json  --libp2p 0.0.0.0:2478  --grpc-address 0.0.0.0:1632  --jsonrpc 0.0.0.0:1541 --seal --log-level ERROR



