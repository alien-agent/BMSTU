int maxarray(void *base, unsigned long nel, unsigned long width, int (*compare)(void *a, void *b)) {
    int max_index = 0;
    int *max_pointer = base;

    for (int i = 1; i < nel; ++i) {
        if (compare((base + i * width), max_pointer) > 0) {
            max_pointer = (base + i * width);
            max_index = i;
        }
    }

    return max_index;
}
