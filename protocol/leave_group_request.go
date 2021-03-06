package protocol

type LeaveGroupRequest struct {
	GroupID  string
	MemberID string
}

func (r *LeaveGroupRequest) Encode(e PacketEncoder) error {
	if err := e.PutString(r.GroupID); err != nil {
		return err
	}
	return e.PutString(r.MemberID)
}

func (r *LeaveGroupRequest) Decode(d PacketDecoder) (err error) {
	if r.GroupID, err = d.String(); err != nil {
		return err
	}
	r.MemberID, err = d.String()
	return err
}

func (r *LeaveGroupRequest) key() int16 {
	return 13
}

func (r *LeaveGroupRequest) version() int16 {
	return 0
}
