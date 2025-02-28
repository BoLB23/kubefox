package controller

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/xigxog/kubefox/api/kubernetes/v1alpha1"
	"github.com/xigxog/kubefox/components/operator/templates"
	"github.com/xigxog/kubefox/core"
	"github.com/xigxog/kubefox/k8s"
	"github.com/xigxog/kubefox/logkf"
	v1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	ErrResourceVersionConflict = errors.New("resource version conflict")
)

type Client struct {
	k8s.Client
}

func (c *Client) ApplyTemplate(ctx context.Context, name string, data *templates.Data, log *logkf.Logger) error {
	objs, err := templates.Render(name, data)
	if err != nil {
		return err
	}

	for _, obj := range objs {
		log.Debugf("applying template resource '%s'", k8s.ToString(obj))
		if err := c.Apply(ctx, obj); err != nil {
			return err
		}
	}

	return nil
}

func (c *Client) DeleteTemplate(ctx context.Context, name string, data *templates.Data, log *logkf.Logger) error {
	objs, err := templates.Render(name, data)
	if err != nil {
		return err
	}

	for _, obj := range objs {
		log.Debugf("deleting template resource '%s'", k8s.ToString(obj))
		if err := c.Delete(ctx, obj); k8s.IgnoreNotFound(err) != nil {
			return err
		}
	}

	return nil
}

func (r *Client) GetPlatform(ctx context.Context, namespace string) (*v1alpha1.Platform, error) {
	ns := &v1.Namespace{}
	if err := r.Get(ctx, k8s.Key("", namespace), ns); err != nil {
		return nil, fmt.Errorf("unable to fetch namespace: %w", err)
	}
	if ns.Status.Phase == v1.NamespaceTerminating {
		return nil, core.ErrNotFound()
	}

	l := &v1alpha1.PlatformList{}
	if err := r.List(ctx, l, client.InNamespace(namespace)); err != nil {
		return nil, fmt.Errorf("unable to fetch platform: %w", err)
	}
	switch c := len(l.Items); c {
	case 0:
		return nil, core.ErrNotFound()
	case 1:
		return &l.Items[0], nil
	default:
		return nil, core.ErrInvalid(fmt.Errorf("too many Platforms"))
	}
}

// IsFailedWebhookErr will return true if error indicates it was caused by
// calling a webhook. This is useful during operator startup when the Pod is not
// marked ready which causes the webhooks to fail.
func IsFailedWebhookErr(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), "failed calling webhook")
}

func IgnoreFailedWebhookErr(err error) error {
	if IsFailedWebhookErr(err) {
		return nil
	}
	return err
}
