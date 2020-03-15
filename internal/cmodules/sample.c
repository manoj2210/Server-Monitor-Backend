#include<stdio.h>
#include<stdlib.h>
void ls(char *out){
    FILE *fp = popen("ls -l", "r");
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