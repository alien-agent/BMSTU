#include <stdio.h>
#include <stdlib.h>
#include <memory.h>

#define min(a, b) (((a)<(b))?(a):(b))
#define max(a, b) (((a)>(b))?(a):(b))

typedef struct Stack_tag {
    int data[100000];
    int size;
} Stack_t;


Stack_t stack;

void stack_push(int value) {
    stack.data[stack.size] = value;
    stack.size++;
}

void stack_add() {
    stack.data[stack.size - 2] += stack.data[stack.size - 1];
    stack.size--;
}

void stack_sub() {
    stack.data[stack.size - 2] = stack.data[stack.size - 1] - stack.data[stack.size - 2];
    stack.size--;
}

void stack_mul() {
    stack.data[stack.size - 2] *= stack.data[stack.size - 1];
    stack.size--;
}

void stack_div() {
    stack.data[stack.size - 1] = stack.data[stack.size - 1] / stack.data[stack.size - 2];
    stack.size--;
}

void stack_max() {
    stack.data[stack.size - 2] = max(stack.data[stack.size - 2], stack.data[stack.size - 1]);
    stack.size--;
}

void stack_min() {
    stack.data[stack.size - 2] = min(stack.data[stack.size - 2], stack.data[stack.size - 1]);
    stack.size--;
}

void stack_neg() {
    stack.data[stack.size - 1] = -stack.data[stack.size - 1];
}

void stack_dup() {
    stack.data[stack.size] = stack.data[stack.size - 1];
    stack.size++;
}

void stack_swap() {
    int temp = stack.data[stack.size - 2];
    stack.data[stack.size - 2] = stack.data[stack.size - 1];
    stack.data[stack.size - 1] = temp;
}


int main() {
    int n;
    char *command = malloc(sizeof(char) * 20);
    int len = -1;
    stack.size = 0;

    scanf("%d", &n);
    fgets(command, 21, stdin); // Skip first input

    for (int i = 0; i < n; i++) {
        fgets(command, 21, stdin);
        command[strlen(command) - 1] = '\0';

        if (strlen(command) >= 7) {
            int val = atoi(command + 5);
            stack_push(val);
        } else if (strcmp(command, "ADD") == 0) {
            stack_add();
        } else if (strcmp(command, "SUB") == 0) {
            stack_sub();
        } else if (strcmp(command, "MUL") == 0) {
            stack_mul();
        } else if (strcmp(command, "DIV") == 0) {
            stack_div();
        } else if (strcmp(command, "MAX") == 0) {
            stack_max();
        } else if (strcmp(command, "MIN") == 0) {
            stack_min();
        } else if (strcmp(command, "NEG") == 0) {
            stack_neg();
        } else if (strcmp(command, "DUP") == 0) {
            stack_dup();
        } else if (strcmp(command, "SWAP") == 0) {
            stack_swap();
        }
    }
    printf("%d", stack.data[stack.size - 1]);
    free(command);
    return 0;
}