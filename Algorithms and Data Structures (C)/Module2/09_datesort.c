#pragma clang diagnostic push
#pragma ide diagnostic ignored "hicpp-multiway-paths-covered"

#include <stdio.h>
#include <malloc.h>

enum SortingKeys {
    Day = 32, Month = 13, Year = 61
};

struct Date {
    int Day, Month, Year;
};

int getDateKey(struct Date date, enum SortingKeys sortingKey) {
    // Get the <Date> object's specific value
    int result;
    switch (sortingKey) {
        case Day:
            result = date.Day;
            break;
        case Month:
            result = date.Month;
            break;
        case Year:
            result = date.Year - 1970;
            break;
    }
    return result;
}

struct Date *sort(struct Date *source, enum SortingKeys sortingKey, int len) {
    // Initialize variables
    struct Date *result = (struct Date *) malloc(sizeof(struct Date) * (len));
    int i, j;
    int count[sortingKey];
    for (i = 0; i < sortingKey; i++) {
        count[i] = 0;
    }
    // Initialize the resulting array and count the number of objects corresponding to each key
    for (i = 0; i < len; i++) {
        result[i] = source[i];
        int k = getDateKey(source[i], sortingKey);
        count[k]++;
    }
    for (i = 1; i < sortingKey; i++) {
        count[i] += count[i - 1];
    }
    // Do the magic
    for (j = len - 1; j >= 0; j--) {
        int k = getDateKey(source[j], sortingKey);
        i = count[k] -= 1;
        result[i] = source[j];
    }
    free(source);
    return result;
}

int main() {
    int n;
    scanf("%d", &n);
    struct Date *dates = (struct Date *) malloc(sizeof(struct Date) * n);
    for (int i = 0; i < n; i++) {
        scanf("%d %d %d", &dates[i].Year, &dates[i].Month, &dates[i].Day);
    }

    dates = sort(dates, Day, n);
    dates = sort(dates, Month, n);
    dates = sort(dates, Year, n);

    for (int i = 0; i < n; i++) {
        printf("%d %02d %02d", dates[i].Year, dates[i].Month, dates[i].Day);
        printf("\n");
    }
    free(dates);
    return 0;
}