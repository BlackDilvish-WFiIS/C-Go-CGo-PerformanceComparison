#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>

void loadData(char ***strings) {
    FILE *f;
    int i = 0;
    size_t len;
    char *line = NULL;

    f = fopen("../strings.txt", "r");
    while (getline(&line, &len, f) != -1) {
        line[strcspn(line, "\n")] = 0;  // remove newline
        *strings = realloc(*strings, (i+1) * sizeof(char *));
        (*strings)[i] = malloc(strlen(line) + 1);
        strcpy((*strings)[i], line);
        i++;
    }
    free(line);
    fclose(f);
}

char* lcs(char *str1, char *str2) {
    int m = strlen(str1);
    int n = strlen(str2);
    int max = 0;
    int end;
    int len[2][n];
    int curr_row = 0;

    for (int i=0; i<=m; i++){
        for (int j=0; j<=n; j++){
            if (i == 0 || j == 0){
                len[curr_row][j] = 0;
            }
            else if (str1[i-1] == str2[j-1]){
                len[curr_row][j] = len[1-curr_row][j-1] + 1;
                if (len[curr_row][j] > max){
                    max = len[curr_row][j];
                    end = i-1;
                }
            }
            else {
                len[curr_row][j] = 0;
            }
        }
        curr_row = 1 - curr_row;
    }
    if (max == 0){
        return "";
    }
    char *resStr = (char*)malloc(sizeof(char)*(max+1));
    strncpy(resStr, &str1[end-max+1], max);
    resStr[max] = '\0';
    return resStr;
}

int main() {
    char **strings = NULL;
    loadData(&strings);

    clock_t start, end;
    double cpu_time_used;
    start = clock();
    char * res = lcs(strings[0], strings[1]);
    end = clock();
    cpu_time_used = ((double) (end - start)) / CLOCKS_PER_SEC;
    printf("LCS execution time: %.6f seconds.\n", cpu_time_used);
    printf("%s\n", res);

    return 0;
}
