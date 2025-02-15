// Copyright (c) Mainflux
// SPDX-License-Identifier: Apache-2.0

package mocks

import (
	notifiers "github.com/mainflux/mainflux/consumers/notifiers"
	"github.com/mainflux/mainflux/pkg/messaging"
)

var _ notifiers.Notifier = (*notifier)(nil)

const invalidSender = "invalid@example.com"

type notifier struct{}

// NewNotifier returns a new Notifier mock.
func NewNotifier() notifiers.Notifier {
	return notifier{}
}

func (n notifier) Notify(from string, to []string, msg *messaging.Message) error {
	for _, t := range to {
		if t == invalidSender {
			return notifiers.ErrNotify
		}
	}
	return nil
}
