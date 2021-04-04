unsigned long peak(unsigned long nel, int (*less)(unsigned long i, unsigned long j)) {
    unsigned long index;
    unsigned long left = 0, right = nel - 1;
    // Let's use binary search :)
    while (left < right) {
        // We can't write (left+right)/2 as (left + right) may overflow, so we first divide them by 2.
        index = left / 2 + right / 2;
        if (less(index, index + 1)) left = index + 1;
        else if (less(index, index - 1)) right = index - 1;
        // If arr[index] is not less than its left and right neighbours, then it is a peak
        else return index;
    }
}