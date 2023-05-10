#include <stdio.h>
int main() {
    FILE *fp;
    char str[20] = "Hello World !";
    fp  = fopen ("data.txt", "w");
    fputs(str, fp);
    fclose(fp);
}