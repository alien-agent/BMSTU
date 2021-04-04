#pragma clang diagnostic push
#pragma ide diagnostic ignored "hicpp-multiway-paths-covered"

#include <stdio.h>
#include <malloc.h>

#define BASE 256

union Int32 {
    int x;
    unsigned char bytes[4];
};

int getKey(union Int32 number, int k) {
    // Get the <Int32> object's k-th byte
    return number.bytes[3 - k];
}

union Int32 *fixNegatives(union Int32 *source, int len) {
    // As negatives starting from the least are placed after all positives, we move the
    // negative part of the array to the beginning.
    union Int32 *result = (union Int32 *) malloc((len + 1) * sizeof(union Int32));
    int i, j = 0, negativeStart = 0;
    for (i = 0; i < len; i++) {
        if (source[i].x < 0) {
            negativeStart = i;
            break;
        }
    }
    // Place negatives at the beginning
    for (i = negativeStart; i < len; i++, j++) {
        result[j] = source[i];
    }
    // Place positives after negatives
    for (i = 0; i < negativeStart; i++, j++) {
        result[j] = source[i];
    }
    free(source);
    return result;
}

union Int32 *sort(union Int32 *source, int sortingBy, int len) {
    // Initialize variables
    union Int32 *result = (union Int32 *) malloc((len + 1) * sizeof(union Int32));
    int i, j, k;
    int count[BASE];
    for (i = 0; i < BASE; i++) {
        count[i] = 0;
    }
    for (i = 0; i < len; i++) {
        result[i] = source[i];
        k = getKey(source[i], sortingBy);
        count[k]++;
    }
    for (i = 1; i < BASE; i++) {
        count[i] += count[i - 1];
    }
    // Do the magic
    for (j = len - 1; j >= 0; j--) {
        k = getKey(source[j], sortingBy);
        i = count[k] -= 1;
        result[i] = source[j];
    }
    free(source);
    return result;
}

int main() {
    // Initialize variables
    int n;
    scanf("%d", &n);
    union Int32 *numbers = (union Int32 *) malloc((n + 1) * sizeof(union Int32));
    for (int i = 0; i < n; i++) {
        scanf("%d", &numbers[i].x);
    }
    // Sort by each byte
    for (int i = 3; i >= 0; i--) {
        numbers = sort(numbers, i, n);
    }
    numbers = fixNegatives(numbers, n);
    // Print out
    for (int i = 0; i < n; i++) {
        printf("%d\n", numbers[i].x);
    }
    free(numbers);
    return 0;
}