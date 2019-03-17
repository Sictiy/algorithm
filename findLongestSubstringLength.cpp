#include<iostream>
typedef std::string string;
int findSameChar(string , int, int, char);
int lengthOfLongestSubstring(string s) {
    int maxLength;
    int begin = 0;
    for(int i= 1; i< s.length(); i++){
        begin = findSameChar(s, begin, i, s[i]);
        maxLength = maxLength > i - begin? maxLength: i-begin;
    }
    return maxLength + 1;
}

int findSameChar(string s, int begin,int end, char c){
    for(int i = begin; i< end; i++){
        if(s[i] == c){
            return i+1;
        }
    }
    return begin;
}

int main(){
    string input;
    std::cin >> input;
    std::cout << lengthOfLongestSubstring(input);
}