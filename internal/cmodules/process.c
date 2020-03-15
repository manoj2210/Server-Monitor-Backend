#include<stdio.h>
#include<stdlib.h>
void process(char *out){
    FILE *fp = popen("top -n 1 -b | head -n 30", "r");
    char ch;
    ch=fgetc(fp);
    int i=0;
    while (ch != EOF)
    {
        out[i]=ch;
        ch=fgetc(fp);
        i++;
    }
    pclose(fp);
}