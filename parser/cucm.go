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

func ParseCUCMCDRs(inputFile string, db *database.DataService, outputDirectory string, deleteOriginal bool) {

	baseFileName := filepath.Base(inputFile)

	if helpers.CMRReg.MatchString(baseFileName) {
		logger.Info("Found CMR file: %s", baseFileName)
		cdrs, err := ParseCucmCMRFile(inputFile)
		if err != nil {
			logger.Error("Error parsing file: %s Error: %s", inputFile, err)
		}

		if len(cdrs) > 0 {

			err := db.CreateCucmCMRs(cdrs)
			if err != nil {
				logger.Error("Error while writing to database: %s", err.Error())
			} else {
				logger.Info("Successfully wrote %s CDRs to database from %s", strconv.Itoa(len(cdrs)), inputFile)
				err := helpers.ChangeFileNameToCompleteAndMoveOrDelete(inputFile, outputDirectory, deleteOriginal)
				if err != nil {
					logger.Error("Error while moving file: %s", err.Error())
				} else {
					logger.Info("Successfully moved file to failed directory: %s", inputFile)
				}
			}
		} else if len(cdrs) == 0 && err == nil {
			logger.Info("No CDRs found in file: %s", inputFile)
			err := helpers.ChangeFileNameToCompleteAndMoveOrDelete(inputFile, outputDirectory, deleteOriginal)
			if err != nil {
				logger.Error("Error while moving file: %s", err.Error())
			} else {
				logger.Info("Successfully moved file to completed directory: %s", inputFile)
			}
		} else {
			logger.Error("Error parsing file: %s Error: %s", inputFile, err)
			err := helpers.ChangeFileNameToFailedAndMove(inputFile, outputDirectory)
			if err != nil {
				logger.Error("Error while moving file: %s", err.Error())
			} else {
				logger.Info("Successfully moved file to failed directory: %s", inputFile)
			}
		}
	}

	if helpers.CDRReg.MatchString(baseFileName) {
		logger.Info("Found CDR file: %s", baseFileName)
		cdrs, err := ParseCucmCDRFile(inputFile)
		if err != nil {
			logger.Error("Error parsing file: %s Error: %s", inputFile, err)
			helpers.ChangeFileNameToFailedAndMove(inputFile, outputDirectory)
		}

		if len(cdrs) > 0 && err == nil {

			err := db.CreateCucmCDRs(cdrs)
			if err != nil {
				logger.Error("Error while writing to database: %s", err.Error())
				err := helpers.ChangeFileNameToFailedAndMove(inputFile, outputDirectory)
				if err != nil {
					logger.Error("Error while moving file: %s", err.Error())
				} else {
					logger.Info("Successfully moved file to failed directory: %s", inputFile)
				}
			} else {
				logger.Info("Successfully wrote %s CDRs to database from %s", strconv.Itoa(len(cdrs)), inputFile)
				err := helpers.ChangeFileNameToCompleteAndMoveOrDelete(inputFile, outputDirectory, deleteOriginal)
				if err != nil {
					logger.Error("Error while moving file: %s", err.Error())
				} else {
					logger.Info("Successfully moved file to completed directory: %s", inputFile)
				}
			}
		} else if len(cdrs) == 0 && err == nil {
			logger.Info("No CDRs found in file: %s", inputFile)
			err := helpers.ChangeFileNameToCompleteAndMoveOrDelete(inputFile, outputDirectory, deleteOriginal)
			if err != nil {
				logger.Error("Error while moving file: %s", err.Error())
			} else {
				logger.Info("Successfully moved file to completed directory: %s", inputFile)
			}
		} else {
			logger.Error("Error parsing file: %s Error: %s", inputFile, err)
			err := helpers.ChangeFileNameToFailedAndMove(inputFile, outputDirectory)
			if err != nil {
				logger.Error("Error while moving file: %s", err.Error())
			} else {
				logger.Info("Successfully moved file to failed directory: %s", inputFile)
			}
		}
	}

}
