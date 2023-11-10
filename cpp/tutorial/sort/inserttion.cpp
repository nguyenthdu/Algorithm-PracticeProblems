#include <iostream>
#include <vector>
void Insertion(std::vector<int>& arr){
	for(int i = 1; i<arr.size(); ++i){
		int key  = arr[i];
		int j=i-1;
		while(j>=0 && key< arr[j]){
			arr[j+1]=arr[j];
			j--;
		}
		arr[j+1]=key;

	}
}
int main() {
    std::vector<int> arr = {64, 34, 25, 12, 22, 11, 90};
    std::cout << "Danh sách trước khi sắp xếp: ";
    for(int i : arr) {
        std::cout << i << " ";
    }
    std::cout << std::endl;
    Insertion(arr);
    std::cout << "Danh sách sau khi sắp xếp: ";
    for(int i : arr) {
        std::cout << i << " ";
    }
    std::cout << std::endl;

    return 0;
}
