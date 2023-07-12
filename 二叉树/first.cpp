/*
二叉树的前序、中序、后序遍历的定义： 前序遍历：对任一子树，先访问根，然
后遍历其左子树，最后遍历其右子树； 中序遍历：对任一子树，先遍历其左子树
，然后访问根，最后遍历其右子树； 后序遍历：对任一子树，先遍历其左子树，
然后遍历其右子树，最后访问根。给定一棵二叉树的前序遍历和中序遍历，求其
后序遍历（提示：给定前序遍历与中序遍历能够唯一确定后序遍历）。（对于面试的
时候没能准确回答的一个回顾）
*/
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
	void create(BTreeNode* &Node)
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
