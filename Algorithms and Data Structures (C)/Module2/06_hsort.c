#include <stdio.h>
#include <malloc.h>
#include <memory.h>

const char TARGET_CHARACTER = 'a';
int *indices;

int compare(const void *a, const void *b) {
    if (a == b) return 0;
    return a < b ? -1 : 1;
}

void swap(int *source, int i, int j) {
    int temp = source[i];
    source[i] = source[j];
    source[j] = temp;
}

void heapify(int source[], int i, int len) {
    int largestIndex = i;
    int leftIndex = 2 * i + 1;
    int rightIndex = 2 * i + 2;
    if (leftIndex < len && source[leftIndex] > source[largestIndex])
        largestIndex = leftIndex;
    if (rightIndex < len && source[rightIndex] > source[largestIndex])
        largestIndex = rightIndex;
    if (largestIndex != i) {
        swap(source, i, largestIndex);
        swap(indices, i, largestIndex);
        heapify(source, largestIndex, len);
    }
}

void heapSort(int arr[], int len) {
    for (int i = len / 2 - 1; i >= 0; i--)
        heapify(arr, i, len);
    for (int i = len - 1; i >= 0; i--) {
        swap(arr, 0, i);
        swap(indices, 0, i);
        heapify(arr, 0, i);
    }
}

void hsort(void *base, size_t nel, size_t width, int (*compare)(const void *a, const void *b)) {
    heapSort((int *) base, nel);
}

char *readline(int *countPtr) {
    // Reads string of arbitrary length and returns it. Sets the value at countPtr
    // to the number of TARGET_CHARACTER in the string.
    int allocated_size = 16;
    int read_size = 0;
    int targetCharacterCount = 0;
    char *buffer = (char *) malloc(allocated_size);

    while (1) {
        char input = getchar();
        if (input == '\n') break;
        // We count the number of letter's occurrences here :)
        if (input == TARGET_CHARACTER) targetCharacterCount++;
        if (read_size == allocated_size - 1) {
            buffer = realloc(buffer, allocated_size * 1.5);
        }
        buffer[read_size++] = input;
    }
    buffer[read_size] = '\0';
    *countPtr = targetCharacterCount;
    return buffer;
}

int main() {
    // strings[] store strings???
    // indices store strings' resulting indices
    // stringValues stores the number of TARGET_CHARACTERs in each string
    // Initialize variables
    int n;
    scanf("%d", &n);
    getchar(); // Skip 'enter' character
    int *stringValues = malloc(sizeof(int) * n);
    char **strings = malloc(sizeof(char *) * n);
    indices = calloc(sizeof(int), n);

    // Read each string into a buffer, allocate memory for it and copy from buffer into array
    for (int i = 0; i < n; i++) {
        strings[i] = readline(&stringValues[i]);
        indices[i] = i;
    }
    heapSort(stringValues, n);
    // Print out the results and !!!free every string
    for (int i = 0; i < n; i++) {
        printf("%s\n", strings[indices[i]]);
        free(strings[indices[i]]);
    }
    free(strings);
    free(stringValues);
    free(indices);
    return 0;
}