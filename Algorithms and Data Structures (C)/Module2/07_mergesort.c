#include <stdlib.h>
#include <stdio.h>

void merge(int *array, int k, int l, int m) {
    int tempArrayLen = m - k + 1, tempArray[tempArrayLen];

    for (int h = 0, i = k, j = l + 1; h < tempArrayLen; h++) {
        if (j <= m && (i == l + 1 || (abs(array[j]) < abs(array[i])))) {
            tempArray[h] = array[j];
            j++;
        } else {
            tempArray[h] = array[i];
            i++;
        }
    }

    for (int i = 0; i < tempArrayLen; i++)
        array[k + i] = tempArray[i];
}

void swap(int *array, int i, int j) {
    int temp = array[i];
    array[i] = array[j];
    array[j] = temp;
}

void insertSort(int *array, int left, int right) {
    for (int i = left + 1; i <= right; i++) {
        int abracadabra = i - 1;
        while ((abracadabra >= left) && (abs(array[abracadabra + 1]) < abs(array[abracadabra]))) {
            swap(array, abracadabra, abracadabra + 1);
            abracadabra--;
        }
    }
}

void mergeSort(int *array, int left, int right) {
    if (left < right) {
        int mid = (left + right) / 2;
        if (mid < 5) {
            insertSort(array, left, mid);
            insertSort(array, mid + 1, right);
        } else {
            mergeSort(array, left, mid);
            mergeSort(array, mid + 1, right);
        }
        merge(array, left, mid, right);
    }
}

int main() {
    int n;
    scanf("%d", &n);
    int array[n];
    for (int i = 0; i < n; i++)
        scanf("%d", &array[i]);

    mergeSort(array, 0, n - 1);

    for (int i = 0; i < n; i++)
        printf("%d ", array[i]);
    return 0;
}