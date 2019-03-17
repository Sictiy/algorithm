#include<iostream>
#include<string>

typedef std::string string;
using namespace std;

int myAtoi(string str) {
    int max = 0x7FFFFFFF/10;
    bool flag = true;
    int result = 0;
    bool begin = false;
    for(int i = 0; i< str.length(); i++) {
        // 前置空格
        if(str[i] == ' ' && !begin){
            continue;
        }
        if(str[i] == '+' && !begin){
            begin = true;;
            continue;
        }
        // 计算数字
        if('0'<=str[i]&&str[i]<='9'){
            begin = true;
            if(result > max){
                result = 0x7FFFFFFF;
                cout << "max > max"<< endl;
                break;
            }
            result =result*10 + (str[i] - '0');
            continue;
        }
        // 判断第一个'-'号
        if(str[i] == (char)'-' && !begin){
            begin = true;
            flag = false;
            continue;
        }
        if(!begin||(begin&&result ==0)){
            return 0;
        }else{
            break;
        }
    }
    if(result <= 0&&flag){
        cout << "max < 0 "<< endl;
        result = 0x7FFFFFFF;
    }
    if(result == 0x7FFFFFFF){
         cout << "max = max"<< endl;
        return flag?result:-1-result;
    }
    return flag?result:0-result;
}

int main(){
    string input;
    cin >> input;
    cout << myAtoi(input) << endl;
}
