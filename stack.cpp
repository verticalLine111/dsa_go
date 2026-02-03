#include <iostream>
#include <vector>
#include <string>

using namespace std;
class Stack
{
private:
    vector<int> elements;     
public:
    void push(int const& element)
    {
        elements.push_back(element);
    }
    void pop()
    {
        if (elements.empty())
        {
            throw out_of_range("Stack<>::pop(): empty stack");
        }
        elements.pop_back();
    }
    int top() const
    {
        if (elements.empty())
        {
            throw out_of_range("Stack<>::top(): empty stack");
        }
        return elements.back();
    }
    bool isEmpty() const
    {
        return elements.empty();
    }
    size_t size() const
    {
        return elements.size();
    }
};


int main()
{
    vector<string> msg {"Hello", "C++", "World", "from", "VS Code", "and the C++ extension!"};
    Stack stack;
    for (const string& word : msg)
    {
        stack.push(word.length());
    }
    cout << "The lengths of the words are: ";   

    for (size_t i = 0; i < msg.size(); ++i)
    {
        cout << stack.top() << " ";
        stack.pop();
    }
    cout << endl;
}
