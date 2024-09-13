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
	"os"
	"path/filepath"

	"github.com/ziondials/go-cdr/database"
	"github.com/ziondials/go-cdr/logger"
)

func ParseFiles(inputDirectory string, outputDirectory string, fileType string, deleteOriginal bool, db *database.DataService) {
	// Get a list of files in the input directory
	files, err := os.ReadDir(inputDirectory)
	if err != nil {
		logger.Error("Error reading directory: %s Error: %s", inputDirectory, err)
		os.Exit(1)
	}

	logger.Info("Parsing files in directory: %s", inputDirectory)

	// TimestampedFilename := fileType + "_" + helpers.FilenameFriendlyTimeStamp() + ".csv"

	// OutputFile := filepath.Join(outputDirectory, TimestampedFilename)

	// writeFile, err := os.OpenFile(OutputFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	// if err != nil {
	//	logger.Error("Error opening file: %s Error: %s", OutputFile, err)
	// }
	// defer writeFile.Close()

	// Loop through the files in the input directory
	for _, file := range files {

		// Check if the file is a directory
		if !file.IsDir() {

			fullFilePath := filepath.Join(inputDirectory, file.Name())
			switch fileType {
			case "cube":
				ParseCUBECDRs(fullFilePath, db, outputDirectory, deleteOriginal)
			case "cucm":
				ParseCUCMCDRs(fullFilePath, db, outputDirectory, deleteOriginal)
			default:
				// Failed to match a file type
				logger.Error("Failed to match file type: %s", fileType)
				os.Exit(1)

			}
		}
	}

	logger.Info("Finished parsing files in directory: %s", inputDirectory)
}
