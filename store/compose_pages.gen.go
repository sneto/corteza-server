package store

// This file is auto-generated.
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//
// Definitions file that controls how this file is generated:
// store/compose_pages.yaml

import (
	"context"
	"github.com/cortezaproject/corteza-server/compose/types"
	"github.com/cortezaproject/corteza-server/pkg/scenario"
)

type (
	ComposePages interface {
		SearchComposePages(ctx context.Context, f types.PageFilter) (types.PageSet, types.PageFilter, error)
		LookupComposePageByHandle(ctx context.Context, handle string) (*types.Page, error)
		CreateComposePage(ctx context.Context, rr ...*types.Page) error
		UpdateComposePage(ctx context.Context, rr ...*types.Page) error
		PartialUpdateComposePage(ctx context.Context, onlyColumns []string, rr ...*types.Page) error
		RemoveComposePage(ctx context.Context, rr ...*types.Page) error
		RemoveComposePageByID(ctx context.Context, ID uint64) error

		TruncateComposePages(ctx context.Context) error
	}

	ComposePagesProvisioned interface {
		ComposePages
		ProvisionComposePages() scenario.Executor
	}
)
