
#include <bits/stdc++.h>
#include <inttypes.h>

using namespace std;

int main(){
	int TC; cin >> TC;
	while(TC--){
		int n; cin>>n;
		int arr[n];
		int res =1000; 
		for(int i = 0; i<n; i++){
			cin >>arr[i];
		}
		// for(int i =0; i<n; i++){
		// 	for(int j=i+1; j<n; j++){
		// 		 int hieu = abs(arr[i]-arr[j]);
		// 		 if(hieu<res ){
		// 		 	res  = hieu;
		// 		 }
		// 	}
		// }
		// cout <<res ;
		sort(arr,arr+n);
		for(int i =1; i< n; i++){
			int  hieu = abs(arr[i]- arr[i-1]);
			if(hieu <res){
				res = hieu; 
			}
		}
		cout <<res;
	}
	return 0;
}	
