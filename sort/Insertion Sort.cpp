#include <iostream>
using namespace std;
//插入排序
void insertionSort(int* arr, int sz)
{
    for (int i = 1; i < sz; i++)
    {
        int key = arr[i];
        int j = 0;
        for (j = i - 1; j >= 0; j--)
        {
            if (arr[j] > key)
            {
                arr[j + 1] = arr[j];
            }
            else
            {
                break;
            }
        }
        arr[j + 1] = key;
    }
}
int main()
{
    int arr[] = { 5, 7, 3, 2, 1 };
    int sz = sizeof(arr) / sizeof(arr[0]);
    insertionSort(arr, sz);
    for (int i = 0; i < sz; i++)
    {
        cout << arr[i] << " ";//cout是输出，<<是输出运算符
    }
    cout << endl;//endl是换行符
    return 0;
}
