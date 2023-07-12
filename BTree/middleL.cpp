//迭代中序遍历
#include <iostream>
using namespace std;
#include<stack>
struct BTreeNode
{
	char data;
	BTreeNode* left;
	BTreeNode* right;
};

class BTree
{
public:
	void create(BTreeNode*& Node)
	{
		char data;
		cin >> data;
		if (data != '0')
		{
			Node = new BTreeNode;
			Node->data = data;
			create(Node->left);
			create(Node->right);
		}
		else
		{
			Node = NULL;
		}
	}

	void inorderTree(BTreeNode* Node)
	{
		stack<BTreeNode*> node;
		while (true)
		{
			while (Node)
			{
				node.push(Node);
				Node = Node->left;
			}

			if (node.empty())
			{
				break;
			}
			Node = node.top();//取出栈顶元素
			node.pop();//弹出栈顶元素
			cout << Node->data << " ";
			Node = Node->right;
		}
	}

	void clear(BTreeNode*& Node)
	{
		if (Node)
		{
			clear(Node->left);
			clear(Node->right);
			delete Node;
		}
	}
};

int main()
{
	BTree tree;
	BTreeNode* root;
	tree.create(root);
	cout << "中序遍历：" << endl;
	tree.inorderTree(root);
	cout << endl;
	tree.clear(root);
	system("pause");
	return 0;
}
