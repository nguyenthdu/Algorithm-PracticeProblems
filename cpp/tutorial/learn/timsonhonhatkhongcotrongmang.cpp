#include <bits/stdc++.h>

using namespace std;
int arr[1000002];
int main(){
	int TC; cin >> TC;
	while(TC--){
		int n; cin >>n;
		memset(arr, 0, sizeof(arr));//gan tat ca phan tu trong mang bang 0
		for(int i=0; i<n; i++){
			int x; cin >>x;
			if(x>0) arr[x]=1;
		}
		for(int i = 1; i<1000001; i++){
			if(arr[i]==0){
				cout <<i<<endl;
				break;
			}
		}
		cout <<endl;
	}
	return 0;
}	
