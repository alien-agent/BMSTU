#include <stdio.h>
#include <math.h>

int main() {
    // Initialize variables
    int n, k;
    int ind_max = -1;
    long long sum = 0;
    long long sum_max = -INFINITY;
    scanf("%d", &n);
    scanf("%d", &k);
    int arr[k];
    // Read first k into array
    for (int i = 0; i < k; i++) {
        scanf("%d", &arr[i]);
        sum += arr[i];
    }
    // Process remaining
    for (int i = k; i < n; i++) {
        if (sum > sum_max) {
            sum_max = sum;
            ind_max = i - k;
        }
        int current;
        scanf("%d", &current);
        sum += current;
        sum -= arr[i % k];
        arr[i % k] = current;
    }
    // Check last (n-th)
    if (sum > sum_max) {
        ind_max = n - k;
    }
    printf("%d", ind_max);
}