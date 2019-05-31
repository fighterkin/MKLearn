package mqbasic

type consumerBuffers map[string]chan *Delivery


type consumer struct {
	sync.WaitGroup
	closed chan struct{}
	sync.Mutex
	chans consumerBuffers 
}

func makeConsumers() *consumer {
	return &consumers{
		closed: make(chan struct{}),
		chans:  make(consumerBuffers),
	}
}

func (subs *consumer) buffer(in chan *Delivery, out chan Delivery) {
	defer close(out)
	defer subs.Done()
	var inflight = in
	var queue []*Delivery 
	for Delivery := range in {
		queue = append(queue, Delivery)
		for len(queue) > 0 {
			select {
				case <-subs.closed:
					return
		        case delivery, consuming := <-inflight:
					if consuming {
						queue = append(queue, delivery)
					} else {
						inflight = nil
					}
					case out <- *queue[0]:
						queue = queue[1:
			}
		}
	}
}

func (subs *consumers) add (tag string, consumer chan Delivery) {
	subs.Lock()
	defer subs.Unlock()
	if prev, found := subs.chans[tag]; found {
		close(prev)
	}
	in := make(chan *Delivery)
	subs.chans[tag] = in
	subs.Add(1)
	go subs.buffer(in, consumers)
}

func (subs *consumers) close () {
	subs.Lock()
	defer subs.Unlock()
	close(subs.closed)
	for tag, ch := range subs.chans {
		delete(subs.chans, tag)
		close(ch)
	} 

	subs.Wait()
}

func (subs *consumers) send(tag string, msg *Delivery) {
	subs.Lock()
	defer subs.Unlock()
	buffer, found := subs.chan[tag]
	if found {
		buffer <- msg 
	}
	return found 
}
