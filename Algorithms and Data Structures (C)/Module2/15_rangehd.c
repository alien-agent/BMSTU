#include <stdio.h>
#include <stdlib.h>

#define max(a, b) ((a)>(b)?(a):(b))

void buildSegmentTree(int *array, int a, int b, int v, int *tree) {
    if (a != b) {
        int mid = (a + b) / 2;
        buildSegmentTree(array, a, mid, 2 * v + 1, tree);
        buildSegmentTree(array, mid + 1, b, v * 2 + 2, tree);
        tree[v] = max(tree[v * 2 + 1], tree[v * 2 + 2]);
    } else tree[v] = array[a];
}

int segmentTreeQuery(int *tree, int left, int right, int segmentLeft, int segmentRight, int v) {
    if (left != segmentLeft || !(right == segmentRight)) {
        int mid = (segmentLeft + segmentRight) / 2;
        return right <= mid ? segmentTreeQuery(tree, left, right, segmentLeft, mid, 2 * v + 1) : left > mid ? 
                              segmentTreeQuery(tree, left, right,mid + 1, segmentRight,2 * v + 2) : max(
                                      segmentTreeQuery(tree, left, mid, segmentLeft, mid, 2 * v + 1),segmentTreeQuery(tree, mid + 1,
                                                                                                                            right, mid + 1,
                                                                                                                            segmentRight,
                                                                                                                            2 * v + 2));
    } else return tree[v];
}

void updateSegmentTree(int v, int a, int b, int i, int value, int *tree) {
    if (a != b) {
        int mid = (a + b) / 2;
        if (i <= mid)
            updateSegmentTree(v * 2 + 1, a, mid, i, value, tree);
        else
            updateSegmentTree(v * 2 + 2, mid + 1, b, i, value, tree);
        tree[v] = max(tree[v * 2 + 1], tree[v * 2 + 2]);
    } else tree[v] = value;
}

int main() {
    int len, commandsCount, a, b;
    scanf("%d", &len);
    int *array = malloc(sizeof(int) * len);
    int *tree = malloc(sizeof(int) * len * 4);
    char command[4];

    for (int i = 0; i < len; i++)
        scanf("%d", &array[i]);
    buildSegmentTree(array, 0, len - 1, 0, tree);

    scanf("%d", &commandsCount);
    for (int i = 0; i < commandsCount; i++) {
        scanf("%s %d %d", command, &a, &b);
        if (command[0] == 'U')
            updateSegmentTree(0, 0, len - 1, a, b, tree);
        else
            printf("%d\n", segmentTreeQuery(tree, a, b, 0, len - 1, 0));
    }

    free(array);
    free(tree);
    return 0;
}