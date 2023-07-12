//后序迭代遍历.
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

	void postorderTree(BTreeNode* Node)
	{
		stack<BTreeNode*> node;
		BTreeNode* cur = Node;
		BTreeNode* pre = NULL;
		while (cur || !node.empty())
		{
			while (cur)
			{
				node.push(cur);
				cur = cur->left;
			}
			cur = node.top();
			if (!cur->right || pre == cur->right)
			{
				cout << cur->data << " ";
				node.pop();
				pre = cur;
				cur = NULL;
			}
			else
			{
				cur = cur->right;
			}
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
	cout << "后续遍历：" << endl;
	tree.postorderTree(root);
	cout << endl;
	tree.clear(root);
	system("pause");
	return 0;
}
