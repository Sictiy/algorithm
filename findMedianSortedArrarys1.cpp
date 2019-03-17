#include<iostream>                                                                                        
#include<vector>                                                                                       
                                                                                                       
using namespace std;                                                                                
                                                                                                       
int findMin(vector<int>& ns1, vector<int>& ns2,  int n1, int n2){                                          
    if(n1 == 0||n1>ns1.size())                                                                         
        return ns2[n2-1];                                                                              
    if(n2 == 0||n2>ns2.size())                                                                         
        return ns1[n1-1];                                                                              
    return ns1[n1-1]>ns2[n2-1]?ns2[n2-1]:ns1[n1-1];                                                    
}                                                                                                      
                                                                                                       
int findMax(vector<int>& ns1, vector<int>& ns2,  int n1, int n2){                                          
    if(n1 == 0||n1>ns1.size())                                                                             
        return ns2[n2-1];                                                                             
    if(n2 == 0||n2>ns2.size())                                                                        
        return ns1[n1-1];                                                                             
    return ns1[n1-1]<ns2[n2-1]?ns2[n2-1]:ns1[n1-1];                                                   
}                                                                                                     
                                                                                                      
double findMedianSortedArrays(vector<int>& nums1, vector<int>& nums2){     
    if(nums1.size() > nums2.size()){
        vector<int> temp = nums1;
        nums1 = nums2;
        nums2 = temp;
    }
    int min = 0;
    int max = nums1.size() + nums2.size();
    if(max == 0){
        cout << "error!";
        return 0;
    }
    if(max == 1){
        return nums2[0];
    }
    max = nums1.size();
    int i = (min + max) / 2;                                                                                        
    int j = (nums1.size() + nums2.size()) / 2 - i;                                                      
    int rightMin = findMin(nums1, nums2, i+1, j+1);
    int leftMax = findMax(nums1,nums2,i,j);                                                    
    while(leftMax > rightMin){
       if(i == 0|| (j!=0 && nums1[i-1] < nums2[j-1])){
           // i 应该变大
           min = i + 1;
       }else{
           max = i - 1;
       }
        i = (min + max) / 2;                                                                                        
        j = (nums1.size() + nums2.size()) / 2 - i;                                                      
        rightMin = findMin(nums1, nums2, i+1, j+1);
        leftMax = findMax(nums1,nums2,i,j);    
    }
    if((nums1.size()+nums2.size())&1){
        return rightMin;
    }else{
        return (rightMin + leftMax)/2.0L;
    }
}

int main(){
    vector<int> lista {1,2,3,4,5};
    vector<int> listb {6,7,8,9,10};
    cout << "result:" << findMedianSortedArrays(lista, listb) << endl;
}
