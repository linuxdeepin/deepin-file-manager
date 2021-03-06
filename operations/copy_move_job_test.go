/**
 * Copyright (C) 2015 Deepin Technology Co., Ltd.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 **/

package operations_test

import (
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"path/filepath"
	"gir/gio-2.0"
	. "pkg.deepin.io/service/file-manager-backend/operations"
	"testing"
)

// TODO: clean test directories
func TestCopyJob(t *testing.T) {
	SkipConvey("copy a file", t, func() {
		os.Setenv("LANGUAGE", "en_US")
		srcFilePath, _ := filepath.Abs("./testdata/copy/src/afile")
		destPath, _ := filepath.Abs("./testdata/copy/dest")

		job := NewCopyJob([]string{srcFilePath}, destPath, "", gio.FileCopyFlagsNone, nil)
		job.Execute()

		copyedFileURL, _ := filepath.Abs("./testdata/copy/dest/afile")
		copyedFile := gio.FileNewForCommandlineArg(copyedFileURL)
		So(copyedFile.QueryExists(nil), ShouldBeTrue)
	})

	SkipConvey("copy a exists file", t, func() {
		os.Setenv("LANGUAGE", "en_US")
		srcFilePath, _ := filepath.Abs("./testdata/copy/src/exsitfile")
		destPath, _ := filepath.Abs("./testdata/copy/dest")

		job := NewCopyJob([]string{srcFilePath}, destPath, "", gio.FileCopyFlagsNone, renameMock)
		job.Execute()

		copyedFileURL, _ := filepath.Abs("./testdata/copy/dest/exsitfile")
		copyedFile := gio.FileNewForCommandlineArg(copyedFileURL)
		So(copyedFile.QueryExists(nil), ShouldBeTrue)
	})

	SkipConvey("copy a dir", t, func() {
		os.Setenv("LANGUAGE", "en_US")
		srcFilePath, _ := filepath.Abs("./testdata/copy/src/adir")
		destPath, _ := filepath.Abs("./testdata/copy/dest")

		job := NewCopyJob([]string{srcFilePath}, destPath, "", gio.FileCopyFlagsNone, nil)
		job.Execute()

		copyedFileURL, _ := filepath.Abs("./testdata/copy/dest/adir")
		copyedFile := gio.FileNewForCommandlineArg(copyedFileURL)
		So(copyedFile.QueryExists(nil), ShouldBeTrue)
	})

	SkipConvey("copy a exists dir", t, func() {
		os.Setenv("LANGUAGE", "en_US")
		srcFilePath, _ := filepath.Abs("./testdata/copy/src/exsitdir")
		destPath, _ := filepath.Abs("./testdata/copy/dest")

		job := NewCopyJob([]string{srcFilePath}, destPath, "", gio.FileCopyFlagsNone, renameMock)
		job.Execute()

		copyedFileURL, _ := filepath.Abs("./testdata/copy/dest/adir")
		copyedFile := gio.FileNewForCommandlineArg(copyedFileURL)
		So(copyedFile.QueryExists(nil), ShouldBeTrue)
	})

	SkipConvey("dup a file", t, func() {
		os.Setenv("LANGUAGE", "en_US")
		srcFilePath, _ := filepath.Abs("./testdata/copy/dest/exsitfile")
		// destPath, _ := filepath.Abs("./testdata/copy/dest")

		job := NewCopyJob([]string{srcFilePath}, "", "", gio.FileCopyFlagsNone, skipMock)
		job.Execute()
	})
}

// TODO
func TestMoveJob(t *testing.T) {
}
