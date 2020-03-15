#include<stdio.h>
#include<string.h>
#include<stdlib.h>
void task(char *out){
    char arr[300];
    FILE *fp = popen("top -n 1 -b | head -n 3", "r");
    char ch;
    ch=fgetc(fp);
    int i=0;
    while (ch != EOF)
    {
        arr[i]=ch;
        ch=fgetc(fp);
        i++;
    }
    i=1;
    int j=0;
    char *token;
    token = strtok(arr, "\n");
    while( token != NULL ) {
      token = strtok(NULL, " ");
      if(i==2){
        strcpy(out,token);
        strcpy(out,";");
      }
      else if(i==4){
        strcpy(out,token);
        strcpy(out,";");
      }
      else if(i==6){
        strcpy(out,token);
        strcpy(out,";");
      }
      else if(i==10){
        strcpy(out,token);
        strcpy(out,";");
      }
      else if(i==12){
        strcpy(out,token);
        strcpy(out,";");
      }
      else if(i==14){
        strcpy(out,token);
        strcpy(out,";");
      }
      else if(i==18){
        strcpy(out,token);
        strcpy(out,";");
      }
      i++;
   }
    pclose(fp);
}