package copyright

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/redish101/reblog/internal/env"
	"github.com/redish101/reblog/internal/store"
	"github.com/sirupsen/logrus"
)

var (
	client   *ethclient.Client
	contract *Copyright
)

func Init() {
	var err error
	client, err = ethclient.Dial(env.CopyrightRPCURL)
	if err != nil {
		logrus.Warnf("[COPYRIGHT] 无法连接到区块链节点: %v", err)
	}

	contractAddr := common.HexToAddress(env.CopyrightContractAddress)
	contract, err = NewCopyright(contractAddr, client)
	if err != nil {
		logrus.Warnf("[COPYRIGHT] 无法实例化合约: %v", err)
	}

	owner, err := contract.Owner(nil)
	if err != nil {
		logrus.Warnf("[COPYRIGHT] 无法获取合约所有者: %v", err)
	}
	logrus.Debugf("[COPYRIGHT] 合约已初始化, 所有者: %s 将启动监听器以监听文章更改", owner.Hex())
	go listenForPostChanges()
}

func listenForPostChanges() {
	eventCh := make(chan *CopyrightPostAddedOrUpdated)

	var err error
	sub, err := contract.WatchPostAddedOrUpdated(nil, eventCh)
	if err != nil {
		logrus.Warnf("[COPYRIGHT] 启动事件订阅失败: %v", err)
		return
	}
	defer sub.Unsubscribe()

	for {
		select {
		case err := <-sub.Err():
			logrus.Warnf("[COPYRIGHT] 事件订阅错误: %v", err)
			// 可以考虑重连逻辑，例如等待几秒后重新启动监听
			return

		case ev := <-eventCh:
			logrus.Debugf("[COPYRIGHT] 文章创建事件: 作者=%s 标题=%s IPFSURL=%s 交易哈希=%s",
				ev.Author.Hex(), ev.Title, ev.IpfsURL, ev.Raw.TxHash.Hex())
			post, err := store.Post.FindBySlug(ev.Slug, env.OwnerEmail)
			if err != nil {
				logrus.Warnf("[COPYRIGHT] 无法找到文章 (slug=%s): %v", ev.Slug, err)
				return
			}
			post.IPFSURL = ev.IpfsURL
			if post.CreateTXHash == "" {
				post.CreateTXHash = ev.Raw.TxHash.Hex()
			} else {
				post.UpdateTXHash = ev.Raw.TxHash.Hex()
			}

			err = store.Post.Update(post)
			if err != nil {
				logrus.Warnf("[COPYRIGHT] 无法更新文章 (slug=%s): %v", ev.Slug, err)
				return
			}
		}
	}
}
