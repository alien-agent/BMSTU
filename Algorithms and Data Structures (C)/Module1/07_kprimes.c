#include <math.h>
#include <malloc.h>
#include "stdio.h"

int main() {
    // Initialize variables
    int k, n;
    scanf("%d", &k);
    scanf("%d", &n);
    // Now divisors_count[i] corresponds to number (i)
    n += 1;
    int *divisors_count = (int*) malloc(n * sizeof(int));
    for (int number = 0; number < n; number++) {
        divisors_count[number] = 1;
    }
    for (int number = 2; number < sqrt(n); number++) {
        if (divisors_count[number] == 1) {
            for (int multiple = 2 * number; multiple < n; multiple += number) {
                divisors_count[multiple] = divisors_count[multiple / number] + 1;
            }
        }
    }
    for (int number = 2; number < n; number++) {
        if(divisors_count[number] == k) printf("%d ", number);
    }
    free(divisors_count);
    return 0;
}