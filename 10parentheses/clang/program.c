#include <stdio.h>
#include <stdlib.h> 
#include <stdbool.h>
#include <string.h>
#include <time.h>

typedef struct {
    char* data;
    int length;
    int cap;
} Stack;

Stack newStack() {
    Stack s;
    s.length = 0;
    s.cap = 5;
    s.data = (char*) malloc(s.cap * sizeof(char));
    if (s.data == NULL) {
        printf("Memory not allocated.\n"); 
        exit(EXIT_FAILURE);
    }
    return s;
}

void deleteStack(Stack* s) {
    free(s->data);
}

void pushStack(Stack* s, char c) {
    if (s->length >= s->cap) {
        // double cap
        s->cap = s->cap * 2;
        // new data
        char* newData = (char*) realloc(s->data, s->cap * sizeof(char));
        if (newData == NULL) {
            newData = (char*) malloc(s->cap * sizeof(char));
            if (newData == NULL) {
                printf("Memory not allocated.\n"); 
                exit(EXIT_FAILURE);
            }
            for (int i = 0 ; i < s->length; i++) {
                newData[i] = s->data[i];
            }
            free(s->data);
        }
        
        s->data = newData;
    }
    s->data[s->length] = c;
    s->length++;
}

void popStack(Stack* s) {
    if (s->length == 0) {
        return;
    }
    s->length--;
}

char topStack(Stack* s) {
    if (s->length == 0) {
        return 0;
    }
    return s->data[s->length - 1];
}

void printStack(Stack* s) {
    printf("Stack length: %d \n", s->length);
    for (int i = 0; i < s->length; i++) {
        printf("[%d] -> %c\n", i, s->data[i]);
    }
}

void printTopStack(Stack* s) {
    printf("Top stack: %c", topStack(s));
}

bool isEmptyStack(Stack* s) {
    return s->length == 0;
}

bool validChar(char c) {
    switch (c) {
        case '{':
        case '}':
        case '[':
        case ']':
        case '(':
        case ')':
            return true;
        }
    return false;
}

char invertChar(char c) {
    switch (c) {
    case '{':
        return '}';
    case '(':
        return ')';
    case '[':
        return ']';
    }
    return 0;
}

bool isValid(char* s) {
    Stack sdata = newStack();
    Stack* stack = &sdata;

    for (int i = 0; ; i++) {
        char c = s[i];
        if (c == '\0') {
            break;
        }
        if (validChar(c)) {
            if (invertChar(topStack(stack)) == c) {
                popStack(stack);
            } else {
                pushStack(stack, c);
            }
        } else {
            return false;
        }
    }
    deleteStack(stack);
    return isEmptyStack(stack);
}

typedef struct {
    char* in;
    bool out;
} Case;

int main()
{
    // char* input = "[[[[()()()()()()()()]]]]";
    // bool valid = isValid(input);
    // printf("Is valid %s is %s", input, valid == true ? "True": "False");
    char* line_buf = NULL;
    size_t line_buf_size = 0;
    int line_count = 0;
    ssize_t line_size;

    FILE *fp;
    fp = fopen("../test.txt", "r"); // read mode
    if (fp == NULL) {
        perror("Error while opening the file. Exit \n");
        exit(EXIT_FAILURE);
    }
    Case cases[11];
    
    /* Get the first line of the file. */
    line_size = getline(&line_buf, &line_buf_size, fp);
    cases[0].in = malloc(line_buf_size * sizeof(char));
    strcpy(cases[0].in, line_buf);
    /* Loop through until we are done with the file. */
    while (line_size >= 0)
    {
        /* Increment our line count */
        line_count++;

        /* Get the next line */
        line_size = getline(&line_buf, &line_buf_size, fp);
        // cases[line_count].size = line_size;
        cases[line_count].in = malloc(line_buf_size * sizeof(char));
        strcpy(cases[line_count].in, line_buf);
    }
    // printf("\nLine count %d\n", line_count);

    for (int i = 0; i < line_count; i++) {
        int minus = 2;
        if (i == line_count - 1) {
            minus = 1;
        }
        cases[i].out = cases[i].in[strlen(cases[i].in) - minus] == '1';
        cases[i].in[strlen(cases[i].in) - minus - 1] = '\0';
        // printf("out %d, contents: %s\n", cases[i].out, cases[i].in);
    }
    struct timespec start, end;
    clock_gettime(CLOCK_MONOTONIC, &start); 
    for (int i = 0; i < line_count; i++) {
        bool got = isValid(cases[i].in);
        if (got != cases[i].out) {
            perror("Error in case: ");
        }
    }
    clock_gettime(CLOCK_MONOTONIC, &end); 
    double time_taken; 
    time_taken = (end.tv_sec - start.tv_sec) * 1e9; 
    time_taken = time_taken + end.tv_nsec - start.tv_nsec;
    printf("Finish test in %.2f µs", time_taken * 1E-3);

    for (int i = 0; i < line_count; i++) {
        free(cases[i].in);
    }
    free(line_buf);
    
    return 0;
}
