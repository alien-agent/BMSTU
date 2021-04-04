#include "stdio.h"

int main(int argc, char **argv) {
    // Initialize variables
    int n;
    long long start, end;
    long long last_start, last_end;
    scanf("%d", &n);
    scanf("%lld", &last_start);
    scanf("%lld", &last_end);
    // Process
    for (int i = 1; i < n; i++) {
        scanf("%lld %lld", &start, &end);
        // If current segment does intersect with previously processed, then we can merge them and continue.
        if (start <= last_end + 1) {
            // last_end = max(last_end, end);
            last_end = last_end > end ? last_end : end;
        // If current segment does NOT intersect with previously processed, then previous segment is finished and
        // we should start a new one.
        } else {
            printf("%lld %lld\n", last_start, last_end);
            last_start = start;
            last_end = end;
        }
    }
    printf("%lld %lld", last_start, last_end);
    return 0;
}