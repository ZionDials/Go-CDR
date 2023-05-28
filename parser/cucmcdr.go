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

func ParseCucmCDRFile(inputFile string) ([]*models.CucmCdr, error) {

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

	rawcdrs := []*models.RawCucmCdr{}
	parsedCDRs := []*models.CucmCdr{}

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

		if len(record) == 133 && lineCount > 2 {
			rawcdrs = append(rawcdrs, &models.RawCucmCdr{
				Cdrrecordtype:                           &record[0],
				Globalcallid_Callmanagerid:              &record[1],
				Globalcallid_Callid:                     &record[2],
				Origlegcallidentifier:                   &record[3],
				Datetimeorigination:                     &record[4],
				Orignodeid:                              &record[5],
				Origspan:                                &record[6],
				Origipaddr:                              &record[7],
				Callingpartynumber:                      &record[8],
				Callingpartyunicodeloginuserid:          &record[9],
				Origcause_Location:                      &record[10],
				Origcause_Value:                         &record[11],
				Origprecedencelevel:                     &record[12],
				Origmediatransportaddress_IP:            &record[13],
				Origmediatransportaddress_Port:          &record[14],
				Origmediacap_Payloadcapability:          &record[15],
				Origmediacap_Maxframesperpacket:         &record[16],
				Origmediacap_G723bitrate:                &record[17],
				Origvideocap_Codec:                      &record[18],
				Origvideocap_Bandwidth:                  &record[19],
				Origvideocap_Resolution:                 &record[20],
				Origvideotransportaddress_IP:            &record[21],
				Origvideotransportaddress_Port:          &record[22],
				Origrsvpaudiostat:                       &record[23],
				Origrsvpvideostat:                       &record[24],
				Destlegcallidentifier:                   &record[25],
				Destnodeid:                              &record[26],
				Destspan:                                &record[27],
				Destipaddr:                              &record[28],
				Originalcalledpartynumber:               &record[29],
				Finalcalledpartynumber:                  &record[30],
				Finalcalledpartyunicodeloginuserid:      &record[31],
				Destcause_Location:                      &record[32],
				Destcause_Value:                         &record[33],
				Destprecedencelevel:                     &record[34],
				Destmediatransportaddress_IP:            &record[35],
				Destmediatransportaddress_Port:          &record[36],
				Destmediacap_Payloadcapability:          &record[37],
				Destmediacap_Maxframesperpacket:         &record[38],
				Destmediacap_G723bitrate:                &record[39],
				Destvideocap_Codec:                      &record[40],
				Destvideocap_Bandwidth:                  &record[41],
				Destvideocap_Resolution:                 &record[42],
				Destvideotransportaddress_IP:            &record[43],
				Destvideotransportaddress_Port:          &record[44],
				Destrsvpaudiostat:                       &record[45],
				Destrsvpvideostat:                       &record[46],
				Datetimeconnect:                         &record[47],
				Datetimedisconnect:                      &record[48],
				Lastredirectdn:                          &record[49],
				Pkid:                                    &record[50],
				Originalcalledpartynumberpartition:      &record[51],
				Callingpartynumberpartition:             &record[52],
				Finalcalledpartynumberpartition:         &record[53],
				Lastredirectdnpartition:                 &record[54],
				Duration:                                &record[55],
				Origdevicename:                          &record[56],
				Destdevicename:                          &record[57],
				Origcallterminationonbehalfof:           &record[58],
				Destcallterminationonbehalfof:           &record[59],
				Origcalledpartyredirectonbehalfof:       &record[60],
				Lastredirectredirectonbehalfof:          &record[61],
				Origcalledpartyredirectreason:           &record[62],
				Lastredirectredirectreason:              &record[63],
				Destconversationid:                      &record[64],
				Globalcallid_Clusterid:                  &record[65],
				Joinonbehalfof:                          &record[66],
				Comment:                                 &record[67],
				Authcodedescription:                     &record[68],
				Authorizationlevel:                      &record[69],
				Clientmattercode:                        &record[70],
				Origdtmfmethod:                          &record[71],
				Destdtmfmethod:                          &record[72],
				Callsecuredstatus:                       &record[73],
				Origconversationid:                      &record[74],
				Origmediacap_Bandwidth:                  &record[75],
				Destmediacap_Bandwidth:                  &record[76],
				Authorizationcodevalue:                  &record[77],
				Outpulsedcallingpartynumber:             &record[78],
				Outpulsedcalledpartynumber:              &record[79],
				Origipv4v6addr:                          &record[80],
				Destipv4v6addr:                          &record[81],
				Origvideocap_Codec_Channel2:             &record[82],
				Origvideocap_Bandwidth_Channel2:         &record[83],
				Origvideocap_Resolution_Channel2:        &record[84],
				Origvideotransportaddress_IP_Channel2:   &record[85],
				Origvideotransportaddress_Port_Channel2: &record[86],
				Origvideochannel_Role_Channel2:          &record[87],
				Destvideocap_Codec_Channel2:             &record[88],
				Destvideocap_Bandwidth_Channel2:         &record[89],
				Destvideocap_Resolution_Channel2:        &record[90],
				Destvideotransportaddress_IP_Channel2:   &record[91],
				Destvideotransportaddress_Port_Channel2: &record[92],
				Destvideochannel_Role_Channel2:          &record[93],
				Incomingprotocolid:                      &record[94],
				Incomingprotocolcallref:                 &record[95],
				Outgoingprotocolid:                      &record[96],
				Outgoingprotocolcallref:                 &record[97],
				Currentroutingreason:                    &record[98],
				Origroutingreason:                       &record[99],
				Lastredirectingroutingreason:            &record[100],
				Huntpilotpartition:                      &record[101],
				Huntpilotdn:                             &record[102],
				Calledpartypatternusage:                 &record[103],
				Incomingicid:                            &record[104],
				Incomingorigioi:                         &record[105],
				Incomingtermioi:                         &record[106],
				Outgoingicid:                            &record[107],
				Outgoingorigioi:                         &record[108],
				Outgoingtermioi:                         &record[109],
				Outpulsedoriginalcalledpartynumber:      &record[110],
				Outpulsedlastredirectingnumber:          &record[111],
				Wascallqueued:                           &record[112],
				Totalwaittimeinqueue:                    &record[113],
				Callingpartynumber_Uri:                  &record[114],
				Originalcalledpartynumber_Uri:           &record[115],
				Finalcalledpartynumber_Uri:              &record[116],
				Lastredirectdn_Uri:                      &record[117],
				Mobilecallingpartynumber:                &record[118],
				Finalmobilecalledpartynumber:            &record[119],
				Origmobiledevicename:                    &record[120],
				Destmobiledevicename:                    &record[121],
				Origmobilecallduration:                  &record[122],
				Destmobilecallduration:                  &record[123],
				Mobilecalltype:                          &record[124],
				Originalcalledpartypattern:              &record[125],
				Finalcalledpartypattern:                 &record[126],
				Lastredirectingpartypattern:             &record[127],
				Huntpilotpattern:                        &record[128],
				FileClusterId:                           &FilenameClusterID,
				FileNodeId:                              &FilenameNodeID,
				FileDateTime:                            FilenameDateTime,
				FileSequenceNumber:                      FilenameSequence,
			})
		} else if len(record) != 1 && lineCount > 2 {
			logger.Error("Error parsing CDR: %s Found %s fields instead of 132", inputFile, strconv.Itoa(len(record)))
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
