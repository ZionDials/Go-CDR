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
	"encoding/csv"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/ziondials/go-cdr/helpers"
	"github.com/ziondials/go-cdr/logger"
	"github.com/ziondials/go-cdr/models"
)

func ParseCucmCMRFile(inputFile string) ([]*models.CucmCmr, error) {

	logger.Info("Parsing file: %s", inputFile)

	readFile, err := os.OpenFile(inputFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		logger.Error("Error opening file: %s Error: %s", inputFile, err)
	}
	defer readFile.Close()

	baseFileName := filepath.Base(inputFile)

	ParsedFilename := strings.Split(baseFileName, "_")

	FilenameClusterID := ParsedFilename[1]
	FilenameNodeID := ParsedFilename[2]
	FilenameDateTime, err := helpers.ParseCUCMFilenameTimestamp(ParsedFilename[3])
	if err != nil {
		logger.Error("Error parsing file date time: %s Error: %s", inputFile, err)
	}
	FilenameSequence, err := helpers.ConvertStringToInt64(&ParsedFilename[4])
	if err != nil {
		logger.Error("Error parsing file sequence: %s Error: %s", inputFile, err)
	}

	rawcdrs := []*models.RawCucmCmr{}
	parsedCDRs := []*models.CucmCmr{}

	lineCount := 0

	reader := csv.NewReader(readFile)
	for {
		lineCount++
		record, err := reader.Read()
		if err == io.EOF {
			logger.Info("Finished parsing file: %s", inputFile)
			break
		}
		if err != nil {
			if perr, ok := err.(*csv.ParseError); ok && perr.Err == csv.ErrFieldCount {
				continue
			}
			logger.Error("Error parsing file: %s Error: %s", inputFile, err)
		}

		if len(record) >= 44 && lineCount > 2 {
			rawcdrs = append(rawcdrs, &models.RawCucmCmr{
				Cdrrecordtype:                       &record[0],
				Globalcallid_Callmanagerid:          &record[1],
				Globalcallid_Callid:                 &record[2],
				Nodeid:                              &record[3],
				Directorynum:                        &record[4],
				Callidentifier:                      &record[5],
				Datetimestamp:                       &record[6],
				Numberpacketssent:                   &record[7],
				Numberoctetssent:                    &record[8],
				Numberpacketsreceived:               &record[9],
				Numberoctetsreceived:                &record[10],
				Numberpacketslost:                   &record[11],
				Jitter:                              &record[12],
				Latency:                             &record[13],
				Pkid:                                &record[14],
				Directorynumpartition:               &record[15],
				Globalcallid_Clusterid:              &record[16],
				Devicename:                          &record[17],
				Varvqmetrics:                        &record[18],
				Duration:                            &record[19],
				Videocontenttype:                    &record[20],
				Videoduration:                       &record[21],
				Numbervideopacketssent:              &record[22],
				Numbervideooctetssent:               &record[23],
				Numbervideopacketsreceived:          &record[24],
				Numbervideooctetsreceived:           &record[25],
				Numbervideopacketslost:              &record[26],
				Videoaveragejitter:                  &record[27],
				Videoroundtriptime:                  &record[28],
				Videoonewaydelay:                    &record[29],
				Videoreceptionmetrics:               &record[30],
				Videotransmissionmetrics:            &record[31],
				Videocontenttype_Channel2:           &record[32],
				Videoduration_Channel2:              &record[33],
				Numbervideopacketssent_Channel2:     &record[34],
				Numbervideooctetssent_Channel2:      &record[35],
				Numbervideopacketsreceived_Channel2: &record[36],
				Numbervideooctetsreceived_Channel2:  &record[37],
				Numbervideopacketslost_Channel2:     &record[38],
				Videoaveragejitter_Channel2:         &record[39],
				Videoroundtriptime_Channel2:         &record[40],
				Videoonewaydelay_Channel2:           &record[41],
				Videoreceptionmetrics_Channel2:      &record[42],
				Videotransmissionmetrics_Channel2:   &record[43],
				FileClusterId:                       &FilenameClusterID,
				FileNodeId:                          &FilenameNodeID,
				FileDateTime:                        FilenameDateTime,
				FileSequenceNumber:                  FilenameSequence,
			})
		} else if len(record) != 1 && lineCount > 2 {
			logger.Error("Error parsing CDR: %s Found %s fields instead of equal to or greater than 44", inputFile, strconv.Itoa(len(record)))
		}
	}

	for _, cdr := range rawcdrs {
		pCDR, err := cdr.Parse(inputFile)
		if err == nil {
			parsedCDRs = append(parsedCDRs, pCDR)
			continue
		}
	}

	return parsedCDRs, nil
}
