package pgsql

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/cortezaproject/corteza-server/corteza/store/provisioner"
	"github.com/cortezaproject/corteza-server/corteza/store/rdbms"
	_ "github.com/lib/pq"
	"net/url"
	"strings"
)

type (
	Store struct {
		*rdbms.Store
	}
)

func New(ctx context.Context, dsn string) (s *Store, err error) {
	var cfg *rdbms.Config

	if cfg, err = ProcDataSourceName(dsn); err != nil {
		return nil, err
	}

	cfg.PlaceholderFormat = squirrel.Dollar

	s = new(Store)
	if s.Store, err = rdbms.New(ctx, cfg); err != nil {
		return nil, err
	}

	return s, nil
}

func (s *Store) Provision() provisioner.Executor {
	return provisioner.Do(
		s.ProvisionCore(),
		s.ProvisionCompose(),
		s.ProvisionMessaging(),
	)
}

// ProcDataSourceName validates given DSN and ensures
// params are present and correct
func ProcDataSourceName(dsn string) (c *rdbms.Config, err error) {
	const (
		validScheme = "postgres"
	)
	var (
		scheme string
		u      *url.URL
	)

	if u, err = url.Parse(dsn); err != nil {
		return nil, err
	}

	if strings.HasPrefix(dsn, "postgres") {
		scheme = u.Scheme
		u.Scheme = validScheme
	}

	return &rdbms.Config{
		DriverName:     scheme,
		DataSourceName: u.String(),
		DBName:         strings.Trim(u.Path, "/"),
	}, nil
}