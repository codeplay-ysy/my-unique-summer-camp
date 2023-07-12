//递归中序遍历
#include <iostream>
using namespace std;
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
		if (Node)
		{
			inorderTree(Node->left);
			cout << Node->data << " ";//访问当前节点.
			inorderTree(Node->right);
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
