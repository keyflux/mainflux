// Copyright (c) Mainflux
// SPDX-License-Identifier: Apache-2.0

package sdk_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mainflux/mainflux/consumers/notifiers"
	httpapi "github.com/mainflux/mainflux/consumers/notifiers/api"
	"github.com/mainflux/mainflux/consumers/notifiers/mocks"
	"github.com/mainflux/mainflux/internal/apiutil"
	"github.com/mainflux/mainflux/logger"
	"github.com/mainflux/mainflux/pkg/errors"
	sdk "github.com/mainflux/mainflux/pkg/sdk/go"
	"github.com/mainflux/mainflux/pkg/uuid"
)

const wrongValue = "wrong_value"

var (
	sub1 = sdk.Subscription{
		Topic:   "topic",
		Contact: "contact",
	}
	emptySubscription = sdk.Subscription{}
	exampleUser1      = "email1@example.com"
	exampleUser2      = "email2@example.com"
	invalidUser       = "invalid@example.com"
)

func newSubscriptionService() notifiers.Service {
	repo := mocks.NewRepo(make(map[string]notifiers.Subscription))
	auth := mocks.NewAuth(map[string]string{exampleUser1: exampleUser1, exampleUser2: exampleUser2, invalidUser: invalidUser})
	notifier := mocks.NewNotifier()
	idp := uuid.NewMock()
	from := "exampleFrom"
	
	return notifiers.New(auth, repo, idp, notifier, from)
}

func newSubscriptionServer(svc notifiers.Service) *httptest.Server {
	logger := logger.NewMock()
	mux := httpapi.MakeHandler(svc, logger)
	
	return httptest.NewServer(mux)
}

func TestCreateSubscription(t *testing.T) {
	svc := newSubscriptionService()
	ts := newSubscriptionServer(svc)
	defer ts.Close()

	sdkConf := sdk.Config{
		UsersURL:        ts.URL,
		MsgContentType:  contentType,
		TLSVerification: false,
	}

	mfsdk := sdk.NewSDK(sdkConf)

	cases := []struct {
		desc         string
		subscription sdk.Subscription
		token        string
		err          errors.SDKError
		empty        bool
	}{
		{
			desc:         "create new subscription",
			subscription: sub1,
			token:        exampleUser1,
			err:          nil,
			empty:        false,
		},
		{
			desc:         "create new subscription with empty token",
			subscription: sub1,
			token:        "",
			err:          errors.NewSDKErrorWithStatus(apiutil.ErrBearerToken, http.StatusUnauthorized),
			empty:        true,
		},
		{
			desc:         "create new subscription with invalid token",
			subscription: sub1,
			token:        wrongValue,
			err:          errors.NewSDKErrorWithStatus(errors.ErrAuthentication, http.StatusUnauthorized),
			empty:        true,
		},
		{
			desc:         "create new empty subscription",
			subscription: emptySubscription,
			token:        token,
			err:          errors.NewSDKErrorWithStatus(apiutil.ErrInvalidTopic, http.StatusBadRequest),
			empty:        true,
		},
	}

	for _, tc := range cases {
		loc, err := mfsdk.CreateSubscription(tc.subscription.Topic, tc.subscription.Contact, tc.token)
		assert.Equal(t, tc.err, err, fmt.Sprintf("%s: expected error %s, got %s", tc.desc, tc.err, err))
		assert.Equal(t, tc.empty, loc == "", fmt.Sprintf("%s: expected empty result location, got: %s", tc.desc, loc))
	}
}

func TestViewSubscription(t *testing.T) {
	svc := newSubscriptionService()
	ts := newSubscriptionServer(svc)
	defer ts.Close()
	sdkConf := sdk.Config{
		UsersURL:        ts.URL,
		MsgContentType:  contentType,
		TLSVerification: false,
	}

	mfsdk := sdk.NewSDK(sdkConf)
	id, err := mfsdk.CreateSubscription("topic", "contact", exampleUser1)
	require.Nil(t, err, fmt.Sprintf("unexpected error during creating subscription: %s", err))

	cases := []struct {
		desc     string
		subID    string
		token    string
		err      errors.SDKError
		response sdk.Subscription
	}{
		{
			desc:     "get existing subscription",
			subID:    id,
			token:    exampleUser1,
			err:      nil,
			response: sub1,
		},
		{
			desc:     "get non-existent subscription",
			subID:    "43",
			token:    exampleUser1,
			err:      errors.NewSDKErrorWithStatus(errors.ErrNotFound, http.StatusNotFound),
			response: sdk.Subscription{},
		},
		{
			desc:     "get subscription with invalid token",
			subID:    id,
			token:    "",
			err:      errors.NewSDKErrorWithStatus(apiutil.ErrBearerToken, http.StatusUnauthorized),
			response: sdk.Subscription{},
		},
	}

	for _, tc := range cases {
		respSub, err := mfsdk.ViewSubscription(tc.subID, tc.token)

		assert.Equal(t, tc.err, err, fmt.Sprintf("%s: expected error %s, got %s", tc.desc, tc.err, err))
		tc.response.ID = respSub.ID
		tc.response.OwnerID = respSub.OwnerID
		assert.Equal(t, tc.response, respSub, fmt.Sprintf("%s: expected response %s, got %s", tc.desc, tc.response, respSub))
	}
}

func TestListSubscription(t *testing.T) {
	svc := newSubscriptionService()
	ts := newSubscriptionServer(svc)
	defer ts.Close()
	sdkConf := sdk.Config{
		UsersURL:        ts.URL,
		MsgContentType:  contentType,
		TLSVerification: false,
	}

	mfsdk := sdk.NewSDK(sdkConf)
	var nSubs = 10
	subs := make([]sdk.Subscription, nSubs)
	for i := 0; i < nSubs; i++ {
		id, err := mfsdk.CreateSubscription(fmt.Sprintf("topic_%d", i), fmt.Sprintf("contact_%d", i), exampleUser1)
		require.Nil(t, err, fmt.Sprintf("unexpected error during creating subscription: %s", err))
		sub, err := mfsdk.ViewSubscription(id, exampleUser1)
		require.Nil(t, err, fmt.Sprintf("unexpected error during getting subscription: %s", err))
		subs[i] = sub
	}

	cases := []struct {
		desc     string
		page     sdk.PageMetadata
		token    string
		err      errors.SDKError
		response []sdk.Subscription
	}{
		{
			desc:     "list all subscription",
			token:    exampleUser1,
			page:     sdk.PageMetadata{Offset: 0, Limit: uint64(nSubs)},
			err:      nil,
			response: subs,
		},
		{
			desc:     "list subscription with specific topic",
			token:    exampleUser1,
			page:     sdk.PageMetadata{Offset: 0, Limit: uint64(nSubs), Topic: "topic_1"},
			err:      nil,
			response: []sdk.Subscription{subs[1]},
		},
		{
			desc:     "list subscription with specific contact",
			token:    exampleUser1,
			page:     sdk.PageMetadata{Offset: 0, Limit: uint64(nSubs), Contact: "contact_1"},
			err:      nil,
			response: []sdk.Subscription{subs[1]},
		},
	}

	for _, tc := range cases {
		subs, err := mfsdk.ListSubscriptions(tc.page, tc.token)
		assert.Equal(t, tc.err, err, fmt.Sprintf("%s: expected error %s, got %s", tc.desc, tc.err, err))
		assert.Equal(t, tc.response, subs.Subscriptions, fmt.Sprintf("%s: expected response %v, got %v", tc.desc, tc.response, subs.Subscriptions))
	}
}

func TestDeleteSubscription(t *testing.T) {
	svc := newSubscriptionService()
	ts := newSubscriptionServer(svc)
	defer ts.Close()
	sdkConf := sdk.Config{
		UsersURL:        ts.URL,
		MsgContentType:  contentType,
		TLSVerification: false,
	}

	mfsdk := sdk.NewSDK(sdkConf)
	id, err := mfsdk.CreateSubscription("topic", "contact", exampleUser1)
	require.Nil(t, err, fmt.Sprintf("unexpected error during creating subscription: %s", err))

	cases := []struct {
		desc     string
		subID    string
		token    string
		err      errors.SDKError
		response sdk.Subscription
	}{
		{
			desc:     "delete existing subscription",
			subID:    id,
			token:    exampleUser1,
			err:      nil,
			response: sub1,
		},
		{
			desc:     "delete non-existent subscription",
			subID:    "43",
			token:    exampleUser1,
			err:      errors.NewSDKErrorWithStatus(errors.ErrNotFound, http.StatusNotFound),
			response: sdk.Subscription{},
		},
		{
			desc:     "delete subscription with invalid token",
			subID:    id,
			token:    "",
			err:      errors.NewSDKErrorWithStatus(apiutil.ErrBearerToken, http.StatusUnauthorized),
			response: sdk.Subscription{},
		},
	}

	for _, tc := range cases {
		err := mfsdk.DeleteSubscription(tc.subID, tc.token)
		assert.Equal(t, tc.err, err, fmt.Sprintf("%s: expected error %s, got %s", tc.desc, tc.err, err))
	}
}
