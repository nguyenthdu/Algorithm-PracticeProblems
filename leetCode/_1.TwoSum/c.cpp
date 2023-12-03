#include <bits/stdc++.h>

using namespace std;
//using Brute Force
vector<int> twoSumBruteForce(vector<int>&nums, int target) {
	vector<int> arr = {};
	int n = nums.size();
	for (int i = 0; i < n; i++) {
		for (int j = 0; j < n; j++) {
			if (nums[i] + nums[j] == target) {
				arr.push_back(i);
				arr.push_back(j);
				break;
			}
		}
	}
	return arr;
}
// vector<int> twoSum(vector<int>& nums, int target) {


// }

int main() {
	vector<int> arr = {2, 7, 11, 15};
	int target = 9;
	vector<int> result;
	result = twoSumBruteForce(arr, target);
	for (int x : result) {
		cout << x;
	}



}