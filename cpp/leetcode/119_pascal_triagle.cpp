#include <iostream>
#include <vector>
using namespace std;
class Solution {
public:
    vector<int> getRow(int rowIndex) {
        vector<int> result = {1};
        for (int i = 1; i<=rowIndex;i++){
            result[i] =  result[i-1]*(rowIndex-i-1);
        }
        return result;
        
    }
};