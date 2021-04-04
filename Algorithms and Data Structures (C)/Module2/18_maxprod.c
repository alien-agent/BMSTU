#include <stdio.h>
#include <malloc.h>
#include <math.h>

void KadaneAlgorithm(const float *source, int len) {
    int i = 0, left = 0, right = 0, lastMinusIndex = 0;
    float currentSum = 0, maxSum = source[0];
    while (i < len) {
        currentSum += source[i];
        if (currentSum > maxSum) {
            maxSum = currentSum;
            left = lastMinusIndex;
            right = i;
        }
        i++;
        if (currentSum < 0) {
            currentSum = 0;
            lastMinusIndex = i;
        }
    }
    printf("%d %d", left, right);
}

/* Maximizing product of fractions is equivalent to maximizing
 * log(product) as log is a continuous and increasing function.
 *
 * As log(a * b) = log(a) + log(b), maximizing log(product) is equivalent
 * to maximizing sum of logarithms of each product term.
 *
 * Therefore, in order to maximize a product of fractions, we can maximize
 * the sum of fractions' logarithms => apply Kadane's algorithm.
 */
int main() {
    int n;
    scanf("%d", &n);
    float *fractionsLog = malloc(sizeof(float) * n);
    float a, b;
    for (int i = 0; i < n; i++) {
        scanf("%f/%f", &a, &b);
        fractionsLog[i] = logf(a / b);
    }
    KadaneAlgorithm(fractionsLog, n);
    free(fractionsLog);
    return 0;
}