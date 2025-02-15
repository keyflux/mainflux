// Copyright (c) Mainflux
// SPDX-License-Identifier: Apache-2.0

package tracing

import (
	"context"

	mfclients "github.com/mainflux/mainflux/pkg/clients"
	"github.com/mainflux/mainflux/things/clients"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

var _ clients.Service = (*tracingMiddleware)(nil)

type tracingMiddleware struct {
	tracer trace.Tracer
	svc    clients.Service
}

// New returns a new group service with tracing capabilities.
func New(svc clients.Service, tracer trace.Tracer) clients.Service {
	return &tracingMiddleware{tracer, svc}
}

// CreateThings traces the "CreateThings" operation of the wrapped policies.Service.
func (tm *tracingMiddleware) CreateThings(ctx context.Context, token string, clis ...mfclients.Client) ([]mfclients.Client, error) {
	ctx, span := tm.tracer.Start(ctx, "svc_create_client")
	defer span.End()

	return tm.svc.CreateThings(ctx, token, clis...)
}

// ViewClient traces the "ViewClient" operation of the wrapped policies.Service.
func (tm *tracingMiddleware) ViewClient(ctx context.Context, token string, id string) (mfclients.Client, error) {
	ctx, span := tm.tracer.Start(ctx, "svc_view_client", trace.WithAttributes(attribute.String("id", id)))
	defer span.End()
	return tm.svc.ViewClient(ctx, token, id)
}

// ListClients traces the "ListClients" operation of the wrapped policies.Service.
func (tm *tracingMiddleware) ListClients(ctx context.Context, token string, pm mfclients.Page) (mfclients.ClientsPage, error) {
	ctx, span := tm.tracer.Start(ctx, "svc_list_clients")
	defer span.End()
	return tm.svc.ListClients(ctx, token, pm)
}

// UpdateClient traces the "UpdateClient" operation of the wrapped policies.Service.
func (tm *tracingMiddleware) UpdateClient(ctx context.Context, token string, cli mfclients.Client) (mfclients.Client, error) {
	ctx, span := tm.tracer.Start(ctx, "svc_update_client_name_and_metadata", trace.WithAttributes(attribute.String("id", cli.ID)))
	defer span.End()

	return tm.svc.UpdateClient(ctx, token, cli)
}

// UpdateClientTags traces the "UpdateClientTags" operation of the wrapped policies.Service.
func (tm *tracingMiddleware) UpdateClientTags(ctx context.Context, token string, cli mfclients.Client) (mfclients.Client, error) {
	ctx, span := tm.tracer.Start(ctx, "svc_update_client_tags", trace.WithAttributes(
		attribute.String("id", cli.ID),
		attribute.StringSlice("tags", cli.Tags),
	))
	defer span.End()

	return tm.svc.UpdateClientTags(ctx, token, cli)
}

// UpdateClientSecret traces the "UpdateClientSecret" operation of the wrapped policies.Service.
func (tm *tracingMiddleware) UpdateClientSecret(ctx context.Context, token, oldSecret, newSecret string) (mfclients.Client, error) {
	ctx, span := tm.tracer.Start(ctx, "svc_update_client_secret")
	defer span.End()

	return tm.svc.UpdateClientSecret(ctx, token, oldSecret, newSecret)

}

// UpdateClientOwner traces the "UpdateClientOwner" operation of the wrapped policies.Service.
func (tm *tracingMiddleware) UpdateClientOwner(ctx context.Context, token string, cli mfclients.Client) (mfclients.Client, error) {
	ctx, span := tm.tracer.Start(ctx, "svc_update_client_owner", trace.WithAttributes(
		attribute.String("id", cli.ID),
		attribute.String("owner", cli.Owner),
	))
	defer span.End()

	return tm.svc.UpdateClientOwner(ctx, token, cli)
}

// EnableClient traces the "EnableClient" operation of the wrapped policies.Service.
func (tm *tracingMiddleware) EnableClient(ctx context.Context, token, id string) (mfclients.Client, error) {
	ctx, span := tm.tracer.Start(ctx, "svc_enable_client", trace.WithAttributes(attribute.String("id", id)))
	defer span.End()

	return tm.svc.EnableClient(ctx, token, id)
}

// DisableClient traces the "DisableClient" operation of the wrapped policies.Service.
func (tm *tracingMiddleware) DisableClient(ctx context.Context, token, id string) (mfclients.Client, error) {
	ctx, span := tm.tracer.Start(ctx, "svc_disable_client", trace.WithAttributes(attribute.String("id", id)))
	defer span.End()

	return tm.svc.DisableClient(ctx, token, id)
}

// ListClientsByGroup traces the "ListClientsByGroup" operation of the wrapped policies.Service.
func (tm *tracingMiddleware) ListClientsByGroup(ctx context.Context, token, groupID string, pm mfclients.Page) (mfclients.MembersPage, error) {
	ctx, span := tm.tracer.Start(ctx, "svc_list_things_by_channel", trace.WithAttributes(attribute.String("groupID", groupID)))
	defer span.End()

	return tm.svc.ListClientsByGroup(ctx, token, groupID, pm)

}

// ListMemberships traces the "ListMemberships" operation of the wrapped policies.Service.
func (tm *tracingMiddleware) Identify(ctx context.Context, key string) (string, error) {
	ctx, span := tm.tracer.Start(ctx, "svc_identify", trace.WithAttributes(attribute.String("key", key)))
	defer span.End()
	return tm.svc.Identify(ctx, key)
}
