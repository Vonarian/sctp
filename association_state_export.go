package sctp

type SerializedState struct {
	MyVerificationTag       uint32
	PeerVerificationTag     uint32
	State                   uint32
	InitialTSN              uint32
	MyNextTSN               uint32
	CumulativeTSNAckPoint   uint32
	AdvancedPeerTSNAckPoint uint32
	MaxPayloadSize          uint32
	Cwnd                    uint32
	Rwnd                    uint32
	Ssthresh                uint32
}

func (a *Association) ExportState() SerializedState {
	a.lock.RLock()
	defer a.lock.RUnlock()

	return SerializedState{
		MyVerificationTag:       a.myVerificationTag,
		PeerVerificationTag:     a.peerVerificationTag,
		State:                   a.state,
		InitialTSN:              a.initialTSN,
		MyNextTSN:               a.myNextTSN,
		CumulativeTSNAckPoint:   a.cumulativeTSNAckPoint,
		AdvancedPeerTSNAckPoint: a.advancedPeerTSNAckPoint,
		MaxPayloadSize:          a.maxPayloadSize,
		Cwnd:                    a.cwnd,
		Rwnd:                    a.rwnd,
		Ssthresh:                a.ssthresh,
	}
}

func (a *Association) ImportState(s SerializedState) {
	a.lock.Lock()
	defer a.lock.Unlock()

	a.myVerificationTag = s.MyVerificationTag
	a.peerVerificationTag = s.PeerVerificationTag
	a.state = s.State
	a.initialTSN = s.InitialTSN
	a.myNextTSN = s.MyNextTSN
	a.cumulativeTSNAckPoint = s.CumulativeTSNAckPoint
	a.advancedPeerTSNAckPoint = s.AdvancedPeerTSNAckPoint
	a.maxPayloadSize = s.MaxPayloadSize
	a.cwnd = s.Cwnd
	a.rwnd = s.Rwnd
	a.ssthresh = s.Ssthresh
}
