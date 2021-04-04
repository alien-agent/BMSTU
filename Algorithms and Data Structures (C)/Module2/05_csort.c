#include <stdio.h>
#include <string.h>
#include <stdlib.h>

void csort(char *src, char *dest, int n, int k, int len) {
    int key[n], i, j, h;

    for (i = 0, j = 0, h = 0; i < len; i++) {
        if (src[i] != ' ')
            j++;
        if (src[i] != ' ' && src[i + 1] == ' ') {
            key[h] = j;
            h++;
            j = 0;
        }
    }

    key[h++] = j;
    int count[n];
    for (i = 0; i < n; i++)
        count[i] = 0;
    for (j = 0; j < n - 1; j++) {
        for (i = j + 1; i < n; i++) {
            if (key[i] < key[j])
                count[j] += key[i] + 1;
            else
                count[i] += key[j] + 1;
        }
    }
    
    for (i = 0; i < n + k; i++)
        dest[i] = ' ';
    for (i = 0, h = 0; i <= len; i++) {
        if (src[i] != ' ') {
            for (j = 0; j < key[h]; j++)
                dest[count[h] + j] = src[i + j];
            i += key[h];
            h++;
        }
    }
    
    dest[n + k - 1] = '\0';
    printf("%s\n", dest);
}

int main() {
    char src[100000];
    gets(src);
    int len = strlen(src), n = 1, k = 0;
    for (int i = 0; i < len; i++) {
        if (src[i] == ' ' && src[i + 1] != ' ')
            n++;
        if (src[i] != ' ')
            k++;
    }
    if (src[0] == ' ')
        n--;
    if (src[len - 1] == ' ')
        n--;

    char *dest = malloc((k + n) * sizeof(char));
    csort(src, dest, n, k, len);

    free(dest);
    return 0;
}