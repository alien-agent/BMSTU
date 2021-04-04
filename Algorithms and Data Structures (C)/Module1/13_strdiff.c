#include <memory.h>
int strdiff(char *a, char *b) {
    int bit_a, bit_b;
    int len = strlen(a) < strlen(b) ? strlen(a) : strlen(b);
    // Special cases
    if(len == 0) return 0;
    else if (strcmp(a, b) == 0) return -1;
    // For every char in a and b
    for (int i = 0; i < len + 1; i++) {
        // For every bit in char
        for (int k = 0; k < 8; k++) {
            bit_a = a[i] & 1 << k;
            bit_b = b[i] & 1 << k;
            if (bit_a != bit_b) return (i * 8 + k);
        }
    }
}