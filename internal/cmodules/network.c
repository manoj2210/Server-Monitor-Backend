#include<stdio.h>

void connections(char *out)
{
    int i=0;
      FILE *fp = popen("nmcli connection show", "r");
      char ch = fgetc(fp);

      while (ch != EOF)
      {
        out[i] = ch;
        ch = fgetc(fp);
        i++;
      }
    //   printf(out);
    pclose(fp);
}

void RouteTable(char *out)
{

    int i=0;
      FILE *fp = popen("netstat -arp", "r");
      char ch = fgetc(fp);

      while (ch != EOF)
      {
        out[i] = ch;
        ch = fgetc(fp);
        i++;
      }
    //   printf(out);
    pclose(fp);
}

void ActiveConn(char *out)
{
    int i=0;
      FILE *fp = popen("netstat -anp | grep LISTEN | grep tcp || udp", "r");
      char ch = fgetc(fp);

      while (ch != EOF)
      {
        out[i] = ch;
        ch = fgetc(fp);
        i++;
      }
    //   printf(out);
    pclose(fp);
}
void NoOfConnections(char *out)
 {
     int i=0;
       FILE *fp = popen("netstat -anp | grep :80 | grep ESTABLISHED|wc -l", "r");
       char ch = fgetc(fp);

       while (ch != EOF)
       {
         out[i] = ch;
         ch = fgetc(fp);
         i++;
       }
     //   printf(out);
     pclose(fp);
 }
