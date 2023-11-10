#include <bits/stdc++.h>

using namespace std;
int arr[1000002];
int main(){
	int TC; cin >> TC;
	while(TC--){
		int n; cin >>n;
		int a[n];
		memset(arr,0,sizeof(arr));
		//neu gia tri co am hoac qua lon thi nen dung map
		//map<int,int>mp;
		for(int i=0; i<n; i++){
			cin >>a[i];
			arr[a[i]]++;//luu so lan lap lai cua cac phan tu trong mang
			//mp[a[i]]++;
		}
		int ans =0;
		for(int i =0; i<n; i++){
			if (arr[a[i]] >=2) ++ans;
		}	//if(mp[a[i]]>=2)
		cout << ans<<endl;

	}
	return 0;
}	
