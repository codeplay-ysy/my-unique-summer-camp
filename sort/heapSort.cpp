#include <iostream>
#include <algorithm>
using namespace std;

// 层序遍历数组表示的堆，从start节点到end节点重新构造大顶堆， 
// 在图解里已经说明了为了节省空间，把最大的元素换到数组最后，
// 所有这个函数调用的时候start一直为0，end递减。表明end后面的元素已经排行顺序 
void max_heapify(int arr[], int start, int end) 
{
    // 堆顶开始，父节点是堆顶 
    int dad = start;
    int son = dad * 2 + 1;  //左子节点
    while (son <= end) 
	{ // 若子节点在范围内才进行后续操作 
        if (son + 1 <= end && arr[son] < arr[son + 1])  
            son++; // 两个子节点选择大的和父节点比较
        if (arr[dad] > arr[son]) //父节点>子节点，不用调整
            return;
        else { // 否则交换父子节点，再和孙节点比较 
            swap(arr[dad], arr[son]);
            dad = son;
            son = dad * 2 + 1;
        }
    }
}

void heap_sort(int arr[], int len) {
    // 初始化，i从最后一个父节点开始调整，
	//因为叶子节点已经是一个堆了，所以从len/2开始 ,直到根节点0，也就是最后下面一行的节点。
    for (int i = len / 2 - 1; i >= 0; i--)
        max_heapify(arr, i, len - 1);
    // 第一个元素已经排序好，放到最后面去，然后调整前面的i-1个元素 
    for (int i = len - 1; i > 0; i--) 
	{
        swap(arr[0], arr[i]);
        max_heapify(arr, 0, i - 1);
    }
}

int main() {
    int arr[] = { 88,51,32,47,55,12,25 };
    int len = (int) sizeof(arr) / sizeof(*arr);
    heap_sort(arr, len);
    for (int i = 0; i < len; i++)
        cout << arr[i] << ' ';//意思是输出arr[i]和空格
    cout << endl;
    return 0;
}
