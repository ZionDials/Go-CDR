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

	"github.com/ziondials/go-cdr/logger"
	"github.com/ziondials/go-cdr/models"
)

func ParseOracleCDRFile(inputFile string) ([]*models.CubeCDR, error) {

	logger.Info("Parsing file: %s", inputFile)

	readFile, err := os.OpenFile(inputFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		logger.Error("Error opening file: %s Error: %s", inputFile, err)
	}
	defer readFile.Close()

	var ParsedFilename *string
	baseFileName := filepath.Base(inputFile)

	Filename := strings.Split(baseFileName, ".")[0]
	if Filename == "" {
		ParsedFilename = nil
	} else {
		ParsedFilename = &Filename
	}
	Hostname := strings.Split(baseFileName, ".")[1]
	FileTimestamp := strings.ReplaceAll(strings.Split(baseFileName, ".")[2]+"."+strings.Split(baseFileName, ".")[3], "_", " ")

	logger.Info("Parsing Gateway %s CDR file", Hostname)

	rawcdrs := []*models.RawCubeCDR{}
	parsedCDRs := []*models.CubeCDR{}

	reader := csv.NewReader(readFile)
	for {
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

		if len(record) >= 129 {
			rawcdrs = append(rawcdrs, &models.RawCubeCDR{
				RecordTimestamp:          &record[0],
				CallId:                   &record[1],
				CdrType:                  &record[2],
				LegType:                  &record[3],
				H323ConfId:               &record[4],
				PeerAddress:              &record[5],
				PeerSubAddress:           &record[6],
				H323SetupTime:            &record[7],
				AlertTime:                &record[8],
				H323ConnectTime:          &record[9],
				H323DisconnectTime:       &record[10],
				H323DisconnectCause:      &record[11],
				DisconnectText:           &record[12],
				H323CallOrigin:           &record[13],
				ChargedUnits:             &record[14],
				InfoType:                 &record[15],
				PaksOut:                  &record[16],
				BytesOut:                 &record[17],
				PaksIn:                   &record[18],
				BytesIn:                  &record[19],
				Username:                 &record[20],
				Clid:                     &record[21],
				Dnis:                     &record[22],
				GtdOrigCic:               &record[23],
				GtdTermCic:               &record[24],
				TxDuration:               &record[25],
				PeerId:                   &record[26],
				PeerIfIndex:              &record[27],
				LogicalIfIndex:           &record[28],
				AcomLevel:                &record[29],
				NoiseLevel:               &record[30],
				VoiceTxDuration:          &record[31],
				AccountCode:              &record[32],
				CodecBytes:               &record[33],
				CodecTypeRate:            &record[34],
				OntimeRvPlayout:          &record[35],
				RemoteUdpPort:            &record[36],
				RemoteMediaUdpPort:       &record[37],
				VadEnable:                &record[38],
				ReceiveDelay:             &record[39],
				RoundTripDelay:           &record[40],
				HiwaterPlayoutDelay:      &record[41],
				LowaterPlayoutDelay:      &record[42],
				GapfillWithInterpolation: &record[43],
				GapfillWithRedundancy:    &record[44],
				GapfillWithSilence:       &record[45],
				GapfillWithPrediction:    &record[46],
				EarlyPackets:             &record[47],
				LatePackets:              &record[48],
				LostPackets:              &record[49],
				MaxBitrate:               &record[50],
				FaxrelayStartTime:        &record[51],
				FaxrelayStopTime:         &record[52],
				FaxrelayMaxJitBufDepth:   &record[53],
				FaxrelayJitBufOvflow:     &record[54],
				FaxrelayInitHsMod:        &record[55],
				FaxrelayMrHsMod:          &record[56],
				FaxrelayNumPages:         &record[57],
				FaxrelayTxPackets:        &record[58],
				FaxrelayRxPackets:        &record[59],
				FaxrelayDirection:        &record[60],
				FaxrelayPktConceal:       &record[61],
				FaxrelayEcmStatus:        &record[62],
				FaxrelayEncapProtocol:    &record[63],
				FaxrelayNsfCountryCode:   &record[64],
				FaxrelayNsfManufCode:     &record[65],
				FaxrelayFaxSuccess:       &record[66],
				OverrideSessionTime:      &record[67],
				H323IvrOut:               &record[68],
				InternalErrorCode:        &record[69],
				H323VoiceQuality:         &record[70],
				RemoteMediaAddress:       &record[71],
				RemoteMediaId:            &record[72],
				CarrierId:                &record[73],
				CallingPartyCategory:     &record[74],
				OriginatingLineInfo:      &record[75],
				ChargeNumber:             &record[76],
				TransmissionMediumReq:    &record[77],
				ServiceDescriptor:        &record[78],
				OutgoingArea:             &record[79],
				IncomingArea:             &record[80],
				OutTrunkgroupLabel:       &record[81],
				OutCarrierId:             &record[82],
				DspId:                    &record[83],
				InTrunkgroupLabel:        &record[84],
				InCarrierId:              &record[85],
				CustBizGrpId:             &record[86],
				SuppSvcXferBy:            &record[87],
				VoiceFeature:             &record[88],
				FeatureOperation:         &record[89],
				FeatureOpStatus:          &record[90],
				FeatureOpTime:            &record[91],
				FeatureId:                &record[92],
				GwRxdCdn:                 &record[93],
				GwRxdCgn:                 &record[94],
				GtdGwRxdOcn:              &record[95],
				GtdGwRxdCnn:              &record[96],
				GwRxdRdn:                 &record[97],
				GwFinalXlatedCdn:         &record[98],
				GwFinalXlatedCgn:         &record[99],
				GwFinalXlatedRdn:         &record[100],
				GkXlatedCdn:              &record[101],
				GkXlatedCgn:              &record[102],
				GwCollectedCdn:           &record[103],
				IPHop:                    &record[104],
				RedirectedStation:        &record[105],
				Subscriber:               &record[106],
				InIntrfcDesc:             &record[107],
				OutIntrfcDesc:            &record[108],
				SessionProtocol:          &record[109],
				LocalHostname:            &record[110],
				BackwardCallId:           &record[111],
				FeatureIdField1:          &record[112],
				FeatureIdField2:          &record[113],
				FeatureIdField3:          &record[114],
				FeatureIdField4:          &record[115],
				FeatureIdField5:          &record[116],
				FeatureIdField6:          &record[117],
				FeatureIdField7:          &record[118],
				FeatureIdField8:          &record[119],
				FeatureIdField9:          &record[120],
				FeatureIdField10:         &record[121],
				FeatureIdField11:         &record[122],
				FeatureIdField12:         &record[123],
				IpPhoneInfo:              &record[124],
				IpPbxMode:                &record[125],
				InLpcorGroup:             &record[126],
				OutLpcorGroup:            &record[127],
				FacDigit:                 &record[128],
				FacStatus:                &record[129],
				Hostname:                 &Hostname,
				Filename:                 ParsedFilename,
				FileTimestamp:            &FileTimestamp,
			})
		} else if len(record) != 1 {
			logger.Error("Error parsing CDR: %s Found %s fields instead of 129", inputFile, strconv.Itoa(len(record)))
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
