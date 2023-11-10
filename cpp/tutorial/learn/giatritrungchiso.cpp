#include<bits/stdc++.h>
#include <cmath>
#include <iomanip>


using namespace std;

int main(){
	int TC; cin>>TC;
	while(TC--){
		int n; cin>>n;
		map<long long, bool>mp;
		for(int i=0; i<n; i++){
			long long x; cin >>x;
			mp[x] = true;// add i into mp
		}
		for(int i=0; i<n; i++){
			if(mp[i]){
				cout <<i<<" ";//if i into mp -> print i
			}
			else cout <<"-1 ";
		}
		cout <<endl;
	}
	return 0;
}