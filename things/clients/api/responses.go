// Copyright (c) Mainflux
// SPDX-License-Identifier: Apache-2.0

package api

import (
	"fmt"
	"net/http"

	"github.com/mainflux/mainflux"
	mfclients "github.com/mainflux/mainflux/pkg/clients"
)

var (
	_ mainflux.Response = (*viewClientRes)(nil)
	_ mainflux.Response = (*createClientRes)(nil)
	_ mainflux.Response = (*deleteClientRes)(nil)
	_ mainflux.Response = (*clientsPageRes)(nil)
	_ mainflux.Response = (*viewMembersRes)(nil)
	_ mainflux.Response = (*memberPageRes)(nil)
)

type pageRes struct {
	Limit  uint64 `json:"limit,omitempty"`
	Offset uint64 `json:"offset,omitempty"`
	Total  uint64 `json:"total,omitempty"`
}

type createClientRes struct {
	mfclients.Client
	created bool
}

func (res createClientRes) Code() int {
	if res.created {
		return http.StatusCreated
	}

	return http.StatusOK
}

func (res createClientRes) Headers() map[string]string {
	if res.created {
		return map[string]string{
			"Location": fmt.Sprintf("/things/%s", res.ID),
		}
	}

	return map[string]string{}
}

func (res createClientRes) Empty() bool {
	return false
}

type updateClientRes struct {
	mfclients.Client
}

func (res updateClientRes) Code() int {
	return http.StatusOK
}

func (res updateClientRes) Headers() map[string]string {
	return map[string]string{}
}

func (res updateClientRes) Empty() bool {
	return false
}

type viewClientRes struct {
	mfclients.Client
}

func (res viewClientRes) Code() int {
	return http.StatusOK
}

func (res viewClientRes) Headers() map[string]string {
	return map[string]string{}
}

func (res viewClientRes) Empty() bool {
	return false
}

type clientsPageRes struct {
	pageRes
	Clients []viewClientRes `json:"things"`
}

func (res clientsPageRes) Code() int {
	return http.StatusOK
}

func (res clientsPageRes) Headers() map[string]string {
	return map[string]string{}
}

func (res clientsPageRes) Empty() bool {
	return false
}

type viewMembersRes struct {
	mfclients.Client
}

func (res viewMembersRes) Code() int {
	return http.StatusOK
}

func (res viewMembersRes) Headers() map[string]string {
	return map[string]string{}
}

func (res viewMembersRes) Empty() bool {
	return false
}

type memberPageRes struct {
	pageRes
	Members []viewMembersRes `json:"things"`
}

func (res memberPageRes) Code() int {
	return http.StatusOK
}

func (res memberPageRes) Headers() map[string]string {
	return map[string]string{}
}

func (res memberPageRes) Empty() bool {
	return false
}

type deleteClientRes struct {
	mfclients.Client
}

func (res deleteClientRes) Code() int {
	return http.StatusOK
}

func (res deleteClientRes) Headers() map[string]string {
	return map[string]string{}
}

func (res deleteClientRes) Empty() bool {
	return false
}
