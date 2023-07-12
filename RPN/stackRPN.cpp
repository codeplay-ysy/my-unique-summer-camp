#include <stack>
#include <string>
#include <cctype>
#include <iostream>
#define priority(c)  (c == '*' || c == '/')? (1) : (0)
//宏定义一个优先级比较函数
using namespace std;

class calc
{
public:
    int eval(const string& s)
	{
		int sz = s.size();
		if (s[0] == '-') //[解释]：对负数的处理，通过减法来模拟负数
			num.push(0);

		for (int i = 0; i < sz; i++)
		{
			if (s[i] == ' ')
				continue;
			else if (isdigit(s[i])) //字符串转数字
			{
				int n = s[i] - '0';
				while (i + 1 < sz && isdigit(s[i + 1]))
					n = n * 10 + (s[++i] - '0'); 
                //括号不能少，防止溢出
				num.push(n);
			}
			else if (s[i] == '(')
			{
				op.push('(');
				flag = true;//标记进入左括号，无脑压栈 
				if (s[i + 1] == '-')
					num.push(0);
			}
			else if (s[i] == ')')
			{
				flag = false;
                // 在遇到括号里第一个运算符后，flag才会置为
                // false,但遇到括号里无运算符的情况就可能会出错，所
                // 以加个双保险
				char _operator = op.top();
				while (_operator != '(')
				{
                   	// [解释]：弹出运算符进行运算的过程
					count_push(_operator);
					op.pop();
					_operator = op.top();
				}
				op.pop();  // 把'('弹出来
			}
			else
			{
				if (op.empty() || flag)
				{
					op.push(s[i]);
					flag = false;
				}
				else 
				{
					char _operator = op.top();
					if (priority(_operator) >= priority(s[i]))
					{
						count_push(_operator);
						op.pop();
					}
					op.push(s[i]);
				}
			}
		}
		while (!op.empty()) //遍历结束后将栈里的操作符依次全部弹出
		{
			count_push(op.top());
			op.pop();
		}
		return num.top();
	}
    
    //进行栈顶两个数据的运算并压入栈中 
	void count_push(char _operator)
	{
		int tmp = 0;
		int right = num.top();
		num.pop();
		int left = num.top();
		num.pop();
		switch (_operator)
		{
		case '+':
			tmp = left + right;
			break;
		case '-':
			tmp = left - right;
			break;
		case '*':
			tmp = left * right;
			break;
		case '/':
			tmp = left / right;
			break;
		}
		num.push(tmp);
	}

	
private:
	bool flag = false;   //[作用]：用来标记是否进入左括号
	stack<int> num;	     //[作用]：存储数据的栈
	stack<char> op;      //[作用]：存储运算符的栈
};

