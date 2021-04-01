package creational_patterns

import "testing"

func TestPrototype(t *testing.T){
	shirtCloner:=GetShirtCloner()
	item1,err:=shirtCloner.GetClone(White)
	if err!=nil{
		t.Fatalf(err.Error())
	}
	shirt1,ok:=item1.(*Shirt)
	if !ok{
		t.Fatalf("error clone of white shirt")
	}
	if shirt1==whitePrototype{
		t.Fatalf("colne can't be same as prototype")
	}
}