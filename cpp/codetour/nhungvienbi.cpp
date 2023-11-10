#include <iostream>
#include <vector>
using namespace std;

int main() {
    int t;
    cin >> t;

    for (int i = 0; i < t; i++) {
        int n;
        cin >> n;

        vector<int> arr(n);
        for (int j = 0; j < n; j++) {
            cin >> arr[j];
        }

        int total_length = 0;
        int ones = 0;
        int one_index = -1;

        for (int j = 0; j < n; j++) {
            if (arr[j] == 1) {
                ones++;
                if (one_index == -1) {
                    one_index = j;
                } else {
                    total_length += (j - one_index - 1);
                    one_index = j;
                }
            }
        }

        int result = ones > 1 ? total_length : 0;
        cout << result << endl;
    }

    return 0;
}
