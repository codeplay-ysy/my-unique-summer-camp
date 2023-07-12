//递归创建二叉树
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
	void create(BTreeNode* &Node)//创建二叉树，先序创建
	{
		char data;
		cin >> data;
		if (data != '0')
		{
			Node = new BTreeNode;
			Node->data = data;
			create(Node->left);//通过递归，先创建左子树，再创建右子树
			create(Node->right);//一直到创建出最左下角的节点，再从下往上依次建右子树
		}
		else
		{
			Node = NULL;
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
	cout << "二叉树创建完成！" << endl;
    system("pause");
	tree.clear(root);
	cout << "二叉树清理完成！" << endl;
	system("pause");
	return 0;
}
