#include <stdio.h>
#include <malloc.h>

int main() {
    // Initialize variables
    int n1, n2;
    int *arr1, *arr2;
    scanf("%d", &n1);
    arr1 = (int *) malloc(n1 * sizeof(int));
    for (int i = 0; i < n1; i++) {
        scanf("%d", &arr1[i]);
    }
    scanf("%d", &n2);
    arr2 = (int *) malloc(n2 * sizeof(int));
    for (int i = 0; i < n2; i++) {
        scanf("%d", &arr2[i]);
    }
    // Process sequences element by element, print the greatest
    int i = 0;
    int j = 0;
    while (i < n1 && j < n2) {
        int val1 = arr1[i];
        int val2 = arr2[j];
        if (val1 < val2) {
            printf("%d ", val1);
            i++;
        } else {
            printf("%d ", val2);
            j++;
        }
    }
    // If one sequence is greater than another, print the rest of it
    for (i; i < n1; i++) {
        printf("%d ", arr1[i]);
    }
    for (j; j < n2; j++) {
        printf("%d ", arr2[j]);
    }
    // Free all memory!!!
    free(arr1);
    free(arr2);
    return 0;
}