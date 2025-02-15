// Copyright (c) Mainflux
// SPDX-License-Identifier: Apache-2.0

package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

const (
	streamID  = "mainflux.mqtt"
	streamLen = 1000
)

type EventStore interface {
	Connect(clientID string) error
	Disconnect(clientID string) error
}

// EventStore is a struct used to store event streams in Redis.
type eventStore struct {
	client   *redis.Client
	instance string
}

// NewEventStore returns wrapper around mProxy service that sends
// events to event store.
func NewEventStore(client *redis.Client, instance string) EventStore {
	return eventStore{
		client:   client,
		instance: instance,
	}
}

func (es eventStore) storeEvent(clientID, eventType string) error {
	event := mqttEvent{
		clientID:  clientID,
		eventType: eventType,
		instance:  es.instance,
	}

	record := &redis.XAddArgs{
		Stream:       streamID,
		MaxLenApprox: streamLen,
		Values:       event.Encode(),
	}

	return es.client.XAdd(context.Background(), record).Err()
}

// Connect issues event on MQTT CONNECT.
func (es eventStore) Connect(clientID string) error {
	return es.storeEvent(clientID, "connect")
}

// Disconnect issues event on MQTT CONNECT.
func (es eventStore) Disconnect(clientID string) error {
	return es.storeEvent(clientID, "disconnect")
}
