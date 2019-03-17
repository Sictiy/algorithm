#include<iostream>

using namespace std;

int reverse(int x){
    bool flag = x < 0;
    if(flag){
        x = -x;
    }
    int max = 0x7FFFFFFF / 10;
    int result(0);
    while(x >= 10){
        int lastNum = x % 10;
        x /= 10;
        result += lastNum;
        // 乘前判断溢出
        if(result > max){
        return 0;
        }
        result *= 10;
        cout << x << " " << lastNum << endl;
    }
    result += x;
    // 加后判断溢出
    if(result < 0){
        return 0;
    }
    return flag?-result:result;
}

int main(){
    int input;
    cin >> input;
    cout << reverse(input) << endl;
}

