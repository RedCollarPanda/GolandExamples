// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package main

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson89aae3efDecodeGolandExamplesEasyJson(in *jlexer.Lexer, out *SenderRecipientAck) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "messageId":
			out.MessageID = string(in.String())
		case "senderDeviceId":
			out.SenderDeviceID = string(in.String())
		case "recipientDeviceId":
			out.RecipientDeviceID = string(in.String())
		case "gmtTimestamp":
			out.GmtTimestamp = int(in.Int())
		case "type":
			out.Type = int(in.Int())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson89aae3efEncodeGolandExamplesEasyJson(out *jwriter.Writer, in SenderRecipientAck) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"messageId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.MessageID))
	}
	{
		const prefix string = ",\"senderDeviceId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.SenderDeviceID))
	}
	{
		const prefix string = ",\"recipientDeviceId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.RecipientDeviceID))
	}
	{
		const prefix string = ",\"gmtTimestamp\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.GmtTimestamp))
	}
	{
		const prefix string = ",\"type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Type))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SenderRecipientAck) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson89aae3efEncodeGolandExamplesEasyJson(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SenderRecipientAck) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson89aae3efEncodeGolandExamplesEasyJson(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SenderRecipientAck) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson89aae3efDecodeGolandExamplesEasyJson(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SenderRecipientAck) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson89aae3efDecodeGolandExamplesEasyJson(l, v)
}
func easyjson89aae3efDecodeGolandExamplesEasyJson1(in *jlexer.Lexer, out *SenderServerAck) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "messageId":
			out.MessageID = string(in.String())
		case "senderDeviceId":
			out.SenderDeviceID = string(in.String())
		case "recipientDeviceId":
			out.RecipientDeviceID = string(in.String())
		case "gmtTimestamp":
			out.GmtTimestamp = int(in.Int())
		case "type":
			out.Type = int(in.Int())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson89aae3efEncodeGolandExamplesEasyJson1(out *jwriter.Writer, in SenderServerAck) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"messageId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.MessageID))
	}
	{
		const prefix string = ",\"senderDeviceId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.SenderDeviceID))
	}
	{
		const prefix string = ",\"recipientDeviceId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.RecipientDeviceID))
	}
	{
		const prefix string = ",\"gmtTimestamp\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.GmtTimestamp))
	}
	{
		const prefix string = ",\"type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Type))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SenderServerAck) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson89aae3efEncodeGolandExamplesEasyJson1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SenderServerAck) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson89aae3efEncodeGolandExamplesEasyJson1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SenderServerAck) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson89aae3efDecodeGolandExamplesEasyJson1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SenderServerAck) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson89aae3efDecodeGolandExamplesEasyJson1(l, v)
}
func easyjson89aae3efDecodeGolandExamplesEasyJson2(in *jlexer.Lexer, out *ServerRecipientAck) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "messageId":
			out.MessageID = string(in.String())
		case "senderDeviceId":
			out.SenderDeviceID = string(in.String())
		case "recipientDeviceId":
			out.RecipientDeviceID = string(in.String())
		case "gmtTimestamp":
			out.GmtTimestamp = int(in.Int())
		case "type":
			out.Type = int(in.Int())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson89aae3efEncodeGolandExamplesEasyJson2(out *jwriter.Writer, in ServerRecipientAck) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"messageId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.MessageID))
	}
	{
		const prefix string = ",\"senderDeviceId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.SenderDeviceID))
	}
	{
		const prefix string = ",\"recipientDeviceId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.RecipientDeviceID))
	}
	{
		const prefix string = ",\"gmtTimestamp\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.GmtTimestamp))
	}
	{
		const prefix string = ",\"type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Type))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ServerRecipientAck) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson89aae3efEncodeGolandExamplesEasyJson2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ServerRecipientAck) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson89aae3efEncodeGolandExamplesEasyJson2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ServerRecipientAck) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson89aae3efDecodeGolandExamplesEasyJson2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ServerRecipientAck) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson89aae3efDecodeGolandExamplesEasyJson2(l, v)
}
func easyjson89aae3efDecodeGolandExamplesEasyJson3(in *jlexer.Lexer, out *SystemMessage) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "systemInfo":
			out.SystemInfo = string(in.String())
		case "recipient":
			out.Recipient = string(in.String())
		case "sender":
			out.Sender = string(in.String())
		case "messageId":
			out.MessageID = string(in.String())
		case "senderDeviceId":
			out.SenderDeviceID = string(in.String())
		case "recipientDeviceId":
			out.RecipientDeviceID = string(in.String())
		case "gmtTimestamp":
			out.GmtTimestamp = int(in.Int())
		case "type":
			out.Type = int(in.Int())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson89aae3efEncodeGolandExamplesEasyJson3(out *jwriter.Writer, in SystemMessage) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"systemInfo\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.SystemInfo))
	}
	{
		const prefix string = ",\"recipient\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Recipient))
	}
	{
		const prefix string = ",\"sender\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Sender))
	}
	{
		const prefix string = ",\"messageId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.MessageID))
	}
	{
		const prefix string = ",\"senderDeviceId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.SenderDeviceID))
	}
	{
		const prefix string = ",\"recipientDeviceId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.RecipientDeviceID))
	}
	{
		const prefix string = ",\"gmtTimestamp\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.GmtTimestamp))
	}
	{
		const prefix string = ",\"type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Type))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SystemMessage) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson89aae3efEncodeGolandExamplesEasyJson3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SystemMessage) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson89aae3efEncodeGolandExamplesEasyJson3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SystemMessage) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson89aae3efDecodeGolandExamplesEasyJson3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SystemMessage) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson89aae3efDecodeGolandExamplesEasyJson3(l, v)
}
func easyjson89aae3efDecodeGolandExamplesEasyJson4(in *jlexer.Lexer, out *GroupChatInvite) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "invitedUser":
			out.InvitedUser = string(in.String())
		case "group":
			out.Group = string(in.String())
		case "recipient":
			out.Recipient = string(in.String())
		case "sender":
			out.Sender = string(in.String())
		case "messageId":
			out.MessageID = string(in.String())
		case "senderDeviceId":
			out.SenderDeviceID = string(in.String())
		case "recipientDeviceId":
			out.RecipientDeviceID = string(in.String())
		case "gmtTimestamp":
			out.GmtTimestamp = int(in.Int())
		case "type":
			out.Type = int(in.Int())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson89aae3efEncodeGolandExamplesEasyJson4(out *jwriter.Writer, in GroupChatInvite) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"invitedUser\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.InvitedUser))
	}
	{
		const prefix string = ",\"group\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Group))
	}
	{
		const prefix string = ",\"recipient\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Recipient))
	}
	{
		const prefix string = ",\"sender\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Sender))
	}
	{
		const prefix string = ",\"messageId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.MessageID))
	}
	{
		const prefix string = ",\"senderDeviceId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.SenderDeviceID))
	}
	{
		const prefix string = ",\"recipientDeviceId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.RecipientDeviceID))
	}
	{
		const prefix string = ",\"gmtTimestamp\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.GmtTimestamp))
	}
	{
		const prefix string = ",\"type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Type))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GroupChatInvite) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson89aae3efEncodeGolandExamplesEasyJson4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GroupChatInvite) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson89aae3efEncodeGolandExamplesEasyJson4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GroupChatInvite) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson89aae3efDecodeGolandExamplesEasyJson4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GroupChatInvite) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson89aae3efDecodeGolandExamplesEasyJson4(l, v)
}
func easyjson89aae3efDecodeGolandExamplesEasyJson5(in *jlexer.Lexer, out *RecipientReadMessage) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "recipient":
			out.Recipient = string(in.String())
		case "sender":
			out.Sender = string(in.String())
		case "messageId":
			out.MessageID = string(in.String())
		case "senderDeviceId":
			out.SenderDeviceID = string(in.String())
		case "recipientDeviceId":
			out.RecipientDeviceID = string(in.String())
		case "gmtTimestamp":
			out.GmtTimestamp = int(in.Int())
		case "type":
			out.Type = int(in.Int())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson89aae3efEncodeGolandExamplesEasyJson5(out *jwriter.Writer, in RecipientReadMessage) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"recipient\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Recipient))
	}
	{
		const prefix string = ",\"sender\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Sender))
	}
	{
		const prefix string = ",\"messageId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.MessageID))
	}
	{
		const prefix string = ",\"senderDeviceId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.SenderDeviceID))
	}
	{
		const prefix string = ",\"recipientDeviceId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.RecipientDeviceID))
	}
	{
		const prefix string = ",\"gmtTimestamp\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.GmtTimestamp))
	}
	{
		const prefix string = ",\"type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Type))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RecipientReadMessage) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson89aae3efEncodeGolandExamplesEasyJson5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RecipientReadMessage) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson89aae3efEncodeGolandExamplesEasyJson5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RecipientReadMessage) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson89aae3efDecodeGolandExamplesEasyJson5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RecipientReadMessage) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson89aae3efDecodeGolandExamplesEasyJson5(l, v)
}
func easyjson89aae3efDecodeGolandExamplesEasyJson6(in *jlexer.Lexer, out *TextMessage) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "message":
			out.Message = string(in.String())
		case "recipient":
			out.Recipient = string(in.String())
		case "sender":
			out.Sender = string(in.String())
		case "messageId":
			out.MessageID = string(in.String())
		case "senderDeviceId":
			out.SenderDeviceID = string(in.String())
		case "recipientDeviceId":
			out.RecipientDeviceID = string(in.String())
		case "gmtTimestamp":
			out.GmtTimestamp = int(in.Int())
		case "type":
			out.Type = int(in.Int())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson89aae3efEncodeGolandExamplesEasyJson6(out *jwriter.Writer, in TextMessage) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"message\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Message))
	}
	{
		const prefix string = ",\"recipient\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Recipient))
	}
	{
		const prefix string = ",\"sender\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Sender))
	}
	{
		const prefix string = ",\"messageId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.MessageID))
	}
	{
		const prefix string = ",\"senderDeviceId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.SenderDeviceID))
	}
	{
		const prefix string = ",\"recipientDeviceId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.RecipientDeviceID))
	}
	{
		const prefix string = ",\"gmtTimestamp\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.GmtTimestamp))
	}
	{
		const prefix string = ",\"type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Type))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v TextMessage) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson89aae3efEncodeGolandExamplesEasyJson6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v TextMessage) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson89aae3efEncodeGolandExamplesEasyJson6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *TextMessage) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson89aae3efDecodeGolandExamplesEasyJson6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *TextMessage) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson89aae3efDecodeGolandExamplesEasyJson6(l, v)
}
func easyjson89aae3efDecodeGolandExamplesEasyJson7(in *jlexer.Lexer, out *BaseMessage) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "message":
			out.Message = string(in.String())
		case "recipient":
			out.Recipient = string(in.String())
		case "sender":
			out.Sender = string(in.String())
		case "messageId":
			out.MessageID = string(in.String())
		case "senderDeviceId":
			out.SenderDeviceID = string(in.String())
		case "recipientDeviceId":
			out.RecipientDeviceID = string(in.String())
		case "gmtTimestamp":
			out.GmtTimestamp = int(in.Int())
		case "type":
			out.Type = int(in.Int())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson89aae3efEncodeGolandExamplesEasyJson7(out *jwriter.Writer, in BaseMessage) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"message\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Message))
	}
	{
		const prefix string = ",\"recipient\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Recipient))
	}
	{
		const prefix string = ",\"sender\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Sender))
	}
	{
		const prefix string = ",\"messageId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.MessageID))
	}
	{
		const prefix string = ",\"senderDeviceId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.SenderDeviceID))
	}
	{
		const prefix string = ",\"recipientDeviceId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.RecipientDeviceID))
	}
	{
		const prefix string = ",\"gmtTimestamp\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.GmtTimestamp))
	}
	{
		const prefix string = ",\"type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Type))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v BaseMessage) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson89aae3efEncodeGolandExamplesEasyJson7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v BaseMessage) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson89aae3efEncodeGolandExamplesEasyJson7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *BaseMessage) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson89aae3efDecodeGolandExamplesEasyJson7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *BaseMessage) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson89aae3efDecodeGolandExamplesEasyJson7(l, v)
}