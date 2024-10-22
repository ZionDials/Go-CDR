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
	"time"

	"github.com/google/uuid"
	"github.com/ziondials/go-cdr/helpers"
	"github.com/ziondials/go-cdr/logger"
)

var (
	TWCCode          = "TWC"
	CFACode          = "CFA"
	CFNACode         = "CFNA"
	CFBYCode         = "CFBY"
	BXFERCode        = "BXFER"
	CXFERCode        = "CXFER"
	HOLDCode         = "HOLD"
	RESUMECode       = "RESUME"
	PhoneNumberSplit = "#:"

	CubeLegTypeTelephony  = 1
	CubeLegTypeVoIP       = 2
	CubeLegTypeMMOIP      = 3
	CubeLegTypeFrameRelay = 4
	CubeLegTypeATM        = 5

	CubeSessionProtocols = []string{"sip", "Cisco", "sipv2", "other"}

	CubeLegTypes = []int{CubeLegTypeTelephony, CubeLegTypeVoIP, CubeLegTypeMMOIP, CubeLegTypeFrameRelay, CubeLegTypeATM}

	H323CauseCodes = []string{
		"1B",
		"43",
		"15",
		"3A",
		"1C",
		"42",
		"45",
		"13",
		"24",
		"21",
		"2D",
		"2B",
		"1F",
		"1",
		"12",
		"11",
		"2C",
		"A",
		"2",
		"38",
		"B",
		"17",
		"8",
		"46",
		"4",
		"3F",
		"29",
		"35",
		"31",
		"1E",
		"3",
		"26",
		"2A",
		"28",
		"2F",
		"10",
		"18",
		"5",
		"39",
		"19",
		"9",
		"32",
		"16",
		"23",
		"27",
		"1D",
		"36",
		"34",
		"33",
		"2E",
		"6",
		"22",
		"44",
		"41",
		"1A",
		"25",
		"7",
		"37",
		"4F",
		"51",
		"52",
		"53",
		"54",
		"55",
		"56",
		"57",
		"58",
		"59",
		"5A",
		"5B",
		"5C",
		"5D",
		"5F",
		"60",
		"61",
		"62",
		"63",
		"64",
		"65",
		"66",
		"67",
		"6F",
		"7F",
		"80",
		"0",
		"14",
	}
	TwoWayCallTypes = []string{TWCCode}

	CallForwardTypes = []string{CFACode, CFNACode, CFBYCode}

	TransferTypes = []string{BXFERCode, CXFERCode}

	HoldTypes = []string{HOLDCode, RESUMECode}

	PhoneNumberSplitType = []string{PhoneNumberSplit}
)

type RawCubeCDRs []*RawCubeCDR

type CubeCDRs []*CubeCDR

type CubeCDR struct {
	ID                  string
	InvalidNTPReference bool
	Hostname            *string // `gorm:"uniqueIndex:cube_cdr_index"`
	Filename            *string
	FileTimestamp       *int64
	RecordTimestamp     *int64
	CallId              *int64 // `gorm:"uniqueIndex:cube_cdr_index"`
	CdrType             *int64

	AccountCode                     *string
	AcomLevel                       *int64
	AlertTime                       *int64
	BackwardCallId                  *string
	BytesIn                         *int64
	BytesOut                        *int64
	CallForwardCount                *string
	CallForwardFeatureCorrelationId *string
	CallForwardFeatureID            *string
	CallForwardFeatureStatus        *string
	CallForwardLegID                *int64
	CallForwardReason               *string
	CallForwardedFromNumber         *string
	CallForwardedNumber             *string
	CallForwardedToNumber           *string
	CallForwardingFromNumber        *string
	CallingPartyCategory            *string
	CarrierId                       *string
	ChargeNumber                    *string
	ChargedUnits                    *int64
	Clid                            *string
	CodecBytes                      *int64
	CodecTypeRate                   *string
	CustBizGrpId                    *string
	DisconnectText                  *string
	Dnis                            *string
	DspId                           *string
	EarlyPackets                    *int64
	FacDigit                        *string
	FacStatus                       *string
	FaxrelayDirection               *string
	FaxrelayEcmStatus               *string
	FaxrelayEncapProtocol           *string
	FaxrelayFaxSuccess              *string
	FaxrelayInitHsMod               *int64
	FaxrelayJitBufOvflow            *int64
	FaxrelayMaxJitBufDepth          *int64
	FaxrelayMrHsMod                 *int64
	FaxrelayNsfCountryCode          *string
	FaxrelayNsfManufCode            *string
	FaxrelayNumPages                *int64
	FaxrelayPktConceal              *int64
	FaxrelayRxPackets               *int64
	FaxrelayStartTime               *string
	FaxrelayStopTime                *string
	FaxrelayTxPackets               *int64
	FeatureId                       *string
	FeatureIdField1                 *string // `gorm:"uniqueIndex:cube_cdr_index"`
	FeatureIdField2                 *int64  // `gorm:"uniqueIndex:cube_cdr_index"`
	FeatureOpStatus                 *string
	FeatureOpTime                   *string
	FeatureOperation                *string
	GapfillWithInterpolation        *int64
	GapfillWithPrediction           *int64
	GapfillWithRedundancy           *int64
	GapfillWithSilence              *int64
	GkXlatedCdn                     *string
	GkXlatedCgn                     *string
	GtdGwRxdCnn                     *string
	GtdGwRxdOcn                     *string
	GtdOrigCic                      *string
	GtdTermCic                      *string
	GwCollectedCdn                  *string
	GwFinalXlatedCdn                *string
	GwFinalXlatedCgn                *string
	GwFinalXlatedRdn                *string
	GwRxdCdn                        *string
	GwRxdCgn                        *string
	GwRxdRdn                        *string
	H323CallOrigin                  *string
	H323ConfId                      *string // `gorm:"uniqueIndex:cube_cdr_index"`
	H323ConnectTime                 *int64
	H323DisconnectCause             *string
	H323DisconnectTime              *int64
	H323IvrOut                      *string
	H323SetupTime                   *int64
	H323VoiceQuality                *int64
	HeldDN                          *int64
	HiwaterPlayoutDelay             *int64
	HoldFeatureCorrelationId        *string
	HoldFeatureID                   *string
	HoldLegID                       *string
	HoldPhoneTag                    *string
	HoldReason                      *string
	HoldSharedLine                  *int64
	HoldStatus                      *string
	HoldUsername                    *string
	HoldingDN                       *int64
	InCarrierId                     *string
	InIntrfcDesc                    *string
	InLpcorGroup                    *string
	InTrunkgroupLabel               *string
	IncomingArea                    *string
	InfoType                        *string
	InternalErrorCode               *string
	IpHop                           *int64
	IpPbxMode                       *string
	IpPhoneInfo                     *string
	LatePackets                     *int64
	LegType                         *int // `gorm:"uniqueIndex:cube_cdr_index"`
	LocalHostname                   *string
	LogicalIfIndex                  *int64
	LostPackets                     *int64
	LowaterPlayoutDelay             *int64
	MaxBitrate                      *string
	NoiseLevel                      *int64
	OntimeRvPlayout                 *int64
	OriginatingLineInfo             *string
	OutCarrierId                    *string
	OutIntrfcDesc                   *string
	OutLpcorGroup                   *string
	OutTrunkgroupLabel              *string
	OutgoingArea                    *string
	OverrideSessionTime             *int64
	PaksIn                          *int64
	PaksOut                         *int64
	PeerAddress                     *string
	PeerId                          *int64
	PeerIfIndex                     *int64
	PeerSubAddress                  *string
	ReceiveDelay                    *int64
	RedirectedStationAddress        *string
	RedirectedStationNOA            *string
	RedirectedStationNPI            *string
	RedirectedStationPI             *string
	RemoteMediaAddress              *string
	RemoteMediaId                   *string
	RemoteMediaUdpPort              *int64
	RemoteUdpPort                   *int64
	RoundTripDelay                  *int64
	ServiceDescriptor               *string
	SessionProtocol                 *string
	Subscriber                      *string
	SuppSvcXferBy                   *string
	TWCCalledNumber                 *string
	TWCCallingNumber                *string
	TWCFeatureCorrelationId         *string
	TWCFeatureID                    *string
	TWCFeatureStatus                *int64
	TWCLegID                        *int64
	TransferConsultationID          *int64
	TransferFeatureCorrelationId    *string
	TransferFeatureID               *string
	TransferFeatureStatus           *string
	TransferForwardingReason        *string
	TransferLegID                   *string
	TransferStatus                  *int64
	TransferredFromPart             *string
	TransferredNumber               *string
	TransferredToParty              *string
	TransmissionMediumReq           *string
	TxDuration                      *int64
	Username                        *string
	VadEnable                       *bool
	VoiceFeature                    *string
	VoiceTxDuration                 *int64
}

func (raw *RawCubeCDR) Parse(filename string) (*CubeCDR, error) {

	var ParsedAccountCode *string
	var ParsedAcomLevel *int64
	var ParsedAlertTime *int64
	var ParsedBytesIn *int64
	var ParsedBytesOut *int64
	var ParsedCallForwardCount *string
	var ParsedCallForwardFeatureCorrelationId *string
	var ParsedCallForwardFeatureID *string
	var ParsedCallForwardFeatureStatus *string
	var ParsedCallForwardLegID *int64
	var ParsedCallForwardReason *string
	var ParsedCallForwardedFromNumber *string
	var ParsedCallForwardedNumber *string
	var ParsedCallForwardedToNumber *string
	var ParsedCallForwardingFromNumber *string
	var ParsedCallId *int64
	var ParsedCallingPartyCategory *string
	var ParsedCarrierId *string
	var ParsedCdrType *int64
	var ParsedChargeNumber *string
	var ParsedChargedUnits *int64
	var ParsedClid *string
	var ParsedCodecBytes *int64
	var ParsedCodecTypeRate *string
	var ParsedCustBizGrpId *string
	var ParsedDnis *string
	var ParsedDspId *string
	var ParsedEarlyPackets *int64
	var ParsedFacDigit *string
	var ParsedFacStatus *string
	var ParsedFaxrelayDirection *string
	var ParsedFaxrelayEcmStatus *string
	var ParsedFaxrelayEncapProtocol *string
	var ParsedFaxrelayFaxSuccess *string
	var ParsedFaxrelayInitHsMod *int64
	var ParsedFaxrelayJitBufOvflow *int64
	var ParsedFaxrelayMaxJitBufDepth *int64
	var ParsedFaxrelayMrHsMod *int64
	var ParsedFaxrelayNsfCountryCode *string
	var ParsedFaxrelayNsfManufCode *string
	var ParsedFaxrelayNumPages *int64
	var ParsedFaxrelayPktConceal *int64
	var ParsedFaxrelayRxPackets *int64
	var ParsedFaxrelayStartTime *string
	var ParsedFaxrelayStopTime *string
	var ParsedFaxrelayTxPackets *int64
	var ParsedFeatureId *string
	var ParsedFeatureIdField1 *string
	var ParsedFeatureIdField2 *int64
	var ParsedFeatureOpStatus *string
	var ParsedFeatureOpTime *string
	var ParsedFeatureOperation *string
	var ParsedFilename *string
	var ParsedFiletimestamp *int64
	var ParsedGapfillWithInterpolation *int64
	var ParsedGapfillWithPrediction *int64
	var ParsedGapfillWithRedundancy *int64
	var ParsedGapfillWithSilence *int64
	var ParsedGkXlatedCdn *string
	var ParsedGkXlatedCgn *string
	var ParsedGtdGwRxdCnn *string
	var ParsedGtdGwRxdOcn *string
	var ParsedGtdOrigCic *string
	var ParsedGtdTermCic *string
	var ParsedGwCollectedCdn *string
	var ParsedGwFinalXlatedCdn *string
	var ParsedGwFinalXlatedCgn *string
	var ParsedGwFinalXlatedRdn *string
	var ParsedGwRxdCdn *string
	var ParsedGwRxdCgn *string
	var ParsedGwRxdRdn *string
	var ParsedH323CallOrigin *string
	var ParsedH323ConnectTime *int64
	var ParsedH323DisconnectCause *string
	var ParsedH323DisconnectTime *int64
	var ParsedH323IvrOut *string
	var ParsedH323SetupTime *int64
	var ParsedH323VoiceQuality *int64
	var ParsedHeldDN *int64
	var ParsedHiwaterPlayoutDelay *int64
	var ParsedHoldFeatureCorrelationId *string
	var ParsedHoldFeatureID *string
	var ParsedHoldLegID *string
	var ParsedHoldPhoneTag *string
	var ParsedHoldReason *string
	var ParsedHoldSharedLine *int64
	var ParsedHoldStatus *string
	var ParsedHoldUsername *string
	var ParsedHoldingDN *int64
	var ParsedHostname *string
	var ParsedIPHop *int64
	var ParsedInCarrierId *string
	var ParsedInLpcorGroup *string
	var ParsedInTrunkgroupLabel *string
	var ParsedIncomingArea *string
	var ParsedInfoType *string
	var ParsedInternalErrorCode *string
	var ParsedIpPbxMode *string
	var ParsedIpPhoneInfo *string
	var ParsedLatePackets *int64
	var ParsedLegType *int
	var ParsedLogicalIfIndex *int64
	var ParsedLostPackets *int64
	var ParsedLowaterPlayoutDelay *int64
	var ParsedMaxBitrate *string
	var ParsedNoiseLevel *int64
	var ParsedOntimeRvPlayout *int64
	var ParsedOriginatingLineInfo *string
	var ParsedOutCarrierId *string
	var ParsedOutLpcorGroup *string
	var ParsedOutTrunkgroupLabel *string
	var ParsedOutgoingArea *string
	var ParsedOverrideSessionTime *int64
	var ParsedPaksIn *int64
	var ParsedPaksOut *int64
	var ParsedPeerAddress *string
	var ParsedPeerId *int64
	var ParsedPeerIfIndex *int64
	var ParsedPeerSubAddress *string
	var ParsedReceiveDelay *int64
	var ParsedRecordTimestamp *int64
	var ParsedRemoteMediaAddress *string
	var ParsedRemoteMediaId *string
	var ParsedRemoteMediaUdpPort *int64
	var ParsedRemoteUdpPort *int64
	var ParsedRoundTripDelay *int64
	var ParsedServiceDescriptor *string
	var ParsedSubscriber *string
	var ParsedSuppSvcXferBy *string
	var ParsedTWCCalledNumber *string
	var ParsedTWCCallingNumber *string
	var ParsedTWCFeatureCorrelationId *string
	var ParsedTWCFeatureID *string
	var ParsedTWCFeatureStatus *int64
	var ParsedTWCLegID *int64
	var ParsedTimeLocation *time.Location
	var ParsedTransferConsultationID *int64
	var ParsedTransferFeatureCorrelationId *string
	var ParsedTransferFeatureID *string
	var ParsedTransferFeatureStatus *string
	var ParsedTransferForwardingReason *string
	var ParsedTransferLegID *string
	var ParsedTransferStatus *int64
	var ParsedTransferredFromPart *string
	var ParsedTransferredNumber *string
	var ParsedTransferredToParty *string
	var ParsedTransmissionMediumReq *string
	var ParsedTxDuration *int64
	var ParsedUsername *string
	var ParsedVadEnable *bool
	var ParsedVoiceFeature *string
	var ParsedVoiceTxDuration *int64
	var ParsedSessionProtocol *string

	ParsedAccountCode = helpers.RemoveSpaceFromString(raw.AccountCode)
	ParsedCallingPartyCategory = helpers.RemoveSpaceFromString(raw.CallingPartyCategory)
	ParsedCarrierId = helpers.RemoveSpaceFromString(raw.CarrierId)
	ParsedChargeNumber = helpers.RemoveSpaceFromString(raw.ChargeNumber)
	ParsedClid = helpers.RemoveSpaceFromString(raw.Clid)
	ParsedCodecTypeRate = helpers.RemoveSpaceFromString(raw.CodecTypeRate)
	ParsedCustBizGrpId = helpers.RemoveSpaceFromString(raw.CustBizGrpId)
	ParsedDnis = helpers.RemoveSpaceFromString(raw.Dnis)
	ParsedDspId = helpers.RemoveSpaceFromString(raw.DspId)
	ParsedFacDigit = helpers.RemoveSpaceFromString(raw.FacDigit)
	ParsedFacStatus = helpers.RemoveSpaceFromString(raw.FacStatus)
	ParsedFaxrelayDirection = helpers.RemoveSpaceFromString(raw.FaxrelayDirection)
	ParsedFaxrelayEcmStatus = helpers.RemoveSpaceFromString(raw.FaxrelayEcmStatus)
	ParsedFaxrelayEncapProtocol = helpers.RemoveSpaceFromString(raw.FaxrelayEncapProtocol)
	ParsedFaxrelayFaxSuccess = helpers.RemoveSpaceFromString(raw.FaxrelayFaxSuccess)
	ParsedFaxrelayNsfCountryCode = helpers.RemoveSpaceFromString(raw.FaxrelayNsfCountryCode)
	ParsedFaxrelayNsfManufCode = helpers.RemoveSpaceFromString(raw.FaxrelayNsfManufCode)
	ParsedFaxrelayStartTime = helpers.RemoveSpaceFromString(raw.FaxrelayStartTime)
	ParsedFaxrelayStopTime = helpers.RemoveSpaceFromString(raw.FaxrelayStopTime)
	ParsedFeatureId = helpers.RemoveSpaceFromString(raw.FeatureId)
	ParsedFeatureIdField1 = helpers.RemoveSpaceFromString(raw.FeatureIdField1)
	ParsedFeatureOpStatus = helpers.RemoveSpaceFromString(raw.FeatureOpStatus)
	ParsedFeatureOpTime = helpers.RemoveSpaceFromString(raw.FeatureOpTime)
	ParsedFeatureOperation = helpers.RemoveSpaceFromString(raw.FeatureOperation)
	ParsedFilename = helpers.RemoveSpaceFromString(raw.Filename)
	ParsedGkXlatedCdn = helpers.ExtractPhoneNumberFromString(raw.GkXlatedCdn)
	ParsedGkXlatedCgn = helpers.ExtractPhoneNumberFromString(raw.GkXlatedCgn)
	ParsedGtdGwRxdCnn = helpers.RemoveSpaceFromString(raw.GtdGwRxdCnn)
	ParsedGtdGwRxdOcn = helpers.RemoveSpaceFromString(raw.GtdGwRxdOcn)
	ParsedGtdOrigCic = helpers.RemoveSpaceFromString(raw.GtdOrigCic)
	ParsedGtdTermCic = helpers.RemoveSpaceFromString(raw.GtdTermCic)
	ParsedGwCollectedCdn = helpers.RemoveSpaceFromString(raw.GwCollectedCdn)
	ParsedGwFinalXlatedCdn = helpers.ExtractPhoneNumberFromString(raw.GwFinalXlatedCdn)
	ParsedGwFinalXlatedCgn = helpers.ExtractPhoneNumberFromString(raw.GwFinalXlatedCgn)
	ParsedGwFinalXlatedRdn = helpers.ExtractPhoneNumberFromString(raw.GwFinalXlatedRdn)
	ParsedGwRxdCdn = helpers.ExtractPhoneNumberFromString(raw.GwRxdCdn)
	ParsedGwRxdCgn = helpers.ExtractPhoneNumberFromString(raw.GwRxdCgn)
	ParsedGwRxdRdn = helpers.ExtractPhoneNumberFromString(raw.GwRxdRdn)
	ParsedH323CallOrigin = helpers.RemoveSpaceFromString(raw.H323CallOrigin)
	ParsedH323IvrOut = helpers.RemoveSpaceFromString(raw.H323IvrOut)
	ParsedHostname = helpers.RemoveSpaceFromString(raw.Hostname)
	ParsedInCarrierId = helpers.RemoveSpaceFromString(raw.InCarrierId)
	ParsedInLpcorGroup = helpers.RemoveSpaceFromString(raw.InLpcorGroup)
	ParsedInTrunkgroupLabel = helpers.RemoveSpaceFromString(raw.InTrunkgroupLabel)
	ParsedIncomingArea = helpers.RemoveSpaceFromString(raw.IncomingArea)
	ParsedInfoType = helpers.RemoveSpaceFromString(raw.InfoType)
	ParsedInternalErrorCode = helpers.RemoveSpaceFromString(raw.InternalErrorCode)
	ParsedIpPbxMode = helpers.RemoveSpaceFromString(raw.IpPbxMode)
	ParsedIpPhoneInfo = helpers.RemoveSpaceFromString(raw.IpPhoneInfo)
	ParsedMaxBitrate = helpers.RemoveSpaceFromString(raw.MaxBitrate)
	ParsedOriginatingLineInfo = helpers.RemoveSpaceFromString(raw.OriginatingLineInfo)
	ParsedOutCarrierId = helpers.RemoveSpaceFromString(raw.OutCarrierId)
	ParsedOutLpcorGroup = helpers.RemoveSpaceFromString(raw.OutLpcorGroup)
	ParsedOutTrunkgroupLabel = helpers.RemoveSpaceFromString(raw.OutTrunkgroupLabel)
	ParsedOutgoingArea = helpers.RemoveSpaceFromString(raw.OutgoingArea)
	ParsedPeerAddress = helpers.RemoveSpaceFromString(raw.PeerAddress)
	ParsedPeerSubAddress = helpers.RemoveSpaceFromString(raw.PeerSubAddress)
	ParsedRemoteMediaAddress = helpers.RemoveSpaceFromString(raw.RemoteMediaAddress)
	ParsedRemoteMediaId = helpers.RemoveSpaceFromString(raw.RemoteMediaId)
	ParsedServiceDescriptor = helpers.RemoveSpaceFromString(raw.ServiceDescriptor)
	ParsedSubscriber = helpers.RemoveSpaceFromString(raw.Subscriber)
	ParsedSuppSvcXferBy = helpers.RemoveSpaceFromString(raw.SuppSvcXferBy)
	ParsedTransmissionMediumReq = helpers.RemoveSpaceFromString(raw.TransmissionMediumReq)
	ParsedUsername = helpers.RemoveSpaceFromString(raw.Username)
	ParsedVoiceFeature = helpers.RemoveSpaceFromString(raw.VoiceFeature)

	InvalidNTPReference := false

	TrimmedVadEnable := helpers.RemoveSpaceFromString(raw.VadEnable)
	if TrimmedVadEnable != nil {
		if *TrimmedVadEnable == "enable" {
			pTrue := true
			ParsedVadEnable = &pTrue
		} else if *TrimmedVadEnable == "disable" {
			pFalse := false
			ParsedVadEnable = &pFalse
		} else {
			logger.Error("Error parsing VadEnable: %s in %s", *TrimmedVadEnable, filename)
		}
	}

	ParsedTimeLocation, err := helpers.ExtractTimeLocationFromString(raw.H323SetupTime)
	if err == helpers.ErrInvalidNTPReferenceAsterisk || err == helpers.ErrInvalidNTPReferencePeriod {
		logger.Error("Error parsing Time Location From H323SetupTime: %s in %s", err, filename)
		InvalidNTPReference = true
	} else if err != nil {
		logger.Error("Error parsing Time Location From H323SetupTime: %s in %s", err, filename)
		InvalidNTPReference = true
	}

	ParsedAlertTime, err = helpers.ConvertStringToUnixTime(raw.AlertTime, ParsedTimeLocation)
	if err != nil {
		logger.Error("Error parsing AlertTime: %s in %s", err, filename)
	}

	ParsedAcomLevel, err = helpers.ConvertStringToInt64(raw.AcomLevel)
	if err != nil {
		logger.Error("Error parsing AcomLevel: %s in %s", err, filename)
	}

	ParsedBytesIn, err = helpers.ConvertStringToInt64(raw.BytesIn)
	if err != nil {
		logger.Error("Error parsing BytesIn: %s in %s", err, filename)
	}

	ParsedBytesOut, err = helpers.ConvertStringToInt64(raw.BytesOut)
	if err != nil {
		logger.Error("Error parsing BytesOut: %s in %s", err, filename)
	}

	ParsedCallId, err = helpers.ConvertStringToInt64(raw.CallId)
	if err != nil {
		logger.Error("Error parsing CallId: %s in %s", err, filename)
	}

	ParsedCdrType, err = helpers.ConvertStringToInt64(raw.CdrType)
	if err != nil {
		logger.Error("Error parsing CdrType: %s in %s", err, filename)
	}

	ParsedChargedUnits, err = helpers.ConvertStringToInt64(raw.ChargedUnits)
	if err != nil {
		logger.Error("Error parsing ChargedUnits: %s in %s", err, filename)
	}

	ParsedCodecBytes, err = helpers.ConvertStringToInt64(raw.CodecBytes)
	if err != nil {
		logger.Error("Error parsing CodecBytes: %s in %s", err, filename)
	}

	ParsedEarlyPackets, err = helpers.ConvertStringToInt64(raw.EarlyPackets)
	if err != nil {
		logger.Error("Error parsing EarlyPackets: %s in %s", err, filename)
	}

	ParsedFaxrelayInitHsMod, err = helpers.ConvertStringToInt64(raw.FaxrelayInitHsMod)
	if err != nil {
		logger.Error("Error parsing FaxrelayInitHsMod: %s in %s", err, filename)
	}

	ParsedFaxrelayJitBufOvflow, err = helpers.ConvertStringToInt64(raw.FaxrelayJitBufOvflow)
	if err != nil {
		logger.Error("Error parsing FaxrelayJitBufOvflow: %s in %s", err, filename)
	}

	ParsedFaxrelayMaxJitBufDepth, err = helpers.ConvertStringToInt64(raw.FaxrelayMaxJitBufDepth)
	if err != nil {
		logger.Error("Error parsing FaxrelayMaxJitBufDepth: %s in %s", err, filename)
	}

	ParsedFaxrelayMrHsMod, err = helpers.ConvertStringToInt64(raw.FaxrelayMrHsMod)
	if err != nil {
		logger.Error("Error parsing FaxrelayMrHsMod: %s in %s", err, filename)
	}

	ParsedFaxrelayNumPages, err = helpers.ConvertStringToInt64(raw.FaxrelayNumPages)
	if err != nil {
		logger.Error("Error parsing FaxrelayNumPages: %s in %s", err, filename)
	}

	ParsedFaxrelayPktConceal, err = helpers.ConvertStringToInt64(raw.FaxrelayPktConceal)
	if err != nil {
		logger.Error("Error parsing FaxrelayPktConceal: %s in %s", err, filename)
	}

	ParsedFaxrelayRxPackets, err = helpers.ConvertStringToInt64(raw.FaxrelayRxPackets)
	if err != nil {
		logger.Error("Error parsing FaxrelayRxPackets: %s in %s", err, filename)
	}

	ParsedFaxrelayTxPackets, err = helpers.ConvertStringToInt64(raw.FaxrelayTxPackets)
	if err != nil {
		logger.Error("Error parsing FaxrelayTxPackets: %s in %s", err, filename)
	}

	ParsedGapfillWithInterpolation, err = helpers.ConvertStringToInt64(raw.GapfillWithInterpolation)
	if err != nil {
		logger.Error("Error parsing GapfillWithInterpolation: %s in %s", err, filename)
	}

	ParsedGapfillWithPrediction, err = helpers.ConvertStringToInt64(raw.GapfillWithPrediction)
	if err != nil {
		logger.Error("Error parsing GapfillWithPrediction: %s in %s", err, filename)
	}

	ParsedGapfillWithRedundancy, err = helpers.ConvertStringToInt64(raw.GapfillWithRedundancy)
	if err != nil {
		logger.Error("Error parsing GapfillWithRedundancy: %s in %s", err, filename)
	}

	ParsedGapfillWithSilence, err = helpers.ConvertStringToInt64(raw.GapfillWithSilence)
	if err != nil {
		logger.Error("Error parsing GapfillWithSilence: %s in %s", err, filename)
	}

	ParsedH323VoiceQuality, err = helpers.ConvertStringToInt64(raw.H323VoiceQuality)
	if err != nil {
		logger.Error("Error parsing H323VoiceQuality: %s in %s", err, filename)
	}

	ParsedHiwaterPlayoutDelay, err = helpers.ConvertStringToInt64(raw.HiwaterPlayoutDelay)
	if err != nil {
		logger.Error("Error parsing HiwaterPlayoutDelay: %s in %s", err, filename)
	}

	ParsedIPHop, err = helpers.ConvertStringToInt64(raw.IPHop)
	if err != nil {
		logger.Error("Error parsing IPHop: %s in %s", err, filename)
	}

	ParsedLatePackets, err = helpers.ConvertStringToInt64(raw.LatePackets)
	if err != nil {
		logger.Error("Error parsing LatePackets: %s in %s", err, filename)
	}

	ConvertedLegType, err := helpers.ConvertStringToInt(raw.LegType)
	if err != nil {
		logger.Error("Error parsing LegType: %d in %s", *ConvertedLegType, filename)
		ParsedLegType = nil
	}

	if ConvertedLegType != nil {
		if helpers.ContainsInt(&CubeLegTypes, ConvertedLegType) {
			logger.Error("Error Invalid LegType: %d in %s", *ConvertedLegType, filename)
			ParsedLegType = nil
		} else {
			ParsedLegType = ConvertedLegType
		}
	} else {
		ParsedLegType = nil
	}

	ParsedLogicalIfIndex, err = helpers.ConvertStringToInt64(raw.LogicalIfIndex)
	if err != nil {
		logger.Error("Error parsing LogicalIfIndex: %s in %s", err, filename)
	}

	ParsedLostPackets, err = helpers.ConvertStringToInt64(raw.LostPackets)
	if err != nil {
		logger.Error("Error parsing LostPackets: %s in %s", err, filename)
	}

	ParsedLowaterPlayoutDelay, err = helpers.ConvertStringToInt64(raw.LowaterPlayoutDelay)
	if err != nil {
		logger.Error("Error parsing LowaterPlayoutDelay: %s in %s", err, filename)
	}

	ParsedNoiseLevel, err = helpers.ConvertStringToInt64(raw.NoiseLevel)
	if err != nil {
		logger.Error("Error parsing NoiseLevel: %s in %s", err, filename)
	}

	ParsedOntimeRvPlayout, err = helpers.ConvertStringToInt64(raw.OntimeRvPlayout)
	if err != nil {
		logger.Error("Error parsing OntimeRvPlayout: %s in %s", err, filename)
	}

	ParsedOverrideSessionTime, err = helpers.ConvertStringToInt64(raw.OverrideSessionTime)
	if err != nil {
		logger.Error("Error parsing OverrideSessionTime: %s in %s", err, filename)
	}

	ParsedPaksIn, err = helpers.ConvertStringToInt64(raw.PaksIn)
	if err != nil {
		logger.Error("Error parsing PaksIn: %s in %s", err, filename)
	}

	ParsedPaksOut, err = helpers.ConvertStringToInt64(raw.PaksOut)
	if err != nil {
		logger.Error("Error parsing PaksOut: %s in %s", err, filename)
	}

	ParsedPeerId, err = helpers.ConvertStringToInt64(raw.PeerId)
	if err != nil {
		logger.Error("Error parsing PeerId: %s in %s", err, filename)
	}

	ParsedPeerIfIndex, err = helpers.ConvertStringToInt64(raw.PeerIfIndex)
	if err != nil {
		logger.Error("Error parsing PeerIfIndex: %s in %s", err, filename)
	}

	ParsedReceiveDelay, err = helpers.ConvertStringToInt64(raw.ReceiveDelay)
	if err != nil {
		logger.Error("Error parsing ReceiveDelay: %s in %s", err, filename)
	}

	ParsedRecordTimestamp, err = helpers.ConvertStringToInt64(raw.RecordTimestamp)
	if err != nil {
		logger.Error("Error parsing RecordTimestamp: %s in %s", err, filename)
	}

	ParsedRemoteMediaUdpPort, err = helpers.ConvertStringToInt64(raw.RemoteMediaUdpPort)
	if err != nil {
		logger.Error("Error parsing RemoteMediaUdpPort: %s in %s", err, filename)
	}

	ParsedRoundTripDelay, err = helpers.ConvertStringToInt64(raw.RoundTripDelay)
	if err != nil {
		logger.Error("Error parsing RoundTripDelay: %s in %s", err, filename)
	}

	ParsedTxDuration, err = helpers.ConvertStringToInt64(raw.TxDuration)
	if err != nil {
		logger.Error("Error parsing TxDuration: %s in %s", err, filename)
	}

	ParsedVoiceTxDuration, err = helpers.ConvertStringToInt64(raw.VoiceTxDuration)
	if err != nil {
		logger.Error("Error parsing VoiceTxDuration: %s in %s", err, filename)
	}

	ParsedFeatureIdField2, err = helpers.ConvertStringToUnixTime(raw.FeatureIdField2, ParsedTimeLocation)
	if err != nil {
		logger.Error("Error parsing FeatureIdField2: %s in %s", err, filename)
	}

	ParsedFiletimestamp, err = helpers.ConvertStringToUnixTime(raw.FileTimestamp, ParsedTimeLocation)
	if err != nil {
		logger.Error("Error parsing FileTimestamp: %s in %s", err, filename)
	}
	ParsedH323ConnectTime, err = helpers.ConvertStringToUnixTime(raw.H323ConnectTime, ParsedTimeLocation)
	if err != nil {
		logger.Error("Error parsing H323ConnectTime: %s in %s", err, filename)
	}

	ParsedH323DisconnectTime, err = helpers.ConvertStringToUnixTime(raw.H323DisconnectTime, ParsedTimeLocation)
	if err != nil {
		logger.Error("Error parsing H323DisconnectTime: %s in %s", err, filename)
	}

	ParsedH323SetupTime, err = helpers.ConvertStringToUnixTime(raw.H323SetupTime, ParsedTimeLocation)
	if err != nil {
		logger.Error("Error parsing H323SetupTime: %s in %s", err, filename)
	}

	TrimmedSessionProtocol := helpers.RemoveSpaceFromString(raw.SessionProtocol)

	if TrimmedSessionProtocol != nil {
		if helpers.ContainsString(&CubeSessionProtocols, TrimmedSessionProtocol) {
			ParsedSessionProtocol = raw.SessionProtocol
		} else {
			ParsedSessionProtocol = nil
			logger.Error("Error parsing SessionProtocol: %s in %s", *TrimmedSessionProtocol, filename)
		}
	}

	TrimmedH323DisconnectCause := helpers.RemoveSpaceFromString(raw.H323DisconnectCause)

	if TrimmedH323DisconnectCause != nil {
		if !(helpers.ContainsString(&H323CauseCodes, TrimmedH323DisconnectCause)) {
			ParsedH323DisconnectCause = nil
			logger.Error("Error parsing H323DisconnectCause: %s in %s", *TrimmedH323DisconnectCause, filename)
		}
	}

	if helpers.ContainsString(&TwoWayCallTypes, raw.FeatureIdField1) {
		TWCCallingNumber := raw.FeatureIdField3
		ParsedTWCCallingNumber = TWCCallingNumber

		TWCCalledNumber := raw.FeatureIdField4
		ParsedTWCCalledNumber = TWCCalledNumber

		ParsedTWCFeatureStatus, err = helpers.ConvertStringToInt64(raw.FeatureIdField5)
		if err != nil {
			logger.Error("Error parsing TWCFeatureStatus: %s in %s", err, filename)
		}

		TWCFeatureCorrelationId := raw.FeatureIdField6
		ParsedTWCFeatureCorrelationId = TWCFeatureCorrelationId

		TWCFeatureID := raw.FeatureIdField7

		ParsedTWCFeatureID = TWCFeatureID

		ParsedTWCLegID = helpers.ConvertHexadecimalToInt(raw.FeatureIdField8)

	}

	if helpers.ContainsString(&CallForwardTypes, raw.FeatureIdField1) {
		CallForwardFeatureStatus := raw.FeatureIdField3
		ParsedCallForwardFeatureStatus = CallForwardFeatureStatus

		CallForwardFeatureID := raw.FeatureIdField4
		ParsedCallForwardFeatureID = CallForwardFeatureID

		CallForwardFeatureCorrelationId := raw.FeatureIdField5
		ParsedCallForwardFeatureCorrelationId = CallForwardFeatureCorrelationId

		ParsedCallForwardLegID, err = helpers.ConvertStringToInt64(raw.FeatureIdField6)
		if err != nil {
			logger.Error("Error parsing CallForwardLegID: %s in %s", err, filename)
		}

		CallForwardReason := raw.FeatureIdField5
		ParsedCallForwardReason = CallForwardReason

		CallForwardCount := raw.FeatureIdField5
		ParsedCallForwardCount = CallForwardCount

		CallForwardedFromNumber := raw.FeatureIdField9
		ParsedCallForwardedFromNumber = CallForwardedFromNumber

		CallForwardedNumber := raw.FeatureIdField10
		ParsedCallForwardedNumber = CallForwardedNumber

		CallForwardedToNumber := raw.FeatureIdField11
		ParsedCallForwardedToNumber = CallForwardedToNumber

		CallForwardingFromNumber := raw.FeatureIdField12
		ParsedCallForwardingFromNumber = CallForwardingFromNumber
	}

	if helpers.ContainsString(&TransferTypes, raw.FeatureIdField1) {
		TransferFeatureStatus := raw.FeatureIdField3
		ParsedTransferFeatureStatus = TransferFeatureStatus

		TransferFeatureID := raw.FeatureIdField4
		ParsedTransferFeatureID = TransferFeatureID

		TransferFeatureCorrelationId := raw.FeatureIdField5
		ParsedTransferFeatureCorrelationId = TransferFeatureCorrelationId

		ParsedTransferConsultationID, err = helpers.ConvertStringToInt64(raw.FeatureIdField6)
		if err != nil {
			logger.Error("Error parsing TransferConsultationID: %s in %s", err, filename)
		}

		TransferLegID := raw.FeatureIdField7
		ParsedTransferLegID = TransferLegID

		TransferForwardingReason := raw.FeatureIdField8
		ParsedTransferForwardingReason = TransferForwardingReason

		ParsedTransferStatus, err = helpers.ConvertStringToInt64(raw.FeatureIdField9)
		if err != nil {
			logger.Error("Error parsing TransferStatus: %s in %s", err, filename)
		}

		TransferredFromPart := raw.FeatureIdField10
		ParsedTransferredFromPart = TransferredFromPart

		TransferredNumber := raw.FeatureIdField11
		ParsedTransferredNumber = TransferredNumber

		TransferredToParty := raw.FeatureIdField12
		ParsedTransferredToParty = TransferredToParty
	}

	if helpers.ContainsString(&HoldTypes, raw.FeatureIdField1) {
		HoldStatus := raw.FeatureIdField3
		ParsedHoldStatus = HoldStatus

		HoldFeatureID := raw.FeatureIdField4
		ParsedHoldFeatureID = HoldFeatureID

		HoldFeatureCorrelationId := raw.FeatureIdField5
		ParsedHoldFeatureCorrelationId = HoldFeatureCorrelationId

		HoldLegID := raw.FeatureIdField6
		ParsedHoldLegID = HoldLegID

		HoldReason := raw.FeatureIdField7
		ParsedHoldReason = HoldReason

		ParsedHoldingDN, err = helpers.ConvertStringToInt64(raw.FeatureIdField8)
		if err != nil {
			logger.Error("Error parsing HoldingDN: %s in %s", err, filename)
		}

		ParsedHeldDN, err = helpers.ConvertStringToInt64(raw.FeatureIdField9)
		if err != nil {
			logger.Error("Error parsing HeldDN: %s in %s", err, filename)
		}

		ParsedHoldSharedLine, err = helpers.ConvertStringToInt64(raw.FeatureIdField10)
		if err != nil {
			logger.Error("Error parsing HoldSharedLine: %s in %s", err, filename)
		}

		HoldUsername := raw.FeatureIdField11
		ParsedHoldUsername = HoldUsername

		HoldPhoneTag := raw.FeatureIdField12
		ParsedHoldPhoneTag = HoldPhoneTag
	}

	return &CubeCDR{
		ID:                              uuid.New().String(),
		FileTimestamp:                   ParsedFiletimestamp,
		RecordTimestamp:                 ParsedRecordTimestamp,
		CallId:                          ParsedCallId,
		SessionProtocol:                 ParsedSessionProtocol,
		CdrType:                         ParsedCdrType,
		AcomLevel:                       ParsedAcomLevel,
		AlertTime:                       ParsedAlertTime,
		BytesIn:                         ParsedBytesIn,
		BytesOut:                        ParsedBytesOut,
		CallForwardCount:                ParsedCallForwardCount,
		CallForwardFeatureCorrelationId: ParsedCallForwardFeatureCorrelationId,
		CallForwardFeatureID:            ParsedCallForwardFeatureID,
		CallForwardFeatureStatus:        ParsedCallForwardFeatureStatus,
		CallForwardLegID:                ParsedCallForwardLegID,
		CallForwardReason:               ParsedCallForwardReason,
		CallForwardedFromNumber:         ParsedCallForwardedFromNumber,
		CallForwardedNumber:             ParsedCallForwardedNumber,
		CallForwardedToNumber:           ParsedCallForwardedToNumber,
		CallForwardingFromNumber:        ParsedCallForwardingFromNumber,
		ChargedUnits:                    ParsedChargedUnits,
		CodecBytes:                      ParsedCodecBytes,
		EarlyPackets:                    ParsedEarlyPackets,
		InvalidNTPReference:             InvalidNTPReference,
		FaxrelayInitHsMod:               ParsedFaxrelayInitHsMod,
		FaxrelayJitBufOvflow:            ParsedFaxrelayJitBufOvflow,
		FaxrelayMaxJitBufDepth:          ParsedFaxrelayMaxJitBufDepth,
		FaxrelayMrHsMod:                 ParsedFaxrelayMrHsMod,
		FaxrelayNumPages:                ParsedFaxrelayNumPages,
		FaxrelayPktConceal:              ParsedFaxrelayPktConceal,
		FaxrelayRxPackets:               ParsedFaxrelayRxPackets,
		FaxrelayTxPackets:               ParsedFaxrelayTxPackets,
		FeatureIdField2:                 ParsedFeatureIdField2,
		GapfillWithInterpolation:        ParsedGapfillWithInterpolation,
		GapfillWithPrediction:           ParsedGapfillWithPrediction,
		GapfillWithRedundancy:           ParsedGapfillWithRedundancy,
		GapfillWithSilence:              ParsedGapfillWithSilence,
		GkXlatedCdn:                     ParsedGkXlatedCdn,
		GkXlatedCgn:                     ParsedGkXlatedCgn,
		GwFinalXlatedCdn:                ParsedGwFinalXlatedCdn,
		GwFinalXlatedCgn:                ParsedGwFinalXlatedCgn,
		GwFinalXlatedRdn:                ParsedGwFinalXlatedRdn,
		GwRxdCdn:                        ParsedGwRxdCdn,
		GwRxdCgn:                        ParsedGwRxdCgn,
		GwRxdRdn:                        ParsedGwRxdRdn,
		H323ConnectTime:                 ParsedH323ConnectTime,
		H323DisconnectCause:             ParsedH323DisconnectCause,
		H323DisconnectTime:              ParsedH323DisconnectTime,
		H323SetupTime:                   ParsedH323SetupTime,
		H323VoiceQuality:                ParsedH323VoiceQuality,
		HeldDN:                          ParsedHeldDN,
		HiwaterPlayoutDelay:             ParsedHiwaterPlayoutDelay,
		HoldFeatureCorrelationId:        ParsedHoldFeatureCorrelationId,
		HoldFeatureID:                   ParsedHoldFeatureID,
		HoldLegID:                       ParsedHoldLegID,
		HoldPhoneTag:                    ParsedHoldPhoneTag,
		HoldReason:                      ParsedHoldReason,
		HoldSharedLine:                  ParsedHoldSharedLine,
		HoldStatus:                      ParsedHoldStatus,
		HoldUsername:                    ParsedHoldUsername,
		HoldingDN:                       ParsedHoldingDN,
		IpHop:                           ParsedIPHop,
		LatePackets:                     ParsedLatePackets,
		LegType:                         ParsedLegType,
		LogicalIfIndex:                  ParsedLogicalIfIndex,
		LostPackets:                     ParsedLostPackets,
		LowaterPlayoutDelay:             ParsedLowaterPlayoutDelay,
		NoiseLevel:                      ParsedNoiseLevel,
		OntimeRvPlayout:                 ParsedOntimeRvPlayout,
		OverrideSessionTime:             ParsedOverrideSessionTime,
		PaksIn:                          ParsedPaksIn,
		PaksOut:                         ParsedPaksOut,
		PeerId:                          ParsedPeerId,
		PeerIfIndex:                     ParsedPeerIfIndex,
		ReceiveDelay:                    ParsedReceiveDelay,
		RemoteMediaUdpPort:              ParsedRemoteMediaUdpPort,
		RemoteUdpPort:                   ParsedRemoteUdpPort,
		RoundTripDelay:                  ParsedRoundTripDelay,
		TWCCalledNumber:                 ParsedTWCCalledNumber,
		TWCCallingNumber:                ParsedTWCCallingNumber,
		TWCFeatureCorrelationId:         ParsedTWCFeatureCorrelationId,
		TWCFeatureID:                    ParsedTWCFeatureID,
		TWCFeatureStatus:                ParsedTWCFeatureStatus,
		TWCLegID:                        ParsedTWCLegID,
		TransferConsultationID:          ParsedTransferConsultationID,
		TransferFeatureCorrelationId:    ParsedTransferFeatureCorrelationId,
		TransferFeatureID:               ParsedTransferFeatureID,
		TransferFeatureStatus:           ParsedTransferFeatureStatus,
		TransferForwardingReason:        ParsedTransferForwardingReason,
		TransferLegID:                   ParsedTransferLegID,
		TransferStatus:                  ParsedTransferStatus,
		TransferredFromPart:             ParsedTransferredFromPart,
		TransferredNumber:               ParsedTransferredNumber,
		TransferredToParty:              ParsedTransferredToParty,
		TxDuration:                      ParsedTxDuration,
		VoiceTxDuration:                 ParsedVoiceTxDuration,
		Hostname:                        ParsedHostname,
		Filename:                        ParsedFilename,
		AccountCode:                     ParsedAccountCode,
		CallingPartyCategory:            ParsedCallingPartyCategory,
		CarrierId:                       ParsedCarrierId,
		ChargeNumber:                    ParsedChargeNumber,
		Clid:                            ParsedClid,
		CodecTypeRate:                   ParsedCodecTypeRate,
		CustBizGrpId:                    ParsedCustBizGrpId,
		DisconnectText:                  raw.DisconnectText,
		Dnis:                            ParsedDnis,
		DspId:                           ParsedDspId,
		FacDigit:                        ParsedFacDigit,
		FacStatus:                       ParsedFacStatus,
		FaxrelayDirection:               ParsedFaxrelayDirection,
		FaxrelayEcmStatus:               ParsedFaxrelayEcmStatus,
		FaxrelayEncapProtocol:           ParsedFaxrelayEncapProtocol,
		FaxrelayFaxSuccess:              ParsedFaxrelayFaxSuccess,
		FaxrelayNsfCountryCode:          ParsedFaxrelayNsfCountryCode,
		FaxrelayNsfManufCode:            ParsedFaxrelayNsfManufCode,
		FaxrelayStartTime:               ParsedFaxrelayStartTime,
		FaxrelayStopTime:                ParsedFaxrelayStopTime,
		FeatureId:                       ParsedFeatureId,
		FeatureIdField1:                 ParsedFeatureIdField1,
		FeatureOpStatus:                 ParsedFeatureOpStatus,
		FeatureOpTime:                   ParsedFeatureOpTime,
		FeatureOperation:                ParsedFeatureOperation,
		GtdGwRxdCnn:                     ParsedGtdGwRxdCnn,
		GtdGwRxdOcn:                     ParsedGtdGwRxdOcn,
		GtdOrigCic:                      ParsedGtdOrigCic,
		GtdTermCic:                      ParsedGtdTermCic,
		GwCollectedCdn:                  ParsedGwCollectedCdn,
		H323CallOrigin:                  ParsedH323CallOrigin,
		H323ConfId:                      raw.H323ConfId,
		H323IvrOut:                      ParsedH323IvrOut,
		InCarrierId:                     ParsedInCarrierId,
		InLpcorGroup:                    ParsedInLpcorGroup,
		InTrunkgroupLabel:               ParsedInTrunkgroupLabel,
		IncomingArea:                    ParsedIncomingArea,
		InfoType:                        ParsedInfoType,
		InternalErrorCode:               ParsedInternalErrorCode,
		IpPbxMode:                       ParsedIpPbxMode,
		IpPhoneInfo:                     ParsedIpPhoneInfo,
		MaxBitrate:                      ParsedMaxBitrate,
		OriginatingLineInfo:             ParsedOriginatingLineInfo,
		OutCarrierId:                    ParsedOutCarrierId,
		OutLpcorGroup:                   ParsedOutLpcorGroup,
		OutTrunkgroupLabel:              ParsedOutTrunkgroupLabel,
		OutgoingArea:                    ParsedOutgoingArea,
		PeerAddress:                     ParsedPeerAddress,
		PeerSubAddress:                  ParsedPeerSubAddress,
		RemoteMediaAddress:              ParsedRemoteMediaAddress,
		RemoteMediaId:                   ParsedRemoteMediaId,
		ServiceDescriptor:               ParsedServiceDescriptor,
		Subscriber:                      ParsedSubscriber,
		SuppSvcXferBy:                   ParsedSuppSvcXferBy,
		TransmissionMediumReq:           ParsedTransmissionMediumReq,
		Username:                        ParsedUsername,
		VadEnable:                       ParsedVadEnable,
		VoiceFeature:                    ParsedVoiceFeature,
	}, nil
}

type RawCubeCDR struct {
	RecordTimestamp          *string
	CallId                   *string
	CdrType                  *string
	LegType                  *string
	H323ConfId               *string
	PeerAddress              *string
	PeerSubAddress           *string
	H323SetupTime            *string
	AlertTime                *string
	H323ConnectTime          *string
	H323DisconnectTime       *string
	H323DisconnectCause      *string
	DisconnectText           *string
	H323CallOrigin           *string
	ChargedUnits             *string
	InfoType                 *string
	PaksOut                  *string
	BytesOut                 *string
	PaksIn                   *string
	BytesIn                  *string
	Username                 *string
	Clid                     *string
	Dnis                     *string
	GtdOrigCic               *string
	GtdTermCic               *string
	TxDuration               *string
	PeerId                   *string
	PeerIfIndex              *string
	LogicalIfIndex           *string
	AcomLevel                *string
	NoiseLevel               *string
	VoiceTxDuration          *string
	AccountCode              *string
	CodecBytes               *string
	CodecTypeRate            *string
	OntimeRvPlayout          *string
	RemoteUdpPort            *string
	RemoteMediaUdpPort       *string
	VadEnable                *string
	ReceiveDelay             *string
	RoundTripDelay           *string
	HiwaterPlayoutDelay      *string
	LowaterPlayoutDelay      *string
	GapfillWithInterpolation *string
	GapfillWithRedundancy    *string
	GapfillWithSilence       *string
	GapfillWithPrediction    *string
	EarlyPackets             *string
	LatePackets              *string
	LostPackets              *string
	MaxBitrate               *string
	FaxrelayStartTime        *string
	FaxrelayStopTime         *string
	FaxrelayMaxJitBufDepth   *string
	FaxrelayJitBufOvflow     *string
	FaxrelayInitHsMod        *string
	FaxrelayMrHsMod          *string
	FaxrelayNumPages         *string
	FaxrelayTxPackets        *string
	FaxrelayRxPackets        *string
	FaxrelayDirection        *string
	FaxrelayPktConceal       *string
	FaxrelayEcmStatus        *string
	FaxrelayEncapProtocol    *string
	FaxrelayNsfCountryCode   *string
	FaxrelayNsfManufCode     *string
	FaxrelayFaxSuccess       *string
	OverrideSessionTime      *string
	H323IvrOut               *string
	InternalErrorCode        *string
	H323VoiceQuality         *string
	RemoteMediaAddress       *string
	RemoteMediaId            *string
	CarrierId                *string
	CallingPartyCategory     *string
	OriginatingLineInfo      *string
	ChargeNumber             *string
	TransmissionMediumReq    *string
	ServiceDescriptor        *string
	OutgoingArea             *string
	IncomingArea             *string
	OutTrunkgroupLabel       *string
	OutCarrierId             *string
	DspId                    *string
	InTrunkgroupLabel        *string
	InCarrierId              *string
	CustBizGrpId             *string
	SuppSvcXferBy            *string
	VoiceFeature             *string
	FeatureOperation         *string
	FeatureOpStatus          *string
	FeatureOpTime            *string
	FeatureId                *string
	GwRxdCdn                 *string
	GwRxdCgn                 *string
	GtdGwRxdOcn              *string
	GtdGwRxdCnn              *string
	GwRxdRdn                 *string
	GwFinalXlatedCdn         *string
	GwFinalXlatedCgn         *string
	GwFinalXlatedRdn         *string
	GkXlatedCdn              *string
	GkXlatedCgn              *string
	GwCollectedCdn           *string
	IPHop                    *string
	RedirectedStation        *string
	Subscriber               *string
	InIntrfcDesc             *string
	OutIntrfcDesc            *string
	SessionProtocol          *string
	LocalHostname            *string
	BackwardCallId           *string
	FeatureIdField1          *string
	FeatureIdField2          *string
	FeatureIdField3          *string
	FeatureIdField4          *string
	FeatureIdField5          *string
	FeatureIdField6          *string
	FeatureIdField7          *string
	FeatureIdField8          *string
	FeatureIdField9          *string
	FeatureIdField10         *string
	FeatureIdField11         *string
	FeatureIdField12         *string
	IpPhoneInfo              *string
	IpPbxMode                *string
	InLpcorGroup             *string
	OutLpcorGroup            *string
	FacDigit                 *string
	FacStatus                *string
	Hostname                 *string
	Filename                 *string
	FileTimestamp            *string
}
