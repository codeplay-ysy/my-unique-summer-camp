//递归先序遍历
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

	void preorderTree(BTreeNode* Node)
	{
		if (Node)
		{
			cout << Node->data << " ";
			preorderTree(Node->left);//通过递归，先遍历左子树
			preorderTree(Node->right);//一直到遍历到最左下角的节点，再从下往上依次先序遍历右子树
		}
		else
		{
			return;
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
	cout << "前序遍历：" << endl;
	tree.preorderTree(root);
	cout << endl;
	tree.clear(root);
	system("pause");
	return 0;
}
