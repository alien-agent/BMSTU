void bubblesort(unsigned long nel,
                int (*compare)(unsigned long i, unsigned long j),
                void (*swap)(unsigned long i, unsigned long j)) {
    int tr = nel - 1, tl = 0, boundRight, boundLeft = 0, forward = 1;

    while (tr - tl != 0) {
        boundRight = tr;
        boundLeft = tl;
        if (forward) {
            tr = boundLeft;
            for (int i = boundLeft; i < boundRight; i++) {
                if (compare(i + 1, i) == -1) {
                    swap(i + 1, i);
                    tr = i;
                }
            }
        } else {
            tl = boundRight;
            for (int i = boundRight; i > boundLeft; i--) {
                if (compare(i - 1, i) == 1) {
                    swap(i - 1, i);
                    tl = i;
                }
            }
        }
        forward = !forward;
    }
}