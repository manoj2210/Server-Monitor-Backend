#include<stdio.h>

int main(){
char numOfCPU[100];
FILE *fp = popen("free -m", "r");
char ch;
int m=0;

fgets(numOfCPU,100,fp);
while (numOfCPU != NULL)
{
        printf ("%s", numOfCPU);
        fgets(numOfCPU,100,fp);
}

pclose(fp);
}

