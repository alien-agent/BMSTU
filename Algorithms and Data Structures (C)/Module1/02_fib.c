#pragma clang diagnostic push
#pragma ide diagnostic ignored "hicpp-multiway-paths-covered"

#include "stdio.h"

/// Idea: read sequence one-by-one and try to find the pattern 1 --> 1 --> 0.
/// If current digit follows the pattern, we note this fact by incrementing variable digits_collected.
/// If pattern breaks, print all saved digits, if pattern is complete, replace "110" with "001" and print it.
int main() {
    // Initialize variables
    int n, first, second, current;
    // 0 - no digits saved, 1 - "1", 2 - "11"
    int digits_collected = 0;
    scanf("%d", &n);
    // Special case
    if (n == 1) {
        scanf("%d", &first);
        if (first == 0) printf("1");
        else printf("0 1");
        return 0;
    }
    // Manage first two digits
    scanf("%d", &first);
    scanf("%d", &second);
    if (first == 0) {
        if (second == 1) digits_collected = 2;
        else printf("1 0 ");
    } else {
        printf("0 ");
        digits_collected = 1;
    }
    for (int i = 2; i < n; i++) {
        scanf("%d", &current);
        switch (digits_collected) {
            case 0:
                if (current == 0) printf("0 ");
                else digits_collected++;
            case 1:
                if (current == 0) {
                    printf("1 0 ");
                    digits_collected = 0;
                } else digits_collected++;
            case 2:
                // We have "110", so we replace it with "001" and print "00". We MUST save the "1" for later use.
                // If we have "10" later, we can get "110" again.
                if (current == 0) {
                    printf("0 0 ");
                    digits_collected = 1;
                }
                // If we have "111", we drop first "1" and next digit MUST be "0", so we can get "110".
                else printf("1 ");
        }
    }
    // Print the rest.
    if (digits_collected == 1) {
        printf("1 ");
    } else if (digits_collected == 2) {
        printf("0 0 1 ");
    }
    return 0;
}

#pragma clang diagnostic pop