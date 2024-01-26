package demo_protobuf3_optional

import (
	"testing"

	"google.golang.org/protobuf/proto"
)

func TestOptionalNilId(t *testing.T) {
	noid := &WithOptional{
		Username: "me",
		Userid:   nil,
	}
	t.Log(noid.String())
	wiremessage, err := proto.Marshal(noid)
	if err != nil {
		t.Fatal("marshal failed: ", err)
	}

	unmarshaled := &WithOptional{}
	if err := proto.Unmarshal(wiremessage, unmarshaled); err != nil {
		t.Fatal("failed to unmarshal message: ", err)
	}
	t.Log("id: ", unmarshaled.Userid)
	if unmarshaled.Userid != nil {
		t.Fatalf("expected userid is nil, got %v", unmarshaled.Userid)
	}
}

func TestOptionalZeroId(t *testing.T) {
	id := int64(0)
	noid := &WithOptional{
		Username: "me",
		Userid:   &id,
	}
	t.Log(noid.String())
	wiremessage, err := proto.Marshal(noid)
	if err != nil {
		t.Fatal("marshal failed: ", err)
	}

	unmarshaled := &WithOptional{}
	if err := proto.Unmarshal(wiremessage, unmarshaled); err != nil {
		t.Fatal("failed to unmarshal message: ", err)
	}
	if unmarshaled.Userid == nil || *unmarshaled.Userid != 0 {
		t.Fatalf("expected userid is non nil 0, got %v", unmarshaled.Userid)
	}
	t.Log("id: ", unmarshaled.Userid)
}

func TestOptionalNonZeroId(t *testing.T) {
	id := int64(1)
	noid := &WithOptional{
		Username: "me",
		Userid:   &id,
	}
	t.Log(noid.String())
	wiremessage, err := proto.Marshal(noid)
	if err != nil {
		t.Fatal("marshal failed: ", err)
	}

	unmarshaled := &WithOptional{}
	if err := proto.Unmarshal(wiremessage, unmarshaled); err != nil {
		t.Fatal("failed to unmarshal message: ", err)
	}
	if unmarshaled.Userid == nil || *unmarshaled.Userid != 1 {
		t.Fatalf("expected userid is non nil 1, got %v", unmarshaled.Userid)
	}
	t.Log("id: ", unmarshaled.Userid)
}
