#include <algorithm>
#include <bits/stdc++.h>

using namespace std;




/*
 * Complete the 'findMedian' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

int findMedian(vector<int> arr) {
    sort(arr.begin(),arr.end());
    long long  n = arr.size();
    return arr[n/2];

}

int main()
{
     vector<int> arr;
     int n; cin >>n;
     for(int i=0; i<n;i++){
        cin >>arr[i];
     }
     int result = findMedian(arr);
     cout <<nre<<endl;
     return 0;

}