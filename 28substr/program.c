#include <stdio.h>
#include <stdlib.h> 
#include <stdbool.h>
#include <string.h>


int strStr(char * haystack, char * needle) {
    char *result = strstr(haystack, needle);
    int pos = result - haystack;
    if (0 <= pos && pos <= strlen(haystack)) {
        return pos;
    }
    return -1;
}

int main() {
    char* str = "hello";
    printf("%d", strStr(str, "ll"));
}