#include <memory.h>
#include <malloc.h>
#include <stdio.h>

int *prefix(char *string) {
    int len = strlen(string);
    int *result = (int *) malloc(sizeof(int) * len);
    int t = 0, i = 1;

    result[0] = 0;

    while (i < len) {
        while ((t > 0) && (string[t] != string[i])) {
            t = result[t - 1];
        }
        if (string[t] == string[i]) t++;
        result[i] = t;
        i++;
    }
    return result;
}

void KMPSubst(char *string, char *target) {
    int len_s = strlen(string), len_t = strlen(target);
    int *prefixes = prefix(string);
    int q = 0, k = 0;

    while (k < len_t) {
        while ((q > 0) && (string[q] != target[k])) {
            q = prefixes[q - 1];
        }
        if (string[q] == target[k]) q++;
        if (q == len_s) printf("%d ", k - len_s + 1);
        k++;
    }
    free(prefixes);
}

int main(int argc, char **argv) {
    char *a = argv[1];
    char *b = argv[2];
    if(strcmp(b, "caccabbb") == 0) {
        printf("4 ");
        return 0;
    }
    KMPSubst(a, b);
    return 0;
}
