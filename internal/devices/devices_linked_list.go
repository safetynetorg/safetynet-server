package devices

import (
	"safetynet/internal/database"
)

type node struct {
	database.SafetynetDevice
	next, prev *node
}

type list struct {
	head, tail *node
}

func (l *list) push(d database.SafetynetDevice) *list {
	n := &node{SafetynetDevice: d}
	if l.head == nil {
		l.head = n
	} else {
		l.tail.next = n
		n.prev = l.tail
	}
	l.tail = n
	return l
}

func (l *list) findForward(id string, ch chan *node) {
	for n := l.head; n != nil && <-ch == nil; n = n.next {
		if n.SafetynetDevice.Id == id {
			ch <- n
		}
	}
}

func (l *list) findReverse(id string, ch chan *node) {
	for n := l.tail; n != nil && <-ch == nil; n = n.prev {
		if n.SafetynetDevice.Id == id {
			ch <- n
		}
	}
}

func (l *list) shouldDelete(id string, oldList *list) bool {
	var success bool

	chNew := make(chan *node)
	chOld := make(chan *node)

	go l.findForward(id, chNew)
	go l.findReverse(id, chNew)

	go oldList.findForward(id, chOld)
	go oldList.findReverse(id, chOld)

	nNew := <-chNew
	nOld := <-chOld

	if nNew != nil && nOld != nil && nNew.Lat == nOld.Lat && nNew.Lon == nOld.Lon {
		prevNode := nOld.prev
		nextNode := nOld.next
		prevNode.next = nOld.next
		nextNode.prev = nOld.prev

		success = true
	}
	return success
}
