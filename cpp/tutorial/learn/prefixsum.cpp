#include <algorithm>
#include <bits/stdc++.h>
#include <future>

using namespace std;
const int MAXN = 200003;
const long long INF = 1e18;

int n, a[MAXN];
long long prefSum[MAXN], prefMin[MAXN], ans = -INF;
//O(n)
// int main(){
// 	int n; cin>>n;
// 	int a[n];
// 	for(int &x :a) cin>>x;
// 	int q;cin>>q;
// 	while (q--){
// 		int l, r; cin>>l>>r;
// 		int sum;
// 		for(int i=l; i<=r; i++){
// 			sum+=a[i];
// 		}
// 		cout<<sum<<endl;
// 	}
// 	return 0;
// }


//prefix sum
int main(){
	// int n; cin>>n;
	// int a[n+1];
	// for(int i=1; i<=n; i++) cin >>a[i];


	// int prefix[n+1];
	// memset(prefix, 0, sizeof(prefix));

	// for(int i=1; i<=n; i++){
	// 	prefix[i]=prefix[i-1] +a[i];
	// }
	// int q;cin>>q;
	// while (q--){
	// 	int l, r; cin>>l>>r;
	// 	int sum;
	// 	for(int i=l; i<=r; i++){
	// 		cout<<prefix[r]- prefix[l-1]<<endl;			
	// 	}
	// }
	// return 0;
	cin >>n;
	for(int i=1; i<=n;i++) cin >> a[i];
	prefSum[0]=prefMin[0]=0;
	for(int i=1; i<=n;i++){
		prefSum[i]=prefSum[i-1]+a[i],
		prefMin[i]=min(prefMin[i-1], prefSum[i]);

	}
	for(int i = 1; i<=n; i++){
		ans=max(ans, prefSum[i]-prefMin[i-1]);
	}
	cout<<ans;
}
