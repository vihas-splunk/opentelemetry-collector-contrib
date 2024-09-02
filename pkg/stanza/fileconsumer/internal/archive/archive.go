// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package archive // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer/internal/archive"

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer/internal/checkpoint"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer/internal/fileset"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer/internal/fingerprint"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer/internal/reader"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
)

const _ = "knownFiles"

type Archive interface {
	SetStorageClient(operator.Persister)
	Match([]*ArchiveFileRecord) ([]*reader.Reader, error)
	Write([]*reader.Metadata) error
}

type ArchiveFileRecord struct {
	file *os.File
	fp   *fingerprint.Fingerprint
}

func NewArchiveRecord(file *os.File, fp *fingerprint.Fingerprint) *ArchiveFileRecord {
	return &ArchiveFileRecord{
		file: file,
		fp:   fp,
	}
}

type archive struct {
	persister      operator.Persister
	pollsToArchive int
	index          int
	_              *fileset.Fileset[*reader.Metadata]
	readerFactory  reader.Factory
}

func NewArchive(pollsToArchive int, readerFactory reader.Factory) Archive {
	return &archive{pollsToArchive: pollsToArchive, readerFactory: readerFactory}
}

func (a *archive) SetStorageClient(persister operator.Persister) {
	a.persister = persister
}

func (a *archive) Match(unmatchedFiles []*ArchiveFileRecord) ([]*reader.Reader, error) {
	// Arguments:
	//		unmatched files
	// Returns:
	//		readers created from old/new metadata
	readers := make([]*reader.Reader, 0)
	var combinedError error
	for i := a.index; i < a.index+a.pollsToArchive; i++ {
		idx := i % a.pollsToArchive
		key := fmt.Sprintf("knownFilesArchive%d", idx)
		archivedMetadata, err := checkpoint.LoadKey(context.Background(), a.persister, key)
		if err != nil {
			return readers, fmt.Errorf("error while loading the archive: %w", err)
		}

		archivedFileset := fileset.New[*reader.Metadata](len(archivedMetadata))
		archivedFileset.Add(archivedMetadata...)

		for j := 0; j < len(unmatchedFiles); {
			record := unmatchedFiles[j]
			if m := archivedFileset.Match(record.fp, fileset.StartsWith); m != nil {
				unmatchedFiles = append(unmatchedFiles[:j], unmatchedFiles[j+1:]...)
				reader, err := a.readerFactory.NewReaderFromMetadata(record.file, m)
				if err != nil {
					combinedError = errors.Join(combinedError, err)
					continue
				}
				readers = append(readers, reader)
				continue
			}
			j += 1
		}

		// rewrite updated archive set to storage
		err = checkpoint.SaveKey(context.Background(), a.persister, archivedFileset.Get(), key)
		if err != nil {
			return readers, fmt.Errorf("error while saving to the archive: %w", err)
		}
	}

	for _, record := range unmatchedFiles {
		r, err := a.readerFactory.NewReader(record.file, record.fp)
		if err != nil {
			combinedError = errors.Join(combinedError, err)
		} else {
			readers = append(readers, r)
		}
	}

	return readers, combinedError
}

func (a *archive) Write(_ []*reader.Metadata) error {
	// TODO:
	// 		Add logic to update the index.
	//	 	Handle rollover logic
	return nil
}
