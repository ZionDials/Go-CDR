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

type CucmCmr struct {
	ID                                  string
	Originpkid                          *string `gorm:"unique;not null"`
	FileClusterId                       *string
	FileNodeId                          *string
	FileDateTime                        *int64
	FileSequenceNumber                  *int64
	Cdrrecordtype                       *int64
	Globalcallid_Callmanagerid          *int64
	Globalcallid_Callid                 *int64
	Nodeid                              *int64
	Directorynum                        *string
	Callidentifier                      *int64
	Datetimestamp                       *int64
	Numberpacketssent                   *int64
	Numberoctetssent                    *int64
	Numberpacketsreceived               *int64
	Numberoctetsreceived                *int64
	Numberpacketslost                   *int64
	Jitter                              *int64
	Latency                             *int64
	Directorynumpartition               *string
	Globalcallid_Clusterid              *string
	Devicename                          *string
	Duration                            *int64
	Videocontenttype                    *string
	Videoduration                       *int64
	Numbervideopacketssent              *int64
	Numbervideooctetssent               *int64
	Numbervideopacketsreceived          *int64
	Numbervideooctetsreceived           *int64
	Numbervideopacketslost              *int64
	Videoaveragejitter                  *int64
	Videoroundtriptime                  *int64
	Videoonewaydelay                    *int64
	Videoreceptionmetrics               *string
	Videotransmissionmetrics            *string
	Videocontenttype_Channel2           *string
	Videoduration_Channel2              *int64
	Numbervideopacketssent_Channel2     *int64
	Numbervideooctetssent_Channel2      *int64
	Numbervideopacketsreceived_Channel2 *int64
	Numbervideooctetsreceived_Channel2  *int64
	Numbervideopacketslost_Channel2     *int64
	Videoaveragejitter_Channel2         *int64
	Videoroundtriptime_Channel2         *int64
	Videoonewaydelay_Channel2           *int64
	Videoreceptionmetrics_Channel2      *string
	Videotransmissionmetrics_Channel2   *string
	Localsessionid                      *string
	Remotesessionid                     *string
	Headsetsn                           *string
	Headsetmetrics                      *string
	VQCCR                               *float64
	VQICR                               *float64
	Vqicrmx                             *float64
	VQCS                                *int64
	VQSCS                               *int64
	Vqver                               *float64
	Vqvorxcodec                         *string
	VQCID                               *int64
	Vqvopktsizems                       *int64
	Vqvopktlost                         *int64
	Vqvopktdis                          *int64
	Vqvoonewaydelayms                   *int64
	Vqmaxjitter                         *int64
	VQMLQK                              *float64
	Vqmlqkav                            *float64
	Vqmlqkmn                            *float64
	Vqmlqkmx                            *float64
	Vqmlqkvr                            *float64
}
