#include<iostream>
using namespace std;
int main(){
    long long r,c,a;
    cin >> r >> c >> a ;
    long long s1=r/a,s2=c/a;
    if(r%a)
        s1++;
    if(c%a)
        s2++;
        cout << s1*s2;
    return 0;
}
