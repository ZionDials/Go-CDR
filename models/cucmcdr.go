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

type CucmCdr struct {
	ID                                      string
	OriginPkid                              *string `gorm:"unique;not null"`
	FileClusterId                           *string
	FileNodeId                              *string
	FileDateTime                            *int64
	FileSequenceNumber                      *int64
	Cdrrecordtype                           *int64
	Globalcallid_Callmanagerid              *int64
	Globalcallid_Callid                     *int64
	Origlegcallidentifier                   *int64
	Datetimeorigination                     *int64
	Orignodeid                              *int64
	Origspan                                *int64
	Origipaddr                              *string
	Callingpartynumber                      *string
	Callingpartyunicodeloginuserid          *string
	Origcause_Location                      *int64
	Origcause_Value                         *int64
	Origprecedencelevel                     *int64
	Origmediatransportaddress_IP            *string
	Origmediatransportaddress_Port          *int64
	Origmediacap_Payloadcapability          *int64
	Origmediacap_Maxframesperpacket         *int64
	Origmediacap_G723bitrate                *int64
	Origvideocap_Codec                      *int64
	Origvideocap_Bandwidth                  *int64
	Origvideocap_Resolution                 *int64
	Origvideotransportaddress_IP            *string
	Origvideotransportaddress_Port          *int64
	Origrsvpaudiostat                       *int64
	Origrsvpvideostat                       *int64
	Destlegcallidentifier                   *int64
	Destnodeid                              *int64
	Destspan                                *int64
	Destipaddr                              *string
	Originalcalledpartynumber               *string
	Finalcalledpartynumber                  *string
	Finalcalledpartyunicodeloginuserid      *string
	Destcause_Location                      *int64
	Destcause_Value                         *int64
	Destprecedencelevel                     *int64
	Destmediatransportaddress_IP            *string
	Destmediatransportaddress_Port          *int64
	Destmediacap_Payloadcapability          *int64
	Destmediacap_Maxframesperpacket         *int64
	Destmediacap_G723bitrate                *int64
	Destvideocap_Codec                      *int64
	Destvideocap_Bandwidth                  *int64
	Destvideocap_Resolution                 *int64
	Destvideotransportaddress_IP            *string
	Destvideotransportaddress_Port          *int64
	Destrsvpaudiostat                       *int64
	Destrsvpvideostat                       *int64
	Datetimeconnect                         *int64
	Datetimedisconnect                      *int64
	Lastredirectdn                          *string
	Originalcalledpartynumberpartition      *string
	Callingpartynumberpartition             *string
	Finalcalledpartynumberpartition         *string
	Lastredirectdnpartition                 *string
	Duration                                *int64
	Origdevicename                          *string
	Destdevicename                          *string
	Origcallterminationonbehalfof           *int64
	Destcallterminationonbehalfof           *int64
	Origcalledpartyredirectonbehalfof       *int64
	Lastredirectredirectonbehalfof          *int64
	Origcalledpartyredirectreason           *int64
	Lastredirectredirectreason              *int64
	Destconversationid                      *int64
	Globalcallid_Clusterid                  *string
	Joinonbehalfof                          *int64
	Comment                                 *string
	Authcodedescription                     *string
	Authorizationlevel                      *int64
	Clientmattercode                        *string
	Origdtmfmethod                          *int64
	Destdtmfmethod                          *int64
	Callsecuredstatus                       *int64
	Origconversationid                      *int64
	Origmediacap_Bandwidth                  *int64
	Destmediacap_Bandwidth                  *int64
	Authorizationcodevalue                  *string
	Outpulsedcallingpartynumber             *string
	Outpulsedcalledpartynumber              *string
	Origipv4v6addr                          *string
	Destipv4v6addr                          *string
	Origvideocap_Codec_Channel2             *int64
	Origvideocap_Bandwidth_Channel2         *int64
	Origvideocap_Resolution_Channel2        *int64
	Origvideotransportaddress_IP_Channel2   *string
	Origvideotransportaddress_Port_Channel2 *int64
	Origvideochannel_Role_Channel2          *int64
	Destvideocap_Codec_Channel2             *int64
	Destvideocap_Bandwidth_Channel2         *int64
	Destvideocap_Resolution_Channel2        *int64
	Destvideotransportaddress_IP_Channel2   *string
	Destvideotransportaddress_Port_Channel2 *int64
	Destvideochannel_Role_Channel2          *int64
	Incomingprotocolid                      *int64
	Incomingprotocolcallref                 *string
	Outgoingprotocolid                      *int64
	Outgoingprotocolcallref                 *string
	Currentroutingreason                    *int64
	Origroutingreason                       *int64
	Lastredirectingroutingreason            *int64
	Huntpilotpartition                      *string
	Huntpilotdn                             *string
	Calledpartypatternusage                 *int64
	Incomingicid                            *string
	Incomingorigioi                         *string
	Incomingtermioi                         *string
	Outgoingicid                            *string
	Outgoingorigioi                         *string
	Outgoingtermioi                         *string
	Outpulsedoriginalcalledpartynumber      *string
	Outpulsedlastredirectingnumber          *string
	Wascallqueued                           *int64
	Totalwaittimeinqueue                    *int64
	Callingpartynumber_Uri                  *string
	Originalcalledpartynumber_Uri           *string
	Finalcalledpartynumber_Uri              *string
	Lastredirectdn_Uri                      *string
	Mobilecallingpartynumber                *string
	Finalmobilecalledpartynumber            *string
	Origmobiledevicename                    *string
	Destmobiledevicename                    *string
	Origmobilecallduration                  *int64
	Destmobilecallduration                  *int64
	Mobilecalltype                          *int64
	Originalcalledpartypattern              *string
	Finalcalledpartypattern                 *string
	Lastredirectingpartypattern             *string
	Huntpilotpattern                        *string
	Origdevicetype                          *string
	Destdevicetype                          *string
	Origdevicesessionid                     *string
	Destdevicesessionid                     *string
}
