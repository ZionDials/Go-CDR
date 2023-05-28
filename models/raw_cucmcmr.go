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

package models

import (
	"github.com/google/uuid"
	"github.com/ziondials/go-cdr/helpers"
	"github.com/ziondials/go-cdr/logger"
)

type RawCucmCmr struct {
	Pkid                                *string
	FileClusterId                       *string
	FileNodeId                          *string
	FileDateTime                        *int64
	FileSequenceNumber                  *int64
	Cdrrecordtype                       *string
	Globalcallid_Callmanagerid          *string
	Globalcallid_Callid                 *string
	Nodeid                              *string
	Directorynum                        *string
	Callidentifier                      *string
	Datetimestamp                       *string
	Numberpacketssent                   *string
	Numberoctetssent                    *string
	Numberpacketsreceived               *string
	Numberoctetsreceived                *string
	Numberpacketslost                   *string
	Jitter                              *string
	Latency                             *string
	Directorynumpartition               *string
	Globalcallid_Clusterid              *string
	Devicename                          *string
	Varvqmetrics                        *string
	Duration                            *string
	Videocontenttype                    *string
	Videoduration                       *string
	Numbervideopacketssent              *string
	Numbervideooctetssent               *string
	Numbervideopacketsreceived          *string
	Numbervideooctetsreceived           *string
	Numbervideopacketslost              *string
	Videoaveragejitter                  *string
	Videoroundtriptime                  *string
	Videoonewaydelay                    *string
	Videoreceptionmetrics               *string
	Videotransmissionmetrics            *string
	Videocontenttype_Channel2           *string
	Videoduration_Channel2              *string
	Numbervideopacketssent_Channel2     *string
	Numbervideooctetssent_Channel2      *string
	Numbervideopacketsreceived_Channel2 *string
	Numbervideooctetsreceived_Channel2  *string
	Numbervideopacketslost_Channel2     *string
	Videoaveragejitter_Channel2         *string
	Videoroundtriptime_Channel2         *string
	Videoonewaydelay_Channel2           *string
	Videoreceptionmetrics_Channel2      *string
	Videotransmissionmetrics_Channel2   *string
	Localsessionid                      *string
	Remotesessionid                     *string
	Headsetsn                           *string
	Headsetmetrics                      *string
}

func (raw *RawCucmCmr) Parse(filename string) (*CucmCmr, error) {
	var ParsedCdrrecordtype *int64
	var ParsedGlobalcallid_Callmanagerid *int64
	var ParsedGlobalcallid_Callid *int64
	var ParsedNodeid *int64
	var ParsedCallidentifier *int64
	var ParsedDatetimestamp *int64
	var ParsedNumberpacketssent *int64
	var ParsedNumberoctetssent *int64
	var ParsedNumberpacketsreceived *int64
	var ParsedNumberoctetsreceived *int64
	var ParsedNumberpacketslost *int64
	var ParsedJitter *int64
	var ParsedLatency *int64
	var ParsedDuration *int64
	var ParsedVideoduration *int64
	var ParsedNumbervideopacketssent *int64
	var ParsedNumbervideooctetssent *int64
	var ParsedNumbervideopacketsreceived *int64
	var ParsedNumbervideooctetsreceived *int64
	var ParsedNumbervideopacketslost *int64
	var ParsedVideoaveragejitter *int64
	var ParsedVideoroundtriptime *int64
	var ParsedVideoonewaydelay *int64
	var ParsedVideoduration_Channel2 *int64
	var ParsedNumbervideopacketssent_Channel2 *int64
	var ParsedNumbervideooctetssent_Channel2 *int64
	var ParsedNumbervideopacketsreceived_Channel2 *int64
	var ParsedNumbervideooctetsreceived_Channel2 *int64
	var ParsedNumbervideopacketslost_Channel2 *int64
	var ParsedVideoaveragejitter_Channel2 *int64
	var ParsedVideoroundtriptime_Channel2 *int64
	var ParsedVideoonewaydelay_Channel2 *int64
	var ParsedVQCS *int64
	var ParsedVQSCS *int64
	var ParsedVQCID *int64
	var ParsedVqvopktsizems *int64
	var ParsedVqvopktlost *int64
	var ParsedVqvopktdis *int64
	var ParsedVqvoonewaydelayms *int64
	var ParsedVqmaxjitter *int64

	var ParsedOriginpkid *string
	var ParsedFileClusterId *string
	var ParsedFileNodeId *string
	var ParsedDirectorynum *string
	var ParsedDirectorynumpartition *string
	var ParsedGlobalcallid_Clusterid *string
	var ParsedDevicename *string
	var ParsedVideocontenttype *string
	var ParsedVideoreceptionmetrics *string
	var ParsedVideotransmissionmetrics *string
	var ParsedVideocontenttype_Channel2 *string
	var ParsedVideoreceptionmetrics_Channel2 *string
	var ParsedVideotransmissionmetrics_Channel2 *string
	var ParsedLocalsessionid *string
	var ParsedRemotesessionid *string
	var ParsedHeadsetsn *string
	var ParsedHeadsetmetrics *string
	var ParsedVqvorxcodec *string

	var ParsedVQCCR *float64
	var ParsedVQICR *float64
	var ParsedVqicrmx *float64
	var ParsedVqver *float64
	var ParsedVQMLQK *float64
	var ParsedVqmlqkav *float64
	var ParsedVqmlqkmn *float64
	var ParsedVqmlqkmx *float64
	var ParsedVqmlqkvr *float64

	ParsedCdrrecordtype, err := helpers.ConvertStringToInt64(raw.Cdrrecordtype)
	if err != nil {
		logger.Error("Error parsing Cdrrecordtype: %s in %s", err, filename)
	}
	ParsedGlobalcallid_Callmanagerid, err = helpers.ConvertStringToInt64(raw.Globalcallid_Callmanagerid)
	if err != nil {
		logger.Error("Error parsing Globalcallid_Callmanagerid: %s in %s", err, filename)
	}
	ParsedGlobalcallid_Callid, err = helpers.ConvertStringToInt64(raw.Globalcallid_Callid)
	if err != nil {
		logger.Error("Error parsing Globalcallid_Callid: %s in %s", err, filename)
	}
	ParsedNodeid, err = helpers.ConvertStringToInt64(raw.Nodeid)
	if err != nil {
		logger.Error("Error parsing Nodeid: %s in %s", err, filename)
	}
	ParsedCallidentifier, err = helpers.ConvertStringToInt64(raw.Callidentifier)
	if err != nil {
		logger.Error("Error parsing Callidentifier: %s in %s", err, filename)
	}
	ParsedDatetimestamp, err = helpers.ConvertStringToInt64(raw.Datetimestamp)
	if err != nil {
		logger.Error("Error parsing Datetimestamp: %s in %s", err, filename)
	}
	ParsedNumberpacketssent, err = helpers.ConvertStringToInt64(raw.Numberpacketssent)
	if err != nil {
		logger.Error("Error parsing Numberpacketssent: %s in %s", err, filename)
	}
	ParsedNumberoctetssent, err = helpers.ConvertStringToInt64(raw.Numberoctetssent)
	if err != nil {
		logger.Error("Error parsing Numberoctetssent: %s in %s", err, filename)
	}
	ParsedNumberpacketsreceived, err = helpers.ConvertStringToInt64(raw.Numberpacketsreceived)
	if err != nil {
		logger.Error("Error parsing Numberpacketsreceived: %s in %s", err, filename)
	}
	ParsedNumberoctetsreceived, err = helpers.ConvertStringToInt64(raw.Numberoctetsreceived)
	if err != nil {
		logger.Error("Error parsing Numberoctetsreceived: %s in %s", err, filename)
	}
	ParsedNumberpacketslost, err = helpers.ConvertStringToInt64(raw.Numberpacketslost)
	if err != nil {
		logger.Error("Error parsing Numberpacketslost: %s in %s", err, filename)
	}
	ParsedJitter, err = helpers.ConvertStringToInt64(raw.Jitter)
	if err != nil {
		logger.Error("Error parsing Jitter: %s in %s", err, filename)
	}
	ParsedLatency, err = helpers.ConvertStringToInt64(raw.Latency)
	if err != nil {
		logger.Error("Error parsing Latency: %s in %s", err, filename)
	}
	ParsedDuration, err = helpers.ConvertStringToInt64(raw.Duration)
	if err != nil {
		logger.Error("Error parsing Duration: %s in %s", err, filename)
	}
	ParsedVideoduration, err = helpers.ConvertStringToInt64(raw.Videoduration)
	if err != nil {
		logger.Error("Error parsing Videoduration: %s in %s", err, filename)
	}
	ParsedNumbervideopacketssent, err = helpers.ConvertStringToInt64(raw.Numbervideopacketssent)
	if err != nil {
		logger.Error("Error parsing Numbervideopacketssent: %s in %s", err, filename)
	}
	ParsedNumbervideooctetssent, err = helpers.ConvertStringToInt64(raw.Numbervideooctetssent)
	if err != nil {
		logger.Error("Error parsing Numbervideooctetssent: %s in %s", err, filename)
	}
	ParsedNumbervideopacketsreceived, err = helpers.ConvertStringToInt64(raw.Numbervideopacketsreceived)
	if err != nil {
		logger.Error("Error parsing Numbervideopacketsreceived: %s in %s", err, filename)
	}
	ParsedNumbervideooctetsreceived, err = helpers.ConvertStringToInt64(raw.Numbervideooctetsreceived)
	if err != nil {
		logger.Error("Error parsing Numbervideooctetsreceived: %s in %s", err, filename)
	}
	ParsedNumbervideopacketslost, err = helpers.ConvertStringToInt64(raw.Numbervideopacketslost)
	if err != nil {
		logger.Error("Error parsing Numbervideopacketslost: %s in %s", err, filename)
	}
	ParsedVideoaveragejitter, err = helpers.ConvertStringToInt64(raw.Videoaveragejitter)
	if err != nil {
		logger.Error("Error parsing Videoaveragejitter: %s in %s", err, filename)
	}
	ParsedVideoroundtriptime, err = helpers.ConvertStringToInt64(raw.Videoroundtriptime)
	if err != nil {
		logger.Error("Error parsing Videoroundtriptime: %s in %s", err, filename)
	}
	ParsedVideoonewaydelay, err = helpers.ConvertStringToInt64(raw.Videoonewaydelay)
	if err != nil {
		logger.Error("Error parsing Videoonewaydelay: %s in %s", err, filename)
	}
	ParsedVideoduration_Channel2, err = helpers.ConvertStringToInt64(raw.Videoduration_Channel2)
	if err != nil {
		logger.Error("Error parsing Videoduration_Channel2: %s in %s", err, filename)
	}
	ParsedNumbervideopacketssent_Channel2, err = helpers.ConvertStringToInt64(raw.Numbervideopacketssent_Channel2)
	if err != nil {
		logger.Error("Error parsing Numbervideopacketssent_Channel2: %s in %s", err, filename)
	}
	ParsedNumbervideooctetssent_Channel2, err = helpers.ConvertStringToInt64(raw.Numbervideooctetssent_Channel2)
	if err != nil {
		logger.Error("Error parsing Numbervideooctetssent_Channel2: %s in %s", err, filename)
	}
	ParsedNumbervideopacketsreceived_Channel2, err = helpers.ConvertStringToInt64(raw.Numbervideopacketsreceived_Channel2)
	if err != nil {
		logger.Error("Error parsing Numbervideopacketsreceived_Channel2: %s in %s", err, filename)
	}
	ParsedNumbervideooctetsreceived_Channel2, err = helpers.ConvertStringToInt64(raw.Numbervideooctetsreceived_Channel2)
	if err != nil {
		logger.Error("Error parsing Numbervideooctetsreceived_Channel2: %s in %s", err, filename)
	}
	ParsedNumbervideopacketslost_Channel2, err = helpers.ConvertStringToInt64(raw.Numbervideopacketslost_Channel2)
	if err != nil {
		logger.Error("Error parsing Numbervideopacketslost_Channel2: %s in %s", err, filename)
	}
	ParsedVideoaveragejitter_Channel2, err = helpers.ConvertStringToInt64(raw.Videoaveragejitter_Channel2)
	if err != nil {
		logger.Error("Error parsing Videoaveragejitter_Channel2: %s in %s", err, filename)
	}
	ParsedVideoroundtriptime_Channel2, err = helpers.ConvertStringToInt64(raw.Videoroundtriptime_Channel2)
	if err != nil {
		logger.Error("Error parsing Videoroundtriptime_Channel2: %s in %s", err, filename)
	}
	ParsedVideoonewaydelay_Channel2, err = helpers.ConvertStringToInt64(raw.Videoonewaydelay_Channel2)
	if err != nil {
		logger.Error("Error parsing Videoonewaydelay_Channel2: %s in %s", err, filename)
	}

	ParsedOriginpkid = helpers.RemoveSpaceFromString(raw.Pkid)
	ParsedFileClusterId = helpers.RemoveSpaceFromString(raw.FileClusterId)
	ParsedFileNodeId = helpers.RemoveSpaceFromString(raw.FileNodeId)
	ParsedDirectorynum = helpers.RemoveSpaceFromString(raw.Directorynum)
	ParsedDirectorynumpartition = helpers.RemoveSpaceFromString(raw.Directorynumpartition)
	ParsedGlobalcallid_Clusterid = helpers.RemoveSpaceFromString(raw.Globalcallid_Clusterid)
	ParsedDevicename = helpers.RemoveSpaceFromString(raw.Devicename)
	ParsedVideocontenttype = helpers.RemoveSpaceFromString(raw.Videocontenttype)
	ParsedVideoreceptionmetrics = helpers.RemoveSpaceFromString(raw.Videoreceptionmetrics)
	ParsedVideotransmissionmetrics = helpers.RemoveSpaceFromString(raw.Videotransmissionmetrics)
	ParsedVideocontenttype_Channel2 = helpers.RemoveSpaceFromString(raw.Videocontenttype_Channel2)
	ParsedVideoreceptionmetrics_Channel2 = helpers.RemoveSpaceFromString(raw.Videoreceptionmetrics_Channel2)
	ParsedVideotransmissionmetrics_Channel2 = helpers.RemoveSpaceFromString(raw.Videotransmissionmetrics_Channel2)
	ParsedLocalsessionid = helpers.RemoveSpaceFromString(raw.Localsessionid)
	ParsedRemotesessionid = helpers.RemoveSpaceFromString(raw.Remotesessionid)
	ParsedHeadsetsn = helpers.RemoveSpaceFromString(raw.Headsetsn)
	ParsedHeadsetmetrics = helpers.RemoveSpaceFromString(raw.Headsetmetrics)

	if raw.Varvqmetrics != nil {
		VarVQMap := helpers.ConvertStringToKeyValuePairs(raw.Varvqmetrics, ";", "=")

		VQMLQK, ok := VarVQMap["MLQK"]
		if ok {
			ParsedVQMLQK, err = helpers.ConvertStringToFloat64(&VQMLQK)
			if err != nil {
				logger.Error("Error parsing VQMLQK: %s in %s", err, filename)
			}
		}
		Vqmlqkav, ok := VarVQMap["MLQKav"]
		if ok {
			ParsedVqmlqkav, err = helpers.ConvertStringToFloat64(&Vqmlqkav)
			if err != nil {
				logger.Error("Error parsing Vqmlqkav: %s in %s", err, filename)
			}
		}
		Vqmlqkmn, ok := VarVQMap["MLQKmn"]
		if ok {
			ParsedVqmlqkmn, err = helpers.ConvertStringToFloat64(&Vqmlqkmn)
			if err != nil {
				logger.Error("Error parsing Vqmlqkmn: %s in %s", err, filename)
			}
		}
		Vqmlqkmx, ok := VarVQMap["MLQKmx"]
		if ok {
			ParsedVqmlqkmx, err = helpers.ConvertStringToFloat64(&Vqmlqkmx)
			if err != nil {
				logger.Error("Error parsing Vqmlqkmx: %s in %s", err, filename)
			}
		}
		Vqmlqkvr, ok := VarVQMap["MLQKvr"]
		if ok {
			ParsedVqmlqkvr, err = helpers.ConvertStringToFloat64(&Vqmlqkvr)
			if err != nil {
				logger.Error("Error parsing Vqmlqkvr: %s in %s", err, filename)
			}
		}
		VQCCR, ok := VarVQMap["CCR"]
		if ok {
			ParsedVQCCR, err = helpers.ConvertStringToFloat64(&VQCCR)
			if err != nil {
				logger.Error("Error parsing VQCCR: %s in %s", err, filename)
			}
		}
		VQICR, ok := VarVQMap["ICR"]
		if ok {
			ParsedVQICR, err = helpers.ConvertStringToFloat64(&VQICR)
			if err != nil {
				logger.Error("Error parsing VQICR: %s in %s", err, filename)
			}
		}
		Vqicrmx, ok := VarVQMap["ICRmx"]
		if ok {
			ParsedVqicrmx, err = helpers.ConvertStringToFloat64(&Vqicrmx)
			if err != nil {
				logger.Error("Error parsing Vqicrmx: %s in %s", err, filename)
			}
		}
		Vqver, ok := VarVQMap["Ver"]
		if ok {
			ParsedVqver, err = helpers.ConvertStringToFloat64(&Vqver)
			if err != nil {
				logger.Error("Error parsing Vqver: %s in %s", err, filename)
			}
		}
		VQCS, ok := VarVQMap["CS"]
		if ok {
			ParsedVQCS, err = helpers.ConvertStringToInt64(&VQCS)
			if err != nil {
				logger.Error("Error parsing VQCS: %s in %s", err, filename)
			}
		}
		VQSCS, ok := VarVQMap["SCS"]
		if ok {
			ParsedVQSCS, err = helpers.ConvertStringToInt64(&VQSCS)
			if err != nil {
				logger.Error("Error parsing VQSCS: %s in %s", err, filename)
			}
		}
		VQCID, ok := VarVQMap["CID"]
		if ok {
			ParsedVQCID, err = helpers.ConvertStringToInt64(&VQCID)
			if err != nil {
				logger.Error("Error parsing VQCID: %s in %s", err, filename)
			}
		}
		Vqvopktsizems, ok := VarVQMap["VoPktSizeMs"]
		if ok {
			ParsedVqvopktsizems, err = helpers.ConvertStringToInt64(&Vqvopktsizems)
			if err != nil {
				logger.Error("Error parsing Vqvopktsizems: %s in %s", err, filename)
			}
		}
		Vqvopktlost, ok := VarVQMap["VoPktLost"]
		if ok {
			ParsedVqvopktlost, err = helpers.ConvertStringToInt64(&Vqvopktlost)
			if err != nil {
				logger.Error("Error parsing Vqvopktlost: %s in %s", err, filename)
			}
		}
		Vqvopktdis, ok := VarVQMap["VoPktDis"]
		if ok {
			ParsedVqvopktdis, err = helpers.ConvertStringToInt64(&Vqvopktdis)
			if err != nil {
				logger.Error("Error parsing Vqvopktdis: %s in %s", err, filename)
			}
		}
		Vqvoonewaydelayms, ok := VarVQMap["VoOneWayDelayMs"]
		if ok {
			ParsedVqvoonewaydelayms, err = helpers.ConvertStringToInt64(&Vqvoonewaydelayms)
			if err != nil {
				logger.Error("Error parsing Vqvoonewaydelayms: %s in %s", err, filename)
			}
		}
		Vqmaxjitter, ok := VarVQMap["maxJitter"]
		if ok {
			ParsedVqmaxjitter, err = helpers.ConvertStringToInt64(&Vqmaxjitter)
			if err != nil {
				logger.Error("Error parsing Vqmaxjitter: %s in %s", err, filename)
			}
		}
		VoRxCodec, ok := VarVQMap["VoRxCodec"]
		if ok {
			ParsedVqvorxcodec = &VoRxCodec
		}
	}

	return &CucmCmr{
		ID:                                  uuid.New().String(),
		Originpkid:                          ParsedOriginpkid,
		FileClusterId:                       ParsedFileClusterId,
		FileNodeId:                          ParsedFileNodeId,
		FileDateTime:                        raw.FileDateTime,
		FileSequenceNumber:                  raw.FileSequenceNumber,
		Cdrrecordtype:                       ParsedCdrrecordtype,
		Globalcallid_Callmanagerid:          ParsedGlobalcallid_Callmanagerid,
		Globalcallid_Callid:                 ParsedGlobalcallid_Callid,
		Nodeid:                              ParsedNodeid,
		Directorynum:                        ParsedDirectorynum,
		Callidentifier:                      ParsedCallidentifier,
		Datetimestamp:                       ParsedDatetimestamp,
		Numberpacketssent:                   ParsedNumberpacketssent,
		Numberoctetssent:                    ParsedNumberoctetssent,
		Numberpacketsreceived:               ParsedNumberpacketsreceived,
		Numberoctetsreceived:                ParsedNumberoctetsreceived,
		Numberpacketslost:                   ParsedNumberpacketslost,
		Jitter:                              ParsedJitter,
		Latency:                             ParsedLatency,
		Directorynumpartition:               ParsedDirectorynumpartition,
		Globalcallid_Clusterid:              ParsedGlobalcallid_Clusterid,
		Devicename:                          ParsedDevicename,
		Duration:                            ParsedDuration,
		Videocontenttype:                    ParsedVideocontenttype,
		Videoduration:                       ParsedVideoduration,
		Numbervideopacketssent:              ParsedNumbervideopacketssent,
		Numbervideooctetssent:               ParsedNumbervideooctetssent,
		Numbervideopacketsreceived:          ParsedNumbervideopacketsreceived,
		Numbervideooctetsreceived:           ParsedNumbervideooctetsreceived,
		Numbervideopacketslost:              ParsedNumbervideopacketslost,
		Videoaveragejitter:                  ParsedVideoaveragejitter,
		Videoroundtriptime:                  ParsedVideoroundtriptime,
		Videoonewaydelay:                    ParsedVideoonewaydelay,
		Videoreceptionmetrics:               ParsedVideoreceptionmetrics,
		Videotransmissionmetrics:            ParsedVideotransmissionmetrics,
		Videocontenttype_Channel2:           ParsedVideocontenttype_Channel2,
		Videoduration_Channel2:              ParsedVideoduration_Channel2,
		Numbervideopacketssent_Channel2:     ParsedNumbervideopacketssent_Channel2,
		Numbervideooctetssent_Channel2:      ParsedNumbervideooctetssent_Channel2,
		Numbervideopacketsreceived_Channel2: ParsedNumbervideopacketsreceived_Channel2,
		Numbervideooctetsreceived_Channel2:  ParsedNumbervideooctetsreceived_Channel2,
		Numbervideopacketslost_Channel2:     ParsedNumbervideopacketslost_Channel2,
		Videoaveragejitter_Channel2:         ParsedVideoaveragejitter_Channel2,
		Videoroundtriptime_Channel2:         ParsedVideoroundtriptime_Channel2,
		Videoonewaydelay_Channel2:           ParsedVideoonewaydelay_Channel2,
		Videoreceptionmetrics_Channel2:      ParsedVideoreceptionmetrics_Channel2,
		Videotransmissionmetrics_Channel2:   ParsedVideotransmissionmetrics_Channel2,
		Localsessionid:                      ParsedLocalsessionid,
		Remotesessionid:                     ParsedRemotesessionid,
		Headsetsn:                           ParsedHeadsetsn,
		Headsetmetrics:                      ParsedHeadsetmetrics,
		VQCCR:                               ParsedVQCCR,
		VQICR:                               ParsedVQICR,
		Vqicrmx:                             ParsedVqicrmx,
		VQCS:                                ParsedVQCS,
		VQSCS:                               ParsedVQSCS,
		Vqver:                               ParsedVqver,
		Vqvorxcodec:                         ParsedVqvorxcodec,
		VQCID:                               ParsedVQCID,
		Vqvopktsizems:                       ParsedVqvopktsizems,
		Vqvopktlost:                         ParsedVqvopktlost,
		Vqvopktdis:                          ParsedVqvopktdis,
		Vqvoonewaydelayms:                   ParsedVqvoonewaydelayms,
		Vqmaxjitter:                         ParsedVqmaxjitter,
		VQMLQK:                              ParsedVQMLQK,
		Vqmlqkav:                            ParsedVqmlqkav,
		Vqmlqkmn:                            ParsedVqmlqkmn,
		Vqmlqkmx:                            ParsedVqmlqkmx,
		Vqmlqkvr:                            ParsedVqmlqkvr,
	}, nil
}
