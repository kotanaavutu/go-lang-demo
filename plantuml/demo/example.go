package demo

//Myinterface only has one method, notice the signature retirn value
type DemoInterface interface {
	demo() bool
}

//MyInterface will implement the foo() bool function so it will have a 'extends' associatio
//with MyInterface
type DemoStruct1 struct{

}

func (s1 *DemoStruct1) demo() bool{
	return true;
}

//Mystruct will be directly composed of Mystruct so it will have a composition relationship with it
type DemoStruct2 struct{
	DemoStruct1
}

//MyStruct3 will have a foo() function but the return value is not a bool, so it will not 
//have any relationship with MyInterface
type DemoStruct3 struct{
	Foo DemoStruct1
}

func (s3 *DemoStruct3) demo() {

}

