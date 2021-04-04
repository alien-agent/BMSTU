#include <stdio.h>
#include <math.h>

int main() {
    // Initialize variables
    int m, n;
    scanf("%d", &m);
    scanf("%d", &n);
    int row_max_values[m], row_max_indices[m], row_max_repeats[m];
    int col_min_values[n], col_min_repeats[n];
    for (int i = 0; i < n; i++) {
        col_min_values[i] = INFINITY;
        col_min_repeats[i] = 0;
    }
    // Process every element
    for (int i = 0; i < m; i++) {
        row_max_values[i] = -INFINITY;
        for (int j = 0; j < n; j++) {
            int x;
            scanf("%d", &x);
            // Check for row maximum
            if (x > row_max_values[i]) {
                row_max_values[i] = x;
                row_max_indices[i] = j;
                row_max_repeats[i] = 0;
            } else if (x == row_max_values[i]) row_max_repeats[i] = 1;
            // Check for col minimum
            if (x < col_min_values[j]) {
                col_min_values[j] = x;
                col_min_repeats[j] = 0;
            } else if (x == col_min_values[j]) col_min_repeats[j] = 1;
        }
    }

    // If row maximum or column minimum repeats, then element is not a saddle
    // Else if row maximum == column minimum, then element is a saddle
    for (int row_index = 0; row_index < m; row_index++) {
        int col_index = row_max_indices[row_index];
        if (!row_max_repeats[row_index] && !col_min_repeats[col_index]){
            if(row_max_values[row_index] == col_min_values[col_index]){
                printf("%d %d", row_index, col_index);
                return 0;
            }
        }
    }
    printf("%s", "none");
    return 0;
}