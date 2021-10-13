#!/bin/bash
#Start the second node on a different port and specify networkid
geth --datadir "$ethereum_home/sifnode_test-a" --port 30304 --nodiscover --networkid 1234 console 2 > console.log