package main

/*
#include <string.h>
#include <stdlib.h>

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
*/
import "C"
import (
	"bufio"
	"fmt"
	"os"
	"time"
	"unsafe"
)

func loadData() ([]string, error) {
	file, err := os.Open("../strings.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	strings, _ := loadData()

	cstr1 := C.CString(strings[0])
	cstr2 := C.CString(strings[1])
	start := time.Now()
	result := C.lcs(cstr1, cstr2)
	elapsed := time.Since(start)
	if result != nil {
		fmt.Println(C.GoString(result))
		C.free(unsafe.Pointer(result))
	} else {
		fmt.Println("No common substring")
	}
	fmt.Printf("LCS execution time: %.6f seconds\n\n", elapsed.Seconds())
	C.free(unsafe.Pointer(cstr1))
	C.free(unsafe.Pointer(cstr2))
}