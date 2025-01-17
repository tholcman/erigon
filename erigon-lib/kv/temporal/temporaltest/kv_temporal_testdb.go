// Copyright 2024 The Erigon Authors
// This file is part of Erigon.
//
// Erigon is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Erigon is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with Erigon. If not, see <http://www.gnu.org/licenses/>.

package temporaltest

import (
	"context"
	"testing"

	"github.com/ledgerwatch/erigon-lib/common/datadir"
	"github.com/ledgerwatch/erigon-lib/config3"
	"github.com/ledgerwatch/erigon-lib/kv"
	"github.com/ledgerwatch/erigon-lib/kv/memdb"
	"github.com/ledgerwatch/erigon-lib/kv/temporal"
	"github.com/ledgerwatch/erigon-lib/log/v3"
	"github.com/ledgerwatch/erigon-lib/state"
)

// nolint:thelper
func NewTestDB(tb testing.TB, dirs datadir.Dirs) (db kv.RwDB, agg *state.Aggregator) {
	if tb != nil {
		tb.Helper()
	}
	logger := log.New()

	if tb != nil {
		db = memdb.NewTestDB(tb)
	} else {
		db = memdb.New(dirs.DataDir)
	}

	var err error
	agg, err = state.NewAggregator(context.Background(), dirs, config3.HistoryV3AggregationStep, db, nil, logger)
	if err != nil {
		panic(err)
	}
	if err := agg.OpenFolder(); err != nil {
		panic(err)
	}

	db, err = temporal.New(db, agg)
	if err != nil {
		panic(err)
	}
	return db, agg
}
