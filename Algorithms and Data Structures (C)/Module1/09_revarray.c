void revarray(void *base, unsigned long nel, unsigned long width) {
    char temp;

    // For every element in array
    for (int i = 0; i < nel / 2; i++) {
        // Swap every bit
        // left - j-th bit of i-th array element
        // right - j-th bit of (n-i)-th array element
        for (int j = 0; j < width; j++) {
            char *left = (char *) (base + width * i + j);
            char *right = (char *) (base + width * (nel - i - 1) + j);
            temp = *left;
            *left = *right;
            *right = temp;
        }
    }
}