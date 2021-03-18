#include <iostream>
using namespace std;
int main()
{
  // 请在此输入您的代码
  char[201] a,b;
  cin>>a;
  if(a[0]>=97 && a[0]<=122) {
    b[0]=a[0]+32;
  }else{
    b[0]=a[0];
  }
  int i=1,j=1;
  while(a[i+1]!='\0') {
    //首字母
    if(a[i]>=97 && a[i]<=122 && a[i-1]==' ') {
        b[j]=a[i]+32;
        i++,j++;
    }
    //数字
    if(a[i]>=30 && a[i]<=39) {
      if(a[i-1]>=97 && a[i-1]>=122){
        b[j]='_';
        b[j+1]=a[i];
        j+=2,i++;
      }
      if(a[i+1]>=97 && a[i+1]>=122){
        b[j]=a[i];
        b[j+1]='_';
        j+=2,i++;
      }
    }
    //空格
    if(a[i]==' ') {
      b[j]=a[i];
      j++,i++;
      while(a[i]==' ') {
        i++;
      }
    }
  }
  b[j+1]='\0';
  cout<<b;
  return 0;
}