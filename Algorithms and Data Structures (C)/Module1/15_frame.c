#include <stdlib.h>
#include "stdio.h"
#include "string.h"

void print_chars(char ch, int times) {
    for (int i = 0; i < times; i++) {
        printf("%c", ch);
    }
}

int main(int argc, char *argv[]) {
    // Check if enough arguments
    if (argc != 4) {
        printf("Usage: frame <height> <width> <text>");
        return 0;
    }
    int height, width;
    height = atoi(argv[1]);
    width = atoi(argv[2]);
    char *text = argv[3];
    int length = strlen(text);
    // Check if it is possible to fit the text in the frame
    if (height < 3 || length > width - 2) {
        printf("Error");
        return 0;
    }
    // Determine frame's position
    int frame_y = height / 2 - !(height % 2);
    int space_count_left = (width - length) / 2 - 1;
    int space_count_right = width - length - space_count_left - 2;
    // Draw frame's top
    print_chars('*', width);
    printf("\n");
    // Draw frame
    for (int i = 0; i < height - 2; i++) {
        printf("*");
        if (i == frame_y - 1) {
            print_chars(' ', space_count_left);
            printf("%s", text);
            print_chars(' ', space_count_right);
        } else {
            printf("%*c", width - 2, ' ');
        }
        printf("*\n");
    }
    // Draw frame's bottom
    print_chars('*', width);
}
