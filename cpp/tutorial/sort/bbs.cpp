#include <iostream>
#include <vector>

/*void bubbleSort(std::vector<int>& arr) {
    int n = arr.size();
    for (int i = 0; i < n-1; i++) {
        for (int j = 0; j < n-i-1; j++) {
            if (arr[j] > arr[j+1]) {
             std::swap(arr[j], arr[j+1]);
            }
        }
        
    }
}*/
void bubbleSort(std::vector<int>& arr) {
    int n = arr.size();
    for (int i = n-1; i>0;i--){
        for( int j =0; j<i; j++){
            if (arr[j]>arr[j+1]){
                std::swap(arr[j+1],arr[j]);
            }
        }
    }
}
int main() {
    std::vector<int> arr = {64, 34, 25, 12, 22, 11, 90};
    std::cout << "Danh sách trước khi sắp xếp: ";
    for(int i : arr) {
        std::cout << i << " ";
    }
    std::cout << std::endl;
    bubbleSort(arr);
    std::cout << "Danh sách sau khi sắp xếp: ";
    for(int i : arr) {
        std::cout << i << " ";
    }
    std::cout << std::endl;

    return 0;
}
