package broker

import (
	"errors"
	"sync"
	"time"
)

type (
	Broker interface {
		// 进行消息的推送，有两个参数即topic、msg，分别是订阅的主题、要传递的消息
		publish(topic string, msg any) error

		// 消息的订阅，传入订阅的主题，即可完成订阅，并返回对应的channel通道用来接收数据
		subscribe(topic string) (<-chan any, error)

		// 取消订阅，传入订阅的主题和对应的通道
		unsubscribe(topic string) error

		// 关闭消息队列
		close()

		// 进行广播，对推送的消息进行广播，保证每一个订阅者都可以收到
		broadcast(msg any, subscribers []chan any)

		// 设置消息队列的容量，这样我们就可以控制消息队列的大小了
		setConditions(capacity int)
	}

	BrokerImpl struct {
		sync.RWMutex // 同步锁

		exit     chan bool
		capacity int

		topics map[string][]chan interface{} // key： topic  value ： queue
	}

	Client struct {
		bro *BrokerImpl
	}
)

func NewClient() *Client {
	return &Client{
		bro: NewBroker(),
		// bro: new(BrokerImpl),
	}
}

func (c *Client) SetConditions(capacity int) {
	c.bro.setConditions(capacity)
}
func (c *Client) Subscribe(topic string) (<-chan any, error) {
	return c.bro.subscribe(topic)
}
func (c *Client) Unsubscribe(topic string, sub <-chan any) error {
	return c.bro.unsubscribe(topic, sub)
}
func (c *Client) Publish(topic string, pub any) error {
	return c.bro.publish(topic, pub)
}
func (c *Client) Close() {
	c.bro.close()
}

func (c *Client) GetPayLoad(sub <-chan any) any {

	for v := range sub {
		if v != nil {
			return v
		}
	}

	return nil
}

func NewBroker() *BrokerImpl {
	return &BrokerImpl{
		topics: map[string][]chan any{},
	}
}
func (b *BrokerImpl) setConditions(capacity int) {
	b.capacity = capacity
}
func (b *BrokerImpl) close() {
	select {
	case <-b.exit:
		return
	default:
		close(b.exit)
		b.Lock()
		b.topics = make(map[string][]chan any)
		b.Unlock()
	}

	return
}

func (b *BrokerImpl) publish(topic string, pub any) error {
	select {
	case <-b.exit:
		return errors.New("broker closed")
	default:
	}

	b.RLock()
	subscribers, ok := b.topics[topic]
	b.RUnlock()

	if !ok {
		return nil
	}

	b.broadcast(pub, subscribers)

	return nil
}
func (b *BrokerImpl) broadcast(msg any, subscribers []chan any) {
	count := len(subscribers)
	concurrency := 1

	switch {
	case count > 1000:
		concurrency = 3
	case count > 100:
		concurrency = 2
	default:
		concurrency = 1
	}

	pub := func(start int) {
		for i := start; i < count; i += concurrency {
			select {
			case subscribers[i] <- msg:
			case <-time.After(time.Millisecond * 5):
			case <-b.exit:
				return
			}
		}
	}

	for j := 0; j < concurrency; j++ {
		go pub(j)
	}
}

func (b *BrokerImpl) subscribe(topic string) (<-chan any, error) {
	select {
	case <-b.exit:
		return nil, errors.New("broker closed")
	default:
	}

	ch := make(chan any, b.capacity)
	b.Lock()
	b.topics[topic] = append(b.topics[topic], ch)
	b.Unlock()

	return ch, nil
}
func (b *BrokerImpl) unsubscribe(topic string, sub <-chan any) error {
	select {
	case <-b.exit:
		return errors.New("broker closed")
	default:
	}

	b.RLock()
	subscribers, ok := b.topics[topic]
	b.RUnlock()

	if !ok {
		return nil
	}

	var newSubs []chan any
	for _, subscriber := range subscribers {
		if subscriber == sub {
			continue
		}
		newSubs = append(newSubs, subscriber)
	}

	b.Lock()
	b.topics[topic] = newSubs
	b.Unlock()

	return nil
}
