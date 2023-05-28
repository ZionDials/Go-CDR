// Copyright (c) 2023 Zion Dials <me@ziondials.com>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

package helpers

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/ziondials/go-cdr/logger"
)

func FilenameFriendlyTimeStamp() string {
	ts := time.Now().UTC().Format(time.RFC3339)
	return strings.Replace(strings.Replace(ts, ":", "", -1), "-", "", -1)
}

func ParseCUCMFilenameTimestamp(input string) (*int64, error) {
	parsedTime, err := time.Parse("200601021504", input)
	if err != nil {
		return nil, err
	}
	ts := parsedTime.Unix()
	return &ts, nil
}

func ChangeFileNameToCompleteAndMove(input string) error {
	OGPath := filepath.Dir(input)
	baseFileName := filepath.Base(input)
	completedFilePath := filepath.Join(OGPath, "complete")
	if _, err := os.Stat(completedFilePath); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(completedFilePath, os.ModePerm)
		if err != nil {
			logger.Error(err.Error())
		}
	}
	NewPath := filepath.Join(OGPath, "complete", (baseFileName + ".complete"))
	err := os.Rename(input, NewPath)
	return err
}

func ChangeFileNameToFailedAndMove(input string) error {
	OGPath := filepath.Dir(input)
	baseFileName := filepath.Base(input)
	failedFilePath := filepath.Join(OGPath, "failed")
	if _, err := os.Stat(failedFilePath); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(failedFilePath, os.ModePerm)
		if err != nil {
			logger.Error(err.Error())
		}
	}
	NewPath := filepath.Join(OGPath, "failed", (baseFileName + ".failed"))
	err := os.Rename(input, NewPath)
	return err
}
