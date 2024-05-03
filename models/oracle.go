package models

type OracleCDR struct {
	Accountingstatus *string
	Nasipaddress *string
	Nasport *string
	Accountingsessionid *string
	Ingresssessionid *string
	Egresssessionid *string
	Sessionprotocoltype *string
	Callingstationid *string
	Calledstationid *string
	Accountingterminationcause *string
	Accountingsessiontime *string
	Ciscosetuptime *string
	Ciscoconnecttime *string
	Ciscodisconnecttime *string
	Ciscodisconnectcause *string
	Egressnetworkinterfaceid *string
	Egressvlantagvalue *string
	Ingressnetworkinterfaceid *string
	Ingressvlantagvalue *string
	Egressrealm *string
	Ingressrealm *string
	Flowidentifier *string
	Flowtype *string
	Flowinputrealm *string
	Flowinputsrcaddr *string
	Flowinputsrcport *string
	Flowinputdestaddress *string
	Flowinputdestport *string
	Flowoutputrealm *string
	Flowoutputsrcaddress *string
	Flowoutputsrcport *string
	Flowoutputdestaddr *string
	Flowoutputdestport *string
	Rtcpcallingpacketslost *string
	Rtcpcallingavgjitter *string
	Rtcpcallingavglatency *string
	Rtcpcallingmaxjitter *string
	Rtcpcallingmaxlatency *string
	Rtpcallingpacketslost *string
	Rtpcallingavgjitter *string
	Rtpcallingmaxjitter *string
	Rtpcallingoctets *string
	Rtpcallingpackets *string
	Callingrfactor *string
	Callingmos *string
	Flowidentifier2 *string
	Flowtype2 *string
	Flowinputrealm2 *string
	Flowinputsrcaddr2 *string
	Flowinputsrcport2 *string
	Flowinputdestaddress2 *string
	Flowinputdestport2 *string
	Flowoutputrealm2 *string
	Flowoutputsrcaddress2 *string
	Flowoutputsrcport2 *string
	Flowoutputdestaddr2 *string
	Flowoutputdestport2 *string
	Rtcpcalledpacketslost *string
	Rtcpcalledavgjitter *string
	Rtcpcalledavglatency *string
	Rtcpcalledmaxjitter *string
	Rtcpcalledmaxlatency *string
	Rtpcalledpacketslost *string
	Rtpcalledavgjitter *string
	Rtpcalledmaxjitter *string
	Rtpcalledoctets *string
	Rtpcalledpackets *string
	Calledrfactor *string
	Calledmos *string
	Firmwareversion *string
	Localtimezone *string
	Postdialdelay *string
	Primaryroutingnumber *string
	Ingresslocaladdress *string
	Ingressremoteaddress *string
	Egresslocaladdress *string
	Egressremoteaddress *string
	Sessiondisposition *string
	Disconnectinitiator *string
	Disconnectcause *string
	Sipstatuscode *string
	Egressroutingnumber *string
	Callingmediastoptime *string
	Calledmediastoptime *string
	Flowmediatype *string
	Flowmediatype2 *string
	Rtpcallingoctetstransmitted *string
	Rtpcallingpacketstransmitted *string
	Rtpcalledoctetstransmitted *string
	Rtpcalledpacketstransmitted *string
	Msrpcalledoctets *string
	Msrpcalledpackets *string
	Msrpcalledoctetstransmitted *string
	Msrpcalledpacketstransmitted *string
	Msrpcallingoctets *string
	Msrpcallingpackets *string
	Msrpcallingoctetstransmitted *string
	Msrpcallingpacketstransmitted *string
	Nodefunctionality *string
	Cdrsequencenumber *string
}
