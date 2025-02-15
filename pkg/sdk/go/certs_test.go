// Copyright (c) Mainflux
// SPDX-License-Identifier: Apache-2.0

package sdk_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	bsmocks "github.com/mainflux/mainflux/bootstrap/mocks"
	"github.com/mainflux/mainflux/certs"
	httpapi "github.com/mainflux/mainflux/certs/api"
	"github.com/mainflux/mainflux/certs/mocks"
	"github.com/mainflux/mainflux/internal/apiutil"
	"github.com/mainflux/mainflux/logger"
	mfclients "github.com/mainflux/mainflux/pkg/clients"
	"github.com/mainflux/mainflux/pkg/errors"
	sdk "github.com/mainflux/mainflux/pkg/sdk/go"
	"github.com/mainflux/mainflux/things/clients"
	"github.com/mainflux/mainflux/things/policies"
	pmocks "github.com/mainflux/mainflux/things/policies/mocks"
	upolicies "github.com/mainflux/mainflux/users/policies"
)

var (
	thingsNum         = 1
	thingKey          = "thingKey"
	thingID           = "1"
	caPath            = "../../../docker/ssl/certs/ca.crt"
	caKeyPath         = "../../../docker/ssl/certs/ca.key"
	cfgAuthTimeout    = "1s"
	cfgSignHoursValid = "24h"
)

func newCertsThingsService(auth upolicies.AuthServiceClient) clients.Service {
	ths := make(map[string]mfclients.Client, thingsNum)
	for i := 0; i < thingsNum; i++ {
		id := strconv.Itoa(i + 1)
		ths[id] = mfclients.Client{
			ID: id,
			Credentials: mfclients.Credentials{
				Secret: thingKey,
			},
			Owner:  adminID,
			Status: mfclients.EnabledStatus,
		}
	}

	return bsmocks.NewThingsService(ths, auth)
}

func newCertService() (certs.Service, error) {
	uauth := bsmocks.NewAuthClient(map[string]string{adminToken: adminID})
	policiesCache := pmocks.NewCache()

	pRepo := new(pmocks.Repository)
	psvc := policies.NewService(uauth, pRepo, policiesCache, idProvider)
	server := newThingsServer(newCertsThingsService(uauth), psvc)

	config := sdk.Config{
		ThingsURL: server.URL,
	}

	mfsdk := sdk.NewSDK(config)
	repo := mocks.NewCertsRepository()

	tlsCert, caCert, err := certs.LoadCertificates(caPath, caKeyPath)
	if err != nil {
		return nil, err
	}

	authTimeout, err := time.ParseDuration(cfgAuthTimeout)
	if err != nil {
		return nil, err
	}

	pki := mocks.NewPkiAgent(tlsCert, caCert, cfgSignHoursValid, authTimeout)

	return certs.New(uauth, repo, mfsdk, pki), nil
}

func newCertServer(svc certs.Service) *httptest.Server {
	logger := logger.NewMock()
	mux := httpapi.MakeHandler(svc, logger)
	return httptest.NewServer(mux)
}

func TestIssueCert(t *testing.T) {
	svc, err := newCertService()
	require.Nil(t, err, fmt.Sprintf("unexpected error during creating service: %s", err))
	ts := newCertServer(svc)
	defer ts.Close()

	sdkConf := sdk.Config{
		CertsURL:        ts.URL,
		MsgContentType:  contentType,
		TLSVerification: false,
	}

	mfsdk := sdk.NewSDK(sdkConf)

	cases := []struct {
		desc     string
		thingID  string
		duration string
		token    string
		err      errors.SDKError
	}{
		{
			desc:     "create new cert with thing id and duration",
			thingID:  thingID,
			duration: "10h",
			token:    adminToken,
			err:      nil,
		},
		{
			desc:     "create new cert with empty thing id and duration",
			thingID:  "",
			duration: "10h",
			token:    adminToken,
			err:      errors.NewSDKErrorWithStatus(apiutil.ErrMissingID, http.StatusBadRequest),
		},
		{
			desc:     "create new cert with invalid thing id and duration",
			thingID:  "ah",
			duration: "10h",
			token:    adminToken,
			err:      errors.NewSDKErrorWithStatus(certs.ErrFailedCertCreation, http.StatusInternalServerError),
		},
		{
			desc:     "create new cert with thing id and empty duration",
			thingID:  thingID,
			duration: "",
			token:    exampleUser1,
			err:      errors.NewSDKErrorWithStatus(apiutil.ErrMissingCertData, http.StatusBadRequest),
		},
		{
			desc:     "create new cert with thing id and malformed duration",
			thingID:  thingID,
			duration: "10g",
			token:    exampleUser1,
			err:      errors.NewSDKErrorWithStatus(apiutil.ErrInvalidCertData, http.StatusBadRequest),
		},
		{
			desc:     "create new cert with empty token",
			thingID:  thingID,
			duration: "10h",
			token:    "",
			err:      errors.NewSDKErrorWithStatus(apiutil.ErrBearerToken, http.StatusUnauthorized),
		},
		{
			desc:     "create new cert with invalid token",
			thingID:  thingID,
			duration: "10h",
			token:    wrongValue,
			err:      errors.NewSDKErrorWithStatus(errors.ErrAuthentication, http.StatusUnauthorized),
		},
		{
			desc:     "create new empty cert",
			thingID:  "",
			duration: "",
			token:    adminToken,
			err:      errors.NewSDKErrorWithStatus(apiutil.ErrMissingID, http.StatusBadRequest),
		},
	}

	for _, tc := range cases {
		cert, err := mfsdk.IssueCert(tc.thingID, tc.duration, tc.token)
		assert.Equal(t, tc.err, err, fmt.Sprintf("%s: expected error %s, got %s", tc.desc, tc.err, err))
		if err == nil {
			assert.NotEmpty(t, cert, fmt.Sprintf("%s: got empty cert", tc.desc))
		}
	}
}

func TestViewCert(t *testing.T) {
	svc, err := newCertService()
	require.Nil(t, err, fmt.Sprintf("unexpected error during creating service: %s", err))
	ts := newCertServer(svc)
	defer ts.Close()

	sdkConf := sdk.Config{
		CertsURL:        ts.URL,
		MsgContentType:  contentType,
		TLSVerification: false,
	}

	mfsdk := sdk.NewSDK(sdkConf)

	cert, err := mfsdk.IssueCert(thingID, "10h", token)
	require.Nil(t, err, fmt.Sprintf("unexpected error during creating cert: %s", err))

	cases := []struct {
		desc     string
		certID   string
		token    string
		err      errors.SDKError
		response sdk.Subscription
	}{
		{
			desc:     "get existing cert",
			certID:   cert.CertSerial,
			token:    token,
			err:      nil,
			response: sub1,
		},
		{
			desc:     "get non-existent cert",
			certID:   "43",
			token:    token,
			err:      errors.NewSDKErrorWithStatus(errors.ErrNotFound, http.StatusInternalServerError),
			response: sdk.Subscription{},
		},
		{
			desc:     "get cert with invalid token",
			certID:   cert.CertSerial,
			token:    "",
			err:      errors.NewSDKErrorWithStatus(apiutil.ErrBearerToken, http.StatusUnauthorized),
			response: sdk.Subscription{},
		},
	}

	for _, tc := range cases {
		cert, err := mfsdk.ViewCert(tc.certID, tc.token)
		assert.Equal(t, tc.err, err, fmt.Sprintf("%s: expected error %s, got %s", tc.desc, tc.err, err))
		if err == nil {
			assert.NotEmpty(t, cert, fmt.Sprintf("%s: got empty cert", tc.desc))
		}
	}
}

func TestViewCertByThing(t *testing.T) {
	svc, err := newCertService()
	require.Nil(t, err, fmt.Sprintf("unexpected error during creating service: %s", err))
	ts := newCertServer(svc)
	defer ts.Close()

	sdkConf := sdk.Config{
		CertsURL:        ts.URL,
		MsgContentType:  contentType,
		TLSVerification: false,
	}

	mfsdk := sdk.NewSDK(sdkConf)

	_, err = mfsdk.IssueCert(thingID, "10h", token)
	require.Nil(t, err, fmt.Sprintf("unexpected error during creating cert: %s", err))

	cases := []struct {
		desc     string
		thingID  string
		token    string
		err      errors.SDKError
		response sdk.Subscription
	}{
		{
			desc:     "get existing cert",
			thingID:  thingID,
			token:    token,
			err:      nil,
			response: sub1,
		},
		{
			desc:     "get non-existent cert",
			thingID:  "43",
			token:    token,
			err:      errors.NewSDKErrorWithStatus(errors.ErrNotFound, http.StatusInternalServerError),
			response: sdk.Subscription{},
		},
		{
			desc:     "get cert with invalid token",
			thingID:  thingID,
			token:    "",
			err:      errors.NewSDKErrorWithStatus(apiutil.ErrBearerToken, http.StatusUnauthorized),
			response: sdk.Subscription{},
		},
	}

	for _, tc := range cases {
		cert, err := mfsdk.ViewCertByThing(tc.thingID, tc.token)
		assert.Equal(t, tc.err, err, fmt.Sprintf("%s: expected error %s, got %s", tc.desc, tc.err, err))
		if err == nil {
			assert.NotEmpty(t, cert, fmt.Sprintf("%s: got empty cert", tc.desc))
		}
	}
}

func TestRevokeCert(t *testing.T) {
	svc, err := newCertService()
	require.Nil(t, err, fmt.Sprintf("unexpected error during creating service: %s", err))
	ts := newCertServer(svc)
	defer ts.Close()

	sdkConf := sdk.Config{
		CertsURL:        ts.URL,
		MsgContentType:  contentType,
		TLSVerification: false,
	}

	mfsdk := sdk.NewSDK(sdkConf)

	_, err = mfsdk.IssueCert(thingID, "10h", adminToken)
	require.Nil(t, err, fmt.Sprintf("unexpected error during creating cert: %s", err))

	cases := []struct {
		desc    string
		thingID string
		token   string
		err     errors.SDKError
	}{
		{
			desc:    "revoke cert with invalid token",
			thingID: thingID,
			token:   wrongValue,
			err:     errors.NewSDKErrorWithStatus(errors.ErrAuthentication, http.StatusUnauthorized),
		},
		{
			desc:    "revoke non-existing cert",
			thingID: "2",
			token:   token,
			err:     errors.NewSDKErrorWithStatus(certs.ErrFailedCertRevocation, http.StatusInternalServerError),
		},
		{
			desc:    "revoke cert with invalid id",
			thingID: "",
			token:   token,
			err:     errors.NewSDKErrorWithStatus(apiutil.ErrMissingID, http.StatusBadRequest),
		},
		{
			desc:    "revoke cert with empty token",
			thingID: thingID,
			token:   "",
			err:     errors.NewSDKErrorWithStatus(apiutil.ErrBearerToken, http.StatusUnauthorized),
		},
		{
			desc:    "revoke existing cert",
			thingID: thingID,
			token:   token,
			err:     nil,
		},
		{
			desc:    "revoke deleted cert",
			thingID: thingID,
			token:   token,
			err:     errors.NewSDKErrorWithStatus(certs.ErrFailedToRemoveCertFromDB, http.StatusInternalServerError),
		},
	}

	for _, tc := range cases {
		response, err := mfsdk.RevokeCert(tc.thingID, tc.token)
		assert.Equal(t, tc.err, err, fmt.Sprintf("%s: expected error %s, got %s", tc.desc, tc.err, err))
		if err == nil {
			assert.NotEmpty(t, response, fmt.Sprintf("%s: got empty revocation time", tc.desc))
		}
	}
}
