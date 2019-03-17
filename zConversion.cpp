#include<iostream>
#include<string>
#include<vector>

using namespace std;
typedef std::string string;

string convert(string s, int n){
    n--;
    if(n == 0){
        return s;
    }
    string result;
    for(int i=0; i<= n; i++){
        int index = i;
        if(i == 0||i==n){
            // 输出单个
            while(index < s.size()){
                //result.push_back(s[index]);
                result+=s[index];
                index += 2*n;
            }
        }
        else{
            // 输出两个
            while(index < s.size()){
                //result.push_back(s[index]);
                result+=s[index];
                if(index < n){
                    index = index + 2*n - 2*i;
                }
                else{
                    if(index + 2*i < s.size()){
                        //result.push_back(s[index + 2*i]);
                        result+=s[index + 2*i];
                    }
                    index += 2*n;
                }
            }
        }
    }
    return result;
}

int main(){
    string input;
    int n;
    cin >> input;
    cin >> n;
    cout << convert(input,n) << endl;
}
