#include <stdio.h>
#include <memory.h>

int gcd(int a, int b) {
    int temp;
    while (b) {
        temp = a % b;
        a = b;
        b = temp;
    }
    return a;
}

void prefix(char *s) {
    int len = strlen(s), t, q, k, flag;
    int pi[len];
    for (int i = 0; i < len; i++) {
        pi[i] = 0;
    }

    pi[0] = t = 0;
    for (int i = 1; i < len; i++) {
        while ((t > 0) && (s[t] != s[i])) {
            t = pi[t - 1];
        }
        if (s[t] == s[i])
            t++;
        pi[i] = t;
    }

    for (int i = 1; i < len + 1; i++) {
        if (pi[i - 1] != 0) {
            q = gcd(i, pi[i - 1]);
            flag = 0;
            for (k = 0; k + q < i; k++) {
                if (s[k] != s[k + q]) {
                    flag = 1;
                }
            }
            if (flag == 0)
                printf("%d %d\n", i, i / gcd(i, pi[i - 1]));
        }
    }
}

int main(int argc, char **argv) {
    prefix(argv[1]);
    return 0;
}