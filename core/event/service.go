package event

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
	"runtime"
	"sync"
)

type Event struct {
	Name string
	Data interface{}
}

type Subscription func(e Event)

type subscriptionID int

type subEntry struct {
	id subscriptionID
	fn Subscription
}

type Service struct {
	subs         map[string][]subEntry
	mu           sync.RWMutex
	currentID    subscriptionID
	logger       *zap.Logger
	eventChan    chan *Event
	pool         sync.Pool
	workerCount  int
	workerWg     sync.WaitGroup
	strategy     PublishStrategy
	shutdownOnce sync.Once
}

type PublishStrategy int

const (
	DiscardNew PublishStrategy = iota
	Block
)

type Option func(*Service)

func NewEvent(logger *zap.Logger, opts ...Option) *Service {
	zap.L().Info("Creating event service")
	es := &Service{
		subs:        make(map[string][]subEntry),
		logger:      logger,
		eventChan:   make(chan *Event, 1000),
		workerCount: runtime.NumCPU() * 2,
		strategy:    DiscardNew,
		pool: sync.Pool{
			New: func() interface{} { return &Event{} },
		},
	}

	for _, opt := range opts {
		opt(es)
	}

	for i := 0; i < es.workerCount; i++ {
		es.workerWg.Add(1)
		go func() {
			defer es.workerWg.Done()
			es.worker()
		}()
	}

	return es
}

func WithWorkerCount(n int) Option {
	return func(es *Service) {
		es.workerCount = n
	}
}

func WithChannelSize(size int) Option {
	return func(es *Service) {
		es.eventChan = make(chan *Event, size)
	}
}

func WithPublishStrategy(strategy PublishStrategy) Option {
	return func(es *Service) {
		es.strategy = strategy
	}
}

func (s *Service) Subscribe(eventName string, fn Subscription) (unsubscribe func()) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.currentID++
	id := s.currentID
	subs := s.subs[eventName]
	subs = append(subs, subEntry{id: id, fn: fn})
	s.subs[eventName] = subs

	s.logger.Debug("New subscription",
		zap.String("event", eventName),
		zap.Int("subscribers", len(subs)),
	)

	return func() {
		s.mu.Lock()
		defer s.mu.Unlock()

		subs := s.subs[eventName]
		for i := 0; i < len(subs); i++ {
			if subs[i].id == id {
				subs = append(subs[:i], subs[i+1:]...)
				s.subs[eventName] = subs
				break
			}
		}
	}
}

func (s *Service) Publish(e Event) {
	s.logger.Debug("Publishing event", zap.String("name", e.Name))

	event := s.pool.Get().(*Event)
	event.Name = e.Name
	event.Data = e.Data

	switch s.strategy {
	case Block:
		s.eventChan <- event
	case DiscardNew:
		select {
		case s.eventChan <- event:
		default:
			s.pool.Put(event)
			s.logger.Warn("Channel full, discarding event",
				zap.String("event", e.Name),
				zap.Int("size", len(s.eventChan)),
			)
		}
	}
}

func (s *Service) worker() {
	for e := range s.eventChan {
		s.mu.RLock()
		subscribers, ok := s.subs[e.Name]
		s.mu.RUnlock()

		if ok {
			var wg sync.WaitGroup
			wg.Add(len(subscribers))

			for _, entry := range subscribers {
				go func(entry subEntry) {
					defer wg.Done()
					defer func() {
						if err := recover(); err != nil {
							s.logger.Error("Subscription panic",
								zap.String("event", e.Name),
								zap.Any("error", err),
								zap.Stack("stack"),
							)
						}
					}()
					entry.fn(*e)
				}(entry)
			}

			wg.Wait()
		}

		e.Name = ""
		e.Data = nil
		s.pool.Put(e)
	}
}

func (s *Service) Shutdown() {
	s.logger.Info("Shutting down event service")
	s.shutdownOnce.Do(func() {
		close(s.eventChan)
		s.workerWg.Wait()
	})
}

var Module = fx.Provide(
	NewEvent,
)
