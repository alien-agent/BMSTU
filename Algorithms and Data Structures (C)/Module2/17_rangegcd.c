#include <math.h>
#include <stdlib.h>
#include <stdio.h>


int *precomputeLogarithms(int n) {
    int k = (1 << (n + 1));
    int *logarithms = (int *) malloc(k * sizeof(int));
    int i = 1;
    int j = 0;
    while (i <= n) {
        k = (1 << i);
        while (j < k) {
            logarithms[j] = i - 1;
            j = j + 1;
        }
        i = i + 1;
    }
    return logarithms;
}

int gcd(int a, int b) {
    return b ? gcd(b, a % b) : abs(a);
}

int **sparseTable_Build(int *array, int len, int *logarithms) {
    int m = logarithms[len] + 1;
    int **sparseTable = (int **) malloc(sizeof(int *) * len);
    for (int i = 0; i < len; i++) {
        sparseTable[i] = (int *) malloc(sizeof(int) * (m + 1));
        sparseTable[i][0] = array[i];
    }

    for (int j = 1; j < m; j++) {
        int k = (1 << j);
        for (int i = 0; i <= len - k; i++) {
            int kl = (1 << (j - 1));
            sparseTable[i][j] = gcd(sparseTable[i][j - 1], sparseTable[i + kl][j - 1]);
        }
    }
    return sparseTable;
}

int sparseTable_Query(int **sparseTable, int l, int r, int *logarithms) {
    int j = logarithms[r - l + 1];
    int k = (1 << j);
    return gcd(sparseTable[l][j], sparseTable[r - k + 1][j]);
}

int main(int argc, char **argv) {
    // Initialize variables
    int n;
    scanf("%d", &n);
    int *numbers = (int *) malloc(sizeof(int) * n);
    for (int i = 0; i < n; i++) scanf("%d", &numbers[i]);
    int* logarithms = precomputeLogarithms(log2(n) + 2);
    int** sparseTable = sparseTable_Build(numbers, n, logarithms);
    // Process queries
    int m, l, r;
    scanf("%d", &m);
    for(int i=0;i<m;i++){
        scanf("%d %d", &l, &r);
        printf("%d\n", sparseTable_Query(sparseTable, l, r, logarithms));
    }
    // Free everything malloc'ed
    for(int i=0;i<n;i++){
        free(sparseTable[i]);
    }
    free(sparseTable);
    free(numbers);
    free(logarithms);
    return 0;
}