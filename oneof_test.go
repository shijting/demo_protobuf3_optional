package demo_protobuf3_optional

import (
	"testing"

	"google.golang.org/protobuf/proto"
)

func TestOneofNilId(t *testing.T) {
	noid := &WithOneof{
		Username:       "me",
		OptionalUserid: nil,
	}
	t.Log(noid.String())
	wiremessage, err := proto.Marshal(noid)
	if err != nil {
		t.Fatal("marshal failed: ", err)
	}

	unmarshaled := &WithOneof{}
	if err := proto.Unmarshal(wiremessage, unmarshaled); err != nil {
		t.Fatal("failed to unmarshal message: ", err)
	}
	t.Log("id: ", unmarshaled.OptionalUserid)
	if unmarshaled.OptionalUserid != nil {
		t.Fatalf("expected userid is nil, got %v", unmarshaled.OptionalUserid)
	}
}

func TestOneofZeroId(t *testing.T) {
	id := &WithOneof_Userid{
		Userid: 0,
	}
	noid := &WithOneof{
		Username:       "me",
		OptionalUserid: id,
	}
	t.Log(noid.String())
	wiremessage, err := proto.Marshal(noid)
	if err != nil {
		t.Fatal("marshal failed: ", err)
	}

	unmarshaled := &WithOneof{}
	if err := proto.Unmarshal(wiremessage, unmarshaled); err != nil {
		t.Fatal("failed to unmarshal message: ", err)
	}

	if unmarshaled.OptionalUserid == nil {
		t.Fatalf("expected userid is not nil, got %v", unmarshaled.OptionalUserid)
	} else if userid, ok := unmarshaled.OptionalUserid.(*WithOneof_Userid); !ok {
		t.Fatalf("expected userid is instance of type %T , got %T", &WithOneof_Userid{}, unmarshaled.OptionalUserid)
	} else if userid.Userid != 0 {
		t.Errorf("expected userid is %v, got %v", id, userid.Userid)
	} else {
		t.Log("id: ", userid.Userid)
	}
}

func TestOneofNonZeroId(t *testing.T) {
	id := &WithOneof_Userid{
		Userid: 1,
	}
	noid := &WithOneof{
		Username:       "me",
		OptionalUserid: id,
	}
	t.Log(noid.String())
	wiremessage, err := proto.Marshal(noid)
	if err != nil {
		t.Fatal("marshal failed: ", err)
	}

	unmarshaled := &WithOneof{}
	if err := proto.Unmarshal(wiremessage, unmarshaled); err != nil {
		t.Fatal("failed to unmarshal message: ", err)
	}

	if unmarshaled.OptionalUserid == nil {
		t.Fatalf("expected userid is not nil, got %v", unmarshaled.OptionalUserid)
	} else if userid, ok := unmarshaled.OptionalUserid.(*WithOneof_Userid); !ok {
		t.Fatalf("expected userid is instance of type %T , got %T", &WithOneof_Userid{}, unmarshaled.OptionalUserid)
	} else if userid.Userid != 1 {
		t.Errorf("expected userid is %v, got %v", id, userid.Userid)
	} else {
		t.Log("id: ", userid.Userid)
	}
}
