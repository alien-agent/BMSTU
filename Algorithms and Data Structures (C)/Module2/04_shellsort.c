void shellsort(unsigned long nel,
               int (*compare)(unsigned long i, unsigned long j),
               void (*swap)(unsigned long i, unsigned long j)) {
    int previousFib = 0, currentFib = 1;
    while (currentFib < nel) {
        currentFib += previousFib;
        previousFib = currentFib - previousFib;
    }
    int i = 0, j = 0;
    while (currentFib > 0) {
        for (i = currentFib; i < nel; i++) {
            j = i;
            while ((j >= currentFib) && compare(j - currentFib, j) == 1) {
                swap(j, j - currentFib);
                j -= currentFib;
            }
        }
        previousFib = currentFib - previousFib;
        currentFib -= previousFib;
        if (previousFib == 0 && currentFib == 1) break;
    }
}
