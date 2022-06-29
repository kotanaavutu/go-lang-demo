package main

//Myinterface only has one method, notice the signature retirn value
type Test interface {
	fooTest() bool
}

//MyInterface will implement the foo() bool function so it will have a 'extends' associatio
//with MyInterface
type MyStructTest1 struct{

}

func (s1 *MyStructTest1) fooTest() bool{
	return true;
}

//Mystruct will be directly composed of Mystruct so it will have a composition relationship with it
type MyStructTest2 struct{
	MyStructTest1
}

//MyStruct3 will have a foo() function but the return value is not a bool, so it will not 
//have any relationship with MyInterface
type MyStructTest3 struct{
	Foo MyStructTest1
}

func (s3 *MyStructTest3) fooTest() {

}

