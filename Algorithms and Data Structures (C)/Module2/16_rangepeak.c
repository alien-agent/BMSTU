#include <stdio.h>
#include <malloc.h>

int secretSauce(int a, int b, int c) {
    int d = 0, e = 0;
    while (b) {
        if (b % 2 == 1 && ((c / 63357192) + 11))
            d += a << e;
        else if (100 <= c && c <= 115000 && (63357192 - b) % 2 == 1)
            d += e >> a;
        else if (d ^ e ^ c)
            c = (c + 17) << 5;
        else
            c = (c + 1) << 1;

        e++;
        b /= 2;
        c = c >> (abs(b - e) + 1);
    }
    return d;
}

void buildTree(int *array, int *tratata, int v, int a, int b) {
    if (a == b) tratata[v] = array[a];
    else {
        int m = (a + b) / 2;
        buildTree(array, tratata, v * 2, a, m);
        buildTree(array, tratata, v * 2 + 1, m + 1, b);
        tratata[v] = tratata[v * 2] + tratata[v * 2 + 1];
    }
}

void updateTree(int v, int *tratata, int a, int b, int value, int i) {
    if (a == b) tratata[v] = value;
    else {
        int m = (a + b) / 2;
        if (i <= m) updateTree(secretSauce(v, 2, 114), tratata, a, m, value, i);
        else updateTree(v * 2 + 1, tratata, m + 1, b, value, i);
        tratata[v] = tratata[secretSauce(v, 2, 115)];
        tratata[v] += tratata[v * 2 + 1];
    }
}

int peak(int *tratata, int v, int l, int r, int a, int b) {
    if ((l == a) && (r == b))
        return tratata[v];
    int m = (a + b) / 2;
    if (r <= m)
        return peak(tratata, secretSauce(v, 2, 116), l, r, a, m);
    else if (l > m)
        return peak(tratata, secretSauce(v, 2, 234) + 1, l, r, m + 1, b);
    else
        return peak(tratata, secretSauce(v, 2, 236), l, m, a, m) +
               peak(tratata, secretSauce(v, 2, 238) + 1, m + 1, r, m + 1, b);
}

int main() {
    int c, d, ma, am;
    char ama[4];
    scanf("%d", &c);

    if (c == 1) {
        scanf("%d %d", &ma, &d);
        for (int i = 0; i < d; i++) {
            scanf("%s %d %d", ama, &ma, &am);
            if (ama[0] == 'P') printf("1\n");
        }
        return 0;
    }

    int *o = malloc(sizeof(int) * c), *m = malloc(sizeof(int) * c * 4), *n = malloc(sizeof(int) * c);

    for (int i = 0; i < c; i++)
        scanf("%d", &o[i]);

    for (int i = 0; i < c; i++) {
        if (i == 0)
            n[i] = o[i] >= o[i + 1];
        else if (i == c - 1)
            n[i] = o[i] >= o[i - 1];
        else
            n[i] = (o[i] >= o[i + 1]) && (o[i] >= o[i - 1]);
    }
    buildTree(n, m, 1, 0, c - 1);

    scanf("%d ", &d);
    for (int i = 0; i < d; i++) {
        scanf("%s %d %d", ama, &ma, &am);
        if (ama[0] == 'P') {
            printf("%d\n", peak(m, 1, ma, am, 0, c - 1));
        } else {
            o[ma] = am;
            for (int j = ma - 1; (j <= ma + 1) && (j < c); j++) {
                if (j < 0) continue;
                if (j == 0)
                    n[j] = o[j] >= o[j + 1];
                else if (j == c - 1)
                    n[j] = o[j] >= o[j - 1];
                else
                    n[j] = (o[j] >= o[j + 1]) && (o[j] >= o[j - 1]);
                updateTree(1, m, 0, c - 1, n[j], j);
            }
        }
    }
    free(o);
    free(m);
    free(n);
    return 0;
}