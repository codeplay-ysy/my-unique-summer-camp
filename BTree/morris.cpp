//Morris遍历.
/*
一些我觉得值得一记的东西：
线索化的实质就是将二叉链表中的空指针改为指向前驱或后继的线索。
由于前驱和后继的信息只有在遍历该二叉树的时候才能得到，所以线
索化的过程就是在遍历的过程中修改空指针的过程。
1.如果cur无左孩子，cur向右移动（cur=cur->right）
2.如果cur有左孩子，找到cur左子树上最右的节点，记为mostright
3.如果pre的right指针指向空，让其指向cur，cur向左移动（cur=cur->left）
4.如果pre的right指针指向cur，让其指向空，cur向右移动（cur=cur->right）
*/
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

	void MorrisTree(BTreeNode* Node)
	{
        BTreeNode* cur = Node;
        BTreeNode* pre = NULL;
		while (cur != NULL)
	{
		pre = cur->left;
		if (pre != NULL)
		{
			while(pre->right != NULL && pre->right != cur)
			{
				pre = pre->right;
			}
			if (pre->right == NULL)
			{
				pre->right = cur;		
				cout<<cur->data<<" ";	
				cur = cur->left;
				continue;
			}
			else
			{
				pre->right = NULL;
			}
		}
		else
		{
			cout<<cur->data<<" ";
		}
		cur = cur->right;
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
	cout << "Morris遍历：" << endl;
	tree.MorrisTree(root);
	cout << endl;
	tree.clear(root);
	system("pause");
	return 0;
}
