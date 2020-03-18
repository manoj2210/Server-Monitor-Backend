#include <stdio.h>
#include <string.h>

void OSInfo(char *out)
{
  FILE *fp = popen("uname -i", "r");
  char ch;
  ch = fgetc(fp);
  int i = 0;
  while (ch != EOF)
  {
    out[i] = ch;
    ch = fgetc(fp);
    i++;
  }
  
  fp = popen("uname -s", "r");
  ch = fgetc(fp);

  while (ch != EOF)
  {
    out[i] = ch;
    ch = fgetc(fp);
    i++;
  }

  fp = popen("uname -n", "r");
  ch = fgetc(fp);

  while (ch != EOF)
  {
    out[i] = ch;
    ch = fgetc(fp);
    i++;
  }

  fp = popen("uname -v", "r");
  ch = fgetc(fp);

  while (ch != EOF)
  {
    out[i] = ch;
    ch = fgetc(fp);
    i++;
  }

  fp = popen("uname -o", "r");
  ch = fgetc(fp);

  while (ch != EOF)
  {
    out[i] = ch;
    ch = fgetc(fp);
    i++;
  }
  // printf(out);

  pclose(fp);
}

void CPUInfo(char *out){
      int i=0;
      FILE *fp = popen("lscpu|head -24", "r");
      char ch = fgetc(fp);

      while (ch != EOF)
      {
        out[i] = ch;
        ch = fgetc(fp);
        i++;
      }
    pclose(fp);
}

void RAMInfo(char *out){
    int i=0;
    FILE *fp = popen("free -m", "r");
    char ch = fgetc(fp);

    while (ch != EOF)
      {
        out[i] = ch;
        ch = fgetc(fp);
        i++;
      }
    pclose(fp);

}

void NetworkInfo(char *out){
    int i=0;
    FILE *fp = popen("ip link show", "r");
    char ch = fgetc(fp);

    while (ch != EOF)
      {
        out[i] = ch;
        ch = fgetc(fp);
        i++;
      }
    pclose(fp);
}

