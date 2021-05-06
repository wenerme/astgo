package ami

import "context"

func SubscribeChan(c chan *Message, names ...string) SubscribeFunc {
	m := make(map[string]struct{}, len(names))
	for _, v := range names {
		m[v] = struct{}{}
	}
	filter := func(ctx context.Context, msg *Message) bool {
		_, ok := m[msg.Name]
		return ok
	}
	if len(names) == 0 {
		filter = func(ctx context.Context, msg *Message) bool {
			return true
		}
	}
	return func(ctx context.Context, msg *Message) bool {
		if filter(ctx, msg) {
			c <- msg
		}

		return true
	}
}

func SubscribeFullyBootedChanOnce(c chan *Message) SubscribeFunc {
	return func(ctx context.Context, msg *Message) bool {
		if msg.Name == "FullyBooted" {
			c <- msg
		}
		return false
	}
}
