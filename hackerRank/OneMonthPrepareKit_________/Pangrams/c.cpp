#include <bits/stdc++.h>

using namespace std;

/*
 * Complete the 'pangrams' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

string pangrams(string s) {
    int n = s.size();
    vector<bool> v(26, false);//
    for(int i=0; i<n; i++){
        if(s[i] >= 'a' && s[i] <= 'z'){
            v[s[i]-'a'] = true;// 
        }
        else if(s[i] >= 'A' && s[i] <= 'Z'){
            v[s[i]-'A'] = true;
        }
    }
    for(int i=0; i<26; i++){
        if(v[i] == false){
            return "not pangram";
        }
    }
    return "pangram";
    //log O(n)

}

int main()
{
    ofstream fout(getenv("OUTPUT_PATH"));

    string s;
    getline(cin, s);

    string result = pangrams(s);

    fout << result << "\n";

    fout.close();

    return 0;
}
