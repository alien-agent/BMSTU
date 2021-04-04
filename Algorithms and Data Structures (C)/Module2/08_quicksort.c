#include <stdio.h>

void swap(int i, int j, int *array) {
    int temp = array[i];
    array[i] = array[j];
    array[j] = temp;
}

void selectsort(int low, int high, int *array) {
    int j = high, i, k;
    while (j > low) {
        k = j;
        i = j - 1;
        while (i >= 0) {
            if (array[k] < array[i])
                k = i;
            i--;
        }
        swap(j, k, array);
        j--;
    }
}

int partition(int low, int high, int *array) {
    int i = low, j = low;
    while (j < high) {
        if (array[j] < array[high]) {
            swap(i, j, array);
            i++;
        }
        j++;
    }
    swap(i, high, array);
    return i;
}

void quicksort(int low, int high, int m, int *array) {
    while (low < high) {
        if (m >= (high - low)) {
            selectsort(low, high, array);
            return;
        }
        int q = partition(low, high, array);
        if ((low - q) < (high - q)) {
            quicksort(low, q - 1, m, array);
            low = q + 1;
        } else {
            quicksort(q + 1, high, m, array);
            high = q - 1;

        }
    }
}

int main() {
    int arraySize, threshold;
    scanf("%d %d", &arraySize, &threshold);
    int array[arraySize];
    for (int i = 0; i < arraySize; i++)
        scanf("%d", &array[i]);

    quicksort(0, arraySize - 1, threshold, array);

    for (int i = 0; i < arraySize; i++)
        printf("%d ", array[i]);
    return 0;
}