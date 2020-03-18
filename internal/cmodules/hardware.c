#include<stdio.h>
#include<string.h>

void hardware(char *out){
    FILE *fp = popen("uname -i", "r");
    char ch;
    ch=fgetc(fp);
    int i=0;
    while (ch != EOF)
    {
        out[i]=ch;
        ch=fgetc(fp);
        i++;
    }
    strcat(out,";");

    fp = popen("uname -i", "r");
    ch=fgetc(fp);

    while (ch != EOF)
    {
      out[i]=ch;
      ch=fgetc(fp);
       i++;
     }



    pclose(fp);
}