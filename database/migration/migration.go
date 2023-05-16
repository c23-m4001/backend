package migration

import (
	"bytes"
	"errors"
	"io"
	"io/fs"
	"strconv"

	"github.com/golang-migrate/migrate/v4/source"
)

var sourceDriver *inMemoryDriver = &inMemoryDriver{migrations: source.NewMigrations()}

func SourceDriver() source.Driver {
	return sourceDriver
}

type inMemoryDriver struct {
	migrations *source.Migrations
}

func (drv *inMemoryDriver) append(version uint, rawUp, rawDown string) bool {
	var (
		migrationUp = &source.Migration{
			Version:   version,
			Direction: source.Up,
			Raw:       rawUp,
		}
		migrationDown = &source.Migration{
			Version:   version,
			Direction: source.Down,
			Raw:       rawDown,
		}
	)

	return drv.migrations.Append(migrationUp) && drv.migrations.Append(migrationDown)
}

// Open is part of source.Driver interface implementation.
// Open cannot be called on the driver.
func (drv *inMemoryDriver) Open(url string) (source.Driver, error) {
	return nil, errors.New("Open() cannot be called on the driver")
}

// Close is part of source.Driver interface implementation.
func (drv *inMemoryDriver) Close() error {
	return nil
}

// First is part of source.Driver interface implementation.
func (drv *inMemoryDriver) First() (version uint, err error) {
	if version, ok := drv.migrations.First(); ok {
		return version, nil
	}
	return 0, &fs.PathError{
		Op:   "first",
		Path: "",
		Err:  fs.ErrNotExist,
	}
}

// Prev is part of source.Driver interface implementation.
func (drv *inMemoryDriver) Prev(version uint) (prevVersion uint, err error) {
	if version, ok := drv.migrations.Prev(version); ok {
		return version, nil
	}
	return 0, &fs.PathError{
		Op:   "prev for version " + strconv.FormatUint(uint64(version), 10),
		Path: "",
		Err:  fs.ErrNotExist,
	}
}

// Next is part of source.Driver interface implementation.
func (drv *inMemoryDriver) Next(version uint) (nextVersion uint, err error) {
	if version, ok := drv.migrations.Next(version); ok {
		return version, nil
	}
	return 0, &fs.PathError{
		Op:   "next for version " + strconv.FormatUint(uint64(version), 10),
		Path: "",
		Err:  fs.ErrNotExist,
	}
}

// ReadUp is part of source.Driver interface implementation.
func (drv *inMemoryDriver) ReadUp(version uint) (r io.ReadCloser, identifier string, err error) {
	if m, ok := drv.migrations.Up(version); ok {
		return newReadCloser(m.Raw), m.Identifier, nil
	}
	return nil, "", &fs.PathError{
		Op:   "read up for version " + strconv.FormatUint(uint64(version), 10),
		Path: "",
		Err:  fs.ErrNotExist,
	}
}

// ReadDown is part of source.Driver interface implementation.
func (drv *inMemoryDriver) ReadDown(version uint) (r io.ReadCloser, identifier string, err error) {
	if m, ok := drv.migrations.Down(version); ok {
		return newReadCloser(m.Raw), m.Identifier, nil
	}
	return nil, "", &fs.PathError{
		Op:   "read down for version " + strconv.FormatUint(uint64(version), 10),
		Path: "",
		Err:  fs.ErrNotExist,
	}
}

func newReadCloser(s string) io.ReadCloser {
	return io.NopCloser(
		bytes.NewBuffer([]byte(s)),
	)
}
