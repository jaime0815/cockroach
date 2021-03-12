// Code generated by "stringer -type=CantCloseReason"; DO NOT EDIT.

package sidetransport

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ReasonUnknown-0]
	_ = x[ReplicaDestroyed-1]
	_ = x[InvalidLease-2]
	_ = x[TargetOverLeaseExpiration-3]
	_ = x[MergeInProgress-4]
	_ = x[ProposalsInFlight-5]
	_ = x[RequestsEvaluatingBelowTarget-6]
	_ = x[MaxReason-7]
}

const _CantCloseReason_name = "ReasonUnknownReplicaDestroyedInvalidLeaseTargetOverLeaseExpirationMergeInProgressProposalsInFlightRequestsEvaluatingBelowTargetMaxReason"

var _CantCloseReason_index = [...]uint8{0, 13, 29, 41, 66, 81, 98, 127, 136}

func (i CantCloseReason) String() string {
	if i < 0 || i >= CantCloseReason(len(_CantCloseReason_index)-1) {
		return "CantCloseReason(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _CantCloseReason_name[_CantCloseReason_index[i]:_CantCloseReason_index[i+1]]
}
