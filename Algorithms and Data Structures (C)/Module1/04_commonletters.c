#include <stdio.h>

int letterIndex(char letter) {
    if (letter < 'a') return letter - 'A';
    else return 26 + (letter - 'a');
}

// 0 = A, 1 = B, ..., 25 = Z, 26 = a, 27 = b, ..., 51 = z
char letterFromIndex(int index) {
    if (index >= 0 && index < 26) return 'A' + index;
    else return 'a' + (index - 26);
}

int main() {
    // Initialize variables
    unsigned long long s1_letters = 0;
    unsigned long long s2_letters = 0;
    // 0 - reading first string, 1 - second string
    int flag = 0;
    // Read strings
    while (1) {
        char letter;
        scanf("%c", &letter);
        if (letter == ' ') flag = 1;
        else if (letter == '\n') break;
        else if (!flag) s1_letters |= 1ULL << letterIndex(letter);
        else s2_letters |= 1ULL << letterIndex(letter);
    }
    // i-th bit == 1 in intersection ---> i'th letter is common for two strings
    unsigned long long intersection = s1_letters & s2_letters;
    for (int i = 0; i < 52; i++) {
        if (intersection & 1ULL << i) {
            printf("%c", letterFromIndex(i));
        }
    }
}