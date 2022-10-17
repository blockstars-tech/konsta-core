# read latest version number
cd ./buildTest/us1

currentKonsta = ls -d konsta*
currentNumber = ${currentKonsta#konsta}
# build with number + 1
cd ../..

baseCommand=go build -o ./macBuilds/konsta

"${baseCommand[@]}""${currentNumber[@]}" 
# clear old data 
cd ./buildTest/us1/
rm -Rf ./data-dir/blockchain ./data-dir/consensus/metadata ./data-dir/consensus/snapshots ./data-dir/trie
rm -Rf konsta*
cd ../..
rm -Rf ./data-dir/blockchain ./data-dir/consensus/metadata ./data-dir/consensus/snapshots ./data-dir/trie
rm -Rf konsta*

cd ../..
# cpy new binaries 
cp ./macBuilds/konsta"${currentNumber[@]}" ./buildTest/us1/
cp ./macBuilds/konsta"${currentNumber[@]}" ./buildTest/us2/

cd ./buildTest/us1/

./konsta"${currentNumber[@]}" server   --data-dir ./data-dir --chain genesis.json  --libp2p 0.0.0.0:2478  --grpc-address 0.0.0.0:1632  --jsonrpc 0.0.0.0:1541 --seal --log-level ERROR



