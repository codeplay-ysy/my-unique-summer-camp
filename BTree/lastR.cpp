//后序递归遍历
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
		if (Node)
		{
			postorderTree(Node->left);
			postorderTree(Node->right);
			cout << Node->data << " ";
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
