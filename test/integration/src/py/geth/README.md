### install Geth on Mac

```shell
brew tap ethereum/ethereum
brew install ethereum
```

### Create ethereum_home

```shell
cd ~
mkdir ethereum
echo 'export ethereum_home=$HOME/ethereum' >> ~/.bash_profile
source ~/.bash_profile
```

### Create Genesis file

```shell
cd $ethereum_home
vi genesis.json
```
copy and save the follow `json`.

```json
{
    "nonce": "0x0000000000000042",
    "timestamp": "0x0",
    "parentHash": "0x0000000000000000000000000000000000000000000000000000000000000000",
    "extraData": "0x",
    "gasLimit": "0x8000000",
    "difficulty": "0x400",
    "config":{},
    "mixhash": "0x0000000000000000000000000000000000000000000000000000000000000000",
    "coinbase": "0x3333333333333333333333333333333333333333",
    "alloc": {}
}
```
### run local test
* start the console `./start_the_console.sh`
* initialise the block `./initialise_the_block.sh`
* create a 2nd node `./create_second_node.sh`
* #Start the second node on a different port and specify networkid `./start_the_second_console.sh`
