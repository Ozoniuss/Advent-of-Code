#include <iostream>
#include <fstream>
#include <vector>
#include <algorithm>
#include <set>

using namespace std;

vector<string> getLines()
{
    ifstream f{"input.txt"};
    string line;

    vector<string> lines{};

    if (f.is_open())
    {
        while (getline(f, line))
        {
            lines.push_back(line);
        }

        f.close();
    }

    return lines;
}

int main()
{
    vector<string> lines = getLines();

    vector<int> left{};
    vector<int> right{};

    for (const auto &l : lines)
    {
        string lefts = l.substr(0, 5);
        string rights = l.substr(8, 5);

        int ll = stoi(lefts);
        int rr = stoi(rights);

        left.push_back(ll);
        right.push_back(rr);
    }

    sort(left.begin(), left.end());
    sort(right.begin(), right.end());

    int total = 0;
    for (uint64_t i = 0; i < left.size(); i++)
    {
        total += abs(left[i] - right[i]);
    }
    cout << total;

    multiset<int> rightelements{};
    for (const auto &r : right)
    {
        rightelements.insert(r);
    }

    total = 0;
    for (const auto &l : left)
    {
        total += rightelements.count(l) * l;
    }
    cout << "\n"
         << total;
}