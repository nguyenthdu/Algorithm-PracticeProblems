#include <iostream>
#include <vector>
class Solution {
public:
    int findNumbers(std::vector<int>& nums) {
        int result = 0;
        for( int i=0; i<nums.size();i++){
            int n =  countN(nums[i]);
            if (n%2==0){
                result++;
            }
        }
        return result;
        
    }
    int countN(int num){
        int count = 0;
        while(num >0){
            num /= 10;
            count++;
        }
        return count;
    }
    int findNumbers2(std::vector<int>& nums){
        int result =0;
        for (int num :nums){
            if (num>9&&num<100||num>999 && num<10000 || num==100000){
                result++;
            }
        }
        return result;
    }

};