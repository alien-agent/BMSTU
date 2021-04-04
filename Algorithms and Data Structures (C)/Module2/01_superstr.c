#include <stdio.h>
#include <string.h>
#include <stdlib.h>

int is_suffix(char const *str, char const *suf, int str_len) {
    for (int i = 0; i < str_len; i++)
        if (strcmp(str + i, suf) == 0) return 1;
    return 0;
}

int overlap(char const *a, char const *b) {
    int a_len = strlen(a), b_len = strlen(b), overlap = 0;
    char suf[b_len + 1];

    for (int i = 1; i < b_len + 1; i++) {
        suf[0] = 0;
        strncat(suf, b, i);
        if (!is_suffix(a, suf, a_len))
            continue;
        overlap = i;
    }
    return overlap;
}

int main() {
    int stringsCount, lenTotal = 0, totalOverlapped = 0;
    scanf("%d ", &stringsCount);
    char **strings = calloc(stringsCount, sizeof(char *));
    for (int i = 0; i < stringsCount; i++) {
        strings[i] = calloc(100, sizeof(char));
        gets(strings[i]);
        lenTotal += strlen(strings[i]);
    }

    int table[stringsCount][stringsCount];
    for (int i = 0; i < stringsCount; i++) {
        for (int j = 0; j < stringsCount; j++) {
            table[i][j] = i != j ? overlap(strings[i], strings[j]) : 0;
        }
    }

    int checked[stringsCount][stringsCount];
    for (int i = 0; i < stringsCount; i++) {
        for (int j = 0; j < stringsCount; j++) {
            checked[i][j] = i == j ? 1 : 0;
        }
    }

    for (int counter = stringsCount - 1; counter--;) {
        int max = -1, maxRowIndex = -1, maxColIndex = -1;

        for (int i = 0; i < stringsCount; i++) {
            for (int j = 0; j < stringsCount; j++) {
                if (!checked[i][j] && table[i][j] > max) {
                    max = table[i][j];
                    maxRowIndex = i;
                    maxColIndex = j;
                }
            }
        }

        if (max == -1) break;

        totalOverlapped += max;
        for (int i = 0; i < stringsCount; i++) {
            checked[maxRowIndex][i] = 1;
            checked[i][maxColIndex] = 1;
        }
    }

    printf("%d", lenTotal - totalOverlapped);
    for (int i = 0; i < stringsCount; i++)
        free(strings[i]);
    free(strings);
    return 0;
}