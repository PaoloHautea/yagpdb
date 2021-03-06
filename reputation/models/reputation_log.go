// This file is generated by SQLBoiler (https://github.com/vattle/sqlboiler)
// and is meant to be re-generated in place and/or deleted at any time.
// DO NOT EDIT

package models

import (
	"bytes"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/queries"
	"github.com/vattle/sqlboiler/queries/qm"
	"github.com/vattle/sqlboiler/strmangle"
)

// ReputationLog is an object representing the database table.
type ReputationLog struct {
	ID             int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt      time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	GuildID        int64     `boil:"guild_id" json:"guild_id" toml:"guild_id" yaml:"guild_id"`
	SenderID       int64     `boil:"sender_id" json:"sender_id" toml:"sender_id" yaml:"sender_id"`
	ReceiverID     int64     `boil:"receiver_id" json:"receiver_id" toml:"receiver_id" yaml:"receiver_id"`
	SetFixedAmount bool      `boil:"set_fixed_amount" json:"set_fixed_amount" toml:"set_fixed_amount" yaml:"set_fixed_amount"`
	Amount         int64     `boil:"amount" json:"amount" toml:"amount" yaml:"amount"`

	R *reputationLogR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L reputationLogL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ReputationLogColumns = struct {
	ID             string
	CreatedAt      string
	GuildID        string
	SenderID       string
	ReceiverID     string
	SetFixedAmount string
	Amount         string
}{
	ID:             "id",
	CreatedAt:      "created_at",
	GuildID:        "guild_id",
	SenderID:       "sender_id",
	ReceiverID:     "receiver_id",
	SetFixedAmount: "set_fixed_amount",
	Amount:         "amount",
}

// reputationLogR is where relationships are stored.
type reputationLogR struct {
}

// reputationLogL is where Load methods for each relationship are stored.
type reputationLogL struct{}

var (
	reputationLogColumns               = []string{"id", "created_at", "guild_id", "sender_id", "receiver_id", "set_fixed_amount", "amount"}
	reputationLogColumnsWithoutDefault = []string{"created_at", "guild_id", "sender_id", "receiver_id", "set_fixed_amount", "amount"}
	reputationLogColumnsWithDefault    = []string{"id"}
	reputationLogPrimaryKeyColumns     = []string{"id"}
)

type (
	// ReputationLogSlice is an alias for a slice of pointers to ReputationLog.
	// This should generally be used opposed to []ReputationLog.
	ReputationLogSlice []*ReputationLog

	reputationLogQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	reputationLogType                 = reflect.TypeOf(&ReputationLog{})
	reputationLogMapping              = queries.MakeStructMapping(reputationLogType)
	reputationLogPrimaryKeyMapping, _ = queries.BindMapping(reputationLogType, reputationLogMapping, reputationLogPrimaryKeyColumns)
	reputationLogInsertCacheMut       sync.RWMutex
	reputationLogInsertCache          = make(map[string]insertCache)
	reputationLogUpdateCacheMut       sync.RWMutex
	reputationLogUpdateCache          = make(map[string]updateCache)
	reputationLogUpsertCacheMut       sync.RWMutex
	reputationLogUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)

// OneP returns a single reputationLog record from the query, and panics on error.
func (q reputationLogQuery) OneP() *ReputationLog {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single reputationLog record from the query.
func (q reputationLogQuery) One() (*ReputationLog, error) {
	o := &ReputationLog{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for reputation_log")
	}

	return o, nil
}

// AllP returns all ReputationLog records from the query, and panics on error.
func (q reputationLogQuery) AllP() ReputationLogSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all ReputationLog records from the query.
func (q reputationLogQuery) All() (ReputationLogSlice, error) {
	var o []*ReputationLog

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to ReputationLog slice")
	}

	return o, nil
}

// CountP returns the count of all ReputationLog records in the query, and panics on error.
func (q reputationLogQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all ReputationLog records in the query.
func (q reputationLogQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count reputation_log rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q reputationLogQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q reputationLogQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if reputation_log exists")
	}

	return count > 0, nil
}

// ReputationLogsG retrieves all records.
func ReputationLogsG(mods ...qm.QueryMod) reputationLogQuery {
	return ReputationLogs(boil.GetDB(), mods...)
}

// ReputationLogs retrieves all the records using an executor.
func ReputationLogs(exec boil.Executor, mods ...qm.QueryMod) reputationLogQuery {
	mods = append(mods, qm.From("\"reputation_log\""))
	return reputationLogQuery{NewQuery(exec, mods...)}
}

// FindReputationLogG retrieves a single record by ID.
func FindReputationLogG(id int64, selectCols ...string) (*ReputationLog, error) {
	return FindReputationLog(boil.GetDB(), id, selectCols...)
}

// FindReputationLogGP retrieves a single record by ID, and panics on error.
func FindReputationLogGP(id int64, selectCols ...string) *ReputationLog {
	retobj, err := FindReputationLog(boil.GetDB(), id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindReputationLog retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindReputationLog(exec boil.Executor, id int64, selectCols ...string) (*ReputationLog, error) {
	reputationLogObj := &ReputationLog{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"reputation_log\" where \"id\"=$1", sel,
	)

	q := queries.Raw(exec, query, id)

	err := q.Bind(reputationLogObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from reputation_log")
	}

	return reputationLogObj, nil
}

// FindReputationLogP retrieves a single record by ID with an executor, and panics on error.
func FindReputationLogP(exec boil.Executor, id int64, selectCols ...string) *ReputationLog {
	retobj, err := FindReputationLog(exec, id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *ReputationLog) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *ReputationLog) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *ReputationLog) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *ReputationLog) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no reputation_log provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(reputationLogColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	reputationLogInsertCacheMut.RLock()
	cache, cached := reputationLogInsertCache[key]
	reputationLogInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			reputationLogColumns,
			reputationLogColumnsWithDefault,
			reputationLogColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(reputationLogType, reputationLogMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(reputationLogType, reputationLogMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"reputation_log\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"reputation_log\" DEFAULT VALUES"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		if len(wl) != 0 {
			cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into reputation_log")
	}

	if !cached {
		reputationLogInsertCacheMut.Lock()
		reputationLogInsertCache[key] = cache
		reputationLogInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single ReputationLog record. See Update for
// whitelist behavior description.
func (o *ReputationLog) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single ReputationLog record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *ReputationLog) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the ReputationLog, and panics on error.
// See Update for whitelist behavior description.
func (o *ReputationLog) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the ReputationLog.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *ReputationLog) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	key := makeCacheKey(whitelist, nil)
	reputationLogUpdateCacheMut.RLock()
	cache, cached := reputationLogUpdateCache[key]
	reputationLogUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(
			reputationLogColumns,
			reputationLogPrimaryKeyColumns,
			whitelist,
		)

		if len(whitelist) == 0 {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("models: unable to update reputation_log, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"reputation_log\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, reputationLogPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(reputationLogType, reputationLogMapping, append(wl, reputationLogPrimaryKeyColumns...))
		if err != nil {
			return err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err = exec.Exec(cache.query, values...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update reputation_log row")
	}

	if !cached {
		reputationLogUpdateCacheMut.Lock()
		reputationLogUpdateCache[key] = cache
		reputationLogUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q reputationLogQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q reputationLogQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for reputation_log")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o ReputationLogSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o ReputationLogSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o ReputationLogSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ReputationLogSlice) UpdateAll(exec boil.Executor, cols M) error {
	ln := int64(len(o))
	if ln == 0 {
		return nil
	}

	if len(cols) == 0 {
		return errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), reputationLogPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"reputation_log\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, reputationLogPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in reputationLog slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *ReputationLog) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *ReputationLog) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *ReputationLog) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *ReputationLog) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no reputation_log provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(reputationLogColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs postgres problems
	buf := strmangle.GetBuffer()

	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range updateColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range whitelist {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	reputationLogUpsertCacheMut.RLock()
	cache, cached := reputationLogUpsertCache[key]
	reputationLogUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := strmangle.InsertColumnSet(
			reputationLogColumns,
			reputationLogColumnsWithDefault,
			reputationLogColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		update := strmangle.UpdateColumnSet(
			reputationLogColumns,
			reputationLogPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert reputation_log, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(reputationLogPrimaryKeyColumns))
			copy(conflict, reputationLogPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"reputation_log\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(reputationLogType, reputationLogMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(reputationLogType, reputationLogMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert reputation_log")
	}

	if !cached {
		reputationLogUpsertCacheMut.Lock()
		reputationLogUpsertCache[key] = cache
		reputationLogUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteP deletes a single ReputationLog record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *ReputationLog) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single ReputationLog record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *ReputationLog) DeleteG() error {
	if o == nil {
		return errors.New("models: no ReputationLog provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single ReputationLog record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *ReputationLog) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single ReputationLog record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ReputationLog) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no ReputationLog provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), reputationLogPrimaryKeyMapping)
	sql := "DELETE FROM \"reputation_log\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from reputation_log")
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q reputationLogQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q reputationLogQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no reputationLogQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from reputation_log")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o ReputationLogSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o ReputationLogSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no ReputationLog slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o ReputationLogSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ReputationLogSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no ReputationLog slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), reputationLogPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"reputation_log\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, reputationLogPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from reputationLog slice")
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *ReputationLog) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *ReputationLog) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *ReputationLog) ReloadG() error {
	if o == nil {
		return errors.New("models: no ReputationLog provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *ReputationLog) Reload(exec boil.Executor) error {
	ret, err := FindReputationLog(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *ReputationLogSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *ReputationLogSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ReputationLogSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty ReputationLogSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ReputationLogSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	reputationLogs := ReputationLogSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), reputationLogPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"reputation_log\".* FROM \"reputation_log\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, reputationLogPrimaryKeyColumns, len(*o))

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&reputationLogs)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ReputationLogSlice")
	}

	*o = reputationLogs

	return nil
}

// ReputationLogExists checks if the ReputationLog row exists.
func ReputationLogExists(exec boil.Executor, id int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"reputation_log\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, id)
	}

	row := exec.QueryRow(sql, id)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if reputation_log exists")
	}

	return exists, nil
}

// ReputationLogExistsG checks if the ReputationLog row exists.
func ReputationLogExistsG(id int64) (bool, error) {
	return ReputationLogExists(boil.GetDB(), id)
}

// ReputationLogExistsGP checks if the ReputationLog row exists. Panics on error.
func ReputationLogExistsGP(id int64) bool {
	e, err := ReputationLogExists(boil.GetDB(), id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// ReputationLogExistsP checks if the ReputationLog row exists. Panics on error.
func ReputationLogExistsP(exec boil.Executor, id int64) bool {
	e, err := ReputationLogExists(exec, id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
