// Code generated by "stringer -type ProtocolViolationErrorLocation -trimprefix ProtocolViolationErrorLocation"; DO NOT EDIT.

package socketcan

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ProtocolViolationErrorLocationUnspecified-0]
	_ = x[ProtocolViolationErrorLocationStartOfFrame-3]
	_ = x[ProtocolViolationErrorLocationID28To21-2]
	_ = x[ProtocolViolationErrorLocationID20To18-6]
	_ = x[ProtocolViolationErrorLocationSubstituteRTR-4]
	_ = x[ProtocolViolationErrorLocationIDExtension-5]
	_ = x[ProtocolViolationErrorLocationIDBits17To13-7]
	_ = x[ProtocolViolationErrorLocationIDBits12To05-15]
	_ = x[ProtocolViolationErrorLocationIDBits04To00-14]
	_ = x[ProtocolViolationErrorLocationRTR-12]
	_ = x[ProtocolViolationErrorLocationReservedBit1-13]
	_ = x[ProtocolViolationErrorLocationReservedBit0-9]
	_ = x[ProtocolViolationErrorLocationDataLengthCode-11]
	_ = x[ProtocolViolationErrorLocationData-10]
	_ = x[ProtocolViolationErrorLocationCRCSequence-8]
	_ = x[ProtocolViolationErrorLocationCRCDelimiter-24]
	_ = x[ProtocolViolationErrorLocationACKSlot-25]
	_ = x[ProtocolViolationErrorLocationACKDelimiter-27]
	_ = x[ProtocolViolationErrorLocationEndOfFrame-26]
	_ = x[ProtocolViolationErrorLocationIntermission-18]
}

const (
	_ProtocolViolationErrorLocation_name_0 = "Unspecified"
	_ProtocolViolationErrorLocation_name_1 = "ID28To21StartOfFrameSubstituteRTRIDExtensionID20To18IDBits17To13CRCSequenceReservedBit0DataDataLengthCodeRTRReservedBit1IDBits04To00IDBits12To05"
	_ProtocolViolationErrorLocation_name_2 = "Intermission"
	_ProtocolViolationErrorLocation_name_3 = "CRCDelimiterACKSlotEndOfFrameACKDelimiter"
)

var (
	_ProtocolViolationErrorLocation_index_1 = [...]uint8{0, 8, 20, 33, 44, 52, 64, 75, 87, 91, 105, 108, 120, 132, 144}
	_ProtocolViolationErrorLocation_index_3 = [...]uint8{0, 12, 19, 29, 41}
)

func (i ProtocolViolationErrorLocation) String() string {
	switch {
	case i == 0:
		return _ProtocolViolationErrorLocation_name_0
	case 2 <= i && i <= 15:
		i -= 2
		return _ProtocolViolationErrorLocation_name_1[_ProtocolViolationErrorLocation_index_1[i]:_ProtocolViolationErrorLocation_index_1[i+1]]
	case i == 18:
		return _ProtocolViolationErrorLocation_name_2
	case 24 <= i && i <= 27:
		i -= 24
		return _ProtocolViolationErrorLocation_name_3[_ProtocolViolationErrorLocation_index_3[i]:_ProtocolViolationErrorLocation_index_3[i+1]]
	default:
		return "ProtocolViolationErrorLocation(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
