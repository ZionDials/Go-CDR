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

type RawCucmCdr struct {
	Pkid                                    *string
	FileClusterId                           *string
	FileNodeId                              *string
	FileDateTime                            *int64
	FileSequenceNumber                      *int64
	Cdrrecordtype                           *string
	Globalcallid_Callmanagerid              *string
	Globalcallid_Callid                     *string
	Origlegcallidentifier                   *string
	Datetimeorigination                     *string
	Orignodeid                              *string
	Origspan                                *string
	Origipaddr                              *string
	Callingpartynumber                      *string
	Callingpartyunicodeloginuserid          *string
	Origcause_Location                      *string
	Origcause_Value                         *string
	Origprecedencelevel                     *string
	Origmediatransportaddress_IP            *string
	Origmediatransportaddress_Port          *string
	Origmediacap_Payloadcapability          *string
	Origmediacap_Maxframesperpacket         *string
	Origmediacap_G723bitrate                *string
	Origvideocap_Codec                      *string
	Origvideocap_Bandwidth                  *string
	Origvideocap_Resolution                 *string
	Origvideotransportaddress_IP            *string
	Origvideotransportaddress_Port          *string
	Origrsvpaudiostat                       *string
	Origrsvpvideostat                       *string
	Destlegcallidentifier                   *string
	Destnodeid                              *string
	Destspan                                *string
	Destipaddr                              *string
	Originalcalledpartynumber               *string
	Finalcalledpartynumber                  *string
	Finalcalledpartyunicodeloginuserid      *string
	Destcause_Location                      *string
	Destcause_Value                         *string
	Destprecedencelevel                     *string
	Destmediatransportaddress_IP            *string
	Destmediatransportaddress_Port          *string
	Destmediacap_Payloadcapability          *string
	Destmediacap_Maxframesperpacket         *string
	Destmediacap_G723bitrate                *string
	Destvideocap_Codec                      *string
	Destvideocap_Bandwidth                  *string
	Destvideocap_Resolution                 *string
	Destvideotransportaddress_IP            *string
	Destvideotransportaddress_Port          *string
	Destrsvpaudiostat                       *string
	Destrsvpvideostat                       *string
	Datetimeconnect                         *string
	Datetimedisconnect                      *string
	Lastredirectdn                          *string
	Originalcalledpartynumberpartition      *string
	Callingpartynumberpartition             *string
	Finalcalledpartynumberpartition         *string
	Lastredirectdnpartition                 *string
	Duration                                *string
	Origdevicename                          *string
	Destdevicename                          *string
	Origcallterminationonbehalfof           *string
	Destcallterminationonbehalfof           *string
	Origcalledpartyredirectonbehalfof       *string
	Lastredirectredirectonbehalfof          *string
	Origcalledpartyredirectreason           *string
	Lastredirectredirectreason              *string
	Destconversationid                      *string
	Globalcallid_Clusterid                  *string
	Joinonbehalfof                          *string
	Comment                                 *string
	Authcodedescription                     *string
	Authorizationlevel                      *string
	Clientmattercode                        *string
	Origdtmfmethod                          *string
	Destdtmfmethod                          *string
	Callsecuredstatus                       *string
	Origconversationid                      *string
	Origmediacap_Bandwidth                  *string
	Destmediacap_Bandwidth                  *string
	Authorizationcodevalue                  *string
	Outpulsedcallingpartynumber             *string
	Outpulsedcalledpartynumber              *string
	Origipv4v6addr                          *string
	Destipv4v6addr                          *string
	Origvideocap_Codec_Channel2             *string
	Origvideocap_Bandwidth_Channel2         *string
	Origvideocap_Resolution_Channel2        *string
	Origvideotransportaddress_IP_Channel2   *string
	Origvideotransportaddress_Port_Channel2 *string
	Origvideochannel_Role_Channel2          *string
	Destvideocap_Codec_Channel2             *string
	Destvideocap_Bandwidth_Channel2         *string
	Destvideocap_Resolution_Channel2        *string
	Destvideotransportaddress_IP_Channel2   *string
	Destvideotransportaddress_Port_Channel2 *string
	Destvideochannel_Role_Channel2          *string
	Incomingprotocolid                      *string
	Incomingprotocolcallref                 *string
	Outgoingprotocolid                      *string
	Outgoingprotocolcallref                 *string
	Currentroutingreason                    *string
	Origroutingreason                       *string
	Lastredirectingroutingreason            *string
	Huntpilotpartition                      *string
	Huntpilotdn                             *string
	Calledpartypatternusage                 *string
	Incomingicid                            *string
	Incomingorigioi                         *string
	Incomingtermioi                         *string
	Outgoingicid                            *string
	Outgoingorigioi                         *string
	Outgoingtermioi                         *string
	Outpulsedoriginalcalledpartynumber      *string
	Outpulsedlastredirectingnumber          *string
	Wascallqueued                           *string
	Totalwaittimeinqueue                    *string
	Callingpartynumber_Uri                  *string
	Originalcalledpartynumber_Uri           *string
	Finalcalledpartynumber_Uri              *string
	Lastredirectdn_Uri                      *string
	Mobilecallingpartynumber                *string
	Finalmobilecalledpartynumber            *string
	Origmobiledevicename                    *string
	Destmobiledevicename                    *string
	Origmobilecallduration                  *string
	Destmobilecallduration                  *string
	Mobilecalltype                          *string
	Originalcalledpartypattern              *string
	Finalcalledpartypattern                 *string
	Lastredirectingpartypattern             *string
	Huntpilotpattern                        *string
	Origdevicetype                          *string
	Destdevicetype                          *string
	Origdevicesessionid                     *string
	Destdevicesessionid                     *string
}

func (raw *RawCucmCdr) Parse(filename string) (*CucmCdr, error) {

	var ParsedCdrrecordtype *int64
	var ParsedGlobalcallid_Callmanagerid *int64
	var ParsedGlobalcallid_Callid *int64
	var ParsedOriglegcallidentifier *int64
	var ParsedDatetimeorigination *int64
	var ParsedOrignodeid *int64
	var ParsedOrigspan *int64
	var ParsedOrigcause_Location *int64
	var ParsedOrigcause_Value *int64
	var ParsedOrigprecedencelevel *int64
	var ParsedOrigmediatransportaddress_Port *int64
	var ParsedOrigmediacap_Payloadcapability *int64
	var ParsedOrigmediacap_Maxframesperpacket *int64
	var ParsedOrigmediacap_G723bitrate *int64
	var ParsedOrigvideocap_Codec *int64
	var ParsedOrigvideocap_Bandwidth *int64
	var ParsedOrigvideocap_Resolution *int64
	var ParsedOrigvideotransportaddress_Port *int64
	var ParsedOrigrsvpaudiostat *int64
	var ParsedOrigrsvpvideostat *int64
	var ParsedDestlegcallidentifier *int64
	var ParsedDestnodeid *int64
	var ParsedDestspan *int64
	var ParsedDestcause_Location *int64
	var ParsedDestcause_Value *int64
	var ParsedDestprecedencelevel *int64
	var ParsedDestmediatransportaddress_Port *int64
	var ParsedDestmediacap_Payloadcapability *int64
	var ParsedDestmediacap_Maxframesperpacket *int64
	var ParsedDestmediacap_G723bitrate *int64
	var ParsedDestvideocap_Codec *int64
	var ParsedDestvideocap_Bandwidth *int64
	var ParsedDestvideocap_Resolution *int64
	var ParsedDestvideotransportaddress_Port *int64
	var ParsedDestrsvpaudiostat *int64
	var ParsedDestrsvpvideostat *int64
	var ParsedDatetimeconnect *int64
	var ParsedDatetimedisconnect *int64
	var ParsedDuration *int64
	var ParsedOrigcallterminationonbehalfof *int64
	var ParsedDestcallterminationonbehalfof *int64
	var ParsedOrigcalledpartyredirectonbehalfof *int64
	var ParsedLastredirectredirectonbehalfof *int64
	var ParsedOrigcalledpartyredirectreason *int64
	var ParsedLastredirectredirectreason *int64
	var ParsedDestconversationid *int64
	var ParsedJoinonbehalfof *int64
	var ParsedAuthorizationlevel *int64
	var ParsedOrigdtmfmethod *int64
	var ParsedDestdtmfmethod *int64
	var ParsedCallsecuredstatus *int64
	var ParsedOrigconversationid *int64
	var ParsedOrigmediacap_Bandwidth *int64
	var ParsedDestmediacap_Bandwidth *int64
	var ParsedOrigvideocap_Codec_Channel2 *int64
	var ParsedOrigvideocap_Bandwidth_Channel2 *int64
	var ParsedOrigvideocap_Resolution_Channel2 *int64
	var ParsedOrigvideotransportaddress_Port_Channel2 *int64
	var ParsedOrigvideochannel_Role_Channel2 *int64
	var ParsedDestvideocap_Codec_Channel2 *int64
	var ParsedDestvideocap_Bandwidth_Channel2 *int64
	var ParsedDestvideocap_Resolution_Channel2 *int64
	var ParsedDestvideotransportaddress_Port_Channel2 *int64
	var ParsedDestvideochannel_Role_Channel2 *int64
	var ParsedIncomingprotocolid *int64
	var ParsedOutgoingprotocolid *int64
	var ParsedCurrentroutingreason *int64
	var ParsedOrigroutingreason *int64
	var ParsedLastredirectingroutingreason *int64
	var ParsedCalledpartypatternusage *int64
	var ParsedWascallqueued *int64
	var ParsedTotalwaittimeinqueue *int64
	var ParsedOrigmobilecallduration *int64
	var ParsedDestmobilecallduration *int64
	var ParsedMobilecalltype *int64

	var ParsedOriginpkid *string
	var ParsedFileClusterId *string
	var ParsedFileNodeId *string
	var ParsedOrigipaddr *string
	var ParsedCallingpartynumber *string
	var ParsedCallingpartyunicodeloginuserid *string
	var ParsedOrigmediatransportaddress_IP *string
	var ParsedOrigvideotransportaddress_IP *string
	var ParsedDestipaddr *string
	var ParsedOriginalcalledpartynumber *string
	var ParsedFinalcalledpartynumber *string
	var ParsedFinalcalledpartyunicodeloginuserid *string
	var ParsedDestmediatransportaddress_IP *string
	var ParsedDestvideotransportaddress_IP *string
	var ParsedLastredirectdn *string
	var ParsedOriginalcalledpartynumberpartition *string
	var ParsedCallingpartynumberpartition *string
	var ParsedFinalcalledpartynumberpartition *string
	var ParsedLastredirectdnpartition *string
	var ParsedOrigdevicename *string
	var ParsedDestdevicename *string
	var ParsedGlobalcallid_Clusterid *string
	var ParsedComment *string
	var ParsedAuthcodedescription *string
	var ParsedClientmattercode *string
	var ParsedAuthorizationcodevalue *string
	var ParsedOutpulsedcallingpartynumber *string
	var ParsedOutpulsedcalledpartynumber *string
	var ParsedOrigipv4v6addr *string
	var ParsedDestipv4v6addr *string
	var ParsedOrigvideotransportaddress_IP_Channel2 *string
	var ParsedDestvideotransportaddress_IP_Channel2 *string
	var ParsedIncomingprotocolcallref *string
	var ParsedOutgoingprotocolcallref *string
	var ParsedHuntpilotpartition *string
	var ParsedHuntpilotdn *string
	var ParsedIncomingicid *string
	var ParsedOutgoingicid *string
	var ParsedOutpulsedoriginalcalledpartynumber *string
	var ParsedOutpulsedlastredirectingnumber *string
	var ParsedCallingpartynumber_Uri *string
	var ParsedOriginalcalledpartynumber_Uri *string
	var ParsedFinalcalledpartynumber_Uri *string
	var ParsedLastredirectdn_Uri *string
	var ParsedMobilecallingpartynumber *string
	var ParsedFinalmobilecalledpartynumber *string
	var ParsedOrigmobiledevicename *string
	var ParsedDestmobiledevicename *string
	var ParsedOriginalcalledpartypattern *string
	var ParsedFinalcalledpartypattern *string
	var ParsedLastredirectingpartypattern *string
	var ParsedHuntpilotpattern *string
	var ParsedOrigdevicetype *string
	var ParsedDestdevicetype *string
	var ParsedOrigdevicesessionid *string
	var ParsedDestdevicesessionid *string

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
	ParsedOriglegcallidentifier, err = helpers.ConvertStringToInt64(raw.Origlegcallidentifier)
	if err != nil {
		logger.Error("Error parsing Origlegcallidentifier: %s in %s", err, filename)
	}
	ParsedDatetimeorigination, err = helpers.ConvertStringToInt64(raw.Datetimeorigination)
	if err != nil {
		logger.Error("Error parsing Datetimeorigination: %s in %s", err, filename)
	}
	ParsedOrignodeid, err = helpers.ConvertStringToInt64(raw.Orignodeid)
	if err != nil {
		logger.Error("Error parsing Orignodeid: %s in %s", err, filename)
	}
	ParsedOrigspan, err = helpers.ConvertStringToInt64(raw.Origspan)
	if err != nil {
		logger.Error("Error parsing Origspan: %s in %s", err, filename)
	}
	ParsedOrigcause_Location, err = helpers.ConvertStringToInt64(raw.Origcause_Location)
	if err != nil {
		logger.Error("Error parsing Origcause_Location: %s in %s", err, filename)
	}
	ParsedOrigcause_Value, err = helpers.ConvertStringToInt64(raw.Origcause_Value)
	if err != nil {
		logger.Error("Error parsing Origcause_Value: %s in %s", err, filename)
	}
	ParsedOrigprecedencelevel, err = helpers.ConvertStringToInt64(raw.Origprecedencelevel)
	if err != nil {
		logger.Error("Error parsing Origprecedencelevel: %s in %s", err, filename)
	}
	ParsedOrigmediatransportaddress_Port, err = helpers.ConvertStringToInt64(raw.Origmediatransportaddress_Port)
	if err != nil {
		logger.Error("Error parsing Origmediatransportaddress_Port: %s in %s", err, filename)
	}
	ParsedOrigmediacap_Payloadcapability, err = helpers.ConvertStringToInt64(raw.Origmediacap_Payloadcapability)
	if err != nil {
		logger.Error("Error parsing Origmediacap_Payloadcapability: %s in %s", err, filename)
	}
	ParsedOrigmediacap_Maxframesperpacket, err = helpers.ConvertStringToInt64(raw.Origmediacap_Maxframesperpacket)
	if err != nil {
		logger.Error("Error parsing Origmediacap_Maxframesperpacket: %s in %s", err, filename)
	}
	ParsedOrigmediacap_G723bitrate, err = helpers.ConvertStringToInt64(raw.Origmediacap_G723bitrate)
	if err != nil {
		logger.Error("Error parsing Origmediacap_G723bitrate: %s in %s", err, filename)
	}
	ParsedOrigvideocap_Codec, err = helpers.ConvertStringToInt64(raw.Origvideocap_Codec)
	if err != nil {
		logger.Error("Error parsing Origvideocap_Codec: %s in %s", err, filename)
	}
	ParsedOrigvideocap_Bandwidth, err = helpers.ConvertStringToInt64(raw.Origvideocap_Bandwidth)
	if err != nil {
		logger.Error("Error parsing Origvideocap_Bandwidth: %s in %s", err, filename)
	}
	ParsedOrigvideocap_Resolution, err = helpers.ConvertStringToInt64(raw.Origvideocap_Resolution)
	if err != nil {
		logger.Error("Error parsing Origvideocap_Resolution: %s in %s", err, filename)
	}
	ParsedOrigvideotransportaddress_Port, err = helpers.ConvertStringToInt64(raw.Origvideotransportaddress_Port)
	if err != nil {
		logger.Error("Error parsing Origvideotransportaddress_Port: %s in %s", err, filename)
	}
	ParsedOrigrsvpaudiostat, err = helpers.ConvertStringToInt64(raw.Origrsvpaudiostat)
	if err != nil {
		logger.Error("Error parsing Origrsvpaudiostat: %s in %s", err, filename)
	}
	ParsedOrigrsvpvideostat, err = helpers.ConvertStringToInt64(raw.Origrsvpvideostat)
	if err != nil {
		logger.Error("Error parsing Origrsvpvideostat: %s in %s", err, filename)
	}
	ParsedDestlegcallidentifier, err = helpers.ConvertStringToInt64(raw.Destlegcallidentifier)
	if err != nil {
		logger.Error("Error parsing Destlegcallidentifier: %s in %s", err, filename)
	}
	ParsedDestnodeid, err = helpers.ConvertStringToInt64(raw.Destnodeid)
	if err != nil {
		logger.Error("Error parsing Destnodeid: %s in %s", err, filename)
	}
	ParsedDestspan, err = helpers.ConvertStringToInt64(raw.Destspan)
	if err != nil {
		logger.Error("Error parsing Destspan: %s in %s", err, filename)
	}
	ParsedDestcause_Location, err = helpers.ConvertStringToInt64(raw.Destcause_Location)
	if err != nil {
		logger.Error("Error parsing Destcause_Location: %s in %s", err, filename)
	}
	ParsedDestcause_Value, err = helpers.ConvertStringToInt64(raw.Destcause_Value)
	if err != nil {
		logger.Error("Error parsing Destcause_Value: %s in %s", err, filename)
	}
	ParsedDestprecedencelevel, err = helpers.ConvertStringToInt64(raw.Destprecedencelevel)
	if err != nil {
		logger.Error("Error parsing Destprecedencelevel: %s in %s", err, filename)
	}
	ParsedDestmediatransportaddress_Port, err = helpers.ConvertStringToInt64(raw.Destmediatransportaddress_Port)
	if err != nil {
		logger.Error("Error parsing Destmediatransportaddress_Port: %s in %s", err, filename)
	}
	ParsedDestmediacap_Payloadcapability, err = helpers.ConvertStringToInt64(raw.Destmediacap_Payloadcapability)
	if err != nil {
		logger.Error("Error parsing Destmediacap_Payloadcapability: %s in %s", err, filename)
	}
	ParsedDestmediacap_Maxframesperpacket, err = helpers.ConvertStringToInt64(raw.Destmediacap_Maxframesperpacket)
	if err != nil {
		logger.Error("Error parsing Destmediacap_Maxframesperpacket: %s in %s", err, filename)
	}
	ParsedDestmediacap_G723bitrate, err = helpers.ConvertStringToInt64(raw.Destmediacap_G723bitrate)
	if err != nil {
		logger.Error("Error parsing Destmediacap_G723bitrate: %s in %s", err, filename)
	}
	ParsedDestvideocap_Codec, err = helpers.ConvertStringToInt64(raw.Destvideocap_Codec)
	if err != nil {
		logger.Error("Error parsing Destvideocap_Codec: %s in %s", err, filename)
	}
	ParsedDestvideocap_Bandwidth, err = helpers.ConvertStringToInt64(raw.Destvideocap_Bandwidth)
	if err != nil {
		logger.Error("Error parsing Destvideocap_Bandwidth: %s in %s", err, filename)
	}
	ParsedDestvideocap_Resolution, err = helpers.ConvertStringToInt64(raw.Destvideocap_Resolution)
	if err != nil {
		logger.Error("Error parsing Destvideocap_Resolution: %s in %s", err, filename)
	}
	ParsedDestvideotransportaddress_Port, err = helpers.ConvertStringToInt64(raw.Destvideotransportaddress_Port)
	if err != nil {
		logger.Error("Error parsing Destvideotransportaddress_Port: %s in %s", err, filename)
	}
	ParsedDestrsvpaudiostat, err = helpers.ConvertStringToInt64(raw.Destrsvpaudiostat)
	if err != nil {
		logger.Error("Error parsing Destrsvpaudiostat: %s in %s", err, filename)
	}
	ParsedDestrsvpvideostat, err = helpers.ConvertStringToInt64(raw.Destrsvpvideostat)
	if err != nil {
		logger.Error("Error parsing Destrsvpvideostat: %s in %s", err, filename)
	}
	ParsedDatetimeconnect, err = helpers.ConvertStringToInt64(raw.Datetimeconnect)
	if err != nil {
		logger.Error("Error parsing Datetimeconnect: %s in %s", err, filename)
	}
	ParsedDatetimedisconnect, err = helpers.ConvertStringToInt64(raw.Datetimedisconnect)
	if err != nil {
		logger.Error("Error parsing Datetimedisconnect: %s in %s", err, filename)
	}
	ParsedDuration, err = helpers.ConvertStringToInt64(raw.Duration)
	if err != nil {
		logger.Error("Error parsing Duration: %s in %s", err, filename)
	}
	ParsedOrigcallterminationonbehalfof, err = helpers.ConvertStringToInt64(raw.Origcallterminationonbehalfof)
	if err != nil {
		logger.Error("Error parsing Origcallterminationonbehalfof: %s in %s", err, filename)
	}
	ParsedDestcallterminationonbehalfof, err = helpers.ConvertStringToInt64(raw.Destcallterminationonbehalfof)
	if err != nil {
		logger.Error("Error parsing Destcallterminationonbehalfof: %s in %s", err, filename)
	}
	ParsedOrigcalledpartyredirectonbehalfof, err = helpers.ConvertStringToInt64(raw.Origcalledpartyredirectonbehalfof)
	if err != nil {
		logger.Error("Error parsing Origcalledpartyredirectonbehalfof: %s in %s", err, filename)
	}
	ParsedLastredirectredirectonbehalfof, err = helpers.ConvertStringToInt64(raw.Lastredirectredirectonbehalfof)
	if err != nil {
		logger.Error("Error parsing Lastredirectredirectonbehalfof: %s in %s", err, filename)
	}
	ParsedOrigcalledpartyredirectreason, err = helpers.ConvertStringToInt64(raw.Origcalledpartyredirectreason)
	if err != nil {
		logger.Error("Error parsing Origcalledpartyredirectreason: %s in %s", err, filename)
	}
	ParsedLastredirectredirectreason, err = helpers.ConvertStringToInt64(raw.Lastredirectredirectreason)
	if err != nil {
		logger.Error("Error parsing Lastredirectredirectreason: %s in %s", err, filename)
	}
	ParsedDestconversationid, err = helpers.ConvertStringToInt64(raw.Destconversationid)
	if err != nil {
		logger.Error("Error parsing Destconversationid: %s in %s", err, filename)
	}
	ParsedJoinonbehalfof, err = helpers.ConvertStringToInt64(raw.Joinonbehalfof)
	if err != nil {
		logger.Error("Error parsing Joinonbehalfof: %s in %s", err, filename)
	}
	ParsedAuthorizationlevel, err = helpers.ConvertStringToInt64(raw.Authorizationlevel)
	if err != nil {
		logger.Error("Error parsing Authorizationlevel: %s in %s", err, filename)
	}
	ParsedOrigdtmfmethod, err = helpers.ConvertStringToInt64(raw.Origdtmfmethod)
	if err != nil {
		logger.Error("Error parsing Origdtmfmethod: %s in %s", err, filename)
	}
	ParsedDestdtmfmethod, err = helpers.ConvertStringToInt64(raw.Destdtmfmethod)
	if err != nil {
		logger.Error("Error parsing Destdtmfmethod: %s in %s", err, filename)
	}
	ParsedCallsecuredstatus, err = helpers.ConvertStringToInt64(raw.Callsecuredstatus)
	if err != nil {
		logger.Error("Error parsing Callsecuredstatus: %s in %s", err, filename)
	}
	ParsedOrigconversationid, err = helpers.ConvertStringToInt64(raw.Origconversationid)
	if err != nil {
		logger.Error("Error parsing Origconversationid: %s in %s", err, filename)
	}
	ParsedOrigmediacap_Bandwidth, err = helpers.ConvertStringToInt64(raw.Origmediacap_Bandwidth)
	if err != nil {
		logger.Error("Error parsing Origmediacap_Bandwidth: %s in %s", err, filename)
	}
	ParsedDestmediacap_Bandwidth, err = helpers.ConvertStringToInt64(raw.Destmediacap_Bandwidth)
	if err != nil {
		logger.Error("Error parsing Destmediacap_Bandwidth: %s in %s", err, filename)
	}
	ParsedOrigvideocap_Codec_Channel2, err = helpers.ConvertStringToInt64(raw.Origvideocap_Codec_Channel2)
	if err != nil {
		logger.Error("Error parsing Origvideocap_Codec_Channel2: %s in %s", err, filename)
	}
	ParsedOrigvideocap_Bandwidth_Channel2, err = helpers.ConvertStringToInt64(raw.Origvideocap_Bandwidth_Channel2)
	if err != nil {
		logger.Error("Error parsing Origvideocap_Bandwidth_Channel2: %s in %s", err, filename)
	}
	ParsedOrigvideocap_Resolution_Channel2, err = helpers.ConvertStringToInt64(raw.Origvideocap_Resolution_Channel2)
	if err != nil {
		logger.Error("Error parsing Origvideocap_Resolution_Channel2: %s in %s", err, filename)
	}
	ParsedOrigvideotransportaddress_Port_Channel2, err = helpers.ConvertStringToInt64(raw.Origvideotransportaddress_Port_Channel2)
	if err != nil {
		logger.Error("Error parsing Origvideotransportaddress_Port_Channel2: %s in %s", err, filename)
	}
	ParsedOrigvideochannel_Role_Channel2, err = helpers.ConvertStringToInt64(raw.Origvideochannel_Role_Channel2)
	if err != nil {
		logger.Error("Error parsing Origvideochannel_Role_Channel2: %s in %s", err, filename)
	}
	ParsedDestvideocap_Codec_Channel2, err = helpers.ConvertStringToInt64(raw.Destvideocap_Codec_Channel2)
	if err != nil {
		logger.Error("Error parsing Destvideocap_Codec_Channel2: %s in %s", err, filename)
	}
	ParsedDestvideocap_Bandwidth_Channel2, err = helpers.ConvertStringToInt64(raw.Destvideocap_Bandwidth_Channel2)
	if err != nil {
		logger.Error("Error parsing Destvideocap_Bandwidth_Channel2: %s in %s", err, filename)
	}
	ParsedDestvideocap_Resolution_Channel2, err = helpers.ConvertStringToInt64(raw.Destvideocap_Resolution_Channel2)
	if err != nil {
		logger.Error("Error parsing Destvideocap_Resolution_Channel2: %s in %s", err, filename)
	}
	ParsedDestvideotransportaddress_Port_Channel2, err = helpers.ConvertStringToInt64(raw.Destvideotransportaddress_Port_Channel2)
	if err != nil {
		logger.Error("Error parsing Destvideotransportaddress_Port_Channel2: %s in %s", err, filename)
	}
	ParsedDestvideochannel_Role_Channel2, err = helpers.ConvertStringToInt64(raw.Destvideochannel_Role_Channel2)
	if err != nil {
		logger.Error("Error parsing Destvideochannel_Role_Channel2: %s in %s", err, filename)
	}
	ParsedIncomingprotocolid, err = helpers.ConvertStringToInt64(raw.Incomingprotocolid)
	if err != nil {
		logger.Error("Error parsing Incomingprotocolid: %s in %s", err, filename)
	}
	ParsedOutgoingprotocolid, err = helpers.ConvertStringToInt64(raw.Outgoingprotocolid)
	if err != nil {
		logger.Error("Error parsing Outgoingprotocolid: %s in %s", err, filename)
	}
	ParsedCurrentroutingreason, err = helpers.ConvertStringToInt64(raw.Currentroutingreason)
	if err != nil {
		logger.Error("Error parsing Currentroutingreason: %s in %s", err, filename)
	}
	ParsedOrigroutingreason, err = helpers.ConvertStringToInt64(raw.Origroutingreason)
	if err != nil {
		logger.Error("Error parsing Origroutingreason: %s in %s", err, filename)
	}
	ParsedLastredirectingroutingreason, err = helpers.ConvertStringToInt64(raw.Lastredirectingroutingreason)
	if err != nil {
		logger.Error("Error parsing Lastredirectingroutingreason: %s in %s", err, filename)
	}
	ParsedCalledpartypatternusage, err = helpers.ConvertStringToInt64(raw.Calledpartypatternusage)
	if err != nil {
		logger.Error("Error parsing Calledpartypatternusage: %s in %s", err, filename)
	}
	ParsedWascallqueued, err = helpers.ConvertStringToInt64(raw.Wascallqueued)
	if err != nil {
		logger.Error("Error parsing Wascallqueued: %s in %s", err, filename)
	}
	ParsedTotalwaittimeinqueue, err = helpers.ConvertStringToInt64(raw.Totalwaittimeinqueue)
	if err != nil {
		logger.Error("Error parsing Totalwaittimeinqueue: %s in %s", err, filename)
	}
	ParsedOrigmobilecallduration, err = helpers.ConvertStringToInt64(raw.Origmobilecallduration)
	if err != nil {
		logger.Error("Error parsing Origmobilecallduration: %s in %s", err, filename)
	}
	ParsedDestmobilecallduration, err = helpers.ConvertStringToInt64(raw.Destmobilecallduration)
	if err != nil {
		logger.Error("Error parsing Destmobilecallduration: %s in %s", err, filename)
	}
	ParsedMobilecalltype, err = helpers.ConvertStringToInt64(raw.Mobilecalltype)
	if err != nil {
		logger.Error("Error parsing Mobilecalltype: %s in %s", err, filename)
	}

	ParsedOriginpkid = helpers.RemoveSpaceFromString(raw.Pkid)
	ParsedFileClusterId = helpers.RemoveSpaceFromString(raw.FileClusterId)
	ParsedFileNodeId = helpers.RemoveSpaceFromString(raw.FileNodeId)
	ParsedOrigipaddr, err = helpers.ConvertStringToIPCisco(raw.Origipaddr)
	if err != nil {
		logger.Error("Error parsing Origipaddr: %s in %s", err, filename)
	}

	ParsedCallingpartynumber = helpers.RemoveSpaceFromString(raw.Callingpartynumber)
	ParsedCallingpartyunicodeloginuserid = helpers.RemoveSpaceFromString(raw.Callingpartyunicodeloginuserid)
	ParsedOrigmediatransportaddress_IP, err = helpers.ConvertStringToIPCisco(raw.Origmediatransportaddress_IP)
	if err != nil {
		logger.Error("Error parsing Origmediatransportaddress_IP: %s in %s", err, filename)
	}
	ParsedOrigvideotransportaddress_IP, err = helpers.ConvertStringToIPCisco(raw.Origvideotransportaddress_IP)
	if err != nil {
		logger.Error("Error parsing Origvideotransportaddress_IP: %s in %s", err, filename)
	}

	ParsedDestipaddr, err = helpers.ConvertStringToIPCisco(raw.Destipaddr)
	if err != nil {
		logger.Error("Error parsing Destipaddr: %s in %s", err, filename)
	}
	ParsedOriginalcalledpartynumber = helpers.RemoveSpaceFromString(raw.Originalcalledpartynumber)
	ParsedFinalcalledpartynumber = helpers.RemoveSpaceFromString(raw.Finalcalledpartynumber)
	ParsedFinalcalledpartyunicodeloginuserid = helpers.RemoveSpaceFromString(raw.Finalcalledpartyunicodeloginuserid)
	ParsedDestmediatransportaddress_IP, err = helpers.ConvertStringToIPCisco(raw.Destmediatransportaddress_IP)
	if err != nil {
		logger.Error("Error parsing Destmediatransportaddress_IP: %s in %s", err, filename)
	}
	ParsedDestvideotransportaddress_IP, err = helpers.ConvertStringToIPCisco(raw.Destvideotransportaddress_IP)
	if err != nil {
		logger.Error("Error parsing Destvideotransportaddress_IP: %s in %s", err, filename)
	}
	ParsedLastredirectdn = helpers.RemoveSpaceFromString(raw.Lastredirectdn)
	ParsedOriginalcalledpartynumberpartition = helpers.RemoveSpaceFromString(raw.Originalcalledpartynumberpartition)
	ParsedCallingpartynumberpartition = helpers.RemoveSpaceFromString(raw.Callingpartynumberpartition)
	ParsedFinalcalledpartynumberpartition = helpers.RemoveSpaceFromString(raw.Finalcalledpartynumberpartition)
	ParsedLastredirectdnpartition = helpers.RemoveSpaceFromString(raw.Lastredirectdnpartition)
	ParsedOrigdevicename = helpers.RemoveSpaceFromString(raw.Origdevicename)
	ParsedDestdevicename = helpers.RemoveSpaceFromString(raw.Destdevicename)
	ParsedGlobalcallid_Clusterid = helpers.RemoveSpaceFromString(raw.Globalcallid_Clusterid)
	ParsedComment = helpers.RemoveSpaceFromString(raw.Comment)
	ParsedAuthcodedescription = helpers.RemoveSpaceFromString(raw.Authcodedescription)
	ParsedClientmattercode = helpers.RemoveSpaceFromString(raw.Clientmattercode)
	ParsedAuthorizationcodevalue = helpers.RemoveSpaceFromString(raw.Authorizationcodevalue)
	ParsedOutpulsedcallingpartynumber = helpers.RemoveSpaceFromString(raw.Outpulsedcallingpartynumber)
	ParsedOutpulsedcalledpartynumber = helpers.RemoveSpaceFromString(raw.Outpulsedcalledpartynumber)
	ParsedOrigipv4v6addr = helpers.RemoveSpaceFromString(raw.Origipv4v6addr)
	ParsedDestipv4v6addr = helpers.RemoveSpaceFromString(raw.Destipv4v6addr)
	ParsedOrigvideotransportaddress_IP_Channel2, err = helpers.ConvertStringToIPCisco(raw.Origvideotransportaddress_IP_Channel2)
	if err != nil {
		logger.Error("Error parsing Origvideotransportaddress_IP_Channel2: %s in %s", err, filename)
	}
	ParsedDestvideotransportaddress_IP_Channel2, err = helpers.ConvertStringToIPCisco(raw.Destvideotransportaddress_IP_Channel2)
	if err != nil {
		logger.Error("Error parsing Destvideotransportaddress_IP_Channel2: %s in %s", err, filename)
	}
	ParsedIncomingprotocolcallref = helpers.RemoveSpaceFromString(raw.Incomingprotocolcallref)
	ParsedOutgoingprotocolcallref = helpers.RemoveSpaceFromString(raw.Outgoingprotocolcallref)
	ParsedHuntpilotpartition = helpers.RemoveSpaceFromString(raw.Huntpilotpartition)
	ParsedHuntpilotdn = helpers.RemoveSpaceFromString(raw.Huntpilotdn)
	ParsedIncomingicid = helpers.RemoveSpaceFromString(raw.Incomingicid)

	ParsedOutgoingicid = helpers.RemoveSpaceFromString(raw.Outgoingicid)
	ParsedOutpulsedoriginalcalledpartynumber = helpers.RemoveSpaceFromString(raw.Outpulsedoriginalcalledpartynumber)
	ParsedOutpulsedlastredirectingnumber = helpers.RemoveSpaceFromString(raw.Outpulsedlastredirectingnumber)
	ParsedCallingpartynumber_Uri = helpers.RemoveSpaceFromString(raw.Callingpartynumber_Uri)
	ParsedOriginalcalledpartynumber_Uri = helpers.RemoveSpaceFromString(raw.Originalcalledpartynumber_Uri)
	ParsedFinalcalledpartynumber_Uri = helpers.RemoveSpaceFromString(raw.Finalcalledpartynumber_Uri)
	ParsedLastredirectdn_Uri = helpers.RemoveSpaceFromString(raw.Lastredirectdn_Uri)
	ParsedMobilecallingpartynumber = helpers.RemoveSpaceFromString(raw.Mobilecallingpartynumber)
	ParsedFinalmobilecalledpartynumber = helpers.RemoveSpaceFromString(raw.Finalmobilecalledpartynumber)
	ParsedOrigmobiledevicename = helpers.RemoveSpaceFromString(raw.Origmobiledevicename)
	ParsedDestmobiledevicename = helpers.RemoveSpaceFromString(raw.Destmobiledevicename)
	ParsedOriginalcalledpartypattern = helpers.RemoveSpaceFromString(raw.Originalcalledpartypattern)
	ParsedFinalcalledpartypattern = helpers.RemoveSpaceFromString(raw.Finalcalledpartypattern)
	ParsedLastredirectingpartypattern = helpers.RemoveSpaceFromString(raw.Lastredirectingpartypattern)
	ParsedHuntpilotpattern = helpers.RemoveSpaceFromString(raw.Huntpilotpattern)
	ParsedOrigdevicetype = helpers.RemoveSpaceFromString(raw.Origdevicetype)
	ParsedDestdevicetype = helpers.RemoveSpaceFromString(raw.Destdevicetype)
	ParsedOrigdevicesessionid = helpers.RemoveSpaceFromString(raw.Origdevicesessionid)
	ParsedDestdevicesessionid = helpers.RemoveSpaceFromString(raw.Destdevicesessionid)

	return &CucmCdr{
		ID:                                      uuid.New().String(),
		OriginPkid:                              ParsedOriginpkid,
		FileClusterId:                           ParsedFileClusterId,
		FileNodeId:                              ParsedFileNodeId,
		FileDateTime:                            raw.FileDateTime,
		FileSequenceNumber:                      raw.FileSequenceNumber,
		Cdrrecordtype:                           ParsedCdrrecordtype,
		Globalcallid_Callmanagerid:              ParsedGlobalcallid_Callmanagerid,
		Globalcallid_Callid:                     ParsedGlobalcallid_Callid,
		Origlegcallidentifier:                   ParsedOriglegcallidentifier,
		Datetimeorigination:                     ParsedDatetimeorigination,
		Orignodeid:                              ParsedOrignodeid,
		Origspan:                                ParsedOrigspan,
		Origipaddr:                              ParsedOrigipaddr,
		Callingpartynumber:                      ParsedCallingpartynumber,
		Callingpartyunicodeloginuserid:          ParsedCallingpartyunicodeloginuserid,
		Origcause_Location:                      ParsedOrigcause_Location,
		Origcause_Value:                         ParsedOrigcause_Value,
		Origprecedencelevel:                     ParsedOrigprecedencelevel,
		Origmediatransportaddress_IP:            ParsedOrigmediatransportaddress_IP,
		Origmediatransportaddress_Port:          ParsedOrigmediatransportaddress_Port,
		Origmediacap_Payloadcapability:          ParsedOrigmediacap_Payloadcapability,
		Origmediacap_Maxframesperpacket:         ParsedOrigmediacap_Maxframesperpacket,
		Origmediacap_G723bitrate:                ParsedOrigmediacap_G723bitrate,
		Origvideocap_Codec:                      ParsedOrigvideocap_Codec,
		Origvideocap_Bandwidth:                  ParsedOrigvideocap_Bandwidth,
		Origvideocap_Resolution:                 ParsedOrigvideocap_Resolution,
		Origvideotransportaddress_IP:            ParsedOrigvideotransportaddress_IP,
		Origvideotransportaddress_Port:          ParsedOrigvideotransportaddress_Port,
		Origrsvpaudiostat:                       ParsedOrigrsvpaudiostat,
		Origrsvpvideostat:                       ParsedOrigrsvpvideostat,
		Destlegcallidentifier:                   ParsedDestlegcallidentifier,
		Destnodeid:                              ParsedDestnodeid,
		Destspan:                                ParsedDestspan,
		Destipaddr:                              ParsedDestipaddr,
		Originalcalledpartynumber:               ParsedOriginalcalledpartynumber,
		Finalcalledpartynumber:                  ParsedFinalcalledpartynumber,
		Finalcalledpartyunicodeloginuserid:      ParsedFinalcalledpartyunicodeloginuserid,
		Destcause_Location:                      ParsedDestcause_Location,
		Destcause_Value:                         ParsedDestcause_Value,
		Destprecedencelevel:                     ParsedDestprecedencelevel,
		Destmediatransportaddress_IP:            ParsedDestmediatransportaddress_IP,
		Destmediatransportaddress_Port:          ParsedDestmediatransportaddress_Port,
		Destmediacap_Payloadcapability:          ParsedDestmediacap_Payloadcapability,
		Destmediacap_Maxframesperpacket:         ParsedDestmediacap_Maxframesperpacket,
		Destmediacap_G723bitrate:                ParsedDestmediacap_G723bitrate,
		Destvideocap_Codec:                      ParsedDestvideocap_Codec,
		Destvideocap_Bandwidth:                  ParsedDestvideocap_Bandwidth,
		Destvideocap_Resolution:                 ParsedDestvideocap_Resolution,
		Destvideotransportaddress_IP:            ParsedDestvideotransportaddress_IP,
		Destvideotransportaddress_Port:          ParsedDestvideotransportaddress_Port,
		Destrsvpaudiostat:                       ParsedDestrsvpaudiostat,
		Destrsvpvideostat:                       ParsedDestrsvpvideostat,
		Datetimeconnect:                         ParsedDatetimeconnect,
		Datetimedisconnect:                      ParsedDatetimedisconnect,
		Lastredirectdn:                          ParsedLastredirectdn,
		Originalcalledpartynumberpartition:      ParsedOriginalcalledpartynumberpartition,
		Callingpartynumberpartition:             ParsedCallingpartynumberpartition,
		Finalcalledpartynumberpartition:         ParsedFinalcalledpartynumberpartition,
		Lastredirectdnpartition:                 ParsedLastredirectdnpartition,
		Duration:                                ParsedDuration,
		Origdevicename:                          ParsedOrigdevicename,
		Destdevicename:                          ParsedDestdevicename,
		Origcallterminationonbehalfof:           ParsedOrigcallterminationonbehalfof,
		Destcallterminationonbehalfof:           ParsedDestcallterminationonbehalfof,
		Origcalledpartyredirectonbehalfof:       ParsedOrigcalledpartyredirectonbehalfof,
		Lastredirectredirectonbehalfof:          ParsedLastredirectredirectonbehalfof,
		Origcalledpartyredirectreason:           ParsedOrigcalledpartyredirectreason,
		Lastredirectredirectreason:              ParsedLastredirectredirectreason,
		Destconversationid:                      ParsedDestconversationid,
		Globalcallid_Clusterid:                  ParsedGlobalcallid_Clusterid,
		Joinonbehalfof:                          ParsedJoinonbehalfof,
		Comment:                                 ParsedComment,
		Authcodedescription:                     ParsedAuthcodedescription,
		Authorizationlevel:                      ParsedAuthorizationlevel,
		Clientmattercode:                        ParsedClientmattercode,
		Origdtmfmethod:                          ParsedOrigdtmfmethod,
		Destdtmfmethod:                          ParsedDestdtmfmethod,
		Callsecuredstatus:                       ParsedCallsecuredstatus,
		Origconversationid:                      ParsedOrigconversationid,
		Origmediacap_Bandwidth:                  ParsedOrigmediacap_Bandwidth,
		Destmediacap_Bandwidth:                  ParsedDestmediacap_Bandwidth,
		Authorizationcodevalue:                  ParsedAuthorizationcodevalue,
		Outpulsedcallingpartynumber:             ParsedOutpulsedcallingpartynumber,
		Outpulsedcalledpartynumber:              ParsedOutpulsedcalledpartynumber,
		Origipv4v6addr:                          ParsedOrigipv4v6addr,
		Destipv4v6addr:                          ParsedDestipv4v6addr,
		Origvideocap_Codec_Channel2:             ParsedOrigvideocap_Codec_Channel2,
		Origvideocap_Bandwidth_Channel2:         ParsedOrigvideocap_Bandwidth_Channel2,
		Origvideocap_Resolution_Channel2:        ParsedOrigvideocap_Resolution_Channel2,
		Origvideotransportaddress_IP_Channel2:   ParsedOrigvideotransportaddress_IP_Channel2,
		Origvideotransportaddress_Port_Channel2: ParsedOrigvideotransportaddress_Port_Channel2,
		Origvideochannel_Role_Channel2:          ParsedOrigvideochannel_Role_Channel2,
		Destvideocap_Codec_Channel2:             ParsedDestvideocap_Codec_Channel2,
		Destvideocap_Bandwidth_Channel2:         ParsedDestvideocap_Bandwidth_Channel2,
		Destvideocap_Resolution_Channel2:        ParsedDestvideocap_Resolution_Channel2,
		Destvideotransportaddress_IP_Channel2:   ParsedDestvideotransportaddress_IP_Channel2,
		Destvideotransportaddress_Port_Channel2: ParsedDestvideotransportaddress_Port_Channel2,
		Destvideochannel_Role_Channel2:          ParsedDestvideochannel_Role_Channel2,
		Incomingprotocolid:                      ParsedIncomingprotocolid,
		Incomingprotocolcallref:                 ParsedIncomingprotocolcallref,
		Outgoingprotocolid:                      ParsedOutgoingprotocolid,
		Outgoingprotocolcallref:                 ParsedOutgoingprotocolcallref,
		Currentroutingreason:                    ParsedCurrentroutingreason,
		Origroutingreason:                       ParsedOrigroutingreason,
		Lastredirectingroutingreason:            ParsedLastredirectingroutingreason,
		Huntpilotpartition:                      ParsedHuntpilotpartition,
		Huntpilotdn:                             ParsedHuntpilotdn,
		Calledpartypatternusage:                 ParsedCalledpartypatternusage,
		Incomingicid:                            ParsedIncomingicid,
		Incomingorigioi:                         raw.Incomingorigioi,
		Incomingtermioi:                         raw.Incomingtermioi,
		Outgoingicid:                            ParsedOutgoingicid,
		Outgoingorigioi:                         raw.Outgoingorigioi,
		Outgoingtermioi:                         raw.Outgoingtermioi,
		Outpulsedoriginalcalledpartynumber:      ParsedOutpulsedoriginalcalledpartynumber,
		Outpulsedlastredirectingnumber:          ParsedOutpulsedlastredirectingnumber,
		Wascallqueued:                           ParsedWascallqueued,
		Totalwaittimeinqueue:                    ParsedTotalwaittimeinqueue,
		Callingpartynumber_Uri:                  ParsedCallingpartynumber_Uri,
		Originalcalledpartynumber_Uri:           ParsedOriginalcalledpartynumber_Uri,
		Finalcalledpartynumber_Uri:              ParsedFinalcalledpartynumber_Uri,
		Lastredirectdn_Uri:                      ParsedLastredirectdn_Uri,
		Mobilecallingpartynumber:                ParsedMobilecallingpartynumber,
		Finalmobilecalledpartynumber:            ParsedFinalmobilecalledpartynumber,
		Origmobiledevicename:                    ParsedOrigmobiledevicename,
		Destmobiledevicename:                    ParsedDestmobiledevicename,
		Origmobilecallduration:                  ParsedOrigmobilecallduration,
		Destmobilecallduration:                  ParsedDestmobilecallduration,
		Mobilecalltype:                          ParsedMobilecalltype,
		Originalcalledpartypattern:              ParsedOriginalcalledpartypattern,
		Finalcalledpartypattern:                 ParsedFinalcalledpartypattern,
		Lastredirectingpartypattern:             ParsedLastredirectingpartypattern,
		Huntpilotpattern:                        ParsedHuntpilotpattern,
		Origdevicetype:                          ParsedOrigdevicetype,
		Destdevicetype:                          ParsedDestdevicetype,
		Origdevicesessionid:                     ParsedOrigdevicesessionid,
		Destdevicesessionid:                     ParsedDestdevicesessionid,
	}, nil
}
