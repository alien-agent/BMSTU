#include <stdio.h>
#include <math.h>

int isPowerOfTwo(int i) {
    return !(i & (i - 1));
}

int solve(int n, int *array) {
    int result = 0;
    for (int j = 1; j < pow(2, n); j++) {
        int sum = 0;
        int p = j;
        for (int i = 0; i < n && p; i++) {
            if (p % 2 == 1) sum += array[i];
            p /= 2;
        }
        if (isPowerOfTwo(sum) && sum > 0) result++;
    }
    return result;
}

int main() {
    int n;
    scanf("%d", &n);
    int array[n];
    for (int i = 0; i < n; i++) scanf("%d", &array[i]);

    printf("%d", solve(n, array));
    return 0;
}
