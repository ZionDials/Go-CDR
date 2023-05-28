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

package parser

import (
	"path/filepath"
	"strconv"

	"github.com/ziondials/go-cdr/database"
	"github.com/ziondials/go-cdr/helpers"
	"github.com/ziondials/go-cdr/logger"
)

func ParseCUBECDRs(inputFile string, db *database.DataService) {

	baseFileName := filepath.Base(inputFile)

	logger.Info("Found CDR file: %s", baseFileName)
	cdrs, err := ParseCubeCDRFile(inputFile)
	if err != nil {
		logger.Error("Error parsing file: %s Error: %s\n", inputFile, err)
		helpers.ChangeFileNameToFailedAndMove(inputFile)
	}

	if len(cdrs) > 0 && err == nil {

		err := db.CreateCubeCDRs(cdrs)
		if err != nil {
			logger.Error("Error while writing to database: %s", err.Error())
			err := helpers.ChangeFileNameToFailedAndMove(inputFile)
			if err != nil {
				logger.Error("Error while moving file: %s", err.Error())
			} else {
				logger.Info("Successfully moved file to failed directory: %s", inputFile)
			}
		} else {
			logger.Info("Successfully wrote %s CDRs to database from %s", strconv.Itoa(len(cdrs)), inputFile)
			err := helpers.ChangeFileNameToCompleteAndMove(inputFile)
			if err != nil {
				logger.Error("Error while moving file: %s", err.Error())
			} else {
				logger.Info("Successfully moved file to completed directory: %s", inputFile)
			}
		}
	}

}
