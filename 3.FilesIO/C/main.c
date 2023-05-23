#include <stdio.h>
#include <time.h>

int main() {
    FILE* file;
    const char* filename = "../input.txt";
    char buffer[100];
    clock_t start, end;
    double cpu_time_used;

    start = clock();
    file = fopen(filename, "r");
    if (file == NULL)
    {
        printf("Error opening file.\n");
        return 1;
    }

    while (fgets(buffer, sizeof(buffer), file) != NULL)
    {
        printf("%s", buffer);
    }

    end = clock();
    fclose(file);

    cpu_time_used = ((double) (end - start)) / CLOCKS_PER_SEC;
    printf("Execution time: %.6f seconds.\n", cpu_time_used);

    return 0;
}