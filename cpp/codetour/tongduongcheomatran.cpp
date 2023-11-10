#include <bits/stdc++.h>
using namespace std;
#define ll long long
ll N;

int main() {
    int sum=0;
	cin >>N;
    for (ll i = 0; i < N; i++) {
        sum = (sum + (N * i + i + 1) % 2027) % 2027;
    }
    cout << sum;
    return 0;
}




// int findMaxSum(vector<vector<int>>& N) {
//     int sum = 0;
//     for (int i = 0; i < N.size(); i++) {
//         sum += N[i][i];
//     }
//     return sum % 2027;
// }

// vector<vector<int>> generateMatrix(int size) {
//     vector<vector<int>> matrix(size, vector<int>(size, 0));
//     for (int i = 0; i < size; i++) {
//         for (int j = 0; j < size; j++) {
//             matrix[i][j] = size * i + j + 1;
//         }
//     }
//     return matrix;
// }

// int main() {
//     int size;
//     cin >> size;
//     vector<vector<int>> matrix = generateMatrix(size);
//     cout << findMaxSum(matrix) << endl;
//     return 0;
// }
