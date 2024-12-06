package services

type ServiceQueues struct {
	PlayerQueues    ServiceQueue
	NavigatorQueues ServiceQueue
	RoomQueues      ServiceQueue

	NumberOfServices int
}

type ServiceQueue struct {
	msgQueue chan Messager
	reqQueue chan Requester
}

func NewServiceQueues(playerQueues, navigatorQueues, roomQueues ServiceQueue) *ServiceQueues {
	return &ServiceQueues{
		PlayerQueues:     playerQueues,
		NavigatorQueues:  navigatorQueues,
		RoomQueues:       roomQueues,
		NumberOfServices: 3,
	}
}

func NewServiceQueue() ServiceQueue {
	return ServiceQueue{
		msgQueue: make(chan Messager),
		reqQueue: make(chan Requester),
	}
}

func (q *ServiceQueue) QueueMessage(msg Messager) {
	q.msgQueue <- msg
}

func (q *ServiceQueue) QueueRequest(req Requester) {
	q.reqQueue <- req
}
