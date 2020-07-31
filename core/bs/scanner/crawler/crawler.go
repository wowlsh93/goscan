/*
Copyright Lemon Corp. All Rights Reserved.

Written by hama
*/

package crawler

import (
	"github.com/sirupsen/logrus"
	"github.com/wowlsh93/goscan/common/flogging"
	"github.com/wowlsh93/goscan/core/bs/scanner/config"
	"github.com/wowlsh93/goscan/core/bs/scanner/crawler/eth"
)

type Crawler struct {
	ethRpc   *eth.EthRPC
	logging  *logrus.Logger
	stopChan chan bool

	startingBlock int
	currentBlock  int
	highestBlock  int
}

func New(conf config.Configuration, stop chan bool, startingBlock int) Crawler {

	crawler := Crawler{eth.New(conf.Scanner.Ethscanner.Node_listen_address),
		flogging.GetLogger(),
		stop, startingBlock, 0, 0}

	return crawler
}

func (c *Crawler) GetBlock() <-chan eth.BlockResult {

	results := make(chan eth.BlockResult)

	startBlock := c.startingBlock

	go func() {
		defer close(results)
		for {

			result := c.getChainData(startBlock)

			select {
			case <-c.stopChan:
				return
			case results <- result:
				startBlock = startBlock + 1
			}
		}
	}()

	return results
}

func (c *Crawler) getChainData(startblock int) eth.BlockResult {

	var result eth.BlockResult
	receivedBlock, err := c.ethRpc.EthGetBlockByNumber(startblock, true)
	result = eth.BlockResult{err, receivedBlock}
	return result

}

func (c *Crawler) GetLastBlockHeight() int {

	height, _ := c.ethRpc.EthLastBlockNumber()
	return height
}
