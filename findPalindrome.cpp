#include<iostream>
#include<string>

using namespace std;
typedef std::string string;

string findPalind(string s, int i, bool isD){
    int j = 1;
    while(true){
        int min = i-j;
        int max;
        if(!isD){
            max = i+j;
        }else{
            max = i+j+1;
        }
        if(min < 0 || max > (s.length() - 1)||s[min] != s[max]){
            cout << min+1 << " " << max-min-1 << endl;
            cout << "min: "<< min<<"max: "<<max<< endl;
            return s.substr(min+1, max-min-1);
        }
        j++;
    }
}

string longestPalindrome(string s){
    string result = "";
    for(int i = 0; i < s.length(); i++){
        string  temp = findPalind(s, i, false);
        if(temp.length() > result.length()){
            result = temp;
        }
        if(i+1<s.length()&&s[i] == s[i+1]){
            temp = findPalind(s, i, true);
            if(temp.length() > result.length()){
                result = temp;
            }
        }
    }
    return result;
}

int main(){
    string input;
    cin >> input;
    cout << longestPalindrome(input) << endl;
}
