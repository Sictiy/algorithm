#include<iostream>
using namespace std;

void testDivision(){
    int b = 0 ;
    try{
        cout<< 100/b << endl;
    }catch(const char* msg){
        cerr << msg << endl;
    }
}

void testNullptr(){
    char* p = new char;
    *p = 10;
    cout << *p << endl;
    delete p;
    try{
        cout << *p << endl;
    }catch(const char* msg){
        cout << msg << endl;
    }   
}

int main(){
    cout << "testStart" << endl;
    testNullptr();
    testDivision();
}

