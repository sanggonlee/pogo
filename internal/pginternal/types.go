package pginternal

import "gopkg.in/guregu/null.v3"

// LSN represetns a pg_lsn type.
type LSN null.String

// OID represents a oid type in Postgres.
type OID null.Int

// BigInt represents a bigint type in Postgres.
// NOTE: null.Int does not span the entire range of Postgres' bigint type.
// TODO: consider big.Int
type BigInt null.Int
