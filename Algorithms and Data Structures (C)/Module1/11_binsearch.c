unsigned long binsearch(unsigned long nel, int (*compare)(unsigned long i)) {
    unsigned long index;
    unsigned long left = 0, right = nel;

    while (left <= right) {
        index = (left + right) / 2;
        switch (compare(index)) {
            case 0:
                return index;
                break;
            case -1:
                left = index + 1;
                break;
            case 1:
                right = index - 1;
                break;
        }
    }
    return nel;
}